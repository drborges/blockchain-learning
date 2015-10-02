package blockchain

type TxOut struct {
	Value        int64
	ScriptLen    uint
	ScriptPubkey []byte
}

type Transaction struct {
	Version  uint32
	InCount  int
	In       []TxIn
	OutCount int
	Out      []TxOut
	LockTime uint32
}
