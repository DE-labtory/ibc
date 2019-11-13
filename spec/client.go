package spec

// ClientType is a set of definitions of the data structures, initialisation logic,
// validity predicate, and misbehaviour predicate required to operate a light client.
type ClientType int

// ClientState keeps arbitrary internal state to track verified roots and past misbehaviours.
type ClientState struct {
	Frozen        bool
	PastPublicKey map[PublicKey]bool
	VerifiedRoot  map[uint64]CommitmentRoot
}

// Header provides information to update a ConsensusState.
type Header struct {
	Sequence       uint64
	CommitmentRoot CommitmentRoot
	Signature      Signature
	NewPublicKey   PublicKey
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
	checkValidityAndUpdateState(Header)
	checkMisbehaviourAndUpdateState([]byte)
	verifyMembership(uint64, CommitmentProof, Path, Value) bool
	verifyNonMembership(uint64, CommitmentProof, Path) bool
	verifyClientConsensusState(uint64, CommitmentProof, Identifier, ConsensusState) bool
	verifyConnectionState(uint64, CommitmentPrefix, CommitmentProof, Identifier, ConnectionEnd) bool
	verifyChannelState(uint64, CommitmentPrefix, CommitmentProof, Identifier, Identifier, ChannelEnd)
	verifyPacketCommitment(uint64, CommitmentPrefix, CommitmentProof, Identifier, Identifier, uint64, []byte) bool
	verifyPacketAcknowledgement(uint64, CommitmentPrefix, CommitmentProof, Identifier, Identifier, uint64, []byte) bool
	verifyPacketAcknowledgementAbsence(uint64, CommitmentPrefix, CommitmentProof, Identifier, Identifier, uint64) bool
	verifyNextSequenceRecv(uint64, CommitmentPrefix, CommitmentProof, Identifier, Identifier, uint64) bool
}
