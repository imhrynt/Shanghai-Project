package encoding

import (
	"math"
)

// writeBool writes true as 1 or false as 0
func (enc *encoder) writeBool(v bool) {
	if v {
		enc.writeByte(1)
	} else {
		enc.writeByte(0)
	}
}

// writeByte writes single byte in range 0 ~ 255
func (enc *encoder) writeByte(v byte) {
	enc.list = append(enc.list, v)
	enc.size += 1
}

// writeBytes writes slice of bytes with padding size for header
func (enc *encoder) writeBytes(v []byte) {
	enc.writeUint32(uint32(len(v)))
	enc.list = append(enc.list, v...)
	enc.size += len(v)
}

// writeFloat32 writes float32
func (enc *encoder) writeFloat32(v float32) {
	enc.writeUint32(math.Float32bits(v))
}

// writeFloat64 writes float64
func (enc *encoder) writeFloat64(v float64) {
	enc.writeUint64(math.Float64bits(v))
}

// writeInt16 writes int16
func (enc *encoder) writeInt16(v int16) {
	enc.writeUint16(uint16(v))
}

// writeInt32 writes int32
func (enc *encoder) writeInt32(v int32) {
	enc.writeUint32(uint32(v))
}

// writeInt64 writes int64
func (enc *encoder) writeInt64(v int64) {
	enc.writeUint64(uint64(v))
}

// writeInt8 writes int8
func (enc *encoder) writeInt8(v int8) {
	enc.writeByte(uint8(v))
}

// writeString writes bytes of string with padding size for header
func (enc *encoder) writeString(v string) {
	enc.writeBytes([]byte(v))
}

// writeUint16 writes uint16
func (enc *encoder) writeUint16(v uint16) {
	enc.list = append(enc.list,
		byte(v), byte(v>>8))
	enc.size += 2
}

// writeUint32 writes uint32
func (enc *encoder) writeUint32(v uint32) {
	enc.list = append(enc.list,
		byte(v), byte(v>>8),
		byte(v>>16), byte(v>>24))
	enc.size += 4
}

// writeUint64 writes uint64
func (enc *encoder) writeUint64(v uint64) {
	enc.list = append(enc.list,
		byte(v), byte(v>>8), byte(v>>16), byte(v>>24),
		byte(v>>32), byte(v>>40), byte(v>>48), byte(v>>56))
	enc.size += 8
}

// writeUint8 writes uint8
func (enc *encoder) writeUint8(v uint8) {
	enc.writeByte(v)
}
