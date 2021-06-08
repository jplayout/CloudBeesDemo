package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var nom string
	fmt.Println("What's your name?")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nom = scanner.Text()
	if nom == "" {
		nom = "You"
	}
	fmt.Println("Hello " + nom)
}
