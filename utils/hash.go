package utils

import (
	"fmt"
	"math/rand"
)

const HashSize = 32

type Hash [HashSize]byte

func Bytes2Hash(b []byte) (Hash, error) {
	var h Hash
	if err := h.SetBytes(b); err != nil {
		return Hash{}, err
	}
	return h, nil
}

func Hex2Hash(s string) (Hash, error) {
	b, err := Decode(s)
	if err != nil {
		return Hash{}, err
	}
	return Bytes2Hash(b)
}

func (h Hash) Bytes() []byte {
	return h[:]
}

func (h Hash) Hex() string {
	return Encode(h.Bytes())
}

func (h Hash) String() string {
	return h.Hex()
}

func (h Hash) TerminalString() string {
	return fmt.Sprintf("0x%x..%x", h[:3], h[HashSize-4:])
}

func (h *Hash) Cmp(cmp Hash) bool {
	return *h == cmp
}

func (h *Hash) Nil() bool {
	return *h == Hash{}
}

func (h *Hash) SetBytes(b []byte) error {
	if len(b) > HashSize {
		return ErrTypesLength.Errorf(len(b), HashSize)
	}
	copy(h[HashSize-len(b):], b)
	return nil
}

// *Generate implements for quick testing
func (h Hash) Generate() Hash {
	for i := HashSize - 1; i > rand.Intn(HashSize); i-- {
		h[i] = byte(rand.Uint32())
	}
	return h
}
