package main

import (
	"bufio"
	"fmt"
	"os"
)

const empty = -1 // empty, ie wait.

type TabString struct {
	notes      []int
	stringName string
}

func setupStrings(stringNames []string, allStrings []TabString) []TabString {
	for i := len(stringNames) - 1; i >= 0; i-- {
		newString := TabString{
			notes:      []int{},
			stringName: stringNames[i],
		}
		allStrings = append(allStrings, newString)
		fmt.Println("added stringName ", newString.stringName)
	}
	return allStrings
}

func main() {

	allStrings := []TabString{}
	allStrings = setupStrings([]string{"E", "A", "D", "G", "B", "e"}, allStrings)

	var reader = bufio.NewReader(os.Stdin)
	message, _ := reader.ReadString('\n')

	fmt.Println(message)
}
