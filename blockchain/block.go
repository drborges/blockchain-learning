package blockchain

import "time"

type BlockHeader struct {
	hash           []byte
	Version        int32
	HashPrevBlock  []byte
	HashMerkleRoot []byte
	Time           time.Time // Needs to be serialized as Unix time
	Bits           uint32
	Nonce          uint
}

func (b *BlockHeader) Hash() []byte {
	if b == nil { // Genesis block
		return []byte{} // Genesis hash == nil hash a.k.a all zeros
	}

	if b.hash == nil {
		b.hash = []byte{} // TODO calculate hash
	}
	return b.hash
}

type BlockPayload struct {
	TxCount int
	Tx      []*Transaction
}

type Block struct {
	Header  BlockHeader
	Payload BlockPayload
}

func NewBlock(prevBlock *Block) *Block {
	return &Block{
		Header: BlockHeader{
			HashPrevBlock: prevBlock.Header.Hash(),
		},
	}
}

func (b *Block) Add(t *Transaction) {
	// TODO trigger trasaction validation
	b.Payload.Tx = append(b.Payload.Tx, t)
}
