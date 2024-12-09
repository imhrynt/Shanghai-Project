package core

import "sample/utils"

type DynamicTx struct {
	Number    uint64
	Fee       uint64
	Amount    uint64
	Recipient *utils.Address
	V, R, S   []byte
}

func (tx *DynamicTx) number() uint64            { return tx.Number }
func (tx *DynamicTx) gas() uint64               { return tx.Fee }
func (tx *DynamicTx) amount() uint64            { return tx.Amount }
func (tx *DynamicTx) recipient() *utils.Address { return tx.Recipient }
