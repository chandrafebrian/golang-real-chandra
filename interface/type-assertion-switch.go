package main

// // type assertion merupakan kemampuan merubah tipe data menjadi tipe data yang di inginkan
// // fitur ini sering sekali di gunakan ketika kita bertemu dengan data interface{} kosong
// // membuat fungsi dengan nama random() paramater kosong dan type data kembalian interface{} kosong {mereturn tipe data int}
// func random() interface{} {
// 	return 10
// }

// // -
// // -
// func main() {
// 	// variable result dengan tipedata interface{} = menampung fungsi random()
// 	// var hasilstring harus sama tipe data nya dengan func random di atas , jika angka berati 'hasilnya int' ,jika kalimat berarti 'hasilnya string'
// 	var result interface{} = random()
// 	var hasilnya int = result.(int)
// 	fmt.Println(hasilnya)

// 	// -
// 	// -
// 	// keyword switch dengan nama varPenampung := menampung value dari variable result.(type)
// 	switch varPenampung := result.(type) {

// 	case string:
// 		fmt.Println("kalimat ", varPenampung, "ini adalah tipe string")
// 	case int:
// 		fmt.Println("angka ", varPenampung, "ini adalah tipe int")

// 	}

// }
