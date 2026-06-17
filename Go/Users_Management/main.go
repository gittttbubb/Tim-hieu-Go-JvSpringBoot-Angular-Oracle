package main

import (
	"log"
	"strconv"

	"go-rbac-system/internal/bootstrap"
	"go-rbac-system/internal/config"
	"go-rbac-system/internal/routes"
)

func main() {
	cfg, err := config.Load("configs/dev.yaml")
	if err != nil {
		log.Fatal(err)
	}
	db, err := bootstrap.InitDatabase(cfg)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 1. container DI
	container := bootstrap.NewContainer(cfg, db)

	// 2. build route config
	routeCfg := bootstrap.BuildRouteConfig(container)

	// 3. fiber app (NO ROUTE INSIDE HERE)
	app := bootstrap.NewApp()

	// 4. register routes ONLY HERE
	routes.RegisterAllRoutes(app, routeCfg)
	port := strconv.Itoa(cfg.App.Port)
	log.Printf("Server started on port %s", port)
	log.Fatal(app.Listen(":" + port))
}