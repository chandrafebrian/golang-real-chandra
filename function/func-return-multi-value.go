package main

import "fmt"

// func dengan multi lebih dari satu return tipe data dan value
func multiDataReturn() (string, string) {

	return "chandra", "febrian"

}

// keyword underscore adalah untuk ingnore ( _ ) value agar tidak di panggil
func main() {

	firstname, lastname := multiDataReturn()
	fmt.Println(firstname, lastname)

}
