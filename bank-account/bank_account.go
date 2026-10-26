package bankaccount

import "sync"

type Account struct {
	amount int64
	isOpen bool
	mutex  sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{
		amount: initialDeposit,
		isOpen: true,
		mutex:  sync.Mutex{},
	}
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.isOpen {
		return 0, false
	}

	return a.amount, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.isOpen {
		return 0, false
	}

	if a.amount+amount < 0 {
		return 0, false
	}

	a.amount += amount

	return a.amount, true
}

func (a *Account) Withdraw(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.isOpen {
		return 0, false
	}

	if amount < 0 {
		return 0, false
	}

	if a.amount-amount < 0 {
		return 0, false
	}

	a.amount -= amount

	return a.amount, true
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()

	if !a.isOpen {
		return 0, false
	}

	a.isOpen = false

	return a.amount, true
}
