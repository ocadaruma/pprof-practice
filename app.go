package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"net/http/pprof"
	"io/ioutil"
)

func blockingHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := http.Get("http://localhost:8081/blocking")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	var body []byte
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	w.Write(body)
}

func cpuIntensiveHandler(w http.ResponseWriter, r *http.Request) {
	str := ""

	for i := 0; i < 10000; i++ {
		if i % 2 == 0 {
			str += "e"
		} else {
			str += "o"
		}
	}

	w.Write([]byte(str[:10]))
}

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/blocking", blockingHandler)
	r.HandleFunc("/cpu-intensive", cpuIntensiveHandler)

	r.HandleFunc("/debug/pprof/", pprof.Index)
	r.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
	r.HandleFunc("/debug/pprof/profile", pprof.Profile)
	r.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	r.HandleFunc("/debug/pprof/trace", pprof.Trace)

	http.ListenAndServe(":8080", r)
}
