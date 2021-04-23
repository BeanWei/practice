package entrest

import (
	"entdemo/entrest/internal"

	"entgo.io/ent/entc/gen"
)

// Support web frameworks
const (
	GF = "gf"
)

var (
	templates = map[string]*gen.Template{
		GF: parse("template/gf.tmpl"),
	}
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -modtime=1 ./template

func parse(path string) *gen.Template {
	text := string(internal.MustAsset(path))
	return gen.MustParse(gen.NewTemplate(path).
		Funcs(gen.Funcs).
		Parse(text))
}
