package handler

import (
	"fmt"
	"net/http"
	"time"
)

// handler function has the handler type signature
func SayHello(w http.ResponseWriter, r *http.Request) {
	timeNow := time.Now().Format(time.RFC850)
	fmt.Fprintf(w, "Hello, now it is the %s", timeNow)
}
