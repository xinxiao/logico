package repository

import (
	"embed"
	"fmt"
	"io/fs"
	"path"

	"github.com/xinxiao/logico/blueprint"
	"github.com/xinxiao/logico/unit"
)

const (
	builtinPath = "builtin"
)

var (
	//go:embed builtin/*.circuit
	//go:embed builtin/**/*.lib
	//go:embed builtin/**/*.circuit
	builtInCircuitSource embed.FS

	baseGates = []unit.Unit{
		&unit.Not{},
		&unit.And{},
		&unit.Or{},
	}

	prebuiltCircuits, _ = LoadPrebuiltCircuit()
)

func LoadPrebuiltCircuit() (map[string]*blueprint.CircuitBlueprint, error) {
	cbpp := blueprint.NewCircuitBlueprintParser(builtInCircuitSource)

	m := make(map[string]*blueprint.CircuitBlueprint)
	err := fs.WalkDir(builtInCircuitSource, builtinPath, func(p string, f fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() || path.Ext(f.Name()) != blueprint.CircuitBlueprintFileExtension {
			return nil
		}

		cbp, err := cbpp.ParseFile(p)
		if err != nil {
			return err
		}

		m[cbp.CircuitName] = cbp
		return nil
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}

type UnitRepository struct {
	cm  map[string]unit.Unit
	bpr BlueprintRepository
}

func NewUnitRepository() *UnitRepository {
	ur := &UnitRepository{
		cm:  make(map[string]unit.Unit),
		bpr: &emptyBlueprintRepository{},
	}

	for _, u := range baseGates {
		ur.cm[u.Name()] = u
	}

	return ur
}

func (ur *UnitRepository) LinkBlueprintRepository(bpr BlueprintRepository) {
	ur.bpr = bpr
}

func (ur *UnitRepository) CheckOutBlueprint(n string) (*blueprint.CircuitBlueprint, error) {
	if cbp, ok := prebuiltCircuits[n]; ok {
		return cbp, nil
	}
	return ur.bpr.CheckOutBlueprint(n)
}

func (ur *UnitRepository) GetUnit(n string) (unit.Unit, error) {
	if p, ok := ur.cm[n]; ok {
		return p, nil
	}

	if bp, err := ur.CheckOutBlueprint(n); err == nil {
		return ur.BuildCircuitFromBlueprint(bp)
	}

	return nil, fmt.Errorf("unit %s not found", n)
}

func (ur *UnitRepository) BuildCircuitFromBlueprint(cbp *blueprint.CircuitBlueprint) (unit.Unit, error) {
	c := &unit.Circuit{
		CircuitName:  cbp.CircuitName,
		UnitMap:      make(map[string]unit.Unit),
		InputPins:    make(map[blueprint.CircuitPin]string),
		ConstantPins: make(map[blueprint.CircuitPin]bool),
		Connectors:   make(map[blueprint.CircuitPin]blueprint.CircuitPin),
		OutputPins:   make(map[string]blueprint.CircuitPin),
	}

	for uid, ut := range cbp.Nodes {
		if u, err := ur.GetUnit(ut); err != nil {
			return nil, err
		} else {
			c.UnitMap[uid] = u
		}
	}

	for n, ipl := range cbp.Inputs {
		for _, ip := range ipl {
			c.InputPins[ip] = n
		}
	}

	for _, cp := range cbp.AlwaysOn {
		c.ConstantPins[cp] = true
	}

	for _, cp := range cbp.AlwaysOff {
		c.ConstantPins[cp] = false
	}

	for _, e := range cbp.Connectors {
		c.Connectors[e.To] = e.From
	}

	for n, op := range cbp.Outputs {
		c.OutputPins[n] = op
	}

	return c, nil
}
