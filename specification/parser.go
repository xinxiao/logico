package specification

import (
	"embed"
	"encoding/json"
	"io"
	"io/fs"
	"path"
	"strings"

	"github.com/google/go-jsonnet"
)

const (
	libPath         = "lib/circuit"
	libImportPrefix = "@"
)

var (
	//go:embed lib/circuit/*.lib
	circuitLib embed.FS
)

type CircuitBlueprintImporter struct {
	fs fs.FS
	c  map[string]jsonnet.Contents
}

func (cfi *CircuitBlueprintImporter) FindImportPath(loc, p string) string {
	if strings.HasPrefix(p, libImportPrefix) {
		return p
	}
	return path.Join(path.Dir(loc), p)
}

func (cfi *CircuitBlueprintImporter) FindSource(p string) ([]byte, error) {
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

func (cfi *CircuitBlueprintImporter) Import(loc, p string) (jsonnet.Contents, string, error) {
	ip := cfi.FindImportPath(loc, p)

	if c, ok := cfi.c[ip]; ok {
		return c, ip, nil
	}

	b, err := cfi.FindSource(ip)
	if err != nil {
		return jsonnet.Contents{}, "", err
	}

	c := jsonnet.MakeContents(string(b))
	cfi.c[ip] = c
	return c, ip, nil
}

type BlueprintParser struct {
	vm *jsonnet.VM
	fs fs.FS
}

func NewBlueprintParser(fs fs.FS) *BlueprintParser {
	vm := jsonnet.MakeVM()
	vm.Importer(&CircuitBlueprintImporter{fs, make(map[string]jsonnet.Contents)})
	return &BlueprintParser{vm, fs}
}

func (bpp *BlueprintParser) ParseFile(f string) (*Blueprint, error) {
	src, err := bpp.vm.EvaluateFile(f)
	if err != nil {
		return nil, err
	}
	return bpp.ParseJson(src)
}

func (bpp *BlueprintParser) ParseJsonnet(src string) (*Blueprint, error) {
	src, err := bpp.vm.EvaluateAnonymousSnippet("src", string(src))
	if err != nil {
		return nil, err
	}
	return bpp.ParseJson(src)
}

func (bpp *BlueprintParser) ParseJson(src string) (*Blueprint, error) {
	cm := &Blueprint{}
	if err := json.Unmarshal([]byte(src), cm); err != nil {
		return nil, err
	}
	return cm, nil
}
