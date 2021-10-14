package account

import "sync"

type Account struct {
	balance int64
	closed  bool
}

var mu sync.Mutex

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	return &Account{
		balance: initialDeposit,
		closed:  false,
	}
}

func (a Account) Balance() (balance int64, ok bool) {
	if a.closed {
		return 0, ok
	}
	return a.balance, true
}

func (a *Account) Close() (payout int64, ok bool) {
	mu.Lock()
	defer mu.Unlock()
	if !a.closed {
		a.closed = true
		payout = a.balance
		ok = true
	}
	return payout, ok
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	mu.Lock()
	defer mu.Unlock()
	if a.closed {
		return 0, ok
	}
	if (a.balance + amount) >= 0 {
		a.balance += amount
		ok = true
	}
	return a.balance, ok
}
