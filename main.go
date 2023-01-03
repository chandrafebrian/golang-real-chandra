package main

// func main() {
// 	// contoh createe nama variable setelah = adalah value yg di simpan
// 	var nameConference = "go conference "
// 	const conferenceTicket = 50
// 	var remainingTickets uint = 30
// 	bookings := []string{}

// 	// ini adalah contoh memanggil fungsi dengan parameter variable string
// 	functionUser(nameConference)

// 	// printf : print mode format (%v artinya:  the value in a default format)
// 	fmt.Printf("we have total of %v ticket and remainingTickets %v are still available \n", conferenceTicket, remainingTickets)

// 	//for adalah looping pengulangan dan menjalankan method apapun di dalam kalibreket {} tanda kurung kurawal
// 	for {
// 		// jika menampung fungsi yg tidak menjelaskan "nama variable nya" , cukup hanya memanggil nama fungsinya saja  fungsiGetUserInput()
// 		username, lastname, email, userTicket := fungsiGetUserInput()
// 		// jika menampung fungsi yg sudah menjelaskan nama variable sebagai parameter , saat memanggil nama fungsi tersebut wajib memanggil juga paramater var nya
// 		isValidateUser, isValidateEmail, isValidTicket := helper.FungsiValidasi(username, lastname, email, userTicket, remainingTickets)

// 		// stetment if ini untuk cek jika tiket yg di input lebih dari available ,
// 		// lalu keyword continue akan melanjutkan perintah program dari atas lagi
// 		if isValidateUser && isValidateEmail && isValidTicket {

// 			fungsiBookingTiket(remainingTickets, userTicket, bookings, username, lastname, email)
// 			sendTicket(userTicket, username, lastname, email)
// 			// contoh memanggil fungsi dengan nama variable parameter nya
// 			fungsicontoh2(bookings)

// 			// untuk menghentikan program jika tiket sudah sold out dengan stetment if dan break
// 			if remainingTickets == 0 {
// 				fmt.Println("our conference is sold out")
// 				break

// 			}

// 			// jika tidak memenuhi syarat program maka else akan dijalankan false
// 		} else {
// 			if !isValidateUser {
// 				fmt.Println("your name is error")
// 			}
// 			if !isValidateEmail {
// 				fmt.Println("your email is error")
// 			}
// 			if !isValidTicket {
// 				fmt.Println("your number ticket is error")
// 			}
// 			fmt.Println("please try again")

// 		}

// 	}
// }

// // contoh membuat fungsi akses private karena huruf kecil penamaan depan nya dan dengan parameter var string, masukan nama variable type string
// func functionUser(varParameter string) {
// 	fmt.Printf("welcome to %v booking application ", varParameter)

// }

// // fungsi akses private dengan nama huruf kecil di depan
// func fungsicontoh2(bookings []string) {
// 	// proses slice membelah array string dengan append , hanya mengambil nama username.
// 	// namapertama adalah variable yg menampung value
// 	// lalu di looping dengan for  , dan menampung value dari range
// 	namapertama := []string{}
// 	// varbooking menyimpan value dari var bookings
// 	for _, varbooking := range bookings {
// 		// var names menyimpan nilai fungsi strings.fields dari varbooking
// 		var names = strings.Fields(varbooking)
// 		// var namapertama menyimpan nilai dan slice(membelah) array dengan fungsi append , dari var names[0]
// 		namapertama = append(namapertama, names[0])
// 	}
// 	// var namapertama di panggil dengan printf format , agar memunculkan slice nama pertamanya saja
// 	fmt.Printf("the first name of bookings are : %v\n ", namapertama)

// }

// // fungsi dengan double parameter () (string, string, string, uint) , artinya untuk return value , yg sudah di tampung variable type tersebut
// // me return value dari variable , username , lastname, email, userTicket
// func fungsiGetUserInput() (string, string, string, uint) {
// 	var username string
// 	var lastname string
// 	var email string
// 	var userTicket uint = 0

// 	// tanda & adalah untuk mengambil value dan di store ke dalam variable username
// 	println("Enter your Name :")
// 	fmt.Scan(&username)

// 	println("Enter your Last Name :")
// 	fmt.Scan(&lastname)

// 	println("Enter your Email :")
// 	fmt.Scan(&email)

// 	println("Enter number of order ticket :")
// 	fmt.Scan(&userTicket)

// 	return username, lastname, email, userTicket

// }

// // fungsi dengan parameter variable yg sama nama dan type var nya , agar bisa di pakai value nya saat fungsi di panggil
// func fungsiBookingTiket(remainingTickets uint, userTicket uint, bookings []string, username string, lastname string, email string) {
// 	// remainingticket dikurang userticket
// 	remainingTickets = remainingTickets - userTicket

// 	var userData = make(map[string]string)
// 	// membuat key dan value
// 	userData["username"] = username
// 	userData["lastname"] = lastname
// 	userData["email"] = email
// 	userData["userTicket"] = strconv.FormatUint(uint64(userTicket), 10)
// 	// append adalah fungsi untuk menanmbahkan value dari variable
// 	bookings = append(bookings, username+" "+lastname)

// 	fmt.Printf("Congratulations! %v %v with email: %v ,has booked %v tikets. \n", username, lastname, email, userTicket)
// 	fmt.Printf("Available ticket right now is %v tickets \n", remainingTickets)

// }

// func sendTicket(userTicket uint, username string, lastname string, email string) {
// 	var ticket = fmt.Sprintf("%v tickets for %v %v", userTicket, username, lastname)
// 	fmt.Printf("############################ ")
// 	fmt.Printf("Sending ticket: \n %v \n to email address %v\n", ticket, email)
// 	fmt.Printf("############################ ")

// }

//
//
// keyword go :
// break     		default		 func    	interface  	select
// case      		defer        go      	map        	struct
// chan      		else         goto    	package    	switch
// const     		fallthrough  if      	range      	type
// continue 		for          import  	return     	var

// contoh testing sum perkalian
func Sum(x int, y int) int {
	return x * y
}
