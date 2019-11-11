package main

import "fmt"

func countHandler(){
	inputString:= "aaoms"

	fmt.Println(inputString)
	vowel:=0
	conso:=0

	for _, value := range inputString {
		switch value {
			case 'a','e','i','o','u','A','E','I','O','U' :
				vowel++
			default:
				conso++
		}
	}

	fmt.Printf("Huruf mati: %d, Huruf hidup: %d ",conso,vowel)

}

func main() {
	countHandler()
}
