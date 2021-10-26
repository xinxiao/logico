package repository

import (
	"embed"
	"io/fs"
	"path"

	"github.com/xinxiao/logico/specification"
	"github.com/xinxiao/logico/unit"
)

const (
	builtinPath = "builtin"
)

var (
	gates = map[string]unit.Unit{
		"nand": &unit.Nand{},
	}

	//go:embed builtin/*.circuit
	//go:embed builtin/**/*.lib
	//go:embed builtin/**/*.circuit
	prebuiltCircuitSource embed.FS

	prebuiltCircuits, _ = LoadPrebuiltCircuit()
)

func LoadPrebuiltCircuit() (map[string]*specification.Blueprint, error) {
	bpp := specification.NewBlueprintParser(prebuiltCircuitSource)

	m := make(map[string]*specification.Blueprint)
	err := fs.WalkDir(prebuiltCircuitSource, builtinPath, func(p string, f fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if f.IsDir() || path.Ext(f.Name()) != specification.CircuitBlueprintFileExtension {
			return nil
		}

		bp, err := bpp.ParseFile(p)
		if err != nil {
			return err
		}

		m[bp.CircuitName] = bp
		return nil
	})

	if err != nil {
		return nil, err
	}

	return m, nil
}

type UnitRepository struct {
	bpr BlueprintRepository
	ubo []UnitBuildOption
}

type UnitBuildOption func(unit.Unit) (unit.Unit, error)

func NewUnitRepository(ubo ...UnitBuildOption) *UnitRepository {
	return &UnitRepository{
		bpr: &emptyBlueprintRepository{},
		ubo: ubo,
	}
}

func (ur *UnitRepository) LinkBlueprintRepository(bpr BlueprintRepository) {
	ur.bpr = bpr
}

func (ur *UnitRepository) GetUnit(n string) (unit.Unit, error) {
	if u, ok := gates[n]; ok {
		return u, nil
	}

	bp, err := ur.loadBlueprint(n)
	if err != nil {
		return nil, err
	}
	return ur.buildCircuitFromBlueprint(bp)
}

func (ur *UnitRepository) loadBlueprint(n string) (*specification.Blueprint, error) {
	if bp, ok := prebuiltCircuits[n]; ok {
		return bp, nil
	}
	return ur.bpr.CheckOutBlueprint(n)
}

func (ur *UnitRepository) buildCircuitFromBlueprint(bp *specification.Blueprint) (unit.Unit, error) {
	c := &unit.Circuit{
		CircuitName:  bp.CircuitName,
		UnitMap:      make(map[string]unit.Unit),
		InputPins:    make(map[specification.Pin]string),
		ConstantPins: make(map[specification.Pin]bool),
		Connections:  make(map[specification.Pin]specification.Pin),
		OutputPins:   make(map[string]specification.Pin),
	}

	for uid, ut := range bp.Nodes {
		u, err := ur.GetUnit(ut)
		if err != nil {
			return nil, err
		}
		c.UnitMap[uid] = u
	}

	for n, ipl := range bp.Inputs {
		for _, ip := range ipl {
			c.InputPins[ip] = n
		}
	}

	for _, cp := range bp.AlwaysOn {
		c.ConstantPins[cp] = true
	}

	for _, cp := range bp.AlwaysOff {
		c.ConstantPins[cp] = false
	}

	for _, e := range bp.Connectors {
		c.Connections[e.To] = e.From
	}

	for n, op := range bp.Outputs {
		c.OutputPins[n] = op
	}

	return c, nil
}
