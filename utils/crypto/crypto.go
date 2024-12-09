package crypto

import (
	"crypto/sha256"
	"sample/utils"
	"hash"
	"sync"
)

var BufferPool = sync.Pool{
	New: func() any {
		return make([]byte, 8)
	},
}

type Sha2State interface {
	hash.Hash
}

func NewSha2State() Sha2State {
	return sha256.New()
}

func H256(data ...[]byte) (h utils.Hash) {
	sha := NewSha2State()
	for _, b := range data {
		sha.Write(b)
	}
	sha.Sum(h[:0])
	return h
}

func Sha256(data ...[]byte) []byte {
	sha := NewSha2State()
	for _, b := range data {
		sha.Write(b)
	}
	return sha.Sum(nil)
}
