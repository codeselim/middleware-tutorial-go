#  Middleware example written in goLang (Step-by-Step tutorial)

To install the project `go get github.com/codeselim/middleware-tutorial-go`

## Step 6
Centralized requests handling.

In order to improve our architecture we wrapped all handlers / handler functions in a `HandlerWrapper`.

The HandlerWrapper will be the unique entry point to execute handler Functions.

In order to have a `Handler` type we need to implement the Handler interface. In other words, our type (HandlerWrapper) must implement a method with signature `ServeHTTP(res http.ResponseWriter, req *http.Request)`

The routing/routing.go file was adapted in order to receive handler type instead of Handler Functions

## Step 5
In this step, we create our custom errors. 

In order to create our own custom errors, with custom structures, our error-types must implement the go `error` interface.
Any go type that implements a method with signature `"Error() string"` respects the `error` interface and thus, can be used as type `error`
ref. https://blog.golang.org/error-handling-and-go 

A 'user id' validation in the `GetUserById` method in `handler/greeter.go` is added in order to demonstrate the usage of our custom errors. 

## Step 4
In this step, we connect our middleware application to an external API `http://jsonplaceholder.typicode.com/users`

Two new endpoints are added to our routing configuration `/users` and `/users/:id`

Calling our `users` endpoint, will make our application fetch data from the External API and return its results

Similarly, calling `/users/:id` will return a user's data based on the specified id

The `handler/greeter.go defines all our handler functions and the logic for connecting and fetching data form the external API

Note that at the moment, we are not handling errors. If an error occurred (e.g. http error while connecting to the external API) our application will crash. In the next steps we will refactor the code and add error handling. 


## Step 3
The handler functions are now extracted from the `main.go` and moved to the handler package

At this step the http server uses a router to decide on routing. In this example we are using the gorilla/mux package.
Ref. https://github.com/gorilla/mux

The routing configuration is done in the `routing/routing.go` file.  

## Step 2
The `main.go` file creates an http server listening on port `8080`

The server handles requests matching the path `"/"`  line 26:  `http.HandleFunc("/", ...`

The servers is bound on the port 8080, listens and serves requests. line 29: `http.ListenAndServe(...`  

## Step: 1
The `main.go` file initializes the application and prints a simple greeting string in the console.
 
Run it `go run main.go`

