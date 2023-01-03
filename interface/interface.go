package main

import "fmt"

// // interface adalah tipe data abstract,dia tidak memiliki implementasi langsung
// // interface adalah contract func , jika ingin menggunakan nya harus memakai fungsi dari contract interface tersebut Getname() tipe kembalian string
// // cara baca : membuat interface kontrak dengan nama hasname dan didalam nya punya method func dengan nama Getname() yg wajib di sebutkan saat di panggil

type Hasname interface {
	Getname() string
}

// // func Sayhello (dengan parameter var hasname dan type interface Hasname)
// // saat mendefind string wajib memanggil var type interface dan method nya
func SayHello(hasname Hasname) {
	fmt.Println("hello ", hasname.Getname())

}

// type struct sama dengan class Person di dart OOP
type Person struct {
	Name string
}

func (person Person) Getname() string {
	return person.Name
}

type Animal struct {
	Jenis string
}

func (hewan Animal) Getname() string {
	return hewan.Jenis
}

func main() {
	judul := Person{
		Name: "chandra",
	}
	SayHello(judul)

	cat := Animal{
		Jenis: "anggora",
	}
	SayHello(cat)

}
