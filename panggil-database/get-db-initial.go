package main

import (
	"fmt"
	"ticket-app/entity"
)

func main() {

	hasil := entity.GetDatabase()
	fmt.Println(hasil)
}
