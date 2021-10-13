package unit

import (
	"embed"
	"encoding/json"
	"fmt"
)

const (
	builtinPath = "builtin"
)

var (
	baseGates = map[string]Unit{
		"not": &SingleOperandGate{Gate{n: "not"}},
		"and": &DoubleOperandGate{Gate{n: "and"}},
		"or":  &DoubleOperandGate{Gate{n: "or"}},
	}

	//go:embed builtin/*.json
	builtInCircuits embed.FS
)

type UnitRegistry map[string]Unit

func NewUnitRegistry() (UnitRegistry, error) {
	ur := UnitRegistry(make(map[string]Unit))
	for n, u := range baseGates {
		ur[n] = u
	}

	dir, err := builtInCircuits.ReadDir(builtinPath)
	if err != nil {
		return nil, err
	}

	for _, f := range dir {
		p := fmt.Sprintf("%s/%s", builtinPath, f.Name())

		b, err := builtInCircuits.ReadFile(p)
		if err != nil {
			return nil, err
		}

		c := &Circuit{}
		if err := json.Unmarshal(b, &c); err != nil {
			return nil, err
		}

		if err := ur.RegisterUnit(c); err != nil {
			return nil, err
		}
	}

	return ur, nil
}

func (ur UnitRegistry) RegisterUnit(p Unit) error {
	n := p.Name()

	if _, ok := ur[n]; ok {
		return fmt.Errorf("unit %s already registered", n)
	}

	ur[n] = p
	return nil
}

func (ur UnitRegistry) GetUnit(n string) (Unit, error) {
	p, ok := ur[n]
	if !ok {
		return nil, fmt.Errorf("unit %s not registered", n)
	}
	return p, nil
}

func (ur UnitRegistry) RemoveUnit(n string) error {
	if _, ok := ur[n]; !ok {
		return fmt.Errorf("unit %s does not exist", n)
	}

	delete(ur, n)
	return nil
}
