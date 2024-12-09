package codex

import (
	"bytes"
	"encoding/binary"
	"io"
	"math"
	"reflect"
	"sample/utils"
	"sync"
)

var encBufferPool = sync.Pool{New: func() interface{} { return new(bytes.Buffer) }}

func getEncBuffer() *bytes.Buffer { return encBufferPool.Get().(*bytes.Buffer) }

func putEncBuffer(buf *bytes.Buffer) {
	buf.Reset()
	encBufferPool.Put(buf)
}

func Encode(w io.Writer, v interface{}) {}

func EncodeToBytes(v interface{}) []byte {
	b := getEncBuffer()
	defer putEncBuffer(b)
	encode(b, reflect.ValueOf(v))
	buffer := make([]byte, b.Len())
	copy(buffer, b.Bytes())
	return buffer
}

func SizeOf(v interface{}) int {
	return 0
}

func encode(buf *bytes.Buffer, rv reflect.Value) {
	fw, _ := getWriter(rv.Type())
	fw(buf, rv)
}

func writeBool(b *bytes.Buffer, rv reflect.Value) {
	if rv.Bool() {
		b.WriteByte(1)
	} else {
		b.WriteByte(0)
	}
}

func writeUint(b *bytes.Buffer, rv reflect.Value) {
	switch rv.Type().Kind() {
	case reflect.Uint8:
		b.WriteByte(byte(rv.Uint()))
	case reflect.Uint16:
		var buf [2]byte
		binary.LittleEndian.PutUint16(buf[:], uint16(rv.Uint()))
		b.Write(buf[:])
	case reflect.Uint32:
		var buf [4]byte
		binary.LittleEndian.PutUint32(buf[:], uint32(rv.Uint()))
		b.Write(buf[:])
	case reflect.Uint64:
		var buf [8]byte
		binary.LittleEndian.PutUint64(buf[:], rv.Uint())
		b.Write(buf[:])
	}
}

func writeInt(buf *bytes.Buffer, rv reflect.Value) {
	switch rv.Type().Kind() {
	case reflect.Int8:
		writeUint(buf, reflect.ValueOf(uint8(rv.Int())))
	case reflect.Int16:
		writeUint(buf, reflect.ValueOf(uint16(rv.Int())))
	case reflect.Int32:
		writeUint(buf, reflect.ValueOf(uint32(rv.Int())))
	case reflect.Int64:
		writeUint(buf, reflect.ValueOf(uint64(rv.Int())))
	}
}

func writeFloat(buf *bytes.Buffer, rv reflect.Value) error {
	if math.IsNaN(rv.Float()) {
		return utils.ErrNaNFloat.Error()
	}
	switch rv.Type().Kind() {
	case reflect.Float32:
		writeUint(buf, reflect.ValueOf(math.Float32bits(float32(rv.Float()))))
	case reflect.Float64:
		writeUint(buf, reflect.ValueOf(math.Float64bits(rv.Float())))
	}
	return nil
}

func writeString(buf *bytes.Buffer, rv reflect.Value) error {
	writeUint(buf, reflect.ValueOf(uint32(rv.Len())))
	buf.WriteString(rv.String())
	return nil
}

func writeMap(buf *bytes.Buffer, rv reflect.Value) {}

func writeStruct(buf *bytes.Buffer, rv reflect.Value) {}

func writeInterface(buf *bytes.Buffer, rv reflect.Value) {}

func writeHead() {}
