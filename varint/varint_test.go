package varint_test

import (
	"bytes"
	"github.com/drborges/blockchain-learning/varint"
	"testing"
)

func TestSize(t *testing.T) {
	tests := []struct {
		val  uint64 // Value to get the serialized size for
		size int    // Expected serialized size
	}{
		// Single byte
		{0, 1},
		// Max single byte
		{0xfc, 1},
		// Min 2-byte
		{0xfd, 3},
		// Max 2-byte
		{0xffff, 3},
		// Min 4-byte
		{0x10000, 5},
		// Max 4-byte
		{0xffffffff, 5},
		// Min 8-byte
		{0x100000000, 9},
		// Max 8-byte
		{0xffffffffffffffff, 9},
	}

	for i, test := range tests {
		serializedSize := varint.Size(test.val)
		if serializedSize != test.size {
			t.Errorf("Size #%d got: %d, want: %d", i, serializedSize, test.size)
			continue
		}
	}
}

func TestVarintSerialize(t *testing.T) {
	tests := []struct {
		val  uint64 // Value to get the serialized size for
		size int    // Expected serialized size
	}{
		// Single byte
		{0, 1},
		// Max single byte
		{0xfc, 1},
		// Min 2-byte
		{0xfd, 3},
		// Max 2-byte
		{0xffff, 3},
		// Min 4-byte
		{0x10000, 5},
		// Max 4-byte
		{0xffffffff, 5},
		// Min 8-byte
		{0x100000000, 9},
		// Max 8-byte
		{0xffffffffffffffff, 9},
	}

	for i, test := range tests {
		serialized := varint.Serialize(test.val)
		if len(serialized) != test.size {
			t.Errorf("Size #%d got: %d, want: %d", i, len(serialized), test.size)
			continue
		}

		value, err := varint.Deserialize(bytes.NewBuffer(serialized))
		if err != nil {
			t.Fatal(err)
		}

		if value != test.val {
			t.Errorf("Value: Expected %v, got %v", i, test.val, value)
			continue
		}
	}
}
