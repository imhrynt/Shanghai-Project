package utils

import (
	"fmt"
	"math/rand"
)

const AddressSize = 20

type Address [AddressSize]byte

func Bytes2Address(b []byte) (Address, error) {
	var a Address
	if err := a.SetBytes(b); err != nil {
		return Address{}, err
	}
	return a, nil
}

func Hex2Address(s string) (Address, error) {
	b, err := Decode(s)
	if err != nil {
		return Address{}, err
	}
	return Bytes2Address(b)
}

func (a Address) Bytes() []byte {
	return a[:]
}

func (a Address) Hex() string {
	return Encode(a.Bytes())
}

func (a Address) String() string {
	return a.Hex()
}

func (a Address) TerminalString() string {
	return fmt.Sprintf("0x%x..%x", a[:3], a[AddressSize-4:])
}

func (a *Address) Cmp(cmp Address) bool {
	return *a == cmp
}

func (a *Address) Nil() bool {
	return *a == Address{}
}

func (a *Address) SetBytes(b []byte) error {
	if len(b) > AddressSize {
		return ErrTypesLength.Errorf(len(b), AddressSize)
	}
	copy(a[AddressSize-len(b):], b)
	return nil
}

// *Generate implements for quick testing
func (a Address) Generate() Address {
	for i := AddressSize - 1; i > rand.Intn(AddressSize); i-- {
		a[i] = byte(rand.Uint32())
	}
	return a
}
