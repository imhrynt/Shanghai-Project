package core

type Account struct {
	Balance uint64
	Number  uint64
}

func (a *Account) ViewBalance() uint64 { return a.Balance }
func (a *Account) ViewNumber() uint64  { return a.Number }
