package spec

// Consensus is a Header generating function which takes the previous ConsensusState with the messages.
type Consensus func(ConsensusState, []Message) Header

// ConsensusState is used by the validity predicate to verify new commits & state roots.
type ConsensusState struct {
	Sequence  uint64
	PublicKey PublicKey
}

type Message interface{}

// Blockchain is a consensus algorithm which generates valid headers.
// It generates a unique list of headers starting from a genesis ConsensusState with arbitrary messages.
type Blockchain interface {
	GetGenesis() Consensus
	GetConsensus() Consensus
}
