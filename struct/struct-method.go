package main

import "fmt"

// struct method adalah tipe data seperti tipe data lainnya , dia bisa di gunakan dalam parameter untuk function
// namun jika ingin menambahkan method(function) ke dalam structs itu bisa, sehingga seakan-akan sebuah struct memiliki function
// method sama artinya dengan function

// struct dengan nama Customer
type Customer struct {
	Name string
}

// nama var nya harus sama dengan nama struct methodnya (customer Customer)
// cara baca: func sayhello() dengan paramater struct methode/function (customer Customer)
func (customer Customer) sayHello(orang string) {
	fmt.Println("hello", orang, "my name is", customer.Name)
}

func (a Customer) FungsiHello() {
	fmt.Println("halo bro ", a.Name)
}

func main() {
	chandra := Customer{Name: "chandra-best-sre"}
	chandra.sayHello("bangsat!")

	febrian := Customer{Name: "febrian-devops"}
	febrian.FungsiHello()

}
