package main

import (
	"fmt"
)

// func main() {

// 	for count := 1; count <= 10; count++ {
// 		fmt.Println("bilangan ke", count)

// 	}

// }
// tanda ; untuk pemisah code selanjutnya
// dalam func for agar code lebih clean dan simple bisa menggunakan 2 statement yaitu
// dengan init statement dan post statement
// yang menjadi init statement adalah " count := 1; " akan di jalankan sekali saat func for
// yang menjadi post statement adalah " count++ " akan dijalankan terus hingga kondisi code nya false
// yang menjadi kondisi code adalah " count <= 10; " sebagai batas function di jalankan

// **************************************//

// slice

func main() {

	// variable slicePotongan menampung nilai map array string dengan isi data { " string " }
	slicePotongan := []string{"chandra", "febrian", "programmer"}
	// ulangi fungsi for di variable 'i' dengan nilai deafault 0; jika i < lebih kecil dari 'len' lenght dengan parameter (slicePotongan); maka tambahkan di var i
	for i := 0; i < len(slicePotongan); i++ {
		fmt.Println(slicePotongan[i])

	}
	for value := range slicePotongan {
		fmt.Println(value)

	}

	// ini cara lain menggunakan keyword make untuk membuat map
	// variable person menampung fungsi make dengan parameter (map [key string]value string)
	// jika tidak ingin menggunakan var key ganti dengan ( _ ), karena tetap harus ada 2 karakter variable yg di define(jelaskan)
	// dan hasil nya hanya value yang di tampilkan
	person := make(map[string]string)
	person["alamat"] = "medan"
	person["name"] = "joko"
	person["position"] = "site realibilty engineer"

	for key, value := range person {
		fmt.Println(key, ":", value)
	}
}
