# 1. Getting Started #
* create a new hello directory and cd into it.
* use "go mod init" to initialise the go module 
    ```go
        go mod init
    ```
* package main : it is the executable package
* import the necessary module : we used inbuilt "fmt" package 
* if external package is used then use "go mod tidy" to fetch the module 
    ```go
        go mod tidy 
    ```
* use the "go run ." command to run the program
    ```go
        go run .
    ```

# 2. Creating a Go module #
* Create a new directory greetings and write the code into it after go mod init.
* This time the package will be something else other than main as it won't be an executable package.
* To import a package form local module , we can use a replace command 
    ```go
        go mod edit -replace elakcode/greetings=../greetings 
    ``` 
* Adding error handling
    * Add basic check if string name is empty then return a new error 
        ```go
            import "errors"
            
            errors.New("error string") // return type is "error"
        ```
    * we will now test what happens in case of error , we will use default logger
        ```go
        import "log"

        log.SetPrefix("greetings: ") // prefix of log
        log.SetFlags(0) // seting time, date etc. to 0 for now , can be changed

        //now we can print the logs 
        log.Fatal(err)  // it will print the error and exit the execution
        ```
* We will add randomness in the greetings.
    * init : golang will automatically execution this function when program starts.
    * slices are basically arrays in golang 
        ```go
        //we can use slices of string as 
        names := []string{
            "elakcode",
            "codelak",
            "elaksetin",
            "viraj",
        }
        ```
* We will use a map for creating greetings for multiple users and use for loop
    ```go

    //To go over a slice we can use a for loop 
    for _, name := range names {  // names := []string
        //...
    }

    //a map can be defined as follows
    messages := make(map[string]string) // map[key_type]value_type
    ```
* Ending a file's name with _test.go tells the go test command that this file contains test functions.
    ```go
        import (
            "testing"
        )

        func TestName(t *testing.T) {
            // we can use t.Fatalf() in case of an error 
            // it will print a message on the screen and end execution
        }
        
        // to run the testcase we can use following commands
        go test     // it executes function whose names begin with Test in test file

        go test -v // verbose, outputs list of all tests and their results. 
    ```
* Building and installing the binaries
    ```go

    // To build
    go build

    // To get the Go install path 
    go list -f '{{.Target}}'

    // The above command will give a target path which we can add to $PATH env variable
    // After that to install use 
    go install 
    ``` 
