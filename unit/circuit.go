package unit

import (
	"encoding/json"
	"fmt"

	"github.com/xinxiao/logico/blueprint"
)

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

func (c *Circuit) GetUnit(uid string) (Unit, error) {
	if u, ok := c.UnitMap[uid]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("cannot find unit %s", uid)
}

func (c *Circuit) AsBlueprint() *blueprint.CircuitBlueprint {
	cbp := &blueprint.CircuitBlueprint{
		CircuitName: c.Name(),
		Nodes:       make(map[string]string),
		AlwaysOn:    make([]blueprint.CircuitPin, 0),
		AlwaysOff:   make([]blueprint.CircuitPin, 0),
		Inputs:      make(map[string][]blueprint.CircuitPin),
		Connectors:  make([]blueprint.CircuitEdge, 0),
		Outputs:     make(map[string]blueprint.CircuitPin),
	}

	for id, u := range c.UnitMap {
		cbp.Nodes[id] = u.Name()
	}

	for p, v := range c.ConstantPins {
		if v {
			cbp.AlwaysOn = append(cbp.AlwaysOn, p)
		} else {
			cbp.AlwaysOff = append(cbp.AlwaysOff, p)
		}
	}

	for p, n := range c.InputPins {
		cbp.Inputs[n] = append(cbp.Inputs[n], p)
	}

	for tp, fp := range c.Connectors {
		cbp.Connectors = append(cbp.Connectors, blueprint.CircuitEdge{From: fp, To: tp})
	}

	cbp.Outputs = c.OutputPins

	return cbp
}

func (c *Circuit) String() string {
	b, _ := json.MarshalIndent(c.AsBlueprint(), "", "  ")
	return string(b)
}

type CircuitSimulationTracker struct {
	Circuit      *Circuit
	InputValue   map[string]bool
	UnitValueMap map[string]map[string]bool
}

func (cst *CircuitSimulationTracker) ReadUnitValue(uid string) (map[string]bool, bool) {
	m, ok := cst.UnitValueMap[uid]
	return m, ok
}

func (cst *CircuitSimulationTracker) SaveUnitValue(uid string, vm map[string]bool) {
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

func (c *Circuit) GetSimulationTracker(args map[string]bool) *CircuitSimulationTracker {
	return &CircuitSimulationTracker{
		Circuit:      c,
		InputValue:   args,
		UnitValueMap: make(map[string]map[string]bool),
	}
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

func ConcatUnitName(pre, suf string) string {
	if pre == "" {
		return suf
	}

	if suf == "" {
		return pre
	}

	return fmt.Sprintf("%s.%s", pre, suf)
}

func (c *Circuit) TraceInputPins(un, pn string) ([]blueprint.CircuitPin, error) {
	im := make(map[string][]blueprint.CircuitPin)
	for p, n := range c.InputPins {
		if l, ok := im[n]; !ok {
			im[n] = []blueprint.CircuitPin{p}
		} else {
			im[n] = append(l, p)
		}
	}

	l, ok := im[pn]
	if !ok {
		return nil, fmt.Errorf("cannot find input pin %s", pn)
	}

	r := make([]blueprint.CircuitPin, 0)
	for _, p := range l {
		if u, ok := c.UnitMap[p.UnitId]; !ok {
			return nil, fmt.Errorf("cannot find unit %s", p.UnitId)
		} else if sc, ok := u.(*Circuit); !ok {
			r = append(r, blueprint.CircuitPin{UnitId: ConcatUnitName(un, p.UnitId), PinId: p.PinId})
		} else if sl, err := sc.TraceInputPins(ConcatUnitName(un, p.UnitId), p.PinId); err != nil {
			return nil, err
		} else {
			r = append(r, sl...)
		}
	}
	return r, nil
}

func (c *Circuit) TraceOutputPin(un, pn string) (blueprint.CircuitPin, error) {
	for n, p := range c.OutputPins {
		if n != pn {
			continue
		}

		su, ok := c.UnitMap[p.UnitId]
		if !ok {
			return blueprint.CircuitPin{}, fmt.Errorf("cannot find unit %s", p.UnitId)
		}
		cn := ConcatUnitName(un, p.UnitId)

		sc, ok := su.(*Circuit)
		if !ok {
			return blueprint.CircuitPin{UnitId: cn, PinId: p.PinId}, nil
		} else if op, err := sc.TraceOutputPin(cn, p.PinId); err != nil {
			return blueprint.CircuitPin{}, err
		} else {
			return op, nil
		}
	}
	return blueprint.CircuitPin{}, fmt.Errorf("cannot find output pin %s", un)
}

func (c *Circuit) Expand() (*Circuit, error) {
	ec := &Circuit{
		CircuitName:  c.CircuitName,
		UnitMap:      make(map[string]Unit),
		InputPins:    make(map[blueprint.CircuitPin]string),
		ConstantPins: make(map[blueprint.CircuitPin]bool),
		Connectors:   make(map[blueprint.CircuitPin]blueprint.CircuitPin),
		OutputPins:   make(map[string]blueprint.CircuitPin),
	}

	for sn, su := range c.UnitMap {
		sc, ok := su.(*Circuit)
		if !ok {
			ec.UnitMap[sn] = su
			continue
		}

		esc, err := sc.Expand()
		if err != nil {
			return nil, err
		}

		for n, u := range esc.UnitMap {
			ec.UnitMap[ConcatUnitName(sn, n)] = u
		}

		for p, v := range esc.ConstantPins {
			cp := blueprint.CircuitPin{UnitId: ConcatUnitName(sn, p.UnitId), PinId: p.PinId}
			ec.ConstantPins[cp] = v
		}

		for ot, of := range esc.Connectors {
			ct := blueprint.CircuitPin{UnitId: ConcatUnitName(sn, ot.UnitId), PinId: ot.PinId}
			cf := blueprint.CircuitPin{UnitId: ConcatUnitName(sn, of.UnitId), PinId: of.PinId}
			ec.Connectors[ct] = cf
		}
	}

	for _, n := range c.Input() {
		il, err := c.TraceInputPins("", n)
		if err != nil {
			return nil, err
		}

		for _, p := range il {
			ec.InputPins[p] = n
		}
	}

	for _, n := range c.Output() {
		op, err := c.TraceOutputPin("", n)
		if err != nil {
			return nil, err
		}
		ec.OutputPins[n] = op
	}

	for cp, v := range c.ConstantPins {
		if cu, err := c.GetUnit(cp.UnitId); err != nil {
			return nil, err
		} else if cc, ok := cu.(*Circuit); !ok {
			ec.ConstantPins[cp] = v
		} else if cpl, err := cc.TraceInputPins(cp.UnitId, cp.PinId); err != nil {
			return nil, err
		} else {
			for _, p := range cpl {
				ec.ConstantPins[p] = v
			}
		}
	}

	for tp, fp := range c.Connectors {
		tpl := []blueprint.CircuitPin{tp}
		if tu, err := c.GetUnit(tp.UnitId); err != nil {
			return nil, err
		} else if tc, ok := tu.(*Circuit); ok {
			tpl, err = tc.TraceInputPins(tp.UnitId, tp.PinId)
			if err != nil {
				return nil, err
			}
		}

		if fu, err := c.GetUnit(fp.UnitId); err != nil {
			return nil, err
		} else if oc, ok := fu.(*Circuit); ok {
			fp, err = oc.TraceOutputPin(fp.UnitId, fp.PinId)
			if err != nil {
				return nil, err
			}
		}

		for _, tp := range tpl {
			ec.Connectors[tp] = fp
		}
	}

	return ec, nil
}
