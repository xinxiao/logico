package unit

type Unit interface {
	Name() string
	Input() []string
	Output() []string
}
