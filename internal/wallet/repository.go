package wallet

type FakeWalletRepository struct {
	wallets []*Wallet
}

func NewWalletRepository() *FakeWalletRepository {
	return &FakeWalletRepository{}
}

func (f *FakeWalletRepository) SaveFakeWalet(w *Wallet) error {
	f.wallets = append(f.wallets, w)
	return nil
}

func (f *FakeWalletRepository) GetWallet(id string, userId string) (*Wallet, error) {
	for _, wallet := range f.wallets {
		if wallet.ID == id && wallet.UserId == userId {
			return wallet, nil
		}
	}
	return nil, nil
}
