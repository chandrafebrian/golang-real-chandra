package main

import "fmt"

// interface kosong adalah interface yang tidak memiliki deklarasi method function, hal ini
// membuat semua tipe data yang di masukan akan otomatis menjadi implementasinya

func Ups(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Ups Bro tidak masuk dalam list int"
	}
}

func main() {

	var test interface{} = Ups(1)
	fmt.Println(test)

}
