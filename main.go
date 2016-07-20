package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	port = "8080"
)

// handler function has the handler type signature
func sayHello(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, "Hello, now it is the %s", timeNow)
}

// init() is always called, regardless if there's main function or not, so if you import a package that has an init function, it will be executed.
func init() {
	fmt.Println("Initializing...")
}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Println("Binding http server on port: " + port)
	log.Fatal(
		http.ListenAndServe(":"+port, nil),
	)
}
