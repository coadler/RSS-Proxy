package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := "80"
	router := NewRouter()

	fmt.Println("Successfully listening on port " + port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}
