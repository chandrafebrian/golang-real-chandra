package main

import (
	"fmt"
)

func main() {

	result := Random()
	// keyword switch dengan nama varPenampung := menampung value dari variable result.(type)
	// var result harus menggunakan .(type) karena dia menampung func Random() yang mereturn interface
	switch penampung := result.(type) {
	case string:
		fmt.Println("kalimat ", penampung, "ini adalah tipe string")
	case int:
		fmt.Println("angka ", penampung, "ini adalah tipe int")

	}
}

// jika return kalimata "" akan menampilkan string case , kalau return angka akan menampilkan int
func Random() interface{} {
	return 84782949

}
