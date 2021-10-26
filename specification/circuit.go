package specification

import (
	"encoding/json"
)

const (
	CircuitBlueprintFileExtension = ".circuit"
)

type Pin struct {
	UnitId string `json:"uid"`
	PinId  string `json:"pid"`
}

type Edge struct {
	From Pin `json:"from"`
	To   Pin `json:"to"`
}

type Blueprint struct {
	CircuitName string            `json:"name"`
	Nodes       map[string]string `json:"nodes"`
	AlwaysOn    []Pin             `json:"always_on,omitempty"`
	AlwaysOff   []Pin             `json:"always_off,omitempty"`
	Inputs      map[string][]Pin  `json:"inputs"`
	Connectors  []Edge            `json:"connectors,omitempty"`
	Outputs     map[string]Pin    `json:"outputs"`
}

func (bp *Blueprint) String() string {
	b, _ := json.MarshalIndent(bp, "", "  ")
	return string(b)
}
