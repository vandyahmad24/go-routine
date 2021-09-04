package go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func RunAsyn(group *sync.WaitGroup, i int){
	defer group.Done()
	group.Add(1)
	fmt.Println("Run Asynnya",i)
	time.Sleep(1*time.Second)
}

//ngerun asyn
func TestWaitGroup(t *testing.T){
	group := &sync.WaitGroup{}
	for i :=0; i<100; i++{
		go RunAsyn(group, i)
	}
	group.Wait()
	fmt.Println("Selesai")
}