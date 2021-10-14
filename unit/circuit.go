package unit

import (
	"github.com/xinxiao/logico/blueprint"
)

type Circuit struct {
	CircuitName  string
	Units        map[string]string
	InputPins    map[blueprint.CircuitPin]string
	ConstantPins map[blueprint.CircuitPin]bool
	Edges        map[blueprint.CircuitPin]blueprint.CircuitPin
	OutputPins   map[string]blueprint.CircuitPin
}

func BuildCircuitFromBlueprint(cbp blueprint.CircuitBlueprint) *Circuit {
	c := &Circuit{
		CircuitName:  cbp.Name,
		Units:        cbp.Nodes,
		InputPins:    make(map[blueprint.CircuitPin]string),
		ConstantPins: make(map[blueprint.CircuitPin]bool),
		Edges:        make(map[blueprint.CircuitPin]blueprint.CircuitPin),
		OutputPins:   make(map[string]blueprint.CircuitPin),
	}

	for n, ipl := range cbp.Input {
		for _, ip := range ipl {
			c.InputPins[ip] = n
		}
	}

	for n, cpl := range cbp.Constant {
		for _, cp := range cpl {
			c.ConstantPins[cp] = n
		}
	}

	for _, e := range cbp.Edges {
		c.Edges[e.To] = e.From
	}

	for n, op := range cbp.Output {
		c.OutputPins[n] = op
	}

	return c
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

	return out
}
