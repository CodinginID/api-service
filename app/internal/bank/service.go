package bank

import "fmt"

type BankAccountService interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

type bankAccountService struct {
	repo BankAccountRepository
}

func NewBankAccountService(r BankAccountRepository) BankAccountService {
	return &bankAccountService{r}
}

var account = &BankAccount{Balance: 1000}

func (s *bankAccountService) GetBalance() int {
	account.Mu.Lock()
	defer account.Mu.Unlock()
	return account.Balance
}

func (s *bankAccountService) Deposit(amount int) {
	account.Mu.Lock()
	defer account.Mu.Unlock()
	account.Balance += amount
}

func (s *bankAccountService) Withdraw(amount int) error {
	account.Mu.Lock()
	defer account.Mu.Unlock()

	if account.Balance < amount {
		return fmt.Errorf("saldo tidak mencukupi")
	}
	account.Balance -= amount
	return nil
}

// BankService is an interface for bank service
type BankService interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}
