package entrest

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"unicode"

	"entgo.io/ent/entc/gen"
	"golang.org/x/tools/imports"
)

const PkgName = "debug"

// Generate .
func Generate(g *gen.Graph, tplName string) error {
	tpl, ok := templates[tplName]
	if !ok {
		return fmt.Errorf("not supports web frameworks: %s", tplName)
	}

	var assets assets
	for _, schema := range g.Schemas {
		assets.dirs = append(assets.dirs, filepath.Join(g.Config.Target, PkgName))
		b := bytes.NewBuffer(nil)
		typ, err := gen.NewType(g.Config, schema)
		if err != nil {
			return fmt.Errorf("get type %q: %w", "", err)
		}
		if err := tpl.Execute(b, typ); err != nil {
			return fmt.Errorf("execute template %q: %w", "", err)
		}

		assets.files = append(assets.files, file{
			path:    filepath.Join(g.Config.Target, PkgName, fmt.Sprintf("%s_service.go", snake(schema.Name))),
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

// snake converts the given struct or field name into a snake_case.
//
//	Username => username
//	FullName => full_name
//	HTTPCode => http_code
//
func snake(s string) string {
	var (
		j int
		b strings.Builder
	)
	for i := 0; i < len(s); i++ {
		r := rune(s[i])
		// Put '_' if it is not a start or end of a word, current letter is uppercase,
		// and previous is lowercase (cases like: "UserInfo"), or next letter is also
		// a lowercase and previous letter is not "_".
		if i > 0 && i < len(s)-1 && unicode.IsUpper(r) {
			if unicode.IsLower(rune(s[i-1])) ||
				j != i-1 && unicode.IsLower(rune(s[i+1])) && unicode.IsLetter(rune(s[i-1])) {
				j = i
				b.WriteString("_")
			}
		}
		b.WriteRune(unicode.ToLower(r))
	}
	return b.String()
}
