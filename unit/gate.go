package unit

const (
	gateOutputName = "out"
)

func GateOutput(v bool) map[string]bool {
	return map[string]bool{gateOutputName: v}
}

type Gate struct{}

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

type Not struct {
	SingleOperandGate
}

func (*Not) Name() string {
	return "not"
}

func (*Not) Simulate(args map[string]bool) (map[string]bool, error) {
	return GateOutput(!args["v"]), nil
}

type And struct {
	DoubleOperandGate
}

func (*And) Name() string {
	return "and"
}

func (*And) Simulate(args map[string]bool) (map[string]bool, error) {
	return GateOutput(args["a"] && args["b"]), nil
}

type Or struct {
	DoubleOperandGate
}

func (*Or) Name() string {
	return "or"
}

func (*Or) Simulate(args map[string]bool) (map[string]bool, error) {
	return GateOutput(args["a"] || args["b"]), nil
}
