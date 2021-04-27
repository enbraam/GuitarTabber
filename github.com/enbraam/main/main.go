package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

const empty = -1 // empty, ie wait.

type TabString struct { // using a map probably would have been better.
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

func codeLine(allStrings []TabString, line string) {
	length := len(allStrings[0].notes) + 1 // the length of everything should now be 1 greater.

	if strings.ToLower(line) == "delete" {
		for i := 0; i < len(allStrings); i++ {
			arrayLen := len(allStrings[i].notes)
			if arrayLen == 1 {
				allStrings[i].notes = []int{}
			} else if arrayLen > 0 {
				allStrings[i].notes = allStrings[i].notes[:arrayLen-1]
			}
		}
		return
	} else if strings.ToLower(line) != "rest" {
		entries := strings.Fields(line)
		for _, entry := range entries {
			entrySplit := strings.Split(entry, ":")
			fmt.Print(entrySplit)
			for i := 0; i < len(allStrings); i++ {
				if entrySplit[0] == allStrings[i].stringName {
					note, _ := strconv.Atoi(entrySplit[1])
					allStrings[i].notes = append(allStrings[i].notes, note) // add the note to the string array
					break                                                   // user could still add 2 notes to same string. This is bad. Need to handle later.
				}
			}
		}
	}

	for i := 0; i < len(allStrings); i++ {
		if len(allStrings[i].notes) < length {
			allStrings[i].notes = append(allStrings[i].notes, empty)
		}
	}
}

func printStrings(allStrings []TabString) { // this is still broken for double didget numbers
	for _, s := range allStrings {
		fmt.Print(s.stringName)
		fmt.Print("|----")
		for _, n := range s.notes {
			if n == empty {
				fmt.Print("-")
			} else {
				note := strconv.Itoa(n)
				fmt.Print(note)
			}
		}
		fmt.Print("\n")
	}
}

func clear() {
	for i := 0; i < 10; i++ {
		fmt.Println("")
	}
}

func main() {

	allStrings := []TabString{}
	allStrings = setupStrings([]string{"E", "A", "D", "G", "B", "e"}, allStrings)

	var reader = bufio.NewReader(os.Stdin)

	for {
		message, _, err := reader.ReadLine()
		messageS := string(message[:])

		fmt.Print("message is ", messageS)

		if strings.ToLower(messageS) == "stop" {
			fmt.Println("BREAK now stop detected")
			break
		}

		if len(messageS) != 0 && err == nil {
			codeLine(allStrings, messageS)
			clear()
			printStrings(allStrings)
		} else {
			break
		}
	}

}
