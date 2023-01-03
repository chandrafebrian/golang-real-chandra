package main

import "fmt"

func main() {
	// break untuk menghentikan seluruh perulangan jika nilai sudah mencapai batas yg di tentukan oleh if
	// keyword for dengan variable i menampung nilai default 0; jika nilai dari variable i lebih kecil < dari 10; maka nilai dari var i di tambah 1
	// dan jika nilai dari var i sama dengan bernilai 5 maka break (hentikan code)
	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}

		fmt.Println("ini angka break ", i)
	}

	// keyword continue akan menghentikan perulangan saat ini dan akan di lanjutkan lagi dengan post statement n++ dan akan di lanjutkan ke perulangan berikutnya
	// keyword n menampung nilai default 0; jika n lebih kecil < dari 10; maka tambah 1 di value pada var n
	// dan jika nilai dari var n dibagi % 2 hasilnya sama dengan 0(ganjil) maka jalankan keyword continue dan munculkan nilai ganjil saja
	for n := 0; n < 10; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println("ini angka continue ", n)
	}
}
