package main

import (
	"fmt"
)

func main() {

	// var booking = "halo bangkeee"

	// fmt.Printf("%v bgst apa kabar lo \n", booking)

	data := NewData("ini adalah isi dari newdata")
	if data != nil {
		fmt.Println("ada datanya bro")
		fmt.Println(data)
	} else {
		fmt.Println("data kosong")
	}

}

// kalau di baca kita membuat sebuah func dengan nama
// newdata(parameter string dengan variable name) me return map[key nya string]value string {}
func NewData(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}
	}
}
