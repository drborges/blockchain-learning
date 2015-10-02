package blockchain_test

import (
	"bytes"
	"github.com/drborges/blockchain-learning/blockchain"
	"reflect"
	"testing"
)

func TestTxOutBtcEncodeDecode(t *testing.T) {
	expected := blockchain.TxOut{
		Value:    123,
		PKScript: []byte("pkscript"),
	}

	buf := bytes.NewBuffer(make([]byte, 0, expected.SerializedSize()))
	if err := expected.BtcEncode(buf); err != nil {
		t.Fatal(err)
	}

	var actual blockchain.TxOut
	if err := actual.BtcDecode(buf); err != nil {
		t.Fatal(err)
	}

	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("expected %+v, got %+v", expected, actual)
	}
}
