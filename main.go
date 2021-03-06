package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/user"

	"github.com/abs-lang/abs/repl"
)

// The ABS interpreter
func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	args := os.Args

	// if we're called without arguments,
	// launch the REPL
	if len(args) == 1 {
		fmt.Printf("Hello %s, welcome to the ABS programming language!\n", user.Username)
		fmt.Printf("Type 'quit' when you're done, 'help' if you get lost!\n")
		repl.Start(os.Stdin, os.Stdout)
		return
	}

	// let's parse our argument as a file
	code, err := ioutil.ReadFile(args[1])

	if err != nil {
		panic(err)
	}

	repl.Run(string(code), false)
}
