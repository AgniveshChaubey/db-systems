package main

import (
	"fmt"
	"log"
	"net/http"
)

func getData(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello world\n")
}

func main() {
	http.HandleFunc("/data", getData)

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal("Error: ", err)
	}
}
