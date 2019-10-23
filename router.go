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
