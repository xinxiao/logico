package unit

import (
	"fmt"
)

func CountTransistor(u Unit) (int, int, int, error) {
	switch u := u.(type) {
	case *Not:
		return 1, 0, 0, nil
	case *And:
		return 0, 1, 0, nil
	case *Or:
		return 0, 0, 1, nil
	case *Circuit:
		n, a, o := 0, 0, 0
		for _, su := range u.UnitMap {
			sn, sa, so, err := CountTransistor(su)
			if err != nil {
				return 0, 0, 0, err
			}
			n, a, o = n+sn, a+sa, o+so
		}
		return n, a, o, nil
	default:
		return 0, 0, 0, fmt.Errorf("unknown unit type: %T", u)
	}
}
