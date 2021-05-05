package entrest

import (
	"entdemo/entrest/internal"

	"entgo.io/ent/entc/gen"
)

var (
	srvTemplate = parse("template/service.tmpl")
)

//go:generate go run github.com/go-bindata/go-bindata/go-bindata -o=internal/bindata.go -pkg=internal -modtime=1 ./template

func parse(path string) *gen.Template {
	text := string(internal.MustAsset(path))
	return gen.MustParse(gen.NewTemplate(path).
		Funcs(gen.Funcs).
		Parse(text))
}
