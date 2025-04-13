package bank

import "sync"

type BankAccount struct {
	Balance int
	Mu      sync.Mutex
}
