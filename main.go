package main

import (
	"log"
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load(".env");
	PortString := os.Getenv("PORT");
	
	if len(PortString) == 0 {
		log.Fatal("PORT env not found...!");
	}

	router := chi.NewRouter();

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedHeaders: []string{"*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		ExposedHeaders: []string{"Link"},
		AllowCredentials: false,
		MaxAge: 300,
	}))

	srv := &http.Server{
		Handler: router,
		Addr: ":" + PortString,
	};

	v1Router := chi.NewRouter()
	v1Router.Get("/healthz", handlerReadiness)

	router.Mount("/v1", v1Router)

	log.Printf("Server starting on port: %v", PortString);

	err := srv.ListenAndServe();
	if err != nil {
		log.Fatal(err);
	}
}