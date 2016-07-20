#  Middleware example written in goLang (Step-by-Step tutorial)

To install the project `go get github.com/codeselim/middleware-tutorial-go`

## Step 2
The main.go file creates an http server listening on port `8080`
The server handles requests matching the path `"/"`  line 26:  `http.HandleFunc("/", ...`
The servers is bound on the port 8080, listens and serves requests. line 29: `http.ListenAndServe(...`  

## Step: 1
The `main.go` file initializes the application and prints a simple greeting string in the console.
 
Run it `go run main.go`

