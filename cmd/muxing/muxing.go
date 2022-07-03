package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods("GET")

	router.HandleFunc("/name/{PARAM}", func(writer http.ResponseWriter, request *http.Request) {
		vv, ok := mux.Vars(request)["PARAM"]
		if !ok {
			http.Error(writer, "empty param", http.StatusBadRequest)
		}
		data := fmt.Sprintf("Hello, %s!", vv)
		fprintf, err := fmt.Fprintf(writer, data)
		if err != nil {
			http.Error(writer, err.Error(), http.StatusBadRequest)
		}
		fmt.Printf("%d bytes were written\n", fprintf)

	}).Methods("GET")

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		body, err := ioutil.ReadAll(r.Body)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		fprintf, err := fmt.Fprintf(w, "I got message:\n%v", string(body))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		fmt.Printf("%d bytes were written\n", fprintf)
	}).Methods("POST")

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		a := r.Header.Get("a")
		b := r.Header.Get("b")

		aI, err := strconv.Atoi(a)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		bI, err := strconv.Atoi(b)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
		}

		res := aI + bI

		w.Header().Set("a+b", strconv.Itoa(res))

	}).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
