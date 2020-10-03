package main

import (
	"fmt"
	"log"
	"math"
	"net/http"
)

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "ok")
}

func handler(w http.ResponseWriter, r *http.Request) {
	/*
		Taken from hpa example that uses php
		https://kubernetes.io/docs/tasks/run-application/horizontal-pod-autoscale-walkthrough/#run-expose-php-apache-server
	*/
	x := 0.0001
	for i := 0; i <= 1000000; i++ {
		x += math.Sqrt(x)
	}
	fmt.Fprint(w, "OK!")
}

func main() {
	http.HandleFunc("/healthz", status)
	http.HandleFunc("/readyz", status)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
