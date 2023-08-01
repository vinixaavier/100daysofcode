// learning resources
// https://www.w3schools.com/go/go_syntax.php
// https://gobyexample.com/hello-world
// https://go.dev/doc/tutorial/getting-started

// A Go file consists of the following parts:

// Package declaration
// Import packages
// Functions
// Statements and expressions

// Package declaration
package main

// Importing Packages
import (
	"fmt"
)

// Functions, statements and expressions
func main() {
	fmt.Println("Hello World!")
}

// Line 1: In Go, every program is part of a package. We define this using the package keyword. In this example, the program belongs to the main package.

// Line 2: import ("fmt") lets us import files included in the fmt package.

// Line 3: A blank line. Go ignores white space. Having white spaces in code makes it more readable.

// Line 4: func main() {} is a function. Any code inside its curly brackets {} will be executed.

// Line 5: fmt.Println() is a function made available from the fmt package. It is used to output/print text. In our example it will output "Hello World!".

// go run <name of the go file>
// go build <name of the go file>
// ./<binary name>

// go run hello-world.go
// go build hello-world.go
// ./hello-world
