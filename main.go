package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router()

	fmt.Printf("Starting server at port 3000\n")
	if err := http.ListenAndServe(":4000", nil); err != nil {
		log.Fatal(err)
	}
}
