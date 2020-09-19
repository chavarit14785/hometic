package main

import (
	"database/sql"
	"encoding/json"
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

type PairDevice struct {
	DeviceID int `json:"DeviceID"`
	UserID   int `json:"UserID"`
}

func PairDeviceHandler(w http.ResponseWriter, r *http.Request) {
	var pd PairDevice
	err := json.NewDecoder(r.Body).Decode(&pd)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	fmt.Printf("result : %#v\n", pd)
	pd.insert()
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"status":"active"}`))
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("hometic"))
}

func (pd PairDevice) insert() {
	db, err := sql.Open("postgres", os.Getenv("DATABASE_URL"))
	//db, err := sql.Open("postgres", "postgres://cxebbtvt:zcvqQofbKqwG-iSwVJEWJecutSonnbDP@arjuna.db.elephantsql.com:5432/cxebbtvt")
	if err != nil {
		log.Fatal("connect to database error", err)
	}
	defer db.Close()
	_, err = db.Exec("INSERT INTO pairs VALUES ($1,$2);", pd.DeviceID, pd.UserID)
	if err != nil {
		log.Fatal("can't insert", err)
	}
	fmt.Println("insert table success.")
}
