package encoding

import (
	"io"
	"sync"
)

// fragmented Binary Object Serialize - FBOS

type subset struct {
	offset int
	size   int
	value  []byte
}

type buffer struct {
	subsetList []subset
	subLisSize int
}

var bufferPool = sync.Pool{
	New: func() interface{} { return new(buffer) },
}

func getBuffer() *buffer {
	buf := bufferPool.Get().(*buffer)
	buf.reset()
	return buf
}

func (enc *buffer) reset() {
	enc.subsetList = enc.subsetList[:0]
	enc.subLisSize = 0
}

func (enc *buffer) writeTo(w io.Writer) error {
	for _, subset := range enc.subsetList {
		_, err := w.Write(subset.value)
		if err != nil {
			return err
		}
	}
	return nil
}

func (enc *buffer) push(v []byte) {
	subset := subset{offset: enc.subLisSize, size: len(v), value: v}
	enc.subsetList = append(enc.subsetList, subset)
	enc.subLisSize += len(v)
}

func (enc *buffer) saveBool(v bool) {
	if v {
		enc.push([]byte{0x01})
	} else {
		enc.push([]byte{0x00})
	}
}

func (enc *buffer) saveString(v string) {
	enc.push([]byte(v))
}

func ForTesting(w io.Writer) *buffer {
	buf := getBuffer()
	defer bufferPool.Put(buf)
	buf.saveBool(false)
	buf.saveString("amirullah heryanto muslan")
	buf.saveBool(true)
	buf.saveString("amirullah heryanto")
	return buf
}
