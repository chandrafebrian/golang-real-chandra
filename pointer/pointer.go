package main

import "fmt"

// keyword pointer adalah kemampuan membuat reference(tujuan) ke lokasi data di memory yang sama, tanpa duplikasi data yang sudah ada
// sederhana nya dengan pointer bisa membuat passing data by reference
// keyword pointer menggunakan operator "&" akan membuat memory baru dengan mengcopy data addrs1 dan merubah value nya sesuai data yg baru untuk addrs2
// address2.city , merubah nama kota nya dari addres1

type Adress struct {
	city, province, country string
}

func main() {

	address1 := Adress{"bandung", "jabar", "indon"}
	address2 := &address1
	// address2.city = "jaktim"

	// keyword bintang * akan memaksa value yg di refer(dituju) menggunakan keyword bintang *address2 = address("update1","up2","up3")
	// penggunaan keyword nya harus sepasang dengan * dan &
	*address2 = Adress{"solo", "jateng", "indon"}

	fmt.Println("add1 :", address1)
	fmt.Println("add2 :", address2)

}
