package main

import "fmt"

func main() {

	// ini proses mengkonversi tipe data dari int64 ke int32
	var nilai32 int32 = 500000000
	// int64 mengambil nilai dari int32 di konversi
	var nilai64 int64 = int64(nilai32)

	fmt.Println(nilai32)
	fmt.Println(nilai64)

}

//noted :
// - jika nama variable di masukan type data nya wajib menggunakan var dan = , contoh= var nilai32 int32 = 100000
// - jika nama varrible tidak dimasukan type data nya wajib menggunakan := tanpa perlu var lagi di belakang nama variable , contoh= nilaiBaru := 5000
