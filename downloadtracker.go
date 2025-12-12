package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func enableCors(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
}

func download(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("downloads.txt")

	if err != nil {
		fmt.Println(err)
		return
	}

	dls, _ := strconv.Atoi(string(data))

	os.WriteFile("downloads.txt", []byte(strconv.Itoa(dls+1)), 0666)
	enableCors(w)
	w.Write([]byte(strconv.Itoa(dls)))

	fmt.Println("New download")
}

func getDownloads(w http.ResponseWriter, r *http.Request) {
	data, err := os.ReadFile("downloads.txt")

	if err != nil {
		w.Write([]byte("0"))
		return
	}

	dls, _ := strconv.Atoi(string(data))
	enableCors(w)
	w.Write([]byte(strconv.Itoa(dls)))

	fmt.Println("Requested DLs")
}

func main() {
	http.HandleFunc("/download", download)
	http.HandleFunc("/getdownloads", getDownloads)

	fmt.Println("Listening on port 8080")

	log.Fatal(http.ListenAndServe(":8080", nil))
}
