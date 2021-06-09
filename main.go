package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("What's your name?")
	name := readName(os.Stdout)
	fmt.Println("Hello " + name)
}

func readName(r io.Reader) string {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	name := scanner.Text()
	if name == "" {
		name = "You"
	}
	return name
}
