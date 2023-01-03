package main

import "fmt"

// nilai default bool di golang adalah false , jika di masukan value true maka akan keluar code error panic
func runApp(error bool) {
	defer endApp()
	if error {
		panic("anjay Aplikasi Error")

	} else {
		fmt.Println("aplikasi berjalan")

	}

}

func endApp() {
	fmt.Println("aplikasi selesai")
}

func main() {
	runApp(false)

}
