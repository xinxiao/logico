package unit

const (
	gateOutputName = "out"
)

func GateOutput(v bool) map[string]bool {
	return map[string]bool{gateOutputName: v}
}

type Gate struct {
	GateName string
}

func (g *Gate) Name() string {
	return g.GateName
}

func (g *Gate) Output() []string {
	return []string{gateOutputName}
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
