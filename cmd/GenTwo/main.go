package main

import (
	"fmt"
	"github.com/GenTwo-Developers/GenTwo-server/handlers"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello GenTwo Developers!")
	fmt.Println("Listening on http://localhost:8080!")
	http.HandleFunc("/", handlers.RootHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
