// +build ignore

package main

import (
	"log"

	"entdemo/entrest"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	err := entc.Generate("./schema", &gen.Config{
		Templates: []*gen.Template{
			entrest.ServiceTemplate,
		},
	})
	if err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}
