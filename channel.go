package ibc

// ChannelState is the current state of the channel end.
type ChannelState int

const (
	// CHANINIT state just started the opening handshake.
	CHANINIT ChannelState = 0

	// CHANOPENTRY state has acknowledged the handshake step on the counterparty chain.
	CHANOPENTRY

	// CHANOPEN state has completed the handshake and is ready to send and receive packets.
	CHANOPEN

	// CHANCLOSED state has been closed and can no longer be used to send or receive packets.
	CHANCLOSED
)

// ChannelOrder means whether packets are ordered or not.
type ChannelOrder int

const (
	// ORDERED channel is a channel where packets are delivered exactly in the order which they were sent.
	ORDERED ChannelOrder = 0

	// UNORDERED channel is a channel where packets can be delivered in any order.
	UNORDERED
)

type Channel interface {
	getState() ChannelState
	getOrdering() ChannelOrder
	getCounterPartyPortIdentifier() Identifier
	getCounterPartyChannelIdentifer() Identifier
	getConnectionHops() []Identifier
	getVersion() string
}

// ChannelImpl is a data structure on one chain storing channel metadata. (temporarily)
type ChannelImpl struct {
	state                         ChannelState
	ordering                      ChannelOrder
	counterPartyPortIdentifier    Identifier
	counterPartyChannelIdentifier Identifier
	connectionHops                []Identifier
	version                       string
}

func (c *ChannelImpl) getState() ChannelState {
	return c.state
}

func (c *ChannelImpl) getOrdering() ChannelOrder {
	return c.ordering
}

func (c *ChannelImpl) getCounterPartyPortIdentifier() Identifier {
	return c.counterPartyPortIdentifier
}

func (c *ChannelImpl) getCounterPartyChannelIdentifer() Identifier {
	return c.counterPartyChannelIdentifier
}

func (c *ChannelImpl) getConnectionHops() []Identifier {
	return c.connectionHops
}

func (c *ChannelImpl) getVersion() string {
	return c.version
}
