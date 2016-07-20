package main

import (
	"fmt"
	"github.com/codeselim/middleware-tutorial-go/routing"
	"log"
	"net/http"
)

const (
	port = "8080"
)

// init() is always called, regardless if there's main function or not, so if you import a package that has an init function, it will be executed.
func init() {
	fmt.Println("Initializing...")
}

func main() {
	r := routing.NewRouter()
	http.Handle("/", r)
	fmt.Println("Binding http server on port: " + port)
	log.Fatal(
		http.ListenAndServe(":"+port, nil),
	)
}
