package spec

// ClientType is a set of definitions of the data structures, initialisation logic,
// validity predicate, and misbehaviour predicate required to operate a light client.
type ClientType int

// ClientState keeps arbitrary internal state to track verified roots and past misbehaviours.
type ClientState struct {
	frozen        bool
	pastPublicKey map[string]bool
	verifiedRoot  map[int]CommitmentRoot
}

// Header provides information to update a ConsensusState.
type Header struct {
	sequence       int
	commitmentRoot CommitmentRoot
	signature      Signature
	newPublicKey   PublicKey
}

// Evidence has candidate headers which may be misbehaviours.
type Evidence struct {
	h1 Header
	h2 Header
}

type Client interface {
	getIdentifier() Identifier
	getClientType() ClientType
	getConsensusState() ConsensusState
}