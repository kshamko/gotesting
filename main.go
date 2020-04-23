package main

import (
	"encoding/json"
	"log"
	"net/http"

	"lect5/fib"

	"github.com/gorilla/mux"

	_ "net/http/pprof"
)

type fiboutput struct {
	N            int64 `json:"n"`
	Nthfibonacci int64 `json:"fibonacci(n)"`
}

func fibHandler(w http.ResponseWriter, r *http.Request) {
	v, _ := fib.Fib(12)
	fibout := fiboutput{N: 12, Nthfibonacci: v}
	json.NewEncoder(w).Encode(fibout)
}

func fibTailHandler(w http.ResponseWriter, r *http.Request) {
	v, _ := fib.FibImproved(10)
	fibout := fiboutput{N: 10, Nthfibonacci: v}
	json.NewEncoder(w).Encode(fibout)
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/fib/", fibHandler)
	router.HandleFunc("/fibtail/", fibTailHandler)
	log.Println("Listening.........")

	//start web server for pprof
	go func() {
		http.ListenAndServe("localhost:6060", nil)
	}()

	//start web server for fib endpoints
	http.ListenAndServe(":8080", router)
}
