package unit

import (
	"fmt"

	"github.com/xinxiao/logico/blueprint"
)

type Circuit struct {
	n    string
	um   map[string]string
	in   map[blueprint.CircuitPin]string
	cst  map[blueprint.CircuitPin]bool
	edge map[blueprint.CircuitPin]blueprint.CircuitPin
	out  map[string]blueprint.CircuitPin
}

func BuildCircuitFromBlueprint(cbp *blueprint.CircuitBlueprint) *Circuit {
	c := &Circuit{
		n:    cbp.CircuitName,
		um:   cbp.Nodes,
		in:   make(map[blueprint.CircuitPin]string),
		cst:  make(map[blueprint.CircuitPin]bool),
		edge: make(map[blueprint.CircuitPin]blueprint.CircuitPin),
		out:  make(map[string]blueprint.CircuitPin),
	}

	for n, ipl := range cbp.Inputs {
		for _, ip := range ipl {
			c.in[ip] = n
		}
	}

	for _, cp := range cbp.AlwaysOn {
		c.cst[cp] = true
	}

	for _, cp := range cbp.AlwaysOff {
		c.cst[cp] = false
	}

	for _, e := range cbp.Connectors {
		c.edge[e.To] = e.From
	}

	for n, op := range cbp.Outputs {
		c.out[n] = op
	}

	return c
}

func (c *Circuit) Name() string {
	return c.n
}

func (c *Circuit) Input() []string {
	im := make(map[string]bool)
	for _, in := range c.in {
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
	for n := range c.out {
		out = append(out, n)
	}
	return out
}

func (c *Circuit) GetUnitType(uid string) (string, error) {
	if u, ok := c.um[uid]; ok {
		return u, nil
	}
	return "", fmt.Errorf("no unit named %s", uid)
}

func (c *Circuit) AssignInputValue(args map[string]bool) map[blueprint.CircuitPin]bool {
	m := make(map[blueprint.CircuitPin]bool)

	for p, n := range c.in {
		m[p] = args[n]
	}

	for p, v := range c.cst {
		m[p] = v
	}

	return m
}

func (c *Circuit) GetFeedingInput(p blueprint.CircuitPin) (blueprint.CircuitPin, error) {
	if lp, ok := c.edge[p]; ok {
		return lp, nil
	}
	return blueprint.CircuitPin{}, fmt.Errorf("%s is not linked to any other pin", p)
}

func (c *Circuit) GetOutputPins(n string) (blueprint.CircuitPin, error) {
	if p, ok := c.out[n]; ok {
		return p, nil
	}
	return blueprint.CircuitPin{}, fmt.Errorf("no output pin named %s", n)
}
