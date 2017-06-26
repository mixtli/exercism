package account

import "sync"

const testVersion = 1

type Account struct {
	balance int64
	closed  bool
	sync.RWMutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}
	account := Account{balance: initialDeposit, closed: false}
	return &account
}

func (a *Account) Balance() (int64, bool) {
	if a.closed {
		return 0, false
	}
	return a.balance, true
}

func (a *Account) Close() (int64, bool) {
	a.Lock()
	if a.closed {
		a.Unlock()
		return 0, false
	}
	a.closed = true
	a.Unlock()
	return a.balance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.Lock()
	if a.closed {
		a.Unlock()
		return 0, false
	}
	if a.balance+amount < 0 {
		a.Unlock()
		return a.balance, false
	}
	a.balance += amount
	a.Unlock()
	return a.balance, true
}
