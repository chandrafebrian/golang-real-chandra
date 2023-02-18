package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// constraint adalah typedata
// constraint nya adalah [any]
// keyword generic bisa di masukan untuk tipe data apapun menggunakan [T any] (parameter T) return T {}
func Lenght[T any](param T) T {
	fmt.Println(param)
	return param
}

func TestSampleGeneric(t *testing.T) {

	var kalimat string = Lenght("chandra febrian")
	assert.Equal(t, "chandra febrian", kalimat, "pesan err: hasil tidak sama bro !") // saat memanggil func assert.equal , harus memasukan paramater yg sesuai dengan perintah nya

	var angka int = Lenght(100)
	assert.Equal(t, 100, angka, "pesan err: hasil gk sama bro!")

}
