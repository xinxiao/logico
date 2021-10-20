package unit

import (
	"fmt"
)

type Nand struct{}

func (*Nand) Name() string {
	return "nand"
}

func (*Nand) Input() []string {
	return []string{"a", "b"}
}
func (*Nand) Output() []string {
	return []string{"out"}
}

func (*Nand) Simulate(args map[string]bool) (map[string]bool, error) {
	if _, ok := args["a"]; !ok {
		return nil, fmt.Errorf("missing input a")
	}

	if _, ok := args["b"]; !ok {
		return nil, fmt.Errorf("missing input b")
	}

	return map[string]bool{"out": !(args["a"] && args["b"])}, nil
}
