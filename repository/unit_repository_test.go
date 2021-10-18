package repository

import (
	"testing"
)

func TestLoadBuiltinCircuit(t *testing.T) {
	if m, err := LoadPrebuiltCircuit(); err != nil {
		t.Error("failed to load builtin circuit blueprint:", err)
	} else if len(m) == 0 {
		t.Error("failed to load builtin circuits: no circuits loaded", m)
	}
}
