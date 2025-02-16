package wallet

import (
	"errors"
	"net/http"

	"github.com/google/uuid"
)

type Wallet struct {
	ID      string  `json: "id"`
	Balance float64 `json: "balance"`
	Name    string  `json: "name"`
	Iban    string  `json: "iban"`
	UserId  string  `json: "userId"`
}

func CreateWallet(userId string, name string) *Wallet {
	return &Wallet{
		ID:      uuid.NewString(),
		Balance: 0.0,
		Name:    name,
		Iban:    uuid.NewString(),
		UserId:  userId,
	}
}

func (w *Wallet) Render(_ http.ResponseWriter, _ *http.Request) error {
	return nil
}

func (w *Wallet) DepositMoney(depositValue float64) *Wallet {
	w.Balance = w.Balance + depositValue
	return w

}

func (w *Wallet) WithdrawMoney(withdrawValue float64) error {
	if w.Balance < withdrawValue {
		return errors.New("not enough money!")

	}
	w.Balance = w.Balance - withdrawValue
	return nil

}
