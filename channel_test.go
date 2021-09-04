package go_routine

import (
	"fmt"
	"strconv"
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
func TestBufferedChannel(t *testing.T) {
	channel := make(chan string,2)
	defer close(channel)
	go func() {
		channel <- "vandy"
		channel <- "ahmad"
	}()
	go func() {
		fmt.Println(<-channel)
		fmt.Println(<-channel)
	}()
	time.Sleep(3*time.Second)
	fmt.Println("selesai")

}

//range channel
func TestRangeChannel(t *testing.T) {
	channel := make(chan string)
	go func() {
		for i:=0; i<100; i++{
			channel <- "Perulangan ke "+strconv.Itoa(i)
		}
		close(channel)
	}()

	for data := range channel{
		fmt.Println("Menerima Data",data)
	}
	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	counter:=0
	for{
		select {
		case data:=<-channel1:
			fmt.Println("Data dari channel 1",data)
			counter++
		case data:=<-channel2:
			fmt.Println("Data dari channel 2",data)
			counter++
		}
		if counter ==2{
			break
		}
	}

	//time.Sleep(2*time.Second)
}

func TestSDefaultelectChannel(t *testing.T) {
	channel1 := make(chan string)
	channel2 := make(chan string)
	defer close(channel1)
	defer close(channel2)
	go GiveMeResponse(channel1)
	go GiveMeResponse(channel2)
	counter:=0
	for{
		select {
		case data:=<-channel1:
			fmt.Println("Data dari channel 1",data)
			counter++
		case data:=<-channel2:
			fmt.Println("Data dari channel 2",data)
			counter++
		default:
			fmt.Println("Menunggu Data")
		}
		if counter ==2{
			break
		}
	}

	//time.Sleep(2*time.Second)
}

//Race Condition
