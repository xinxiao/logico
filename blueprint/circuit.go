package blueprint

type CircuitPin struct {
	UnitId string `json:"uid"`
	PinId  string `json:"pid"`
}

type CircuitEdge struct {
	From CircuitPin `json:"from"`
	To   CircuitPin `json:"to"`
}

type CircuitBlueprint struct {
	Name     string                  `json:"name"`
	Nodes    map[string]string       `json:"nodes"`
	Input    map[string][]CircuitPin `json:"input"`
	Constant map[bool][]CircuitPin   `json:"constant_input,omitempty"`
	Edges    []CircuitEdge           `json:"edges,omitempty"`
	Output   map[string]CircuitPin   `json:"output"`
}
