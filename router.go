package ibc

import "github.com/DE-labtory/ibc/spec"

type Router struct {
}

func (r *Router) onChanOpenInit(
	order spec.ChannelOrder,
	connectionHops []spec.Identifier,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	counterpartyPortIdentifier spec.Identifier,
	counterpartyChannelIdentifer spec.Identifier,
	version string) {
	panic("implement me")
}

func (r *Router) onChanOpenTry(
	order spec.ChannelOrder,
	connectionHops []spec.Identifier,
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	counterpartyPortIdentifier spec.Identifier,
	counterpartyChannelIdentifer spec.Identifier,
	version string) {
	panic("implement me")
}

func (r *Router) onOpenAck(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier,
	version string) {
	panic("implement me")
}

func (r *Router) onChanOpenConfirm(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier) {
	panic("implement me")
}

func (r *Router) onChanCloseInit(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier) {
	panic("implement me")
}

func (r *Router) onChanCloseConfirm(
	portIdentifier spec.Identifier,
	channelIdentifier spec.Identifier) {
	panic("implement me")
}

func (r *Router) onRecvPacket(packet interface{}) []byte {
	panic("implement me")
}

func (r *Router) onTimeoutPacket(packet interface{}) {
	panic("implement me")
}

func (r *Router) onAcknowledgePacket(packet interface{}) {
	panic("implement me")
}

func (r *Router) onTimeoutPacketClose(packet interface{}) {
	panic("implement me")
}
