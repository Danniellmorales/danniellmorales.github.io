package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	router()

	fmt.Printf("Starting server at port 4000\n")
	if err := http.ListenAndServe(":3000", nil); err != nil {
		log.Fatal(err)
	}
}
