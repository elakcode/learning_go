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
