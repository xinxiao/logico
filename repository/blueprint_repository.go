package repository

import (
	"fmt"

	"github.com/xinxiao/logico/specification"
)

type BlueprintRepository interface {
	CheckInBlueprint(*specification.Blueprint) error
	CheckOutBlueprint(string) (*specification.Blueprint, error)
	DeleteBlueprint(string) error
	ListAllBlueprints() ([]string, error)
}

type emptyBlueprintRepository struct{}

func (*emptyBlueprintRepository) CheckInBlueprint(_ *specification.Blueprint) error {
	return fmt.Errorf("should not check in any blueprint")
}

func (*emptyBlueprintRepository) CheckOutBlueprint(_ string) (*specification.Blueprint, error) {
	return nil, fmt.Errorf("should not check out any blueprint")
}

func (*emptyBlueprintRepository) DeleteBlueprint(n string) error {
	return fmt.Errorf("should not delete any blueprint")
}

func (*emptyBlueprintRepository) ListAllBlueprints() ([]string, error) {
	return []string{}, nil
}
