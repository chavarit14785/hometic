package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// git init, git add server.go go.mod, git commit -m "[Nong] init project"
func main() {
	fmt.Println("hello hometic : I'm Gopher!!")
	r := mux.NewRouter()
	r.HandleFunc("/", IndexHandler).Methods(http.MethodGet)
	r.HandleFunc("/pair-device", PairDeviceHandler).Methods(http.MethodPost)
	server := http.Server{
		Addr:    fmt.Sprintf("0.0.0.0:%s", os.Getenv("PORT")),
		Handler: r,
	}
	log.Println("starting...")
	log.Fatal(server.ListenAndServe())
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(`{"status":"active"}`))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hometic"))
}
