package ibc

// Router defines callback used by IBC handler.
type Router interface {
	onChanOpenInit(
		order ChannelOrder,
		connectionHops []Identifier,
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterpartyPortIdentifier Identifier,
		counterpartyChannelIdentifer Identifier,
		version string)

	onChanOpenTry(
		order ChannelOrder,
		connectionHops []Identifier,
		portIdentifier Identifier,
		channelIdentifier Identifier,
		counterpartyPortIdentifier Identifier,
		counterpartyChannelIdentifer Identifier,
		version string)

	onOpenAck(
		portIdentifier Identifier,
		channelIdentifier Identifier,
		version string)

	onChanOpenConfirm(
		portIdentifier Identifier,
		channelIdentifier Identifier)

	onChanCloseInit(
		portIdentifier Identifier,
		channelIdentifier Identifier)

	onChanCloseConfirm(
		portIdentifier Identifier,
		channelIdentifier Identifier)

	onRecvPacket(packet interface{}) []byte

	onTimeoutPacket(packet interface{})

	onAcknowledgePacket(packet interface{})

	onTimeoutPacketClose(packet interface{})
}

type RouterModule struct {
}

func (r *RouterModule) onChanOpenInit(
	order ChannelOrder,
	connectionHops []Identifier,
	portIdentifier Identifier,
	channelIdentifier Identifier,
	counterpartyPortIdentifier Identifier,
	counterpartyChannelIdentifer Identifier,
	version string) {
	panic("implement me")
}

func (r *RouterModule) onChanOpenTry(
	order ChannelOrder,
	connectionHops []Identifier,
	portIdentifier Identifier,
	channelIdentifier Identifier,
	counterpartyPortIdentifier Identifier,
	counterpartyChannelIdentifer Identifier,
	version string) {
	panic("implement me")
}

func (r *RouterModule) onOpenAck(
	portIdentifier Identifier,
	channelIdentifier Identifier,
	version string) {
	panic("implement me")
}

func (r *RouterModule) onChanOpenConfirm(
	portIdentifier Identifier,
	channelIdentifier Identifier) {
	panic("implement me")
}

func (r *RouterModule) onChanCloseInit(
	portIdentifier Identifier,
	channelIdentifier Identifier) {
	panic("implement me")
}

func (r *RouterModule) onChanCloseConfirm(
	portIdentifier Identifier,
	channelIdentifier Identifier) {
	panic("implement me")
}

func (r *RouterModule) onRecvPacket(packet interface{}) []byte {
	panic("implement me")
}

func (r *RouterModule) onTimeoutPacket(packet interface{}) {
	panic("implement me")
}

func (r *RouterModule) onAcknowledgePacket(packet interface{}) {
	panic("implement me")
}

func (r *RouterModule) onTimeoutPacketClose(packet interface{}) {
	panic("implement me")
}
