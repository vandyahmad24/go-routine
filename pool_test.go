package go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)


func TestPool(t *testing.T) {
	pool := sync.Pool{
		New:func() interface{}{
			return "New"
		},
	}
	pool.Put("Ahmad")
	pool.Put("Misry")
	pool.Put("ar razy")
	for i:=1; i<10; i++{
		go func() {
			data := pool.Get()
			fmt.Println(data)
			time.Sleep(1*time.Second)
			pool.Put(data)
		}()
	}
	time.Sleep(11*time.Second)
	fmt.Println("selesai")
}