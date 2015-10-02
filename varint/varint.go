package varint

import (
	"encoding/binary"
	"io"
	"math"
	"errors"
)

// CODE EXTRACTED AND ADAPTED FROM:
// https://github.com/btcsuite/btcd/blob/master/wire/common.go
//
// For more information on variable length integer, see the bitcoin specs:
// https://en.bitcoin.it/wiki/Protocol_documentation#Variable_length_integer

func Size(val uint64) int {
	// The value is small enough to be represented by itself, so it's
	// just 1 byte.
	if val < 0xfd {
		return 1
	}

	// Discriminant 1 byte plus 2 bytes for the uint16.
	if val <= math.MaxUint16 {
		return 3
	}

	// Discriminant 1 byte plus 4 bytes for the uint32.
	if val <= math.MaxUint32 {
		return 5
	}

	// Discriminant 1 byte plus 8 bytes for the uint64.
	return 9
}

// Serialize serializes the given int into 1, 3, 5 or 9 bytes depending on whether
// the value can be represented with a corresponding amount of bytes
func Serialize(val uint64) []byte {
	switch Size(val) {
	case 1:
		return []byte{uint8(val)}
	case 3:
		var buf [3]byte
		buf[0] = 0xfd
		binary.LittleEndian.PutUint16(buf[1:], uint16(val))
		return buf[:]
	case 5:
		var buf [5]byte
		buf[0] = 0xfe
		binary.LittleEndian.PutUint32(buf[1:], uint32(val))
		return buf[:]
	default:
		var buf [9]byte
		buf[0] = 0xff
		binary.LittleEndian.PutUint64(buf[1:], val)
		return buf[:]
	}
}

// Deserialize deserializes a variable length int from an io.Reader by:
// 1. Read the int length as the first byte in the reader;
// 2. Read the amount of bytes pointed by the extracted length in 1 corresponding to the actual int value
func Deserialize(r io.Reader) (uint64, error) {
	var b [8]byte
	_, err := io.ReadFull(r, b[0:1])
	if err != nil {
		return 0, err
	}

	var rv uint64
	discriminant := uint8(b[0])
	switch discriminant {
	case 0xff:
		_, err := io.ReadFull(r, b[:])
		if err != nil {
			return 0, err
		}
		rv = binary.LittleEndian.Uint64(b[:])

		// The encoding is not canonical if the value could have been
		// encoded using fewer bytes.
		min := uint64(0x100000000)
		if rv < min {
			return 0, errors.New("Invalid encoded varint")
		}

	case 0xfe:
		_, err := io.ReadFull(r, b[0:4])
		if err != nil {
			return 0, err
		}
		rv = uint64(binary.LittleEndian.Uint32(b[:]))

		// The encoding is not canonical if the value could have been
		// encoded using fewer bytes.
		min := uint64(0x10000)
		if rv < min {
			return 0, errors.New("Invalid encoded varint")
		}

	case 0xfd:
		_, err := io.ReadFull(r, b[0:2])
		if err != nil {
			return 0, err
		}
		rv = uint64(binary.LittleEndian.Uint16(b[:]))

		// The encoding is not canonical if the value could have been
		// encoded using fewer bytes.
		min := uint64(0xfd)
		if rv < min {
			return 0, errors.New("Invalid encoded varint")
		}

	default:
		rv = uint64(discriminant)
	}

	return rv, nil
}
