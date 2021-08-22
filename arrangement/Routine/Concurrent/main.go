// 동시성 프로그래밍

//

package main

import (
	"fmt"
	"sync"
	"time"
)

type Account struct {
	Balance int
}

func main() {
	var wg sync.WaitGroup

	account := &Account{10}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		// 10번을 돌릴때 까지 무한히 돌린다.
		go func() {
			for {
				// 1000원씩 입, 출금 함수
				DepositAndWithdraw(account)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}

func DepositAndWithdraw(account *Account) {
	if account.Balance < 0 {
		// acount.Balance (잔고)는 음수가 될 수 없다.
		panic(fmt.Sprintf("Balance should not be negative value : %d", account.Balance))
	}
	account.Balance += 1000
	time.Sleep(time.Millisecond) // 0.001초
	account.Balance -= 1000
}
