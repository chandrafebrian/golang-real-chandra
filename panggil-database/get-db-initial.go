package main

import (
	"fmt"

	"github.com/chandrafebrian/golang-real-chandra/entity"
)

func main() {

	hasil := entity.GetDatabase()
	fmt.Println(hasil)
}
