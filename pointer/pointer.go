package main

import "fmt"

// keyword pointer adalah kemampuan membuat reference(tujuan) ke lokasi data di memory yang sama, tanpa duplikasi data yang sudah ada
// sederhana nya dengan pointer bisa membuat passing data by reference
// keyword pointer menggunakan operator "&" akan membuat memory baru dengan mengcopy data addrs1 dan merubah value nya sesuai data yg baru untuk addrs2
// address2.city , merubah nama kota nya dari addres1

// keyword * = sebagai penunjuk arah ke mana akan di rubah
// keyword & = sebagai wadah tujuan untuk perubahan yang baru

type Adress struct {
	city, province, country string
}

type AlamatBaru struct {
	jalan, kota string
}

func main() {
	//noted : penggunaan nama var dengan * harus sama seperti nam var yg akan di rubah
	aspal := AlamatBaru{"sm.raja", "medan"}
	aspalBaru := &aspal
	*aspalBaru = AlamatBaru{"klender", "jakarta"}

	fmt.Println("dari aspal :", aspalBaru)

	//-------------------------------------------------------
	address1 := Adress{"bandung", "jabar", "indon"}
	address2 := &address1
	// address2.city = "jaktim"

	// keyword bintang *address2 akan memaksa merubah var value yg di refer(dituju) ke variable address2 menjadi Address("solo","jateng","wakanda") penggunaan keyword nya harus sepasang dengan * dan &
	*address2 = Adress{"solo", "jateng", "wakanda"}

	fmt.Println("add1 :", address1)
	fmt.Println("add2 :", address2)

}
