package main

import (
	"fmt"
	"testing"
	"time"
)

func HelloWorld() {
	fmt.Println("hallo world")
}

// membuat testiing dengan goroutine
func TestGoroutine(t *testing.T) {
	//go di depan function adalah keyword goroutine
	//saat menjalankan function dengan goroutine , maka fucntion tersebut akan otomatis mennjadi mode asyncronous(tidak di tunggu sampai selesai proses dari function tersebut) dan langsung menjalankan proses function lain di bawah nya
	go HelloWorld()
	fmt.Println("ups motherfucker")

	time.Sleep(1 * time.Second)
}

func Iterasi(number int) {

	fmt.Println("angka", number)
}

func TestPerulangan(t *testing.T) {
	for i := 0; i < 100; i++ {
		go Iterasi(i)
	}

	time.Sleep(6 * time.Second)
}
