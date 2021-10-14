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
	Name      string                  `json:"name"`
	Nodes     map[string]string       `json:"nodes"`
	AlwaysOn  []CircuitPin            `json:"always_on,omitempty"`
	AlwaysOff []CircuitPin            `json:"always_off,omitempty"`
	Inputs    map[string][]CircuitPin `json:"inputs"`
	Edges     []CircuitEdge           `json:"edges,omitempty"`
	Outputs   map[string]CircuitPin   `json:"outputs"`
}
