package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.Handle("/", http.FileServer(http.Dir(".")))

	fmt.Print("Listening on: https://localhost:8080\n")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
