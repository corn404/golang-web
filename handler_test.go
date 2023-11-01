package golang_web

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T) {
	var handler http.HandlerFunc = func(writer http.ResponseWriter, reques *http.Request) {
		fmt.Fprint(writer, "Hello World")
	}
	server := http.Server{
		Addr:    "localhost:5000",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServerMux(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Hello WOrld")
	})

	mux.HandleFunc("/hello", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Fprint(writer, "Belajar Golang Web")
	})

	server := http.Server{
		Addr:    "localhost:5000",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}

}
