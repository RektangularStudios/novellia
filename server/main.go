package main

import (
	"fmt"
	"net/http"
)

func dummy(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "dummy");
}

func main() {
	port := 8080
	
	fmt.Printf("Starting server on port %d", port);

	http.HandleFunc("/dummy", dummy);

	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil);
	if err != nil {
		fmt.Printf("Error starting server")
	}
}
