package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

func Greet(writer io.Writer, name string) {
	// fmt.Printf("Hello, %s", name)
	fmt.Fprintf(writer, "Hello, %s", name)
}

func myGreeterHandler(w http.ResponseWriter, r *http.Request) {
	Greet(w, "world")
}

func main() {
	log.Fatal(http.ListenAndServe(":5001", http.HandlerFunc(myGreeterHandler)))
}
