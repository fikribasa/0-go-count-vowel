package main

import (
	"fmt"
)

func countHandler(w string) (vowel int, conso int) {

	fmt.Println("Input =", w)
	vowel = 0
	conso = 0

	for _, value := range w {
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
