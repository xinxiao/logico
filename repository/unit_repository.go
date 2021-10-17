package repository

import (
	"embed"
	"fmt"
	"io/fs"
	"log"
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
		&unit.SingleOperandGate{Gate: unit.Gate{GateName: "not"}},
		&unit.DoubleOperandGate{Gate: unit.Gate{GateName: "and"}},
		&unit.DoubleOperandGate{Gate: unit.Gate{GateName: "or"}},
	}

	prebuiltCircuits, _ = LoadPrebuiltCircuit()

	builtinCircuits = append(baseGates, prebuiltCircuits...)
)

func LoadPrebuiltCircuit() ([]unit.Unit, error) {
	cbpp := blueprint.NewCircuitBlueprintParser(builtInCircuitSource)
	l := make([]unit.Unit, 0)
	if err := fs.WalkDir(builtInCircuitSource, builtinPath, func(p string, f fs.DirEntry, err error) error {
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

		l = append(l, unit.BuildCircuitFromBlueprint(cbp))
		return nil
	}); err != nil {
		log.Fatal(err)
		return nil, err
	}
	return l, nil
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

	for _, u := range builtinCircuits {
		ur.cm[u.Name()] = u
	}

	return ur
}

func (ur *UnitRepository) LinkBlueprintRepository(bpr BlueprintRepository) {
	ur.bpr = bpr
}

func (ur *UnitRepository) GetUnit(n string) (unit.Unit, error) {
	if p, ok := ur.cm[n]; ok {
		return p, nil
	}

	if bp, err := ur.bpr.CheckOutBlueprint(n); err == nil {
		return unit.BuildCircuitFromBlueprint(bp), err
	}

	return nil, fmt.Errorf("unit %s not found", n)
}

func (ur *UnitRepository) ListAllUnits() (chan string, chan error) {
	sc := make(chan string)
	ec := make(chan error)
	go ur.iterateThroughAllUnits(sc, ec)
	return sc, ec
}

func (ur *UnitRepository) iterateThroughAllUnits(sc chan string, ec chan error) {
	defer close(sc)
	defer close(ec)

	for _, u := range ur.cm {
		sc <- u.Name()
	}

	l, err := ur.bpr.ListAllBlueprints()
	if err != nil {
		ec <- err
		return
	}

	for _, bp := range l {
		sc <- bp
	}
}
