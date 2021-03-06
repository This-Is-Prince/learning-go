package ch1

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

func Server3() {
	// http.HandleFunc("/", handler3)
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Print(err)
		}
		cyclesParam := r.Form.Get("cycles")
		log.Print(cyclesParam)
		cycles, err := strconv.Atoi(cyclesParam)
		if err != nil {
			log.Print(err)
			cycles = 5
		}
		lissajous(w, float64(cycles))
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// handler3 echoes the HTTP request.
func handler3(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s %s %s\n", r.Method, r.URL, r.Proto)
	for k, v := range r.Header {
		fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
	fmt.Fprintf(w, "Host = %q\n", r.Host)
	fmt.Fprintf(w, "RemoteAddr = %q\n", r.RemoteAddr)
	if err := r.ParseForm(); err != nil {
		log.Print(err)
	}
	for k, v := range r.Form {
		fmt.Fprintf(w, "Form[%q] = %q\n", k, v)
	}
}
