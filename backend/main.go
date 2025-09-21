package main

import (
	"log"
	"net/http"
	"os"

	"github.com/AmiyoKm/basic_http/config"
	"github.com/AmiyoKm/basic_http/infra"
	"github.com/AmiyoKm/basic_http/middleware"
	"github.com/AmiyoKm/basic_http/repo"
	productHandler "github.com/AmiyoKm/basic_http/rest/product"
	userHandler "github.com/AmiyoKm/basic_http/rest/user"
	productService "github.com/AmiyoKm/basic_http/service/product"
	userService "github.com/AmiyoKm/basic_http/service/user"
)

func main() {
	cfg := config.NewConfig()

	db, err := infra.InitDB(cfg.DbConfig)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	userRepo := repo.NewUserRepo(db)
	productRepo := repo.NewProductRepo(db)

	userService := userService.NewService(userRepo)
	productService := productService.NewService(productRepo)

	uHandler := userHandler.NewHandler(cfg, userService)
	pHandler := productHandler.NewHandler(cfg, productService)

	mux := http.NewServeMux()
	manager := middleware.NewManager(cfg)
	manager.Use(manager.CorsMiddleware, manager.Logger)

	uHandler.HttpRoutes(mux, manager)
	pHandler.HttpRoutes(mux, manager)

	server := http.Server{
		Addr:    cfg.Addr,
		Handler: mux,
	}

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
