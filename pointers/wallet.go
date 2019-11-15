package main

type Wallet struct {
	balance int
}

func (w *Wallet) Deposit(val int) {
	w.balance += val
}

func (w *Wallet) Balance() int {
	return w.balance
}
