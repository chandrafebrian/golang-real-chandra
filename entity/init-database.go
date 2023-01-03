package entity

// func dengan nama init akan di execute atau dijalankan duluan jika di panggi panggil dari folder lain
// init cocok untuk membuat koneksi ke database
var connection string

func init() {
	connection = "MySQL"
}

func GetDatabase() string {
	return connection
}
