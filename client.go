package ibc

import "github.com/DE-labtory/ibc/spec"

// Client is temporarily implementation of Client interface.
type Client struct {
	identifier     spec.Identifier
	clientType     spec.ClientType
	consensusState spec.ConsensusState
	clientState    spec.ClientState
}

func (c *Client) getIdentifier() spec.Identifier {
	return c.identifier
}

func (c *Client) getClientType() spec.ClientType {
	return c.clientType
}

func (c *Client) getConsensusState() spec.ConsensusState {
	return c.consensusState
}

func (c *Client) initialize(cs spec.ConsensusState) spec.ClientState {
	return spec.ClientState{
		Frozen: false,
		PastPublicKey: map[spec.PublicKey]bool{
			cs.PublicKey: true,
		},
		VerifiedRoot: make(map[uint64]spec.CommitmentRoot),
	}
}

func (c *Client) checkValidityAndUpdateState(header spec.Header) {
	if c.consensusState.Sequence+1 != header.Sequence {
		panic("Header's sequence must same as consensus state's sequence.")
	}

	if !c.consensusState.PublicKey.Verify(header.Signature) {
		panic("Signature verifying failed.")
	}

	c.consensusState.PublicKey = header.NewPublicKey
	c.clientState.PastPublicKey[header.NewPublicKey] = true
	c.consensusState.Sequence = header.Sequence
	c.clientState.VerifiedRoot[c.consensusState.Sequence] = header.CommitmentRoot
}

func (c *Client) checkMisbehaviourAndUpdateState(b []byte) {
	panic("implement me!")
}

func (c *Client) verifyMembership(sequence uint64, proof spec.CommitmentProof, path spec.Path, value spec.Value) bool {
	if c.clientState.Frozen {
		panic("Client state is frozen.")
	}
	return c.clientState.VerifiedRoot[sequence].VerifyMemebership(path, value, proof)
}

func (c *Client) verifyNonMembership(sequence uint64, proof spec.CommitmentProof, path spec.Path) bool {
	if c.clientState.Frozen {
		panic("Client state is frozen")
	}
	return c.clientState.VerifiedRoot[sequence].VerifyNonMembership(path, proof)
}

func (c *Client) verifyClientConsensusState(height uint64,
	proof spec.CommitmentProof,
	clientIdentifier spec.Identifier,
	consensusState spec.ConsensusState) bool {
	return false
}

func (c *Client) verifyConnectionState(height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	connectionIdentifier spec.Identifier,
	connectionEnd spec.ConnectionEnd) bool {
	return false
}

func (c *Client) verifyChannelState(height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	protIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	channelEnd spec.ChannelEnd) {
	panic("implement me!")
}

func (c *Client) verifyPacketCommitment(height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	sequence uint64,
	commitment []byte) bool {
	return false
}

func (c *Client) verifyPacketAcknowledgement(height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	sequence uint64,
	acknowledgement []byte) bool {
	return false
}

func (c *Client) verifyPacketAcknowledgementAbsence(height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	sequence uint64) bool {
	return false
}

func (c *Client) verifyNextSequenceRecv(height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	nextSequenceRecv uint64) bool {
	return false
}
