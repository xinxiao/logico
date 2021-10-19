package unit

const (
	GateOutputName = "out"
)

func GateOutput(v bool) map[string]bool {
	return map[string]bool{GateOutputName: v}
}

type Gate struct{}

func (*Gate) Output() []string {
	return []string{GateOutputName}
}

type SingleOperandGate struct {
	Gate
}

func (*SingleOperandGate) Input() []string {
	return []string{"v"}
}

type DoubleOperandGate struct {
	Gate
}

func (*DoubleOperandGate) Input() []string {
	return []string{"a", "b"}
}

type Not struct {
	SingleOperandGate
}

func (*Not) Name() string {
	return "not"
}

func (g *Not) Simulate(args map[string]bool) (map[string]bool, error) {
	return GateOutput(!args["v"]), nil
}

type And struct {
	DoubleOperandGate
}

func (*And) Name() string {
	return "and"
}

func (g *And) Simulate(args map[string]bool) (map[string]bool, error) {
	return GateOutput(args["a"] && args["b"]), nil
}

type Or struct {
	DoubleOperandGate
}

func (*Or) Name() string {
	return "or"
}

func (g *Or) Simulate(args map[string]bool) (map[string]bool, error) {
	return GateOutput(args["a"] || args["b"]), nil
}
