package main // way to group functions 

import (
	"fmt" // format package 
    "rsc.io/quote"
	"example.com/greetings"
	"log"
)

func main() { // entrypoint 
	log.SetPrefix("greetings: ")
    log.SetFlags(0)
    fmt.Println("hello")
    fmt.Println(quote.Go())
	message, _ := greetings.Hello("Jerry")
	fmt.Println(message)
	message2, err2 := greetings.Hello("")
	if err2 != nil {
		log.Fatal(err2)
	}
	fmt.Println(message2)
}