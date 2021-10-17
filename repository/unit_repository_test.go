package repository

import (
	"testing"
)

func TestLoadBuiltinCircuit(t *testing.T) {
	if l, err := LoadPrebuiltCircuit(); err != nil {
		t.Error("Failed to load builtin circuit blueprint:", err)
	} else if len(l) == 0 {
		t.Error("failed to load builtin circuits: no circuits loaded", l)
	}
}
