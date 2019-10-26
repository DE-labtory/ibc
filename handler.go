package ibc

import "github.com/DE-labtory/ibc/spec"

// Handler is implementation of Handler interface.
type Handler struct {
	privateStore  *PrivateStore
	provableStore *ProvableStore
}

func (h *Handler) CreateClient(id spec.Identifier, clientType spec.ClientType, consensusState spec.ConsensusState) {
	panic("implement me")
}

func (h *Handler) UpdateClient(id spec.Identifier, header spec.Header) {
	panic("implement me")
}

// QueryClientConsensusState provides specific client consensus state using identifier.
func (h *Handler) QueryClientConsensusState(id spec.Identifier) spec.ConsensusState {
	return h.privateStore.get(spec.ClientStatePath(id)).(spec.ConsensusState)
}

func (h *Handler) QueryClient(id spec.Identifier) spec.ClientState {
	panic("implement me")
}

func (h *Handler) ConnOpenInit(identifier spec.Identifier,
	desiredCounterpartyConnectionIdentifier spec.Identifier,
	counterpartyPrefix spec.CommitmentPrefix,
	clientIdentifier spec.Identifier,
	counterPartyClientIdentifer spec.Identifier) {
	panic("implement me")
}

func (h *Handler) ConnOpenTry(
	desiredIdentifer spec.Identifier,
	counterPartyConnectionIdentifier spec.Identifier,
	counterPartyPrefix spec.CommitmentPrefix,
	counterPartyClientIdentifier spec.Identifier,
	clientIdentifier spec.Identifier,
	counterPartyVersions []string,
	proofInit spec.CommitmentProof,
	proofHeight int,
	consensusHeight int) {
	panic("implement me")
}

func (h *Handler) ChanOpenAck(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	counterPartyVersion string,
	proofTry spec.CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *Handler) ChanOpenConfirm(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	proofAck spec.CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *Handler) ChanCloseInit(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier) {
	panic("implement me")
}

func (h *Handler) ChanCloseConfirm(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	proofInit spec.CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *Handler) QueryChannel() {
	panic("implement me")
}

func (h *Handler) SendPacket(packet spec.Packet) {
	panic("implement me")
}

func (h *Handler) RecvPacket(
	packet spec.OpaquePacket,
	proof spec.CommitmentProof,
	proofHeight int,
	acknowledgement []byte) {
	panic("implement me")
}

func (h *Handler) AcknowledgePacket(
	packet spec.OpaquePacket,
	acknowledgement []byte,
	proof spec.CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *Handler) TimeoutPacket(
	packet spec.OpaquePacket,
	proof spec.CommitmentProof,
	proofHeight int,
	nextSequenceRecv int) {
	panic("implement me")
}

func (h *Handler) TimeoutOnClose(
	packet spec.Packet,
	proofNonMembership spec.CommitmentProof,
	proofClosed spec.CommitmentProof,
	proofHeight int) {
	panic("implement me")
}

func (h *Handler) CleanupPacket(
	packet spec.OpaquePacket,
	proof spec.CommitmentProof,
	proofHeight int,
	// TODO : modify it
	nextSequenceRecvOrAcknowledgement interface{}) {
	panic("implement me")
}
