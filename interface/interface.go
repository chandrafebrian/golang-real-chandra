package main

// import "fmt"

// // // interface adalah tipe data abstract,dia tidak memiliki implementasi langsung
// // // interface adalah contract func , jika ingin menggunakan nya harus memakai fungsi dari contract interface tersebut Getname() tipe kembalian string
// // // cara baca : membuat interface kontrak dengan nama Hasname dan didalam nya ada method func dengan nama Getname() yg wajib di sebutkan saat memanggil interface name

// type Hasname interface {
// 	Getname() string
// }

// // // func Sayhello (dengan parameter var hasname dan type interface Hasname)
// // // saat mendefind(menjelaskan) string wajib memanggil var type interface dan method(function) nya
// func SayHello(hasname Hasname) {
// 	fmt.Println("hello ", hasname.Getname())

// }

// // --------------------------------------------------------

// type Engineer struct {
// 	Develepor string
// }

// func (engineer Engineer) Getname() string {
// 	return engineer.Develepor
// }

// // type struct sama dengan class Person di dart OOP
// type Person struct {
// 	Name string
// }

// func (person Person) Getname() string {
// 	return person.Name
// }

// // -------------------------------------------------------

// type Animal struct {
// 	Jenis string
// }

// func (hewan Animal) Getname() string {
// 	return hewan.Jenis
// }

// // -------------------------------------------------------

// func main() {
// 	judul := Person{
// 		Name: "chandra",
// 	}
// 	SayHello(judul)

// 	cat := Animal{
// 		Jenis: "anggora",
// 	}
// 	SayHello(cat)

// 	orangDev := Engineer{
// 		Develepor: "chandra best of the best",
// 	}
// 	SayHello(orangDev)
// }
