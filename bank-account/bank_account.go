package account

import "sync"

type Account struct {
	currentBalance int64
	isClosed       bool
	mutex          *sync.Mutex
}

func Open(initialDeposit int64) *Account {
	if initialDeposit < 0 {
		return nil
	}

	return &Account{currentBalance: initialDeposit, mutex: &sync.Mutex{}}
}

func (a *Account) Close() (payout int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.isClosed {
		return 0, false
	}

	a.isClosed = true

	return a.currentBalance, true
}

func (a *Account) Balance() (balance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.isClosed {
		return 0, false
	}

	return a.currentBalance, true
}

func (a *Account) Deposit(amount int64) (newBalance int64, ok bool) {
	a.mutex.Lock()
	defer a.mutex.Unlock()
	if a.isClosed {
		return 0, false
	}

	if amount > 0 {
		a.currentBalance += amount
	}

	if amount < 0 {
		afterWithdrawalBalance := a.currentBalance + amount
		if afterWithdrawalBalance < 0 {
			return a.currentBalance, false
		}

		a.currentBalance = afterWithdrawalBalance
	}

	return a.currentBalance, true
}
