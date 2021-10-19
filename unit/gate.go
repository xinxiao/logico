package unit

const (
	gateOutputName = "out"
)

func gateOutput(v bool) map[string]bool {
	return map[string]bool{gateOutputName: v}
}

type gate struct{}

func (*gate) Output() []string {
	return []string{gateOutputName}
}

type singleOperandGate struct {
	gate
}

func (*singleOperandGate) Input() []string {
	return []string{"v"}
}

type doubleOperandGate struct {
	gate
}

func (*doubleOperandGate) Input() []string {
	return []string{"a", "b"}
}

type Not struct {
	singleOperandGate
}

func (*Not) Name() string {
	return "not"
}

func (g *Not) Simulate(args map[string]bool) (map[string]bool, error) {
	return gateOutput(!args["v"]), nil
}

type And struct {
	doubleOperandGate
}

func (*And) Name() string {
	return "and"
}

func (g *And) Simulate(args map[string]bool) (map[string]bool, error) {
	return gateOutput(args["a"] && args["b"]), nil
}

type Or struct {
	doubleOperandGate
}

func (*Or) Name() string {
	return "or"
}

func (g *Or) Simulate(args map[string]bool) (map[string]bool, error) {
	return gateOutput(args["a"] || args["b"]), nil
}
