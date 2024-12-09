package core

import (
	"sample/utils"
	"time"
)

type Transaction struct {
	payload Payload
	time    uint64
	hash    *utils.Hash
	sender  *utils.Address
}

type Payload interface {
	number() uint64
	gas() uint64
	amount() uint64
	recipient() *utils.Address
}

// Transaction value accessors

func (tx *Transaction) GetTime() uint64           { return tx.time }
func (tx *Transaction) GetHash() utils.Hash       { return *tx.hash }
func (tx *Transaction) GetSender() *utils.Address { return tx.sender }

// Transaction function

func NewTransaction(payload Payload) *Transaction {
	return &Transaction{
		payload: payload,
		time:    uint64(time.Now().UnixNano()),
	}
}

type Transactions []*Transaction

// Transactions function

func (txs Transactions) Len() int { return len(txs) }

func (txs Transactions) Transaction(hash utils.Hash) *Transaction {
	for _, transaction := range txs {
		if transaction.GetHash() == hash {
			return transaction
		}
	}
	return nil
}
