package ibc

import "github.com/DE-labtory/ibc/spec"

// Client is temporarily implementation of Client interface.
type Client struct {
	identifier     spec.Identifier
	clientType     spec.ClientType
	consensusState spec.ConsensusState
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
	panic("implement me!")
}

func (c *Client) checkValidityAndUpdateState(h spec.Header) {
	panic("implement me!")
}

func (c *Client) checkMisbehaviourAndUpdateState(b []byte) {
	panic("implement me!")
}
func (c *Client) verifyClientConsensusState(clientState spec.ClientState,
	height uint64,
	proof spec.CommitmentProof,
	clientIdentifier spec.Identifier,
	consensusState spec.ConsensusState) bool {
	return false
}

func (c *Client) verifyConnectionState(clientState spec.ClientState,
	height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	connectionIdentifier spec.Identifier,
	connectionEnd spec.ConnectionEnd) bool {
	return false
}

func (c *Client) verifyChannelState(clientState spec.ClientState,
	height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	protIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	channelEnd spec.ChannelEnd) {
	panic("implement me!")
}

func (c *Client) verifyPacketCommitment(clientState spec.ClientState,
	height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	sequence uint64,
	commitment []byte) bool {
	return false
}

func (c *Client) verifyPacketAcknowledgement(clientState spec.ClientState,
	height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	sequence uint64,
	acknowledgement []byte) bool {
	return false
}

func (c *Client) verifyPacketAcknowledgementAbsence(clientState spec.ClientState,
	height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	sequence uint64) bool {
	return false
}

func (c *Client) verifyNextSequenceRecv(clientState spec.ClientState,
	height uint64,
	prefix spec.CommitmentPrefix,
	proof spec.CommitmentProof,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	nextSequenceRecv uint64) bool {
	return false
}
