package unit

type CircuitPin struct {
	UnitId string `json:"uid"`
	PinId  string `json:"pid"`
}

func NewCircuitPin(uid, pid string) CircuitPin {
	return CircuitPin{uid, pid}
}

type CircuitUnitNode struct {
	UnitType string `json:"unit_type"`
}

type CircuitEdge struct {
	From CircuitPin `json:"from"`
	To   CircuitPin `json:"to"`
}

type Circuit struct {
	CircuitName      string                     `json:"name"`
	CircuitUnitNodes map[string]CircuitUnitNode `json:"nodes"`
	InputPinMap      map[string][]CircuitPin    `json:"input"`
	InputConstantMap map[bool][]CircuitPin      `json:"constant_input,omitempty"`
	InteriorEdges    []CircuitEdge              `json:"edges"`
	OutputPinMap     map[string]CircuitPin      `json:"output"`
}

func (c *Circuit) Name() string {
	return c.CircuitName
}

func (c *Circuit) Input() []string {
	in := make([]string, 0)
	for n := range c.InputPinMap {
		in = append(in, n)
	}
	return in
}

func (c *Circuit) Output() []string {
	out := make([]string, 0)
	for n := range c.OutputPinMap {
		out = append(out, n)
	}
	return out
}
