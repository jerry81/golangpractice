package greetings

import (
	"fmt"
	"errors"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) { // so the return types are given after the args
    // Return a greeting that embeds the name in a message.
	if name == "" {
        return "", errors.New("empty name")
    }
    message := fmt.Sprintf("Hi, %v. Welcome!", name) // := declares and initializes in a single line 
    return message, nil 
}