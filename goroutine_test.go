package go_routine

import (
	"fmt"
	"testing"
	"time"
)

func RunHelloWorld(){
	fmt.Println("Hello world")
}

func TestCreateGoroutine(t *testing.T){
	go RunHelloWorld()
	fmt.Println("Hello dari test")

	//	menunggu go routine
	time.Sleep(2 *time.Second)
}

func DisplayNumber(number int){
	fmt.Println("Display",number)
}

func TestManyGoroutine(t *testing.T){
	for i := 0; i<100000; i++{
		go DisplayNumber(i)
	}

	time.Sleep(5 * time.Second)
}