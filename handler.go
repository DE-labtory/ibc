package ibc

// Handler interface defines client-semantics, connection-semantice, channel and packet-semantics, port-allocation.
type Handler interface {
	// Client Lifecycle Management : ICS 2
	createClient(
		id Identifier,
		clientType ClientType,
		consensusState ConsensusState)
	updateClient(id Identifier, header Header)
	queryClientConsensusState(id Identifier) ConsensusState
	queryClient(id Identifier) ClientState

	// Connection Lifecycle Management : ICS 3
	connOpenInit(
		identifier Identifier,
		desiredCounterpartyConnectionIdentifier Identifier,
		counterpartyPrefix CommitmentPrefix,
		clientIdentifier Identifier,
		counterPartyClientIdentifer Identifier)
	connOpenTry(
		desiredIdentifer Identifier,
		counterPartyConnectionIdentifier Identifier,
		counterPartyPrefix CommitmentPrefix,
		counterPartyClientIdentifier Identifier,
		clientIdentifier Identifier,
		counterPartyVersions []string,
		proofInit CommitmentProof,
		proofHeight int,
		consensusHeight int)
	connOpenAck(
		identifier Identifier,
		version string,
		proofTry CommitmentProof,
		proofHeight int,
		consensusHeight int)
	connOpenConfirm(
		identifier Identifier,
		proofAck CommitmentProof,
		proofHeight int)
	queryConnection(id Identifier) (ConnectionEnd, error)

	// Channel Lifecycle Management : ICS 4
	chanOpenInit(
		order ChannelOrder,
		connectionHops []Identifier,
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterPartyPortIdentifier Identifier,
		counterPartyChannelIdentifier Identifier,
		version string)
	chanOpenTry(
		order ChannelOrder,
		connectionHops []Identifier,
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterPartyPortIdentifier Identifier,
		counterPartyChannelIdentifier Identifier,
		version string,
		counterPartyVersion string,
		proofInit CommitmentProof,
		proofHeight int)
	chanOpenAck(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterPartyVersion string,
		proofTry CommitmentProof,
		proofHeight int)
	chanOpenConfirm(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		proofAck CommitmentProof,
		proofHeight int)
	chanCloseInit(
		portIdentifier Identifier,
		channelIdentifier Identifier)
	chanCloseConfirm(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		proofInit CommitmentProof,
		proofHeight int)
	queryChannel()

	// Packet Relay : ICS 4
	sendPacket(packet Packet)
	recvPacket(
		packet OpaquePacket,
		proof CommitmentProof,
		proofHeight int,
		acknowledgement []byte)
	acknowledgePacket(
		packet OpaquePacket,
		acknowledgement []byte,
		proof CommitmentProof,
		proofHeight int)
	timeoutPacket(
		packet OpaquePacket,
		proof CommitmentProof,
		proofHeight int,
		nextSequenceRecv int)
	timeoutOnClose(
		packet Packet,
		proofNonMembership CommitmentProof,
		proofClosed CommitmentProof,
		proofHeight int)
	cleanupPacket(
		packet OpaquePacket,
		proof CommitmentProof,
		proofHeight int,
		// TODO : modify it
		nextSequenceRecvOrAcknowledgement interface{})
}

// HandlerModule is implementation of Handler interface.
type HandlerModule struct {
	privateStore  *PrivateStore
	provableStore *ProvableStore
}
