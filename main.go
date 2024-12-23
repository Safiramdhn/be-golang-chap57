package main

import (
	"be-chap56/infra"
	"be-chap56/routes"
	"log"
)

func main() {
	ctx, err := infra.NewServiceContext()
	if err != nil {
		log.Fatal("can't init service context %w", err)
	}

	routes.NewRoutes(*ctx)
}
