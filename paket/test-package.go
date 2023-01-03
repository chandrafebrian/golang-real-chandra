package paket

import (
	"fmt"
)

// setiap membuat nama func untuk public agar bisa di import keluaar harus pakai Uppercase name (huruf besar awalan nya)
// jika huruf kecil hanya bisa di pakai untuk private func di paket dan file yg sama
func SayHello(nama, belakang string) {

	fmt.Println("hallo", nama, belakang)
}
