package main

import (
	"fmt"
	"net/http"
)

func download(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Connection recieved!")
}

func main() {
	http.HandleFunc("/download", download)

	fmt.Println("Listening on port 8080")

	http.ListenAndServe(":8080", nil)
}
