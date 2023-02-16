package main

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T) {

	varChannel := make(chan string)

	//go = artinya goroutine membuat func annonymous
	go func() {
		time.Sleep(2 * time.Second)     // jalankan perintah waiting sleep selama 2 detik
		varChannel <- "chandra febrian" //lalu mengirim isi data string 'chandra febrian' ke varChannel
		fmt.Println("Ping")
	}()

	data := <-varChannel //lalu isi data string dari varChannel di tampung oleh variable data
	fmt.Println(data)    //print hasil dari variable data
	time.Sleep(5 * time.Second)
	fmt.Println("Pong")

	defer close(varChannel)

}
