package main

import (
	"fmt"
	"log"
	"net/http"
)

func download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connection recieved!")
}

func main() {
	http.HandleFunc("/download", download)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
