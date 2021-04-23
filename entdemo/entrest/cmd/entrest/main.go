package main

import (
	"entdemo/entrest"
	"flag"
	"log"

	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	var (
		schemaPath = flag.String("p", "", "path to schema directory")
		tplName    = flag.String("f", "", "web frameworks: supports gf")
	)
	flag.Parse()
	if *schemaPath == "" {
		log.Fatal("entrest: must specify schema path. use entrest -path ./ent/schema")
	}
	graph, err := entc.LoadGraph(*schemaPath, &gen.Config{})
	if err != nil {
		log.Fatalf("entrest: failed loading ent graph: %v", err)
	}
	if err := entrest.Generate(graph, *tplName); err != nil {
		log.Fatalf("entrest: failed generating entrest-services: %s", err)
	}
}
