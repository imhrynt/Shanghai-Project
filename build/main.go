package main

import (
	"fmt"
	"sample/core"
	"sample/utils"
	"math/rand/v2"
	"reflect"
	"time"
	"unsafe"
)

var header = core.Header{
	Time:      uint64(time.Now().Unix()),
	Ancestors: utils.Hash{}.Generate(),
	State:     utils.Hash{}.Generate(),
	Root:      utils.Hash{}.Generate(),
	Author:    utils.Address{}.Generate(),
}

var recipient = utils.Address{}.Generate()

var dynamicTx = &core.DynamicTx{
	Number:    uint64(rand.Uint64()),
	Fee:       uint64(rand.Uint64()),
	Amount:    uint64(rand.Uint64()),
	Recipient: &recipient,
}

var txs = append(core.Transactions{}, core.NewTransaction(dynamicTx))

var block = core.NewBlock(&header, txs)

var slice = []byte{255, 123, 34, 65, 34, 34, 156, 45, 75}

func writeFloat32(b []byte, f float32) {
	v := *(*uint32)(unsafe.Pointer(&f))
	b = append(b, byte(v), byte(v>>8),
		byte(v>>16), byte(v>>24))
}

func fieldsCalc(x any) int {
	rv := reflect.TypeOf(x)
	f := rv.Elem().NumField()
	return f
}

func main() {
	fmt.Print(utils.Hash{}.TerminalString())
}
