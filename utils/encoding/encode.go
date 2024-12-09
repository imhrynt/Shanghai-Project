package encoding

import (
	"io"
	"reflect"
)

type encoder struct {
	list []byte
	size int // sum size of all bytes
}

func getPool() *encoder {
	buf := bufferPool.Get().(*encoder)
	buf.reset()
	return buf
}

func (enc *encoder) reset() {
	enc.list = enc.list[:0]
	enc.size = 0
}

func (enc *encoder) writeTo(w io.Writer) error {
	n, err := w.Write(enc.list)
	if err != nil {
		return err
	}
	if n != enc.size {
		return io.ErrShortBuffer
	}
	return nil
}

func (enc *encoder) Size() int     { return enc.size }
func (enc *encoder) Bytes() []byte { return enc.list }

func Encode(v any, w io.Writer) error {
	buf := getPool()
	defer bufferPool.Put(buf)
	buf.encode(reflect.ValueOf(v))
	err := buf.writeTo(w)
	return err
}

func (enc *encoder) encode(rv reflect.Value) {
	switch rv.Kind() {
	case reflect.Bool:
		enc.writeBool(rv.Bool())
	case reflect.String:
		enc.writeString(rv.String())
	case reflect.Float32:
		enc.writeFloat32(float32(rv.Float()))
	case reflect.Float64:
		enc.writeFloat64(rv.Float())
	case reflect.Int8:
		enc.writeInt8(int8(rv.Int()))
	case reflect.Int16:
		enc.writeInt16(int16(rv.Int()))
	case reflect.Int32:
		enc.writeInt32(int32(rv.Int()))
	case reflect.Int64:
		enc.writeInt64(rv.Int())
	case reflect.Uint8:
		enc.writeUint8(uint8(rv.Uint()))
	case reflect.Uint16:
		enc.writeUint16(uint16(rv.Uint()))
	case reflect.Uint32:
		enc.writeUint32(uint32(rv.Uint()))
	case reflect.Uint64:
		enc.writeUint64(rv.Uint())
	}
}
