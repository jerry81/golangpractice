package main // way to group functions 

import "fmt" // format package 

import "rsc.io/quote"

func main() { // entrypoint 
    fmt.Println("hello")
    fmt.Println(quote.Go())
}