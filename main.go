package main

import (
	"duolingo-api/api"
	"log"
	"net/http"
)

func main(){
	log.Println("Server at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", http.HandlerFunc(api.Handler)))
}

