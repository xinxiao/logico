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
	gateProviders = map[string]func() unit.Unit{
		"nand": func() unit.Unit { return &unit.Nand{} },
	}

	//go:embed builtin/*.circuit
	//go:embed builtin/**/*.lib
	//go:embed builtin/**/*.circuit
	prebuiltCircuitSource embed.FS

	prebuiltCircuits, _ = LoadPrebuiltCircuit()
)

func LoadPrebuiltCircuit() (map[string]*blueprint.CircuitBlueprint, error) {
	cbpp := blueprint.NewCircuitBlueprintParser(prebuiltCircuitSource)

	m := make(map[string]*blueprint.CircuitBlueprint)
	err := fs.WalkDir(prebuiltCircuitSource, builtinPath, func(p string, f fs.DirEntry, err error) error {
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
	bpr BlueprintRepository
}

func NewUnitRepository() *UnitRepository {
	return &UnitRepository{
		bpr: &emptyBlueprintRepository{},
	}
}

func (ur *UnitRepository) LinkBlueprintRepository(bpr BlueprintRepository) {
	ur.bpr = bpr
}

func (ur *UnitRepository) GetUnit(n string) (unit.Unit, error) {
	if u, ok := gateProviders[n]; ok {
		return u(), nil
	}

	if cbp, ok := prebuiltCircuits[n]; ok {
		return ur.BuildCircuitFromBlueprint(cbp)
	}

	if cbp, err := ur.bpr.CheckOutBlueprint(n); err == nil {
		return ur.BuildCircuitFromBlueprint(cbp)
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
		u, err := ur.GetUnit(ut)
		if err != nil {
			return nil, err
		}
		c.UnitMap[uid] = u
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
