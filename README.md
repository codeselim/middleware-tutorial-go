#  Middleware example written in goLang (Step-by-Step tutorial)

To install the project `go get github.com/codeselim/middleware-tutorial-go`

## Step 9
Contracts with external APIs and requests transformation.

At this step, we want our API to return a user resource in a different structure than what is returned from the external API. Thus, the data fetched from the external API is mapped/converted in order to suit our API contract. The implementation is found in the `mapper` package.

`Contracts` between the client and our API and between our API and the external API are in the package `contract`.

## Step 8
Refactoring. 

In this step all connection to the external APIs are extracted from the `handler` package and moved to the connection package.

Furthermore, functions are wrapped in structures and dependencies in properties. In this way, isolating components when testing becomes easier.

Additionally, the `connection` package defines interfaces. All the implementations in the `connection` packages respect those interfaces and thus, mocking the handlers dependencies is now possible.

## Step 7
Centralized error Handling.

Since in the last step we implemented a centralized requests handling, we can now handle all errors in one place and thus having a centralized error handling.

In this step we added `error` return type to our handlers.  `error == nil` signifies that no errors are present (goLang convention)
All handlers return the error object as is (if present). Then, the HandlerWrapper manages the logic of error handling.

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

