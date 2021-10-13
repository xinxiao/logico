package simulation

import (
	"fmt"

	"github.com/xinxiao/logico/unit"
)

type BlockingSimulator struct {
	ur unit.UnitRegistry
}

func NewBlockingSimulator(ur unit.UnitRegistry) *BlockingSimulator {
	return &BlockingSimulator{ur}
}

func (bs *BlockingSimulator) Simulate(ut string, args map[string]bool) (map[string]bool, error) {
	u, err := bs.ur.GetUnit(ut)
	if err != nil {
		return nil, err
	}

	for _, i := range u.Input() {
		if _, ok := args[i]; !ok {
			return nil, fmt.Errorf("%s is not in the input", i)
		}
	}

	out, err := bs.simulateUnit(u, args)
	if err != nil {
		return nil, err
	}

	for _, o := range u.Output() {
		if _, ok := out[o]; !ok {
			return nil, fmt.Errorf("%s is not in the output", o)
		}
	}
	return out, nil
}

func (bs *BlockingSimulator) simulateUnit(u unit.Unit, args map[string]bool) (map[string]bool, error) {
	in := u.Input()

	switch u := u.(type) {
	case *unit.SingleOperandGate:
		v := args[in[0]]
		switch u.Name() {
		case "not":
			return map[string]bool{unit.GateOutput: !v}, nil
		default:
			return nil, fmt.Errorf("unknown single operand gate %s", u.Name())
		}
	case *unit.DoubleOperandGate:
		a, b := args[in[0]], args[in[1]]
		switch u.Name() {
		case "and":
			return map[string]bool{unit.GateOutput: a && b}, nil
		case "or":
			return map[string]bool{unit.GateOutput: a || b}, nil
		default:
			return nil, fmt.Errorf("unknown double operand gate %s", u.Name())
		}
	case *unit.Circuit:
		return bs.simulateCircuit(u, args)
	default:
		return nil, fmt.Errorf("unknown unit type %T", u)
	}
}

type BlockingSimulationTracker struct {
	nt     map[string]unit.CircuitUnitNode
	in     map[unit.CircuitPin]bool
	consts map[unit.CircuitPin]bool
	edge   map[unit.CircuitPin]unit.CircuitPin
	uvm    map[string]map[string]bool
}

func NewBlockingSimulationTracker(c *unit.Circuit, args map[string]bool) *BlockingSimulationTracker {
	bst := &BlockingSimulationTracker{
		nt:     c.CircuitUnitNodes,
		in:     make(map[unit.CircuitPin]bool),
		consts: make(map[unit.CircuitPin]bool),
		edge:   make(map[unit.CircuitPin]unit.CircuitPin),
		uvm:    make(map[string]map[string]bool),
	}

	for iv, pl := range c.InputPinMap {
		for _, p := range pl {
			bst.in[p] = args[iv]
		}
	}

	for cv, pl := range c.InputConstantMap {
		for _, p := range pl {
			bst.consts[p] = cv
		}
	}

	for _, e := range c.InteriorEdges {
		bst.edge[e.To] = e.From
	}

	return bst
}

func (bst *BlockingSimulationTracker) input(bs *BlockingSimulator, p unit.CircuitPin) (bool, error) {
	if v, ok := bst.in[p]; ok {
		return v, nil
	}

	if v, ok := bst.consts[p]; ok {
		return v, nil
	}

	if l, ok := bst.edge[p]; ok {
		return bst.output(bs, l)
	}

	return false, fmt.Errorf("input pin %s not found", p)
}

func (bst *BlockingSimulationTracker) output(bs *BlockingSimulator, p unit.CircuitPin) (bool, error) {
	if om, ok := bst.uvm[p.UnitId]; ok {
		if v, ok := om[p.PinId]; ok {
			return v, nil
		}
		return false, fmt.Errorf("output pin %s not found", p)
	}

	u, err := bs.ur.GetUnit(bst.nt[p.UnitId].UnitType)
	if err != nil {
		return false, err
	}

	args := make(map[string]bool)
	for _, in := range u.Input() {
		v, err := bst.input(bs, unit.CircuitPin{UnitId: p.UnitId, PinId: in})
		if err != nil {
			return false, err
		}
		args[in] = v
	}

	out, err := bs.Simulate(u.Name(), args)
	if err != nil {
		return false, err
	}
	bst.uvm[p.UnitId] = out

	return bst.output(bs, p)
}

func (bs *BlockingSimulator) simulateCircuit(c *unit.Circuit, args map[string]bool) (map[string]bool, error) {
	bst := NewBlockingSimulationTracker(c, args)

	out := make(map[string]bool)
	for on, op := range c.OutputPinMap {
		ov, err := bst.output(bs, op)
		if err != nil {
			return nil, err
		}

		out[on] = ov
	}
	return out, nil
}
