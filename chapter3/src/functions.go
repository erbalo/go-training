package src

import (
	"errors"
	"fmt"
	"sync"
)

type Balance struct {
	mu     sync.Mutex
	amount int
}

func (b *Balance) Balance() (int, error) {
	if b.amount == 0 {
		return 0, errors.New("you don't have any money")
	}
	return b.amount, nil
}

func (b *Balance) BalanceNamedNaked() (balance int, err error) {
	if b.amount == 0 {
		err = errors.New("you don't have any money")
	} else {
		balance = b.amount
	}
	return
}

func (b *Balance) Update(amount int) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.amount = amount
}

func main() {
}

func lifo() {
	defer fmt.Print(3)
	defer fmt.Print(2)
	defer fmt.Print(1)
}

func panicRecover() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered %v", r)
		}
	}()
	panic(1)
}
