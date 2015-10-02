package blockchain_test

import (
	"bytes"
	"github.com/drborges/blockchain-learning/blockchain"
	"reflect"
	"testing"
)

func TestTxInBtcEncodeDecode(t *testing.T) {
	expected := blockchain.TxIn{
		PreviousTxOut: blockchain.Outpoint{
			Hash:  [32]byte{'1', '2', '3'},
			Index: 2,
		},
		ScriptSignature: []byte("abcd"),
		Sequence:        2,
	}

	buf := bytes.NewBuffer(make([]byte, 0, expected.SerializedSize()))
	if err := expected.BtcEncode(buf); err != nil {
		t.Fatal(err)
	}

	var actual blockchain.TxIn
	if err := actual.BtcDecode(buf); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
