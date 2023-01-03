package folderTesting

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// contoh table unit testing dynamic subtesting
func TestTableHelloWorld(t *testing.T) {
	vSub := []struct {
		namaSub   string
		request   string
		ekspetasi string
		pesanErr  string
	}{
		{
			namaSub:   "subChandra",
			request:   "chandra febrian",
			ekspetasi: "halo chandra febrian",
			pesanErr:  "anjay SubChandra gk sesuai ekspetasi bro cek lagi !",
		},
		{
			namaSub:   "subBudi",
			request:   "budi",
			ekspetasi: "halo budi",
			pesanErr:  "anjay subBudi gak sesuai ekspetasi bro cek lagi !",
		},
	}

	for _, vTes := range vSub {
		t.Run(vTes.namaSub, func(t *testing.T) {
			vRequest := HelloWorld(vTes.request)
			require.Equal(t, vTes.ekspetasi, vRequest, vTes.pesanErr)

		})

	}
}

// --
//--
// --
// ---
// ---

// ini contoh unit testing yang biasa kurang lengkap error detail nya dr go
func TestHelloWorld(test *testing.T) {
	hasil := HelloWorld("chandra")
	if hasil != "halo chandra" {
		test.Error("anjay salah ini bro!")
	} else {
		println("nah ini baru bener bro chandra !")
	}

}

// contoh testing pakai package library high level external : testify assert
func TestHelloJoko(t *testing.T) {
	hasil := HelloWorld("joko")
	// cara baca (t=var untuk pointer ke *testing, 'halo joko'=ekspetasi yang di harapkan, hasil= value dari req func HelloWorld, "Error hasil nya harus halo eko bro !!"=message kalau terjadi eror)
	assert.Equal(t, "halo joko", hasil, "Error hasil nya harus halo eko bro !!")
	// fmt.Println("nah ini baru bener bro !")
}

// go test -v -run TestHanyaFungsiTertentu : untuk test nama fungsi yang spesifik
// go test -v : untuk test semua fungsi di folder yang sama
// go test -v ./... : untuk test fungsi di semua folder tanpa terkecuali

// package high level assert go= go get github.com/stretchr/testify
