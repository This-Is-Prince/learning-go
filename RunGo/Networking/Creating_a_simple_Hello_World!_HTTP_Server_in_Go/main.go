package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Creating a simple hello world http server in golang")

	// writing_Hello_World_Server()
	// understanding_Mux()
	// using_Http_DefaultServeMux()
	returning_A_Better_Response()
}

/*
=================
Writing a "Hello World!" Server
=================
*/
// create a handler struct
type HttpHandler struct{}

// implement `ServeHTTP` method on `HttpHandler` struct
func (h HttpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	/*
	   // create response binary data
	   	data := []byte("<h1>Hello World!</h1>") // slice of bytes

	   	// write `data` to response
	   	res.Write(data)
	*/

	// write `Hello` using `io.WriteString` function
	io.WriteString(res, "<h1>Hello")

	// write `World` using `fmt.Fprint` function
	fmt.Fprint(res, " World! ")

	// write `❤️` using simple `Write` call
	res.Write([]byte("❤️</h1>"))
}

func writing_Hello_World_Server() {
	// create a new handler
	handler := HttpHandler{}

	// listen and serve
	log.Fatal(http.ListenAndServe(":9000", handler))
}

/*
=================
Understanding ServeMux
=================
*/
func understanding_Mux() {
	// create a new `ServeMux``
	mux := http.NewServeMux()

	// handle `/` route
	mux.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// handle `/hello/golang` route
	mux.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve using `ServeMux`
	log.Fatal(http.ListenAndServe(":9000", mux))
}

/*
=================
Using http.DefaultServeMux
=================
*/

func using_Http_DefaultServeMux() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello World!")
	})

	// handle `/hello/golang` route to `http.DefaultServeMux`
	http.HandleFunc("/hello/golang", func(res http.ResponseWriter, req *http.Request) {
		fmt.Fprint(res, "Hello Golang!")
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))
}

/*
=================
Returning a better response
=================
*/

func returning_A_Better_Response() {
	// handle `/` route to `http.DefaultServeMux`
	http.HandleFunc("/", func(res http.ResponseWriter, req *http.Request) {
		// get response headers
		header := res.Header()

		// set content type header
		header.Set("Content-Type", "application/json")

		// reset date header (inline call)
		res.Header().Set("Date", "18-03-2022")

		// set status header
		res.WriteHeader(http.StatusBadRequest) // http.StatusBadRequest == 400

		// respond with a JSON string
		fmt.Fprint(res, `{"status": "FAILURE"}`)
	})

	// listen and serve using `http.DefaultServeMux`
	log.Fatal(http.ListenAndServe(":9000", nil))
}
