package main

import (
	"log"
	"net/http"
	"os"

	"github.com/builders-lab/trailblazer-frontend/internal/handlers"
	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env")
	

	portString := os.Getenv("PORT")

	webHookSecret := os.Getenv("WEBHOOK_SECRET")

	// Handler configuration structure
	apiCfg := &handlers.ApiConfig{
		WHSecret: webHookSecret,
	}

	// Initializing router
	router := chi.NewRouter()


	// Liberal cors configuration for development
	router.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	v1Router := chi.NewRouter()

	v1Router.Get("/healthz", handlers.HandlerReadiness)
	v1Router.Get("/err", handlers.HandlerErr)
	v1Router.Post("/webhook", apiCfg.HandleWebhook)

	// API version control mounting
	router.Mount("/v1", v1Router)

	srv := &http.Server{
		Handler: router,
		Addr:    ":" + portString,
	}

	log.Printf("Server Starting on port %v", portString)
	
	err := srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
