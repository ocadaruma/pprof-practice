package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/blocking", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(500 * time.Millisecond)

		w.Write([]byte("OK"))
	})

	http.ListenAndServe(":8081", r)
}
