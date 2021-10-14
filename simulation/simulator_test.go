package simulation

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xinxiao/logico/repository"
)

const (
	maxTestCasesSize = 10000
)

func TestSimulate_Not(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
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
			t.Errorf("tc %d: not(v: %t): %s", i, tc.v, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_And(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
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
			t.Errorf("tc %d: and(a: %t, b: %t): %s", i, tc.a, tc.b, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Or(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
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
			t.Errorf("tc %d: or(a: %t, b: %t): %s", i, tc.a, tc.b, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Nand(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
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
			t.Errorf("tc %d: nand(a: %t, b: %t): %s", i, tc.a, tc.b, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Nor(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: false},
		{a: false, b: true, out: false},
		{a: true, b: false, out: false},
		{a: false, b: false, out: true},
	} {
		got, err := sim.Simulate("nor", map[string]bool{"a": tc.a, "b": tc.b})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("tc %d: nand(a: %t, b: %t): %s", i, tc.a, tc.b, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Xor(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
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
			t.Errorf("tc %d: xor(a: %t, b: %t): %s", i, tc.a, tc.b, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_If(t *testing.T) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range []struct {
		a, b, cond bool
		out        bool
	}{
		{a: false, b: false, cond: false, out: false},
		{a: false, b: false, cond: true, out: false},
		{a: false, b: true, cond: false, out: true},
		{a: false, b: true, cond: true, out: false},
		{a: true, b: false, cond: false, out: false},
		{a: true, b: false, cond: true, out: true},
		{a: true, b: true, cond: false, out: true},
		{a: true, b: true, cond: true, out: true},
	} {
		got, err := sim.Simulate("if", map[string]bool{"a": tc.a, "b": tc.b, "cond": tc.cond})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("tc %d: if(a: %t, b: %t, c: %t): %s", i, tc.a, tc.b, tc.cond, cmp.Diff(got, expected))
		}
	}
}

func getMaskForSize(bs int) int64 {
	return (1 << bs) - 1
}

type flipTestCase struct {
	v   int64
	out int64
}

func generateFlipTestCases(bs int) []*flipTestCase {
	m := getMaskForSize(bs)

	l := make([]*flipTestCase, 0)
	for i := int64(0); i <= m; i++ {
		l = append(l, &flipTestCase{
			v:   i,
			out: (m ^ i) & m,
		})
	}
	return l
}

func testSimulate_Flip(t *testing.T, bs int) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range generateFlipTestCases(bs) {
		flip := fmt.Sprintf("flip_%dbit", bs)

		in := map[string]bool{}
		expected := map[string]bool{}

		for i := 0; i < bs; i++ {
			m := int64(0b1) << i

			in[fmt.Sprintf("v_%d", i)] = (tc.v & m) != 0

			expected[fmt.Sprintf("out_%d", i)] = (tc.out & m) != 0
		}

		got, err := sim.Simulate(flip, in)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if !cmp.Equal(got, expected) {
			t.Errorf("tc %d: %s(v: %d): %s", i, flip, tc.v, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Flip8Bit(t *testing.T) {
	testSimulate_Flip(t, 8)
}

func TestSimulate_Flip16Bit(t *testing.T) {
	testSimulate_Flip(t, 16)
}

type addTestCase struct {
	a, b int64
	cIn  bool
	sum  int64
	cOut bool
}

func generateAddTestCases(bs int) []*addTestCase {
	m := getMaskForSize(bs)

	l := make([]*addTestCase, 0)
	for a := m; a >= 0; a-- {
		for b := m; b >= 0; b-- {
			s := a + b

			l = append(l,
				&addTestCase{
					a: a, b: b, cIn: false,
					sum: s & m, cOut: s > m,
				},
				&addTestCase{
					a: a, b: b, cIn: true,
					sum: (s + 1) & m, cOut: s >= m,
				})

			if len(l) >= maxTestCasesSize {
				return l
			}
		}
	}

	return l
}

func testSimulate_Add(t *testing.T, bs int) {
	ur, err := repository.NewUnitRepository()
	if err != nil {
		t.Fatalf("failed to create unit repository: %s", err)
	}

	sim := NewSimulator(ur)

	for i, tc := range generateAddTestCases(bs) {
		add := fmt.Sprintf("add_%dbit", bs)

		in := map[string]bool{"c_in": tc.cIn}
		expected := map[string]bool{"c_out": tc.cOut}

		for i := 0; i < bs; i++ {
			m := int64(0b1) << i

			in[fmt.Sprintf("a_%d", i)] = (tc.a & m) != 0
			in[fmt.Sprintf("b_%d", i)] = (tc.b & m) != 0

			expected[fmt.Sprintf("sum_%d", i)] = (tc.sum & m) != 0
		}

		got, err := sim.Simulate(add, in)
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		if !cmp.Equal(got, expected) {
			t.Errorf("tc %d: %s(a: %d, b: %d, cIn: %t): %s", i, add, tc.a, tc.b, tc.cIn, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Add1Bit(t *testing.T) {
	testSimulate_Add(t, 1)
}

func TestSimulate_Add2Bit(t *testing.T) {
	testSimulate_Add(t, 2)
}

func TestSimulate_Add4Bit(t *testing.T) {
	testSimulate_Add(t, 4)
}

func TestSimulate_Add8Bit(t *testing.T) {
	testSimulate_Add(t, 8)
}

func TestSimulate_Add16Bit(t *testing.T) {
	testSimulate_Add(t, 16)
}
