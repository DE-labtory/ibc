package ibc

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

// HandlerModule is implementation of Handler interface.
type HandlerModule struct {
	privateStore  *PrivateStore
	provableStore *ProvableStore
}

func (h *HandlerModule) CreateClient(id Identifier, clientType ClientType, consensusState ConsensusState) {
	panic("implement me")
}

func (h *HandlerModule) UpdateClient(id Identifier, header Header) {
	panic("implement me")
}

// QueryClientConsensusState provides specific client consensus state using identifier.
func (h *HandlerModule) QueryClientConsensusState(id Identifier) ConsensusState {
	return h.privateStore.get(clientStatePath(id)).(ConsensusState)
}

func (h *HandlerModule) QueryClient(id Identifier) ClientState {
	panic("implement me")
}

func (h *HandlerModule) ConnOpenInit(identifier Identifier,
	desiredCounterpartyConnectionIdentifier Identifier,
	counterpartyPrefix CommitmentPrefix,
	clientIdentifier Identifier,
	counterPartyClientIdentifer Identifier) {
	panic("implement me")
}

func (h *HandlerModule) ConnOpenTry(
	desiredIdentifer Identifier,
	counterPartyConnectionIdentifier Identifier,
	counterPartyPrefix CommitmentPrefix,
	counterPartyClientIdentifier Identifier,
	clientIdentifier Identifier,
	counterPartyVersions []string,
	proofInit CommitmentProof,
	proofHeight int,
	consensusHeight int) {
	panic("implement me")
}

func (h *HandlerModule) ChanOpenAck(
	portIdentifier Identifier,
	channelIdentifier Identifier,
	counterPartyVersion string,
	proofTry CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *HandlerModule) ChanOpenConfirm(
	portIdentifier Identifier,
	channelIdentifier Identifier,
	proofAck CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *HandlerModule) ChanCloseInit(
	portIdentifier Identifier,
	channelIdentifier Identifier) {
	panic("implement me")
}

func (h *HandlerModule) ChanCloseConfirm(
	portIdentifier Identifier,
	channelIdentifier Identifier,
	proofInit CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *HandlerModule) QueryChannel() {
	panic("implement me")
}

func (h *HandlerModule) SendPacket(packet Packet) {
	panic("implement me")
}

func (h *HandlerModule) RecvPacket(
	packet OpaquePacket,
	proof CommitmentProof,
	proofHeight int,
	acknowledgement []byte) {
	panic("implement me")
}

func (h *HandlerModule) AcknowledgePacket(
	packet OpaquePacket,
	acknowledgement []byte,
	proof CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *HandlerModule) TimeoutPacket(
	packet OpaquePacket,
	proof CommitmentProof,
	proofHeight int,
	nextSequenceRecv int) {
	panic("implement me")
}

func (h *HandlerModule) TimeoutOnClose(
	packet Packet,
	proofNonMembership CommitmentProof,
	proofClosed CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *HandlerModule) CleanupPacket(
	packet OpaquePacket,
	proof CommitmentProof,
	proofHeight int,
	// TODO : modify it
	nextSequenceRecvOrAcknowledgement interface{}) {
	panic("implement me")
}
