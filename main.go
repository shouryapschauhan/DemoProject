package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/home", HelloWorld)
	http.ListenAndServe(":8080", r)

}

func HelloWorld(w http.ResponseWriter, r *http.Request) {
	log.Println("hello!")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hello World"))
}
