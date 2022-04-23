package main // way to group functions 

import (
	"fmt" // format package 
    "rsc.io/quote"
	"example.com/greetings"
)

func main() { // entrypoint 
    fmt.Println("hello")
    fmt.Println(quote.Go())
	message := greetings.Hello("Jerry")
	fmt.Println(message)
}