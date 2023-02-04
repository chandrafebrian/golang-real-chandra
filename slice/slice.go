package main

import (
	"fmt"
)

// tipe data slice adalah potongan data dari array
// slice mirip dengan array, yang membedakan adalah ukuran slice bisa berubah misal tadi nya 10 index bisa di ubah menjadi 20 sesuai kebutuhan
// slice dan array selalu terkoneksi , dimana slice adalah data yang mengakses sebagian atau seluruh data di  array
// tipe detail data slice ada 3 = yaitu pointer , length , capacity
// pointer = adalah penunjuk data pertama di array pada slice
// len= adalah panjang dari index slice
// cap = adalah kapasitas index dari slice dimana length tidak boleh lebih panjang dari capacity

func main() {

	// tanda titik 3x dibaca ,jumlah penampung index data array bisa kita set tak terbatas sampai berapa banyak bisa menampung data array , tipe data array nya string di contoh ini
	var bulan = [...]string{ // ini adalah array
		"januari",
		"feb",
		"maret",
		"april",
		"mei",
		"juni",
		"juli",
		"agustus",
		"september",
		"oktober",
		"november",
		"desember",
	}
	var slice1 = bulan[4:7]
	fmt.Println(slice1)      //menampilkan nama bulan dari mei -juli
	fmt.Println(len(slice1)) //len akan mengambil index data dari mei - juli = 3
	fmt.Println(cap(slice1)) // cap(capacity) akan menapilkan jumlah index data dari mei - habis = 8

	var slice2 = bulan[10:]
	fmt.Println(slice2) // slice2 memanggil slice mulai dari index 10 sampai akhir dari array bulan

	var slice3 = append(slice2, "chandra", "febrian")
	fmt.Println(slice3) //append untuk membuat slice baru di dalam slice dengan menambah(replace) data di posisi terakhir slice

	// $$$$$$
	// cara lain dengan memanggil fungsi make untuk membuat slice
	// cara baca= buat slice array dengan memanggil fungsi make(paramter arraykosong[]typedata string. lengt 10, capacity 15,)
	newSlice := make([]string, 10, 15) // ini adalah slice , len=10 , cap=15
	newSlice[0] = "hello"
	newSlice[1] = "world"
	newSlice[2] = "chandra"
	newSlice[3] = "febrian"
	fmt.Println("result dari var newslice:", newSlice)

	// $$$$$$$
	// jangan salah membuat array atau slice
	// array = ciri" kalau ingin membuat array  harus ada titik3 [...] atau [5] [harus ada nilai nya]
	// slice = ciri" [] kalau ingin membuat slice cukup kotak[] kosong saja
	iniArray := [5]int{1, 2, 4, 6, 3} //array
	iniSlice := []int{1, 2, 4, 6, 3}  //slice
	fmt.Println("ini array:", iniArray)
	fmt.Println("ini slice:", iniSlice)

	//
	// $$$$
	// ini cara copy slice to slice
	copySlice := make([]int, len(iniSlice), cap(iniSlice))
	copy(copySlice, iniSlice)
	fmt.Println("copySlice:", copySlice)

	// varkopislice = menampung slice type string dengan len=3 , cap=semua nya dari newSlice
	varkopislice := make([]string, 3, cap(newSlice))
	copy(varkopislice, newSlice) //jalankan func copy ke (dst.tujuan varkopislice , dari asalnya newslice)
	fmt.Println("hasilcopy:", varkopislice)

}
