package repository

import (
	"fmt"

	"github.com/xinxiao/logico/blueprint"
	"github.com/xinxiao/logico/unit"
)

var (
	baseGates = []unit.Unit{
		&unit.SingleOperandGate{Gate: unit.Gate{GateName: "not"}},
		&unit.DoubleOperandGate{Gate: unit.Gate{GateName: "and"}},
		&unit.DoubleOperandGate{Gate: unit.Gate{GateName: "or"}},
	}
)

type UnitRepository map[string]unit.Unit

func NewUnitRepository() (UnitRepository, error) {
	ur := UnitRepository(make(map[string]unit.Unit))

	for _, u := range baseGates {
		ur.RegisterUnit(u)
	}

	cbpl, err := blueprint.LoadBuiltinCircuitBlueprint()
	if err != nil {
		return nil, err
	}

	for _, cbp := range cbpl {
		c := unit.BuildCircuitFromBlueprint(cbp)
		ur.RegisterUnit(c)
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
