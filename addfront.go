package main

import "fmt"

func AddFront(s string, slice []string) []string {
	slice = Append(slice, s)
	// runes := []rune{}
	// l := len(slice) + 1
	// runes := []rune{}
	l := len(slice)
	// for i, value := range s {
	// for i, value := range slice {
	// }

	fmt.Println(slice[l-1])
	// for i := 0; i < 1; i++ {
	slice[l-1], slice[0] = slice[0], slice[l-1]
	// slice[i], slice[l-1-i] = slice[l-1-i], slice[i]
	// return slice
	// }
	// }

	// newSlice := []string{}
	// for i, value := range slice {
	// 	// runes = []rune(value)
	// 	newSlice[i] = string(value)
	// }

	// return slice
	// for i := 0; i < len(runes)/2; i++ {
	// 	runes[i], runes[l-1-i] = runes[l-1-i], runes[i]
	// }
	// fmt.Println(runes)
	// for i, v := range runes {
	// 	newSlice[i] = string(v)
	// }
	return slice
}

func main() {
	fmt.Println(AddFront("Hello", []string{"world"}))
	fmt.Println(AddFront("Hello", []string{"world", "!"}))
	fmt.Println(AddFront("Hello", []string{}))
}
