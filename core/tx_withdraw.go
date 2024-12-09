package core

import "sample/utils"

type WithdrawTx struct {
	Amount    uint64
	Recipient utils.Address
}
