package unit

const (
	GateOutput = "out"
)

type Gate struct {
	n string
}

func (g *Gate) Name() string {
	return g.n
}

func (g *Gate) Output() []string {
	return []string{GateOutput}
}

type SingleOperandGate struct {
	Gate
}

func (g *SingleOperandGate) Input() []string {
	return []string{"v"}
}

type DoubleOperandGate struct {
	Gate
}

func (g *DoubleOperandGate) Input() []string {
	return []string{"a", "b"}
}
