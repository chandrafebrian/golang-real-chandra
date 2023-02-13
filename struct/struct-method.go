package main

import (
	"fmt"
	"time"
)

// struct method adalah tipe data seperti tipe data lainnya , dia bisa di gunakan dalam parameter untuk function
// namun jika ingin menambahkan method(function) ke dalam structs itu bisa, sehingga seakan-akan sebuah struct memiliki function
// method sama artinya dengan function

// struct dengan nama Customer

// type Customer struct {
// 	Name string
// }

// nama var nya harus sama dengan nama struct methodnya (customer Customer)
// cara baca: func sayhello(typedata string) dengan paramater struct(Customer)

// func (customer Customer) sayHello(orang string) {
// 	fmt.Println("hello", orang, "my name is", customer.Name)
// }

func main() {
	// chandra := Customer{Name: "chandra best sre"}
	// chandra.sayHello("bangsat!")

	sat := AclTokenPreProduction()
	fmt.Println(sat)
}

// ACLtoken for new cluster pre-production nomad
type ACLToken struct {
	AccessorID     string
	SecretID       string
	Name           string
	Type           string
	Policies       []string
	Global         bool
	Hash           []byte
	CreateTime     time.Time
	ExpirationTime *time.Time
	ExpirationTTL  time.Duration
	CreateIndex    uint64
	ModifyIndex    uint64
}

func AclTokenPreProduction() *ACLToken {
	return &ACLToken{
		AccessorID: ("f38b0983-eecc-d840-ab15-660d0f30d673"),
		SecretID:   ("0279cba2-e070-bd11-68eb-b8d7ff60d816"),
		Name:       ("Token-Pre-Production"),
		Type:       ("Management"),
		CreateTime: time.Now(),
	}
}
