package blueprint

import (
	"embed"
	"encoding/json"
	"fmt"
	"io"
	"io/fs"
	"path"
	"strings"

	"github.com/google/go-jsonnet"
)

const (
	builtinPath     = "circuit/builtin"
	libPath         = "circuit/lib"
	libImportPrefix = "~"
)

var (
	//go:embed circuit/lib/*.libsonnet
	circuitLib embed.FS

	//go:embed circuit/builtin/*.jsonnet
	builtInCircuits embed.FS
)

type CircuitPin struct {
	UnitId string `json:"uid"`
	PinId  string `json:"pid"`
}

type CircuitEdge struct {
	From CircuitPin `json:"from"`
	To   CircuitPin `json:"to"`
}

type CircuitBlueprint struct {
	Name      string                  `json:"name"`
	Nodes     map[string]string       `json:"nodes"`
	AlwaysOn  []CircuitPin            `json:"always_on,omitempty"`
	AlwaysOff []CircuitPin            `json:"always_off,omitempty"`
	Inputs    map[string][]CircuitPin `json:"inputs"`
	Edges     []CircuitEdge           `json:"edges,omitempty"`
	Outputs   map[string]CircuitPin   `json:"outputs"`
}

func LoadBuiltinCircuitBlueprint() ([]*CircuitBlueprint, error) {
	dir, err := builtInCircuits.ReadDir(builtinPath)
	if err != nil {
		return nil, err
	}

	l := make([]*CircuitBlueprint, 0)
	cbpp := NewCircuitBlueprintParser(builtInCircuits)
	for _, f := range dir {
		cbp, err := cbpp.ParseFile(fmt.Sprintf("%s/%s", builtinPath, f.Name()))
		if err != nil {
			return nil, err
		}
		l = append(l, cbp)
	}
	return l, nil
}

type circuitBlueprintImporter struct {
	fs fs.FS
	c  map[string]jsonnet.Contents
}

func (cfi *circuitBlueprintImporter) findSource(p string) ([]byte, error) {
	if strings.HasPrefix(p, libImportPrefix) {
		return circuitLib.ReadFile(path.Join(libPath, strings.TrimPrefix(p, libImportPrefix)))
	}

	f, err := cfi.fs.Open(p)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}

func (cfi *circuitBlueprintImporter) Import(_, p string) (jsonnet.Contents, string, error) {
	if c, ok := cfi.c[p]; ok {
		return c, p, nil
	}

	b, err := cfi.findSource(p)
	if err != nil {
		return jsonnet.Contents{}, "", err
	}

	c := jsonnet.MakeContents(string(b))
	cfi.c[p] = c
	return c, p, nil
}

type CircuitBlueprintParser struct {
	vm *jsonnet.VM
	fs fs.FS
}

func NewCircuitBlueprintParser(fs fs.FS) *CircuitBlueprintParser {
	vm := jsonnet.MakeVM()
	vm.Importer(&circuitBlueprintImporter{fs, make(map[string]jsonnet.Contents)})
	return &CircuitBlueprintParser{vm, fs}
}

func (cbpp *CircuitBlueprintParser) ParseFile(f string) (*CircuitBlueprint, error) {
	src, err := cbpp.vm.EvaluateFile(f)
	if err != nil {
		return nil, err
	}
	return cbpp.ParseJsonSource(src)
}

func (cbpp *CircuitBlueprintParser) ParseJsonnetSource(src string) (*CircuitBlueprint, error) {
	src, err := cbpp.vm.EvaluateAnonymousSnippet("src", string(src))
	if err != nil {
		return nil, err
	}
	return cbpp.ParseJsonSource(src)
}

func (cbpp *CircuitBlueprintParser) ParseJsonSource(src string) (*CircuitBlueprint, error) {
	cm := &CircuitBlueprint{}
	if err := json.Unmarshal([]byte(src), cm); err != nil {
		return nil, err
	}
	return cm, nil
}
