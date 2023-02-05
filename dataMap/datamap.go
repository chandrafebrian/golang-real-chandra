package main

import "fmt"

// Map adalah tipe data lain yang berisikan kumpulan data yang sama,
// namun kita bisa menentukan jenis tipe data index yang akan kita gunakan.
// sederhananya ,map adalah tipe daata kumpulan key-value (kata kunci dan nilai),dimana kata kuncinya bersifat unik ,tidak boleh sama
// berbeda dengan array dan slice , jumlah data yang kita masukan ke dalam map boleh sebanyak banyak nya ,
// asalkan kata kunci nya berbeda , jika kita gunakan kata kunci yang sama ,
// maka secata otomatis data sebelumnya akan di ganti dengan data yg baru

func main() {
	// cara baca buat map di var person [keynya: string] , valuenya: string
	// key "name" , value "chandra"
	//

	// keyword map untuk mengambil data di map dengan key
	person := map[string]string{
		"name":    "chandra",
		"address": "subang",
	}
	// proses dibwah ini untuk menambahkan data ke dalam map person ,
	//perhatikan jika key nya sama kembar dengan key person yg lama ,
	//maka otomatis data key yg lama akan di hapus lalu di timpa replace oleh key yang baru
	person["title"] = "programmer"
	person["address"] = "jakarta"

	fmt.Println(person)
	fmt.Println(person["name"]) //ini proses memanggil map , dengan mengetikan key nya saja ,nanti yang muncul adalah value nya
	fmt.Println(person["address"])
	fmt.Println(person["title"])

	// ****
	//
	// make untuk membuat map baru
	mapBaru := make(map[string]string)
	mapBaru["title"] = "buku golang"
	mapBaru["author"] = "chandra"
	mapBaru["hapus"] = "salah"
	fmt.Println(mapBaru, "len sebelum ada key yang dihapus:", len(mapBaru))
	// delete(nama map yang akan di hapus , keymap:"hapus")
	delete(mapBaru, "hapus")
	fmt.Println(mapBaru, "len sudah hapus key:", len(mapBaru))

	fmt.Println(mapBaru["title"]) //ini cara memanggil data nya dengan spesifik key tertentu

}
