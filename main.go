package main

import (
	"fmt"
)

func countHandler() (vowel int, conso int) {
	fmt.Println("Go String Vowel-Consonant Counter")
	inputString := readLine("Input Word to Count: ")

	fmt.Println(inputString)
	vowel = 0
	conso = 0

	for _, value := range inputString {
		switch value {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			vowel++
		default:
			conso++
		}
	}
	fmt.Printf("Huruf mati: %d, Huruf hidup: %d ", conso, vowel)
	return
}

func main() {
	countHandler()

}

func readLine(word string) string {
	fmt.Print(word)
	var input string
	fmt.Scanln(&input)
	return input
}
