package go_routine

import (
	"fmt"
	"testing"
	"time"
)

func TestCreateChannel(t *testing.T)  {
	//untuk mengirim dan menerima data dari goroutine
	channel := make(chan string)
////	mengirim data
//	channel <-"Vandy"
//	//menerima data
//	data := <- channel
//	fmt.Println(data)
////	langsung ke paramter
//	fmt.Println(<- channel)
////	tutup channelnya
//	close(channel)

	go func() {
		time.Sleep(2*time.Second)
		channel <- "Vandy Ahmad"
		fmt.Println("selesai mengirim data ke channel")
	}()
	data := <- channel
	fmt.Println(data)
	time.Sleep(5 *time.Second)

}
func GiveMeResponse(channel chan string)  {
	time.Sleep(2*time.Second)
	channel <- "vandy ahmad"
}

func TestChannelAsParameter(t *testing.T){
	channel := make(chan string)

	go GiveMeResponse(channel)
	data := <- channel
	fmt.Println(data)
	time.Sleep(5 *time.Second)
	close(channel)
}

func OnlyIn(channel chan<-string){
	time.Sleep(2*time.Second)
	channel <- "vandy ahmad"
	//data := <-channel
}

func OnlyOut(channel <-chan string){
	data := <- channel
	fmt.Println("Dari only out",data)
}

func TestInOutChannel(t *testing.T){
	channel := make(chan string)

	defer close(channel)
	go OnlyIn(channel)
	go OnlyOut(channel)

	//close(channel)
	time.Sleep(2 *time.Second)

}

//buffered channel menampung data antrian di channel