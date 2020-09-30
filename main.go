package main

import (
	"log"
	"math"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("ok"))
}

func handler(w http.ResponseWriter, r *http.Request) {
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	w.Write([]byte("OK!"))
}

func main() {
	http.HandleFunc("/healthz", status)
	http.HandleFunc("/readyz", status)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
