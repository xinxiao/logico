package simulation

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/xinxiao/logico/repository"
)

var (
	ur = repository.NewUnitRepository()
)

func MaskForBits(bs int) int64 {
	return (1 << bs) - 1
}

func TestSimulate_Not(t *testing.T) {
	u, err := ur.GetUnit("not")
	if err != nil {
		t.Fatalf("failed to get not unit: %s", err)
	}

	for i, tc := range []struct {
		v   bool
		out bool
	}{
		{v: true, out: false},
		{v: false, out: true},
	} {
		got, err := u.Simulate(map[string]bool{"v": tc.v})
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
	u, err := ur.GetUnit("and")
	if err != nil {
		t.Fatalf("failed to get and unit: %s", err)
	}

	for i, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: true},
		{a: false, b: true, out: false},
		{a: true, b: false, out: false},
		{a: false, b: false, out: false},
	} {
		got, err := u.Simulate(map[string]bool{"a": tc.a, "b": tc.b})
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
	u, err := ur.GetUnit("or")
	if err != nil {
		t.Fatalf("failed to get ot unit: %s", err)
	}

	for i, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: true},
		{a: false, b: true, out: true},
		{a: true, b: false, out: true},
		{a: false, b: false, out: false},
	} {
		got, err := u.Simulate(map[string]bool{"a": tc.a, "b": tc.b})
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
	u, err := ur.GetUnit("nand")
	if err != nil {
		t.Fatalf("failed to get nand unit: %s", err)
	}

	for i, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: false},
		{a: false, b: true, out: true},
		{a: true, b: false, out: true},
		{a: false, b: false, out: true},
	} {
		got, err := u.Simulate(map[string]bool{"a": tc.a, "b": tc.b})
		if got == nil {
			t.Fatalf("unexpected nil and error: %v, %s", got, err)
		}

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
	u, err := ur.GetUnit("nor")
	if err != nil {
		t.Fatalf("failed to get nor unit: %s", err)
	}

	for i, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: false},
		{a: false, b: true, out: false},
		{a: true, b: false, out: false},
		{a: false, b: false, out: true},
	} {
		got, err := u.Simulate(map[string]bool{"a": tc.a, "b": tc.b})
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
	u, err := ur.GetUnit("xor")
	if err != nil {
		t.Fatalf("failed to get xor unit: %s", err)
	}

	for i, tc := range []struct {
		a, b bool
		out  bool
	}{
		{a: true, b: true, out: false},
		{a: false, b: true, out: true},
		{a: true, b: false, out: true},
		{a: false, b: false, out: false},
	} {
		got, err := u.Simulate(map[string]bool{"a": tc.a, "b": tc.b})
		if err != nil {
			t.Fatalf("unexpected error: %s", err)
		}

		expected := map[string]bool{"out": tc.out}
		if !cmp.Equal(got, expected) {
			t.Errorf("tc %d: xor(a: %t, b: %t): %s", i, tc.a, tc.b, cmp.Diff(got, expected))
		}
	}
}

func TestSimulate_Mux(t *testing.T) {
	type TestCase struct {
		v    int64
		cond int64
		out  bool
	}

	for _, cbs := range []int{
		1, 2, 3,
	} {
		bs := 1 << cbs
		n := fmt.Sprintf("mux_%dbit", bs)

		u, err := ur.GetUnit(n)
		if err != nil {
			t.Fatalf("failed to get %s unit: %s", n, err)
		}

		tcl := make([]TestCase, 0)
		for v := 0; v <= int(MaskForBits(bs)); v++ {
			for cond := 0; cond <= int(MaskForBits(cbs)); cond++ {
				tcl = append(tcl, TestCase{
					v:    int64(v),
					cond: int64(cond),
					out:  (v>>cond)&1 == 1,
				})
			}
		}

		t.Run(n, func(t *testing.T) {
			for i, tc := range tcl {
				in := make(map[string]bool)

				for i := 0; i < bs; i++ {
					in[fmt.Sprintf("v_%d", i)] = (tc.v>>i)&1 == 1
				}

				for i := 0; i < cbs; i++ {
					in[fmt.Sprintf("cond_%d", i)] = (int(tc.cond)>>i)&1 == 1
				}

				got, err := u.Simulate(in)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				expected := map[string]bool{"out": tc.out}
				if !cmp.Equal(got, expected) {
					t.Errorf("tc %d: %s(v: %b, cond: %b): %s", i, n, tc.v, tc.cond, cmp.Diff(got, expected))
				}
			}
		})
	}
}

func TestSimulate_Flip(t *testing.T) {
	type TestCase struct {
		v   int64
		out int64
	}

	for _, bs := range []int{
		4, 8,
	} {
		n := fmt.Sprintf("flip_%dbit", bs)
		u, err := ur.GetUnit(n)
		if err != nil {
			t.Fatalf("failed to get %s unit: %s", n, err)
		}

		t.Run(n, func(t *testing.T) {
			m := MaskForBits(bs)

			tcl := make([]*TestCase, 0)
			for i := int64(0); i <= m; i++ {
				tcl = append(tcl, &TestCase{
					v:   i,
					out: (m ^ i) & m,
				})
			}

			for i, tc := range tcl {
				in := map[string]bool{}
				expected := map[string]bool{}

				for i := 0; i < bs; i++ {
					m := int64(0b1) << i

					in[fmt.Sprintf("v_%d", i)] = (tc.v & m) != 0

					expected[fmt.Sprintf("out_%d", i)] = (tc.out & m) != 0
				}

				got, err := u.Simulate(in)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if !cmp.Equal(got, expected) {
					t.Errorf("tc %d: %s(v: %d): %s", i, n, tc.v, cmp.Diff(got, expected))
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

	for _, bs := range []int{
		4, 8,
	} {
		n := fmt.Sprintf("negate_%dbit", bs)
		u, err := ur.GetUnit(n)
		if err != nil {
			t.Fatalf("failed to get %s unit: %s", n, err)
		}

		t.Run(n, func(t *testing.T) {
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
				in := map[string]bool{}
				expected := map[string]bool{}

				for i := 0; i < bs; i++ {
					m := int64(0b1) << i

					in[fmt.Sprintf("v_%d", i)] = (tc.v & m) != 0

					expected[fmt.Sprintf("out_%d", i)] = (tc.out & m) != 0
				}

				got, err := u.Simulate(in)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if !cmp.Equal(got, expected) {
					t.Errorf("tc %d: %s(v: %d): %s", i, n, tc.v, cmp.Diff(got, expected))
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

	for _, bs := range []int{
		1, 2, 4,
	} {
		n := fmt.Sprintf("add_%dbit", bs)
		u, err := ur.GetUnit(n)
		if err != nil {
			t.Fatalf("failed to get %s unit: %s", n, err)
		}

		t.Run(n, func(t *testing.T) {
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

			for i, tc := range tcl {
				in := map[string]bool{"c_in": tc.cIn}
				expected := map[string]bool{"c_out": tc.cOut}

				for i := 0; i < bs; i++ {
					m := int64(0b1) << i

					in[fmt.Sprintf("a_%d", i)] = (tc.a & m) != 0
					in[fmt.Sprintf("b_%d", i)] = (tc.b & m) != 0

					expected[fmt.Sprintf("sum_%d", i)] = (tc.sum & m) != 0
				}

				got, err := u.Simulate(in)
				if err != nil {
					t.Fatalf("unexpected error: %s", err)
				}

				if !cmp.Equal(got, expected) {
					t.Errorf("tc %d: %s(a: %d, b: %d, cIn: %t): %s", i, n, tc.a, tc.b, tc.cIn, cmp.Diff(got, expected))
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

	for _, c := range []int64{
		1, 4,
	} {
		for _, bs := range []int{
			4,
		} {
			n := fmt.Sprintf("add%d_%dbit", c, bs)
			u, err := ur.GetUnit(n)
			if err != nil {
				t.Fatalf("failed to get %s unit: %s", n, err)
			}

			t.Run(n, func(t *testing.T) {
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

				for i, tc := range tcl {
					in := map[string]bool{}
					expected := map[string]bool{"c_out": tc.cOut}

					for i := 0; i < bs; i++ {
						m := int64(0b1) << i

						in[fmt.Sprintf("a_%d", i)] = (tc.a & m) != 0

						expected[fmt.Sprintf("sum_%d", i)] = (tc.sum & m) != 0
					}

					got, err := u.Simulate(in)
					if err != nil {
						t.Fatalf("unexpected error: %s", err)
					}

					if !cmp.Equal(got, expected) {
						t.Errorf("tc %d: %s(a: %d): %s", i, n, tc.a, cmp.Diff(got, expected))
					}
				}
			})
		}
	}
}
