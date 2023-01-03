package main

// import (
// 	"flag"
// 	"fmt"
// )

// // package flag untuk parse(menguraikan) dan   membuat argument dengan benar , jadi ada nilai default , helper messege jika error

// func main() {

// 	// var host = menampung func flag.string( "var config nya " , "nilai default jika tidak di isi", "command helper pesan jika terjadi error" )
// 	host := flag.String("host", "localhost", "put your database host")
// 	username := flag.String("username", "root", "put your database username")
// 	password := flag.String("password", "root", "put your database password")

// 	// note: untuk menggunakan package flag harus panggil flag.Parse() kalau tidak akan error
// 	flag.Parse()

// 	fmt.Println("host :", *host, "username :", *username, "password :", *password)
// 	// cara memanggil nya
// 	// go run package-jenis/flag.go -host=chandraLocal -username=userfebrian -password=root
// }
