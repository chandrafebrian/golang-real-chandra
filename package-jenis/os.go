package main

// // package os.Args = adalah untuk mengambil data argument
// func main() {
// 	argument := os.Args

// 	fmt.Println("args bro", argument)
// 	// jika memanggil denga argument slice seperti ini harus dengan command run nya
// 	// go run package-jenis/os.go chandra febrian devops
// 	fmt.Println(argument[1])
// 	// maka result nya
// 	// args bro [/var/folders/96/7p9f1mxd5ys80qg1svmrfvt40000gn/T/go-build1423212715/b001/exe/os chandra febrian devops]
// 	// chandra
// 	// argument os ke 0 , chandra ke 1 , febrian ke 2 , devops ke 3

// 	// -
// 	// -
// 	// ********************************************

// 	// os.Hostname untuk memanggil nama hostname device
// 	hostname, err := os.Hostname()
// 	// cara baca nya jika error nya samadengan nil : berarti yg akan di return code yg benar dan tidak error
// 	if err == nil {
// 		fmt.Println("hostname: ", hostname)
// 	} else {
// 		fmt.Println("Error", err.Error())
// 	}

// 	// **************************************
// 	// os.Getenv untuk mengambil data environment
// 	username := os.Getenv("APP_USERNAME")
// 	password := os.Getenv("APP_PASSWORD")

// 	fmt.Println(username)
// 	fmt.Println(password)

// 	// cara manggil nya di cli ketik ini dulu , baru bisa jalankan go run
// 	// export APP_USERNAME=chandra9201
// 	// export APP_PASSWORD=root
// }
