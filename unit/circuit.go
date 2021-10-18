package unit

import (
	"fmt"
	"sync"

	"github.com/xinxiao/logico/blueprint"
)

type CircuitSimulationTracker struct {
	Circuit          *Circuit
	InputValue       map[string]bool
	UnitValueMap     map[string]map[string]bool
	UnitValueMapLock sync.RWMutex
}

func (cst *CircuitSimulationTracker) ReadUnitValue(uid string) (map[string]bool, bool) {
	cst.UnitValueMapLock.RLock()
	defer cst.UnitValueMapLock.RUnlock()

	m, ok := cst.UnitValueMap[uid]
	return m, ok
}

func (cst *CircuitSimulationTracker) SaveUnitValue(uid string, vm map[string]bool) {
	cst.UnitValueMapLock.Lock()
	defer cst.UnitValueMapLock.Unlock()

	cst.UnitValueMap[uid] = vm
}

func (cst *CircuitSimulationTracker) GetInputValue(p blueprint.CircuitPin) (bool, error) {
	if v, ok := cst.Circuit.ConstantPins[p]; ok {
		return v, nil
	}

	if n, ok := cst.Circuit.InputPins[p]; ok {
		if v, ok := cst.InputValue[n]; ok {
			return v, nil
		}
	}

	if fp, ok := cst.Circuit.Connectors[p]; ok {
		return cst.GetOutputValue(fp)
	}

	return false, fmt.Errorf("cannot find input value for pin %s", p)
}

func (cst *CircuitSimulationTracker) GetOutputValue(p blueprint.CircuitPin) (bool, error) {
	if m, ok := cst.ReadUnitValue(p.UnitId); ok {
		if v, ok := m[p.PinId]; ok {
			return v, nil
		}
		return false, fmt.Errorf("cannot find output value for pin %s", p)
	}

	u, ok := cst.Circuit.UnitMap[p.UnitId]
	if !ok {
		return false, fmt.Errorf("cannot find unit %s", p.UnitId)
	}

	im := make(map[string]bool)
	for _, in := range u.Input() {
		iv, err := cst.GetInputValue(blueprint.CircuitPin{UnitId: p.UnitId, PinId: in})
		if err != nil {
			return false, err
		}
		im[in] = iv
	}

	out, err := u.Simulate(im)
	if err != nil {
		return false, err
	}
	cst.SaveUnitValue(p.UnitId, out)

	return cst.GetOutputValue(p)
}

type Circuit struct {
	CircuitName  string
	UnitMap      map[string]Unit
	InputPins    map[blueprint.CircuitPin]string
	ConstantPins map[blueprint.CircuitPin]bool
	Connectors   map[blueprint.CircuitPin]blueprint.CircuitPin
	OutputPins   map[string]blueprint.CircuitPin
}

func (c *Circuit) Name() string {
	return c.CircuitName
}

func (c *Circuit) Input() []string {
	im := make(map[string]bool)
	for _, in := range c.InputPins {
		im[in] = true
	}

	in := make([]string, 0)
	for n := range im {
		in = append(in, n)
	}
	return in
}

func (c *Circuit) Output() []string {
	out := make([]string, 0)
	for n := range c.OutputPins {
		out = append(out, n)
	}
	return out
}

func (c *Circuit) GetSimulationTracker(args map[string]bool) *CircuitSimulationTracker {
	return &CircuitSimulationTracker{
		Circuit:          c,
		InputValue:       args,
		UnitValueMap:     make(map[string]map[string]bool),
		UnitValueMapLock: sync.RWMutex{},
	}
}

type CircuitValueEntry struct {
	Key   string
	Value bool
}

func (c *Circuit) Simulate(args map[string]bool) (map[string]bool, error) {
	cst := c.GetSimulationTracker(args)
	out := make(map[string]bool)
	for on, op := range c.OutputPins {
		ov, err := cst.GetOutputValue(op)
		if err != nil {
			return nil, err
		}
		out[on] = ov
	}
	return out, nil
}
