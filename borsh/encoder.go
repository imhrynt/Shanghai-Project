package borsh

import (
	"io"
	"reflect"
)

const (
	unsignIntType = 0x80
	sliceByteType = 0x00
	mapType       = 0x00
)

type Encoder struct {
	len int
	dst io.Writer
	buf [9]byte
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{
		dst: w,
	}
}

func (enc *Encoder) writeTo(b []byte) error {
	n, err := enc.dst.Write(b)
	enc.len += n
	return err
}

func isUint(rk reflect.Kind) bool {
	return rk >= reflect.Uint && rk <= reflect.Uintptr
}

func (enc *Encoder) WriteUint(i uint64) error {
	switch {
	case i < (1 << 7):
		return enc.writeTo([]byte{byte(i)})
	case i < (1 << 8):
		enc.buf[0] = 0x81
		enc.buf[1] = byte(i)
		return enc.writeTo(enc.buf[:2])
	case i < (1 << 16):
		enc.buf[0] = 0x82
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		return enc.writeTo(enc.buf[:3])
	case i < (1 << 24):
		enc.buf[0] = 0x83
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		enc.buf[3] = byte(i >> 16)
		return enc.writeTo(enc.buf[:4])
	case i < (1 << 32):
		enc.buf[0] = 0x84
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		enc.buf[3] = byte(i >> 16)
		enc.buf[4] = byte(i >> 24)
		return enc.writeTo(enc.buf[:5])
	case i < (1 << 40):
		enc.buf[0] = 0x85
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		enc.buf[3] = byte(i >> 16)
		enc.buf[4] = byte(i >> 24)
		enc.buf[5] = byte(i >> 40)
		return enc.writeTo(enc.buf[:6])
	case i < (1 << 48):
		enc.buf[0] = 0x86
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		enc.buf[3] = byte(i >> 16)
		enc.buf[4] = byte(i >> 24)
		enc.buf[5] = byte(i >> 32)
		enc.buf[6] = byte(i >> 40)
		return enc.writeTo(enc.buf[:7])
	case i < (1 << 56):
		enc.buf[0] = 0x87
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		enc.buf[3] = byte(i >> 16)
		enc.buf[4] = byte(i >> 24)
		enc.buf[5] = byte(i >> 32)
		enc.buf[6] = byte(i >> 40)
		enc.buf[7] = byte(i >> 48)
		return enc.writeTo(enc.buf[:8])
	default:
		enc.buf[0] = 0x88
		enc.buf[1] = byte(i)
		enc.buf[2] = byte(i >> 8)
		enc.buf[3] = byte(i >> 16)
		enc.buf[4] = byte(i >> 24)
		enc.buf[5] = byte(i >> 32)
		enc.buf[6] = byte(i >> 40)
		enc.buf[7] = byte(i >> 48)
		enc.buf[8] = byte(i >> 56)
		return enc.writeTo(enc.buf[:9])
	}
}

/*func Encode(w io.Writer, s interface{}) {

}

func (enc *Encoder) Size() int {
	return enc.len
}

func (enc *Encoder) serialize(rv reflect.Value) (err error) {
	switch rv.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			n, err = enc.dst.Write([]byte{1})
		} else {
			n, err = enc.dst.Write([]byte{0})
		}
		enc.length += n

	case reflect.Int8:

	case reflect.Int16:
		err = enc.writeInt16(int16(rv.Int()))

	case reflect.Int32:
		err = enc.writeInt32(int32(rv.Int()))

	case reflect.Int64:
		err = enc.writeInt64(rv.Int())

	case reflect.Uint8:
		err = enc.writeUint8(uint8(rv.Uint()))

	case reflect.Uint16:
		err = enc.writeUint16(uint16(rv.Uint()))

	case reflect.Uint32:
		err = enc.writeUint32(uint32(rv.Uint()))

	case reflect.Uint64:
		err = enc.writeUint64(rv.Uint())

	case reflect.Float32:
		err = enc.writeFloat32(float32(rv.Float()))

	case reflect.Float64:
		err = enc.writeFloat64(rv.Float())

	case reflect.String:
		var size [4]byte
		binary.LittleEndian.PutUint32(size[:], uint32(len(rv.String())))
		n, err := enc.dst.Write(size[:] + []byte(rv.String()))
		enc.len += n
	case reflect.Array:
		for i := 0; i < rv.Len(); i++ {
			err = enc.serialize(rv.Index(i))
		}

	case reflect.Slice:
		err = enc.writeUint32(uint32(len(rv.Bytes())))
		for i := 0; i < rv.Len(); i++ {
			err = enc.serialize(rv.Index(i))
		}

	case reflect.Pointer:

	case reflect.Struct:

	case reflect.Map:

	}
	return err
}

func (enc *Encoder) serializeStruct(rv reflect.Value) error {
	rt := reflect.TypeOf(rv)
	if rt.NumField() > 0 {
		firstField := rt.Field(0)
		if isEnum(firstField.Type) && firstField.Tag.Get("borsh_enum") == "true" {
			return enc.serializeComplexEnum(rv)
		}
	}
	for i := 0; i < rt.NumField(); i++ {
		field := rt.Field(i)
		if field.Tag.Get("borsh_skip") == "true" {
			continue
		}
		if err := enc.serialize(rv.Field(i)); err != nil {
			return err
		}
	}
	return nil
}

func (enc *Encoder) serializeComplexEnum(rv reflect.Value) error {
	rt := rv.Type()
	enum := Enum(rv.Field(0).Uint())
	if err := enc.writeTo([]byte{byte(enum)}); err != nil {
		return err
	}
	if int(enum)+1 >= rt.NumField() {
		return errors.New("complex enum too large")
	}
	field := rv.Field(int(enum) + 1)
	if field.Kind() == reflect.Struct {
		return enc.serializeStruct(field)
	}
	return nil
}

// Primitive

func (enc *Encoder) writeBool(b bool) error {
	var v byte
	if b {
		v += 1
	}
	n, err := enc.dst.Write([]byte{v})
	if err != nil {
		return err
	}
	enc.len += n
	return nil
}*/
