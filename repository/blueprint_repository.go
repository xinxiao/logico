package repository

import (
	"fmt"

	"github.com/xinxiao/logico/blueprint"
)

type BlueprintRepository interface {
	CheckInBlueprint(*blueprint.CircuitBlueprint) error
	CheckOutBlueprint(string) (*blueprint.CircuitBlueprint, error)
	DeleteBlueprint(string) error
	ListAllBlueprints() ([]string, error)
}

type emptyBlueprintRepository struct{}

func (*emptyBlueprintRepository) CheckInBlueprint(_ *blueprint.CircuitBlueprint) error {
	return fmt.Errorf("should not check in any blueprint")
}

func (*emptyBlueprintRepository) CheckOutBlueprint(_ string) (*blueprint.CircuitBlueprint, error) {
	return nil, fmt.Errorf("should not check out any blueprint")
}

func (*emptyBlueprintRepository) DeleteBlueprint(n string) error {
	return fmt.Errorf("should not delete any blueprint")
}

func (*emptyBlueprintRepository) ListAllBlueprints() ([]string, error) {
	return []string{}, nil
}
