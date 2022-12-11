package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Hello(name string) (string, error) {
    if name == "" {
        return "", errors.New("empty name")
    }   

    //message := fmt.Sprintf(randomFormat()) // uncomment for testing break 
    message := fmt.Sprintf(randomFormat(), name)
    return message, nil
}

// return greeting for multiple people
func Hellos(names []string) (map[string]string, error) {
    messages := make(map[string]string)

    for _, name := range(names) { // we can loop through slices like this
        message, err := Hello(name)

        if err != nil {
            return nil, err
        }

        messages[name] = message 
    }

    return messages, nil
}





//init functions will run in the beginning
func init() {
    rand.Seed(time.Now().UnixNano())
}

// it returns one of the random string from a list called slice in golang
func randomFormat() string {
    formats := []string {
        "Hi, %v. Welcome!",
        "Great to see you, %v",
        "Welcome dude, %v !",
        "Nice to meet you %v !",
    }

    return formats[rand.Intn(len(formats))]
}
