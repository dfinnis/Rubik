package main

import (
	rubik "Rubik/src"
	"fmt"//
)

func main() {
	fmt.Printf("oh hi!\n")/////////
	rubik.RunRubik()
	fmt.Printf("\nEND!!\n")//////////
}

// ## To run enter either command:
// go run main.go "$(< mix/subject.txt)"
// go build; ./Rubik "$(< mix/subject.txt)"