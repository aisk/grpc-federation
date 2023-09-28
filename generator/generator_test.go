package generator_test

import (
	"context"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"

	"github.com/mercari/grpc-federation/generator"
)

func TestGenerateAll(t *testing.T) {
	binDir, err := filepath.Abs("../bin")
	if err != nil {
		t.Fatal(err)
	}
	path := os.Getenv("PATH")
	t.Setenv("PATH", fmt.Sprintf("%s:%s", binDir, path))

	const standardPluginNum = 3

	t.Run("explicit_declare_standard_plugins", func(t *testing.T) {
		content := `
imports:
  - ../testdata
src:
  - ../testdata
out: .
plugins:
  - plugin: go
    opt: paths=source_relative
  - plugin: go-grpc
    opt: paths=source_relative
  - plugin: grpc-federation
    opt: paths=source_relative
`
		cfg, err := generator.LoadConfigFromReader(strings.NewReader(content))
		if err != nil {
			t.Fatal(err)
		}
		g := generator.New(cfg)
		protoToRespMap, err := g.GenerateAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for name, res := range protoToRespMap {
			if len(res) != standardPluginNum {
				t.Fatalf("failed to generate standard plugin for %s. code generator response number is %d", name, len(res))
			}
		}
	})
	t.Run("implicit_declare_standard_plugins", func(t *testing.T) {
		content := `
imports:
  - ../testdata
src:
  - ../testdata
out: .
`
		cfg, err := generator.LoadConfigFromReader(strings.NewReader(content))
		if err != nil {
			t.Fatal(err)
		}
		g := generator.New(cfg)
		protoToRespMap, err := g.GenerateAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		for name, res := range protoToRespMap {
			if len(res) != standardPluginNum {
				t.Fatalf("failed to generate standard plugin for %s. code generator response number is %d", name, len(res))
			}
		}
	})
	t.Run("additional_plugin", func(t *testing.T) {
		content := `
imports:
  - ../testdata
src:
  - ../testdata
out: .
plugins:
  - plugin: validate-go
`
		cfg, err := generator.LoadConfigFromReader(strings.NewReader(content))
		if err != nil {
			t.Fatal(err)
		}
		g := generator.New(cfg)
		protoToRespMap, err := g.GenerateAll(context.Background())
		if err != nil {
			t.Fatal(err)
		}
		pluginNum := standardPluginNum + 1 // validate-go
		for name, res := range protoToRespMap {
			if len(res) != pluginNum {
				t.Fatalf("failed to generate standard plugin for %s. code generator response number is %d", name, len(res))
			}
		}
	})
}
