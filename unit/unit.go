package unit

type Unit interface {
	Name() string
	Input() []string
	Output() []string
	Simulate(map[string]bool) (map[string]bool, error)
}
