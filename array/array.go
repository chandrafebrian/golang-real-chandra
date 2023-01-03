package main

import "fmt"

// array adalah kumpulan tipe data yang sama contoh tipedata (string) atau int dll.
// saat membuat array . kita perlu mementukan jumlah data yang bisa di tampung oleh array tersebut
// daya tampung array tidak bisa bertambah setelah array dibuat
// len untuk mengecek panjang penampung array , walaupun value nya masih kosong jika penampung data array nya 10 , maka saat di print len , hasil nya =10
func main() {

	var name [3]string

	name[0] = "chandra"
	name[1] = "febrian"
	name[2] = "devops engineer"

	// array [...] untuk tidak membatasi jumlah data array nya
	var nilai = [...]int{
		10, 23, 320, 32, 121, 5345, 545, 45, 54, 54, 54, 54, 54, 5,
	}

	fmt.Println(name[1])
	fmt.Println(nilai[9])
	fmt.Println("panjang len pada array nilai: ", len(nilai))

	a := [5]int{
		3, 3, 1,
	}

	fmt.Println(len(a))

}
