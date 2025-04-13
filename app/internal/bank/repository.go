package bank

import (
	"gorm.io/gorm"
)

type BankAccountRepository interface {
	GetBalance() int
	Deposit(amount int)
	Withdraw(amount int) error
}

type bankAccountRepo struct {
	db *gorm.DB
}

func NewBankAccountRepository(db *gorm.DB) BankAccountRepository {
	return &bankAccountRepo{db}
}

func (r *bankAccountRepo) GetBalance() int {
	var balance int
	if err := r.db.Table("bank_accounts").Select("balance").Row().Scan(&balance); err != nil {
		return 0
	}
	return balance
}
func (r *bankAccountRepo) Deposit(amount int) {
	if err := r.db.Table("bank_accounts").Update("balance", gorm.Expr("balance + ?", amount)).Error; err != nil {
		return
	}
}
func (r *bankAccountRepo) Withdraw(amount int) error {
	if err := r.db.Table("bank_accounts").Update("balance", gorm.Expr("balance - ?", amount)).Error; err != nil {
		return err
	}
	return nil
}
