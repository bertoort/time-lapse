package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file: ", err.Error())
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	router := httprouter.New()
	router.ServeFiles("/public/*filepath", http.Dir("public"))

	// Navigation routes
	router.GET("/", index)
	router.GET("/devteam", index)
	router.GET("/aws-s3", awsS3)
	router.NotFound = http.HandlerFunc(notFound)

	log.Printf("Running on port: %v", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
