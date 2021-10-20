package unit

func CountTransistors(u Unit) int {
	c, ok := u.(*Circuit)
	if !ok {
		return 1
	}

	t := 0
	for _, u := range c.UnitMap {
		t += CountTransistors(u)
	}
	return t
}
