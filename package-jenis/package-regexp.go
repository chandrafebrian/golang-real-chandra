package main

import (
	"fmt"
	"regexp"
)

// package regexp untuk mencari string atau memvalidasi berdasarkan aturan regexnya
func main() {
	// cara baca nya
	// var regexp tolong tampung = value package regexp.Mustcompile(harus ada regexp dengan 1huruf depan e, 1huruf tengah antara a-z, 1huruf akhir o ) dengan parameter(`e([a-z])o`)
	regexp := regexp.MustCompile(`e([a-z])o`)
	// cari apakah ada dan sama string ("eko") menggunakan regexp.MatchString
	fmt.Println(regexp.MatchString("eko"))
	fmt.Println(regexp.MatchString("chandra febrian"))

	fmt.Println(regexp.FindAllString("eko eto edo epo", 4))
}
