package main

import "fmt"

// constan adalah tipe data yg tidak bisa di rubah jika sudah di deklarasi
func main() {
	const (
		nama     = "chandra"
		belakang = "febrian"
		alamat   = "medan"
	)

	fmt.Println(nama, belakang, alamat)
}

//noted :
// - jika nama variable di masukan type data nya wajib menggunakan var dan = , contoh= var nilai32 int32 = 100000
// - jika nama varrible tidak dimasukan type data nya wajib menggunakan := tanpa perlu var lagi di belakang nama variable , contoh= nilaiBaru := 5000
