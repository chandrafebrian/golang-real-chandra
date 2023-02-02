package main

import "fmt"

// keyword struct adalah template data sama dengan class kalau di dart OOP ,
// dengan menuliskan type NamaStructnya struct
type Customer struct {
	Nama, Alamat string
	Age          int
}

type Human struct {
	orang string
	umur  int
}

// cara memanggil struct dengan membuat variable baru lalu panggil nama Struct nya
func main() {
	chandra := Customer{
		Age:    30,
		Alamat: "medan",
		Nama:   "chandra febrian",
	}

	manusia := Human{
		orang: "chandra nih bos",
		umur:  33,
	}

	fmt.Println(chandra)
	fmt.Println(chandra.Nama)
	fmt.Println(chandra.Alamat)
	fmt.Println(chandra.Age)
	fmt.Println("halo : ", manusia)
	fmt.Println(manusia.orang)
	fmt.Println(manusia.umur)
}
