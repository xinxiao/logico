package blueprint

import (
	"encoding/json"
	"io/ioutil"
	"path"
)

const (
	CircuitBlueprintFileExtension = ".circuit"
)

type CircuitPin struct {
	UnitId string `json:"uid"`
	PinId  string `json:"pid"`
}

type CircuitEdge struct {
	From CircuitPin `json:"from"`
	To   CircuitPin `json:"to"`
}

type CircuitBlueprint struct {
	CircuitName string                  `json:"name"`
	Nodes       map[string]string       `json:"nodes"`
	AlwaysOn    []CircuitPin            `json:"always_on,omitempty"`
	AlwaysOff   []CircuitPin            `json:"always_off,omitempty"`
	Inputs      map[string][]CircuitPin `json:"inputs"`
	Connectors  []CircuitEdge           `json:"connectors,omitempty"`
	Outputs     map[string]CircuitPin   `json:"outputs"`
}

func (cbp *CircuitBlueprint) String() string {
	b, _ := json.MarshalIndent(cbp, "", "  ")
	return string(b)
}

func (cbp *CircuitBlueprint) SaveAsFile(p string) error {
	f := path.Join(p, cbp.CircuitName+CircuitBlueprintFileExtension)
	return ioutil.WriteFile(f, []byte(cbp.String()), 0644)
}
