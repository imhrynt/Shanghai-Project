package core

import (
	"io"
	"sample/utils"
)

type Header struct {
	Time      uint64
	Ancestors utils.Hash
	Root      utils.Hash
	State     utils.Hash
	Author    utils.Address
}

// Header function

func NewHeader(time uint64, ancestors, root, state utils.Hash, author utils.Address) *Header {
	return &Header{
		Time:      time,
		Ancestors: ancestors,
		Root:      root,
		State:     state,
		Author:    author,
	}
}

func (h *Header) Size() int {
	return 8 /* Time */ + 32 /* ParentHash */ + 32 /* StateHash */ + 32 /* RootHash */ + 20 /* Author */
}

func (h *Header) Encode(w io.Writer) error {
	return nil
}

func (h *Header) Decode(r io.Reader) error {
	return nil
}

func (h *Header) Check() error {
	if h.Author.Nil() {
	}
	return nil
}

type Block struct {
	header       *Header
	transactions Transactions
	successor    *utils.Hash
}

// Block value accessors

func (b *Block) GetHeader() *Header            { return b.header }
func (b *Block) GetTransactions() Transactions { return b.transactions }
func (b *Block) GetSuccessor() utils.Hash      { return *b.successor }

// Block function

func NewBlock(header *Header, transactions Transactions) *Block {
	return &Block{
		header:       header,
		transactions: transactions,
	}
}

func (b *Block) Size() int {
	size := b.header.Size()

	return size
}

func (b *Block) Encode(w io.Writer) error {
	return nil
}

func (b *Block) Decode(r io.Reader) error {
	return nil
}

func (b *Block) Check() {}
