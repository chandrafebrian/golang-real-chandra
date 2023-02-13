package main

import "fmt"

// func dengan parameter  wajib memasukan value dari type data nya jika memanggil func tersebut
func FungsiParameter(firstName, middleName, lastName string) {
	fmt.Println("hallo bro", firstName, middleName, lastName)
}

func main() {
	FungsiParameter("chandra", "devops", "febrian")
	Testingrr("dsds")
}

func Testingrr(firstName string) {
	fmt.Println(firstName)
}
