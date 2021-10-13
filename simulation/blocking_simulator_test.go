package simulation

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xinxiao/logico/unit"
)

func TestSimulate_And(t *testing.T) {
	ur, err := unit.NewUnitRegistry()
	if err != nil {
		t.Fatalf("failed to create unit registry: %s", err)
	}

	sim := NewBlockingSimulator(ur)

	for _, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: true},
		{a: false, b: true, out: false},
		{a: true, b: false, out: false},
		{a: false, b: false, out: false},
	} {
		got, err := sim.Simulate("and", map[string]bool{"a": tc.a, "b": tc.b})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("and: %v: got %v, want %v", tc, got, expected)
		}
	}
}

func TestSimulate_Or(t *testing.T) {
	ur, err := unit.NewUnitRegistry()
	if err != nil {
		t.Fatalf("failed to create unit registry: %s", err)
	}

	sim := NewBlockingSimulator(ur)

	for _, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: true},
		{a: false, b: true, out: true},
		{a: true, b: false, out: true},
		{a: false, b: false, out: false},
	} {
		got, err := sim.Simulate("or", map[string]bool{"a": tc.a, "b": tc.b})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("or: %v: got %v, want %v", tc, got, expected)
		}
	}
}

func TestSimulate_Not(t *testing.T) {
	ur, err := unit.NewUnitRegistry()
	if err != nil {
		t.Fatalf("failed to create unit registry: %s", err)
	}

	sim := NewBlockingSimulator(ur)

	for _, tc := range []struct {
		v   bool
		out bool
	}{
		{v: true, out: false},
		{v: false, out: true},
	} {
		got, err := sim.Simulate("not", map[string]bool{"v": tc.v})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("not: %v: got %v, want %v", tc, got, expected)
		}
	}
}

func TestSimulate_Nand(t *testing.T) {
	ur, err := unit.NewUnitRegistry()
	if err != nil {
		t.Fatalf("failed to create unit registry: %s", err)
	}

	sim := NewBlockingSimulator(ur)

	for _, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: false},
		{a: false, b: true, out: true},
		{a: true, b: false, out: true},
		{a: false, b: false, out: true},
	} {
		got, err := sim.Simulate("nand", map[string]bool{"a": tc.a, "b": tc.b})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("nand: %v: got %v, want %v", tc, got, expected)
		}
	}
}

func TestSimulate_Xor(t *testing.T) {
	ur, err := unit.NewUnitRegistry()
	if err != nil {
		t.Fatalf("failed to create unit registry: %s", err)
	}

	sim := NewBlockingSimulator(ur)

	for _, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: false},
		{a: false, b: true, out: true},
		{a: true, b: false, out: true},
		{a: false, b: false, out: false},
	} {
		got, err := sim.Simulate("xor", map[string]bool{"a": tc.a, "b": tc.b})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("xor: %v: got %v, want %v", tc, got, expected)
		}
	}
}

func TestSimulate_Add(t *testing.T) {
	ur, err := unit.NewUnitRegistry()
	if err != nil {
		t.Fatalf("failed to create unit registry: %s", err)
	}

	sim := NewBlockingSimulator(ur)

	for _, tc := range []struct {
		a, b, c_in bool
		sum, c_out bool
	}{
		{a: false, b: false, c_in: false, sum: false, c_out: false},
		{a: false, b: false, c_in: true, sum: true, c_out: false},
		{a: false, b: true, c_in: false, sum: true, c_out: false},
		{a: false, b: true, c_in: true, sum: false, c_out: true},
		{a: true, b: false, c_in: false, sum: true, c_out: false},
		{a: true, b: false, c_in: true, sum: false, c_out: true},
		{a: true, b: true, c_in: false, sum: false, c_out: true},
		{a: true, b: true, c_in: true, sum: true, c_out: true},
	} {
		got, err := sim.Simulate("add", map[string]bool{"a": tc.a, "b": tc.b, "c_in": tc.c_in})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"sum": tc.sum, "c_out": tc.c_out}
		if !cmp.Equal(got, expected) {
			t.Errorf("xor: %v: got %v, want %v", tc, got, expected)
		}
	}
}
