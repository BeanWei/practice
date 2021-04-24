package entrest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"entgo.io/ent/entc/gen"
	"golang.org/x/tools/imports"
)

const PkgName = "debug"

// var (
// 	snake = gen.Funcs["snake"].(func(string) string)
// 	pascal = gen.Funcs["pascal"].(func(string) string)
// 	camel  = gen.Funcs["camel"].(func(string) string)
// )

// Generate .
func Generate(g *gen.Graph, tplName string) error {
	tpl, ok := templates[tplName]
	if !ok {
		return fmt.Errorf("not supports web frameworks: %s", tplName)
	}

	var assets assets
	for _, n := range g.Nodes {
		if annotation, ok := n.Annotations[EntRESTAPI]; ok {
			apiAnnotation := annotation.(map[string]interface{})
			if fmt.Sprintf("%v", apiAnnotation["SkipGen"]) == "true" {
				continue
			}
		}

		assets.dirs = append(assets.dirs, filepath.Join(g.Config.Target, PkgName))
		b := bytes.NewBuffer(nil)
		if err := tpl.Execute(b, n); err != nil {
			return fmt.Errorf("execute template %q: %w", "", err)
		}

		assets.files = append(assets.files, file{
			path:    filepath.Join(g.Config.Target, PkgName, fmt.Sprintf("%s_service.go", strings.ToLower(n.Name))),
			content: b.Bytes(),
		})
	}

	if err := assets.write(); err != nil {
		return err
	}

	return assets.format()
}

type (
	file struct {
		path    string
		content []byte
	}
	assets struct {
		dirs  []string
		files []file
	}
)

// write files and dirs in the assets.
func (a assets) write() error {
	for _, dir := range a.dirs {
		if err := os.MkdirAll(dir, os.ModePerm); err != nil {
			return fmt.Errorf("create dir %q: %w", dir, err)
		}
	}
	for _, file := range a.files {
		if err := ioutil.WriteFile(file.path, file.content, 0644); err != nil {
			return fmt.Errorf("write file %q: %w", file.path, err)
		}
	}
	return nil
}

// format runs "goimports" on all assets.
func (a assets) format() error {
	for _, file := range a.files {
		path := file.path
		src, err := imports.Process(path, file.content, nil)
		if err != nil {
			return fmt.Errorf("format file %s: %w", path, err)
		}
		if err := ioutil.WriteFile(path, src, 0644); err != nil {
			return fmt.Errorf("write file %s: %w", path, err)
		}
	}
	return nil
}
