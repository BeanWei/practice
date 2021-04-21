package main

import (
	"context"
	"entdemo/ent"
	"entdemo/ent/entrest"
	"log"

	"github.com/gogf/gf/frame/g"
)

func main() {
	client, err := ent.Open("mysql", "root:1234567890@tcp(localhost:3306)/antadmin")
	if err != nil {
		log.Fatalf("failed connecting to mysql: %v", err)
	}
	defer client.Close()
	ctx := context.Background()
	// Run migration.
	err = client.Schema.Create(ctx)
	if err != nil {
		g.Log().Fatalf("failed creating schema resources: %v", err)
	}

	entrest.NewPetServiceHandler(client)

	s := g.Server()
	s.Run()
}
