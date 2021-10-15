package simulation

import (
	"fmt"

	"github.com/xinxiao/logico/blueprint"
	"github.com/xinxiao/logico/repository"
	"github.com/xinxiao/logico/unit"
)

type Simulator struct {
	ur repository.UnitRepository
}

func NewSimulator(ur repository.UnitRepository) *Simulator {
	return &Simulator{ur}
}

func (bs *Simulator) Simulate(ut string, args map[string]bool) (map[string]bool, error) {
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

func (bs *Simulator) simulateUnit(u unit.Unit, args map[string]bool) (map[string]bool, error) {
	in := u.Input()

	switch u := u.(type) {
	case *unit.SingleOperandGate:
		v := args[in[0]]
		switch u.Name() {
		case "not":
			return unit.GateOutput(!v), nil
		default:
			return nil, fmt.Errorf("unknown single operand gate %s", u.Name())
		}
	case *unit.DoubleOperandGate:
		a, b := args[in[0]], args[in[1]]
		switch u.Name() {
		case "and":
			return unit.GateOutput(a && b), nil
		case "or":
			return unit.GateOutput(a || b), nil
		default:
			return nil, fmt.Errorf("unknown double operand gate %s", u.Name())
		}
	case *unit.Circuit:
		return bs.simulateCircuit(u, args)
	default:
		return nil, fmt.Errorf("unknown unit type %T", u)
	}
}

type simulationTracker struct {
	utf func(string) (string, error)
	fpm func(blueprint.CircuitPin) (blueprint.CircuitPin, error)
	in  map[blueprint.CircuitPin]bool
	uvm map[string]map[string]bool
}

func newSimulationTracker(c *unit.Circuit, args map[string]bool) *simulationTracker {
	return &simulationTracker{
		utf: c.GetUnitType,
		fpm: c.GetFeedingInput,
		in:  c.AssignInputPinWithValue(args),
		uvm: make(map[string]map[string]bool),
	}
}

func (st *simulationTracker) findPinInput(bs *Simulator, p blueprint.CircuitPin) (bool, error) {
	if v, ok := st.in[p]; ok {
		return v, nil
	}

	if l, err := st.fpm(p); err == nil {
		return st.finPinOutput(bs, l)
	}

	return false, fmt.Errorf("input pin %s not found", p)
}

func (st *simulationTracker) finPinOutput(bs *Simulator, p blueprint.CircuitPin) (bool, error) {
	if om, ok := st.uvm[p.UnitId]; ok {
		if v, ok := om[p.PinId]; ok {
			return v, nil
		}
		return false, fmt.Errorf("output pin %s not found", p)
	}

	ut, err := st.utf(p.UnitId)
	if err != nil {
		return false, err
	}

	u, err := bs.ur.GetUnit(ut)
	if err != nil {
		return false, err
	}

	args := make(map[string]bool)
	for _, in := range u.Input() {
		v, err := st.findPinInput(bs, blueprint.CircuitPin{UnitId: p.UnitId, PinId: in})
		if err != nil {
			return false, err
		}
		args[in] = v
	}

	out, err := bs.Simulate(u.Name(), args)
	if err != nil {
		return false, err
	}
	st.uvm[p.UnitId] = out

	return st.finPinOutput(bs, p)
}

func (bs *Simulator) simulateCircuit(c *unit.Circuit, args map[string]bool) (map[string]bool, error) {
	st := newSimulationTracker(c, args)

	out := make(map[string]bool)
	for _, on := range c.Output() {
		op, err := c.GetOutputPins(on)
		if err != nil {
			return nil, err
		}

		ov, err := st.finPinOutput(bs, op)
		if err != nil {
			return nil, err
		}

		out[on] = ov
	}
	return out, nil
}
