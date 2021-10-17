package simulation

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xinxiao/logico/repository"
)

func MaskForBits(bs int) int64 {
	return (1 << bs) - 1
}

func TestSimulate_Not(t *testing.T) {
	ur := repository.NewUnitRepository()
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
	ur := repository.NewUnitRepository()
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
	ur := repository.NewUnitRepository()
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
	ur := repository.NewUnitRepository()
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
	ur := repository.NewUnitRepository()
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
	ur := repository.NewUnitRepository()
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
	ur := repository.NewUnitRepository()
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

func TestSimulate_Flip(t *testing.T) {
	type TestCase struct {
		v   int64
		out int64
	}

	ur := repository.NewUnitRepository()
	sim := NewSimulator(ur)

	for _, bs := range []int{
		4, 8,
	} {
		t.Run(fmt.Sprintf("%dbit", bs), func(t *testing.T) {
			m := MaskForBits(bs)

			tcl := make([]*TestCase, 0)
			for i := int64(0); i <= m; i++ {
				tcl = append(tcl, &TestCase{
					v:   i,
					out: (m ^ i) & m,
				})
			}

			for i, tc := range tcl {
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
		})
	}
}

func TestSimulate_Negate(t *testing.T) {
	type TestCase struct {
		v   int64
		out int64
	}

	ur := repository.NewUnitRepository()
	sim := NewSimulator(ur)

	for _, bs := range []int{
		4, 8,
	} {
		t.Run(fmt.Sprintf("%dbit", bs), func(t *testing.T) {
			tcl := []*TestCase{{v: 0, out: 0}}
			for i := int64(1); i <= MaskForBits(bs-1); i++ {
				tcl = append(tcl,
					&TestCase{
						v:   i,
						out: -i,
					},
					&TestCase{
						v:   -i,
						out: i,
					})
			}

			for i, tc := range tcl {
				negate := fmt.Sprintf("negate_%dbit", bs)

				in := map[string]bool{}
				expected := map[string]bool{}

				for i := 0; i < bs; i++ {
					m := int64(0b1) << i

					in[fmt.Sprintf("v_%d", i)] = (tc.v & m) != 0

					expected[fmt.Sprintf("out_%d", i)] = (tc.out & m) != 0
				}

				got, err := sim.Simulate(negate, in)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if !cmp.Equal(got, expected) {
					t.Errorf("tc %d: %s(v: %d): %s", i, negate, tc.v, cmp.Diff(got, expected))
				}
			}

		})
	}
}

func TestSimulate_Add(t *testing.T) {
	type TestCase struct {
		a, b int64
		cIn  bool
		sum  int64
		cOut bool
	}

	ur := repository.NewUnitRepository()
	sim := NewSimulator(ur)

	for _, bs := range []int{
		1, 2, 4,
	} {
		m := MaskForBits(bs)

		tcl := make([]*TestCase, 0)
		for a := m; a >= 0; a-- {
			for b := m; b >= 0; b-- {
				s := a + b

				tcl = append(tcl,
					&TestCase{
						a: a, b: b, cIn: false,
						sum: s & m, cOut: s > m,
					},
					&TestCase{
						a: a, b: b, cIn: true,
						sum: (s + 1) & m, cOut: s >= m,
					})
			}
		}

		t.Run(fmt.Sprintf("%dbit", bs), func(t *testing.T) {
			for i, tc := range tcl {
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
		})
	}
}

func TestSimulate_AddConstant(t *testing.T) {
	type TestCase struct {
		a    int64
		sum  int64
		cOut bool
	}

	ur := repository.NewUnitRepository()
	sim := NewSimulator(ur)

	for _, c := range []int64{
		1, 4,
	} {
		for _, bs := range []int{
			4,
		} {
			m := MaskForBits(bs)

			tcl := make([]*TestCase, 0)
			for a := m; a >= 0; a-- {
				for b := m; b >= 0; b-- {
					s := a + c
					tcl = append(tcl,
						&TestCase{
							a: a, sum: s & m, cOut: s > m,
						})
				}
			}

			t.Run(fmt.Sprintf("%d_%dbit", c, bs), func(t *testing.T) {
				for i, tc := range tcl {
					add := fmt.Sprintf("add%d_%dbit", c, bs)

					in := map[string]bool{}
					expected := map[string]bool{"c_out": tc.cOut}

					for i := 0; i < bs; i++ {
						m := int64(0b1) << i

						in[fmt.Sprintf("a_%d", i)] = (tc.a & m) != 0

						expected[fmt.Sprintf("sum_%d", i)] = (tc.sum & m) != 0
					}

					got, err := sim.Simulate(add, in)
					if err != nil {
						t.Fatalf("unexpected error: %s", err)
					}

					if !cmp.Equal(got, expected) {
						t.Errorf("tc %d: %s(a: %d): %s", i, add, tc.a, cmp.Diff(got, expected))
					}
				}
			})
		}
	}
}
