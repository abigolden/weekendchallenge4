package main

import (
	"algorithms"
	"fmt"
)

func main() {

	//Segundo ejercicio

	text1 := "amazon"
	text2 := "dad"

	result1 := algorithms.IsPalindrome(text1)
	result2 := algorithms.IsPalindrome(text2)

	fmt.Printf("Is '%s' a palindrome? %t\n", text1, result1)
	fmt.Printf("Is '%s' a palindrome? %t\n", text2, result2)

	// primer ejercicio
	arrBooks := []string{
		"Cien años de soledad",
		"1984",
		"Matar un ruiseñor",
		"El Gran Gatsby",
		"En busca del tiempo perdido",
		"Moby-Dick",
		"Don Quijote de la Mancha",
		"Ulises",
		"El amor en los tiempos del cólera",
		"Mujer en punto cero",
		"Los juegos del hambre",
		"El Señor de los Anillos",
		"La Odisea",
		"Crimen y castigo",
		"Orgullo y prejuicio",
		"Las aventuras de Sherlock Holmes",
		"La metamorfosis",
		"El Principito",
		"Los pilares de la Tierra",
		"Los hombres me explican cosas",
	}

	sortedBooks := SortBooks(arrBooks)
	fmt.Println(sortedBooks)

	// Ejercicio 5

	inputList := []int{10, 4, 5, 8, 2, 9}
	filteredList := algorithms.FilterLt5(inputList)
	fmt.Printf("Original List: %v\n", inputList)
	fmt.Printf("Filtered List: %v\n", filteredList)
}
