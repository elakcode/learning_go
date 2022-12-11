package main

import (
	"elakcode/greetings"
	"fmt"
	"log"
)

func main(){
    //setting properties of the predefined logger
    log.SetPrefix("greetings: ")
    log.SetFlags(0)
    
    //message, err := greetings.Hello("") //uncomment to see error
    message, err := greetings.Hello("Elakcode")

    if err != nil {
        log.Fatal(err)
    }

    fmt.Println(message)

    //for a number of users

    //A slice of names
    names := []string{"ElakA", "ElakB", "ElakC"}
    
    messages, err := greetings.Hellos(names)
    if err != nil {
        log.Fatal(err)
    }
    fmt.Println(messages)
    
}
