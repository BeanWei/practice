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
	gen.Funcs["trimRightStr"] = trimRightStr
	return gen.MustParse(gen.NewTemplate(path).
		Funcs(gen.Funcs).
		Parse(text))
}

// trimRightStr strips all of the given <cut> string from the end of a string.
// Note that it does not strips the whitespaces of its end.
//
//	aaa/bbb/ccc => aaa/bbb
//
func trimRightStr(str string, cut string, count ...int) string {
	var (
		lenStr   = len(str)
		lenCut   = len(cut)
		cutCount = 0
	)
	for lenStr >= lenCut && str[lenStr-lenCut:lenStr] == cut {
		lenStr = lenStr - lenCut
		str = str[:lenStr]
		cutCount++
		if len(count) > 0 && count[0] != -1 && cutCount >= count[0] {
			break
		}
	}
	return str
}
