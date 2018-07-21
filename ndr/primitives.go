package ndr

import (
	"math"
)

const (
	SizeBool   = 1
	SizeChar   = 1
	SizeUint8  = 1
	SizeUint16 = 2
	SizeUint32 = 4
	SizeUint64 = 8
	SizeEnum   = 2
	SizeSingle = 4
	SizeDouble = 8
	SizePtr    = 4
)

// Bool is an NDR Boolean which is a logical quantity that assumes one of two values: TRUE or FALSE.
// NDR represents a Boolean as one octet.
// It represents a value of FALSE as a zero octet, an octet in which every bit is reset.
// It represents a value of TRUE as a non-zero octet, an octet in which one or more bits are set.

// Char is an NDR character.
// NDR represents a character as one octet.
// Characters have two representation formats: ASCII and EBCDIC.

// USmall is an unsigned 8 bit integer

// UShort is an unsigned 16 bit integer

// ULong is an unsigned 32 bit integer

// UHyper is an unsigned 64 bit integer

// Small is an signed 8 bit integer

// Short is an signed 16 bit integer

// Long is an signed 32 bit integer

// Hyper is an signed 64 bit integer

// Enum is the NDR representation of enumerated types as signed short integers (2 octets)

// Single is an NDR defined single-precision floating-point data type

// Double is an NDR defined double-precision floating-point data type

// readBool reads a byte representing a boolean.
// NDR represents a Boolean as one octet.
// It represents a value of FALSE as a zero octet, an octet in which every bit is reset.
// It represents a value of TRUE as a non-zero octet, an octet in which one or more bits are set.
func (dec *Decoder) readBool() (bool, error) {
	i, err := dec.readUint8()
	if err != nil {
		return false, err
	}
	if i != 0 {
		return true, nil
	}
	return false, nil
}

// readUint8 reads bytes representing a 8bit unsigned integer.
func (dec *Decoder) readUint8() (uint8, error) {
	b, err := dec.r.ReadByte()
	if err != nil {
		return uint8(0), err
	}
	dec.ensureAlignment()
	return uint8(b), nil
}

// readUint16 reads bytes representing a 16bit unsigned integer.
func (dec *Decoder) readUint16() (uint16, error) {
	b, err := dec.readBytes(SizeUint16)
	if err != nil {
		return uint16(0), err
	}
	dec.ensureAlignment()
	return dec.ch.Endianness.Uint16(b), nil
}

// readUint32 reads bytes representing a 32bit unsigned integer.
func (dec *Decoder) readUint32() (uint32, error) {
	b, err := dec.readBytes(SizeUint32)
	if err != nil {
		return uint32(0), err
	}
	dec.ensureAlignment()
	return dec.ch.Endianness.Uint32(b), nil
}

// readUint32 reads bytes representing a 32bit unsigned integer.
func (dec *Decoder) readUint64() (uint64, error) {
	b, err := dec.readBytes(SizeUint64)
	if err != nil {
		return uint64(0), err
	}
	dec.ensureAlignment()
	return dec.ch.Endianness.Uint64(b), nil
}

// https://en.wikipedia.org/wiki/IEEE_754-1985
func (dec *Decoder) readFloat32() (f float32, err error) {
	b, err := dec.readBytes(SizeSingle)
	if err != nil {
		return
	}
	bits := dec.ch.Endianness.Uint32(b)
	f = math.Float32frombits(bits)
	dec.ensureAlignment()
	return
}

func (dec *Decoder) readFloat64() (f float64, err error) {
	b, err := dec.readBytes(SizeDouble)
	if err != nil {
		return
	}
	bits := dec.ch.Endianness.Uint64(b)
	f = math.Float64frombits(bits)
	dec.ensureAlignment()
	return
}