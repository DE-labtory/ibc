package spec

// Handler interface defines client-semantics, connection-semantice, channel and packet-semantics, port-allocation.
type Handler interface {
	// Client Lifecycle Management : ICS 2
	CreateClient(
		id Identifier,
		clientType ClientType,
		consensusState ConsensusState)
	UpdateClient(id Identifier, header Header)
	QueryClientConsensusState(id Identifier) ConsensusState
	QueryClient(id Identifier) ClientState
	ClientStatePath(id Identifier) Path
	ClientTypePath(id Identifier) Path
	ConsensusStatePath(id Identifier) Path

	// Connection Lifecycle Management : ICS 3
	ConnOpenInit(
		identifier Identifier,
		desiredCounterpartyConnectionIdentifier Identifier,
		counterpartyPrefix CommitmentPrefix,
		clientIdentifier Identifier,
		counterPartyClientIdentifer Identifier)
	ConnOpenTry(
		desiredIdentifer Identifier,
		counterPartyConnectionIdentifier Identifier,
		counterPartyPrefix CommitmentPrefix,
		counterPartyClientIdentifier Identifier,
		clientIdentifier Identifier,
		counterPartyVersions []string,
		proofInit CommitmentProof,
		proofHeight int,
		consensusHeight int)
	ConnOpenAck(
		identifier Identifier,
		version string,
		proofTry CommitmentProof,
		proofHeight int,
		consensusHeight int)
	ConnOpenConfirm(
		identifier Identifier,
		proofAck CommitmentProof,
		proofHeight int)
	QueryConnection(id Identifier) (ConnectionEnd, error)

	// Channel Lifecycle Management : ICS 4
	ChanOpenInit(
		order ChannelOrder,
		connectionHops []Identifier,
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterPartyPortIdentifier Identifier,
		counterPartyChannelIdentifier Identifier,
		version string)
	ChanOpenTry(
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
	ChanOpenAck(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterPartyVersion string,
		proofTry CommitmentProof,
		proofHeight int)
	ChanOpenConfirm(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		proofAck CommitmentProof,
		proofHeight int)
	ChanCloseInit(
		portIdentifier Identifier,
		channelIdentifier Identifier)
	ChanCloseConfirm(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		proofInit CommitmentProof,
		proofHeight int)
	QueryChannel()

	// Packet Relay : ICS 4
	SendPacket(packet Packet)
	RecvPacket(
		packet OpaquePacket,
		proof CommitmentProof,
		proofHeight int,
		acknowledgement []byte)
	AcknowledgePacket(
		packet OpaquePacket,
		acknowledgement []byte,
		proof CommitmentProof,
		proofHeight int)
	TimeoutPacket(
		packet OpaquePacket,
		proof CommitmentProof,
		proofHeight int,
		nextSequenceRecv int)
	TimeoutOnClose(
		packet Packet,
		proofNonMembership CommitmentProof,
		proofClosed CommitmentProof,
		proofHeight int)
	CleanupPacket(
		packet OpaquePacket,
		proof CommitmentProof,
		proofHeight int,
		// TODO : modify it
		nextSequenceRecvOrAcknowledgement interface{})
}
