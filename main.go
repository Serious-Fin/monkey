package main

import (
	"fmt"
	"monkey/repl"
	"os"
)

func main() {
	fmt.Println("Monkey Language 0.1")
	repl.Start(os.Stdin, os.Stdout)
}
