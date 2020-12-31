package main

import (
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "pong")
}

func main() {
	http.HandleFunc("/ping", homePage)
	log.Fatal(http.ListenAndServe(":80", nil))
}
