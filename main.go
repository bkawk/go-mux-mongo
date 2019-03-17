package main

import (
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
)


func main() {
	godotenv.Load()
	port := os.Getenv("PORT")
	router := NewRouter()

	if err := http.ListenAndServe(":"+port, router); err != nil {
		log.Fatal(err)
	}
}
