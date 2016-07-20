package main

import (
	"fmt"
	"time"
)

func sayHello(name string) (string, string){
	greeting := fmt.Sprintf("Hello %s", name)
	timeNow := time.Now().Format(time.RFC850)
	return greeting, timeNow
}

// init() is always called, regardless if there's main function or not, so if you import a package that has an init function, it will be executed.
func init() {
	fmt.Println("Initializing...")
}

func main(){
	greetingString, timeNowString := sayHello("Go")
	fmt.Println(fmt.Sprintf("Main says: %s, time: %s", greetingString, timeNowString))
}

