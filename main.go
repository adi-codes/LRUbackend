package main

import (
	"LRUbackend/controller"
	"LRUbackend/service"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	cacheService := service.NewCacheService(1024)
	cacheController := controller.NewCacheController(cacheService)

	r.HandleFunc("/api/cache/set", cacheController.Set).Methods("POST")
	r.HandleFunc("/api/cache/get", cacheController.Get).Methods("GET")

	headers := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"})
	methods := handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"})
	origins := handlers.AllowedOrigins([]string{"*"})

	log.Println("Listening on :8080...")

	log.Fatal(http.ListenAndServe(":8080", handlers.CORS(headers, methods, origins)(r)))
}
