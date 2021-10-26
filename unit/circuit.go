package unit

import (
	"encoding/json"
	"fmt"

	"github.com/xinxiao/logico/specification"
)

type Circuit struct {
	CircuitName  string
	UnitMap      map[string]Unit
	InputPins    map[specification.Pin]string
	ConstantPins map[specification.Pin]bool
	Connections  map[specification.Pin]specification.Pin
	OutputPins   map[string]specification.Pin
}

func ExpandCircuit(u Unit) (Unit, error) {
	if c, ok := u.(*Circuit); ok {
		return c.expand()
	}
	return u, nil
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

func (c *Circuit) getUnit(uid string) (Unit, error) {
	if u, ok := c.UnitMap[uid]; ok {
		return u, nil
	}
	return nil, fmt.Errorf("cannot find unit %s", uid)
}

func (c *Circuit) AsBlueprint() *specification.Blueprint {
	bp := &specification.Blueprint{
		CircuitName: c.Name(),
		Nodes:       make(map[string]string),
		AlwaysOn:    make([]specification.Pin, 0),
		AlwaysOff:   make([]specification.Pin, 0),
		Inputs:      make(map[string][]specification.Pin),
		Connectors:  make([]specification.Edge, 0),
		Outputs:     make(map[string]specification.Pin),
	}

	for id, u := range c.UnitMap {
		bp.Nodes[id] = u.Name()
	}

	for p, v := range c.ConstantPins {
		if v {
			bp.AlwaysOn = append(bp.AlwaysOn, p)
		} else {
			bp.AlwaysOff = append(bp.AlwaysOff, p)
		}
	}

	for p, n := range c.InputPins {
		bp.Inputs[n] = append(bp.Inputs[n], p)
	}

	for tp, fp := range c.Connections {
		bp.Connectors = append(bp.Connectors, specification.Edge{From: fp, To: tp})
	}

	bp.Outputs = c.OutputPins

	return bp
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

func (cst *CircuitSimulationTracker) readUnitValue(uid string) (map[string]bool, bool) {
	m, ok := cst.UnitValueMap[uid]
	return m, ok
}

func (cst *CircuitSimulationTracker) saveUnitValue(uid string, vm map[string]bool) {
	cst.UnitValueMap[uid] = vm
}

func (cst *CircuitSimulationTracker) getInputValue(p specification.Pin) (bool, error) {
	if v, ok := cst.Circuit.ConstantPins[p]; ok {
		return v, nil
	}

	if n, ok := cst.Circuit.InputPins[p]; ok {
		if v, ok := cst.InputValue[n]; ok {
			return v, nil
		}
	}

	if fp, ok := cst.Circuit.Connections[p]; ok {
		return cst.getOutputValue(fp)
	}

	return false, fmt.Errorf("cannot find input value for pin %s", p)
}

func (cst *CircuitSimulationTracker) getOutputValue(p specification.Pin) (bool, error) {
	if m, ok := cst.readUnitValue(p.UnitId); ok {
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
		iv, err := cst.getInputValue(specification.Pin{UnitId: p.UnitId, PinId: in})
		if err != nil {
			return false, err
		}
		im[in] = iv
	}

	out, err := u.Simulate(im)
	if err != nil {
		return false, err
	}
	cst.saveUnitValue(p.UnitId, out)

	return cst.getOutputValue(p)
}

func (c *Circuit) getSimulationTracker(args map[string]bool) *CircuitSimulationTracker {
	return &CircuitSimulationTracker{
		Circuit:      c,
		InputValue:   args,
		UnitValueMap: make(map[string]map[string]bool),
	}
}

func (c *Circuit) Simulate(args map[string]bool) (map[string]bool, error) {
	cst := c.getSimulationTracker(args)
	out := make(map[string]bool)
	for on, op := range c.OutputPins {
		ov, err := cst.getOutputValue(op)
		if err != nil {
			return nil, err
		}
		out[on] = ov
	}
	return out, nil
}

func concatUnitName(pre, suf string) string {
	if pre == "" {
		return suf
	}

	if suf == "" {
		return pre
	}

	return fmt.Sprintf("%s.%s", pre, suf)
}

func (c *Circuit) traceInputPins(uid, pid string) ([]specification.Pin, error) {
	r := make([]specification.Pin, 0)
	for p, n := range c.InputPins {
		if n != pid {
			continue
		}

		if u, ok := c.UnitMap[p.UnitId]; !ok {
			return nil, fmt.Errorf("cannot find unit %s", p.UnitId)
		} else if sc, ok := u.(*Circuit); !ok {
			r = append(r, specification.Pin{UnitId: concatUnitName(uid, p.UnitId), PinId: p.PinId})
		} else if sl, err := sc.traceInputPins(concatUnitName(uid, p.UnitId), p.PinId); err != nil {
			return nil, err
		} else {
			r = append(r, sl...)
		}
	}
	return r, nil
}

func (c *Circuit) traceOutputPin(uid, pid string) (specification.Pin, error) {
	for n, p := range c.OutputPins {
		if n != pid {
			continue
		}

		su, ok := c.UnitMap[p.UnitId]
		if !ok {
			return specification.Pin{}, fmt.Errorf("cannot find unit %s", p.UnitId)
		}
		cn := concatUnitName(uid, p.UnitId)

		if sc, ok := su.(*Circuit); !ok {
			return specification.Pin{UnitId: cn, PinId: p.PinId}, nil
		} else if op, err := sc.traceOutputPin(cn, p.PinId); err != nil {
			return specification.Pin{}, err
		} else {
			return op, nil
		}
	}
	return specification.Pin{}, fmt.Errorf("cannot find output pin %s", uid)
}

func (c *Circuit) expand() (*Circuit, error) {
	ec := &Circuit{
		CircuitName:  c.CircuitName,
		UnitMap:      make(map[string]Unit),
		InputPins:    make(map[specification.Pin]string),
		ConstantPins: make(map[specification.Pin]bool),
		Connections:  make(map[specification.Pin]specification.Pin),
		OutputPins:   make(map[string]specification.Pin),
	}

	for sn, su := range c.UnitMap {
		sc, ok := su.(*Circuit)
		if !ok {
			ec.UnitMap[sn] = su
			continue
		}

		esc, err := sc.expand()
		if err != nil {
			return nil, err
		}

		for n, u := range esc.UnitMap {
			ec.UnitMap[concatUnitName(sn, n)] = u
		}

		for p, v := range esc.ConstantPins {
			cp := specification.Pin{UnitId: concatUnitName(sn, p.UnitId), PinId: p.PinId}
			ec.ConstantPins[cp] = v
		}

		for ot, of := range esc.Connections {
			ct := specification.Pin{UnitId: concatUnitName(sn, ot.UnitId), PinId: ot.PinId}
			cf := specification.Pin{UnitId: concatUnitName(sn, of.UnitId), PinId: of.PinId}
			ec.Connections[ct] = cf
		}
	}

	for _, n := range c.Input() {
		il, err := c.traceInputPins("", n)
		if err != nil {
			return nil, err
		}

		for _, p := range il {
			ec.InputPins[p] = n
		}
	}

	for _, n := range c.Output() {
		op, err := c.traceOutputPin("", n)
		if err != nil {
			return nil, err
		}
		ec.OutputPins[n] = op
	}

	for cp, v := range c.ConstantPins {
		if cu, err := c.getUnit(cp.UnitId); err != nil {
			return nil, err
		} else if cc, ok := cu.(*Circuit); !ok {
			ec.ConstantPins[cp] = v
		} else if cpl, err := cc.traceInputPins(cp.UnitId, cp.PinId); err != nil {
			return nil, err
		} else {
			for _, p := range cpl {
				ec.ConstantPins[p] = v
			}
		}
	}

	for tp, fp := range c.Connections {
		tpl := []specification.Pin{tp}
		if tu, err := c.getUnit(tp.UnitId); err != nil {
			return nil, err
		} else if tc, ok := tu.(*Circuit); ok {
			tpl, err = tc.traceInputPins(tp.UnitId, tp.PinId)
			if err != nil {
				return nil, err
			}
		}

		if fu, err := c.getUnit(fp.UnitId); err != nil {
			return nil, err
		} else if oc, ok := fu.(*Circuit); ok {
			fp, err = oc.traceOutputPin(fp.UnitId, fp.PinId)
			if err != nil {
				return nil, err
			}
		}

		for _, tp := range tpl {
			ec.Connections[tp] = fp
		}
	}

	return ec, nil
}
