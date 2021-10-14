package repository

import (
	"embed"
	"encoding/json"
	"fmt"

	"github.com/xinxiao/logico/blueprint"
	"github.com/xinxiao/logico/unit"
)

const (
	builtinPath = "builtin"
)

var (
	baseGates = map[string]unit.Unit{
		"not": &unit.SingleOperandGate{Gate: unit.Gate{GateName: "not"}},
		"and": &unit.DoubleOperandGate{Gate: unit.Gate{GateName: "and"}},
		"or":  &unit.DoubleOperandGate{Gate: unit.Gate{GateName: "or"}},
	}

	//go:embed builtin/*.json
	builtInCircuits embed.FS
)

type UnitRepository map[string]unit.Unit

func NewUnitRepository() (UnitRepository, error) {
	ur := UnitRepository(make(map[string]unit.Unit))
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

		cm := blueprint.CircuitBlueprint{}
		if err := json.Unmarshal(b, &cm); err != nil {
			return nil, fmt.Errorf("error parsing %s: %s", p, err)
		}

		c := unit.BuildCircuitFromBlueprint(cm)
		if err := ur.RegisterUnit(c); err != nil {
			return nil, err
		}
	}

	return ur, nil
}

func (ur UnitRepository) RegisterUnit(p unit.Unit) error {
	n := p.Name()

	if _, ok := ur[n]; ok {
		return fmt.Errorf("unit %s already registered", n)
	}

	ur[n] = p
	return nil
}

func (ur UnitRepository) GetUnit(n string) (unit.Unit, error) {
	p, ok := ur[n]
	if !ok {
		return nil, fmt.Errorf("unit %s not registered", n)
	}
	return p, nil
}

func (ur UnitRepository) RemoveUnit(n string) error {
	if _, ok := ur[n]; !ok {
		return fmt.Errorf("unit %s does not exist", n)
	}

	delete(ur, n)
	return nil
}
