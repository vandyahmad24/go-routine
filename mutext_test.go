package go_routine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestMutext(t *testing.T) {
	x :=0
	var mutext sync.Mutex
	for i:=1; i<=1000; i++{
		go func() {
			for j:=1; j<=100; j++{
				mutext.Lock()
				x =  x+1
				//fmt.Println(x)
				mutext.Unlock()
			}
		}()
	}

	time.Sleep(5*time.Second)
	fmt.Println("Counter",x)
}
//RWMutext
type BankAccount struct {
	RWMutex sync.RWMutex
	Balance int
}


func (account *BankAccount) AddBalance(amount int){
	account.RWMutex.Lock()
	account.Balance=account.Balance+amount
	account.RWMutex.Unlock()
}

func (account *BankAccount) GetBalance() int {
		account.RWMutex.RLock()
		balance := account.Balance
		account.RWMutex.RUnlock()
		return balance
}

func TestReadWriteMutext(t *testing.T) {
	account := BankAccount{}
	for i := 0; i<100; i++{
		go func() {
			for j :=0;  j<100; j++{
				account.AddBalance(1)
				fmt.Println(account.GetBalance())
			}
		}()
	}
	time.Sleep(5*time.Second)
	fmt.Println("Total balance",account.GetBalance())
}

type UserBalance struct{
	sync.Mutex
	Name string
	Balance int
}

func (user *UserBalance)Lock()  {
	user.Mutex.Lock()
}

func (user *UserBalance)Unlock(){
	user.Mutex.Unlock()
}

func (user *UserBalance)Change(amount int){
	user.Balance= user.Balance+amount
}

func Transfer(user1 *UserBalance, user2 *UserBalance, amount int){
	user1.Lock()
	fmt.Println("Lock 1", user1.Name)
	user1.Change(-amount)
	time.Sleep(1*time.Second)

	user2.Lock()
	fmt.Println("Lock 2", user2.Name)
	user2.Change(amount)
	time.Sleep(1*time.Second)

	user1.Unlock()
	user2.Unlock()

}

func TestDeadlock(t *testing.T){
	user1:= UserBalance{
		Name: "vandy",
		Balance: 1000,
	}
	user2:= UserBalance{
		Name: "ahmad",
		Balance: 1000,
	}

	go Transfer(&user1, &user2, 1000)
	go Transfer(&user2, &user1, 2000)
	time.Sleep(3*time.Second)
	fmt.Println("User 1 ",user1.Name, "Balance",user1.Balance)
	fmt.Println("User 2 ",user2.Name, "Balance",user2.Balance)

}