package main

import (
	"log"
	"net/http"

	adapters "github.com/crazyvan/go-specs-greet/adapters"
)

func main() {
	handler := http.HandlerFunc(adapters.GreetHandler)
	if err := http.ListenAndServe(":8080", handler); err != nil {
		log.Fatal(err)
	}
}
