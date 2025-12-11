package main

import (
	"fmt"
	"net/http"
)

func download(w http.ResponseWriter, r *http.Request) {
	data := []byte{}
	n, _ := (r.Body).Read(data)
	fmt.Println(string(data), n)
}

func main() {
	http.HandleFunc("/download", download)
	http.ListenAndServe(":8080", nil)
}
