package main

import (
	"fmt"
)

func countHandler(w string) (vowel int, conso int) {
	vowel = 0
	conso = 0
	// var savedWord rune
	var filteredString []rune
	runes := []rune(w)

	for i := 0; i < len(runes); i++ {
		// Scan slice for a previous element of the same value.
		exists := false
		for v := 0; v < i; v++ {
			if runes[v] == runes[i] {
				exists = true
				break
			}
		}
		// If no previous element exists, append this one.
		if !exists {
			filteredString = append(filteredString, runes[i])
		}
	}

	for _, value := range filteredString {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel++
		default:
			conso++
		}
	}
	fmt.Printf("Huruf mati: %d, Huruf hidup: %d \n \n", conso, vowel)
	return
}

func main() {
	fmt.Println("Go String Vowel-Consonant Counter")
	inputString := readLine("Input Word to Count: ")

	countHandler(inputString)

}

func readLine(word string) string {
	fmt.Print(word)
	var input string
	fmt.Scanln(&input)
	return input
}
