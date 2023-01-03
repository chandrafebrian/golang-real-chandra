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
	fmt.Println(slice1)
	fmt.Println(len(slice1)) //len akan mengambil index data dari mei - juli = 3
	fmt.Println(cap(slice1)) // cap(capacity) akan mengambil index data dari mei - habis = 8

	var slice2 = bulan[10:] // slice2 memanggil slice mulai dari index 10 sampai akhir dari array bulan
	fmt.Println(slice2)

	var slice3 = append(slice2, "chandra", "febrian") //append untuk membuat slice baru di dalam slice dengan menambah(replace) data di posisi terakhir slice
	slice3[1] = "kosong"
	fmt.Println(slice3)

	// $$$$$$
	// cara lain dengan memanggil fungsi make untuk membuat slice
	// cara baca membuat slice arrta dengan memanggil fungsi make(paramter arraykosong[]type string. lengt 2, capacity 5,)
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

	// varkopislice = menampung slice type string dengan len=3 , cap=semua nya dari slice newSlice
	varkopislice := make([]string, 3, cap(newSlice))
	copy(varkopislice, newSlice)
	fmt.Println("hasilcopy:", varkopislice)

}
