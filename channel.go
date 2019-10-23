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
	GetState() ChannelState
	GetOrdering() ChannelOrder
	GetCounterPartyPortIdentifier() Identifier
	GetCounterPartyChannelIdentifer() Identifier
	GetConnectionHops() []Identifier
	GetVersion() string
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

func (c *ChannelImpl) GetState() ChannelState {
	return c.state
}

func (c *ChannelImpl) GetOrdering() ChannelOrder {
	return c.ordering
}

func (c *ChannelImpl) GetCounterPartyPortIdentifier() Identifier {
	return c.counterPartyPortIdentifier
}

func (c *ChannelImpl) GetCounterPartyChannelIdentifer() Identifier {
	return c.counterPartyChannelIdentifier
}

func (c *ChannelImpl) GetConnectionHops() []Identifier {
	return c.connectionHops
}

func (c *ChannelImpl) GetVersion() string {
	return c.version
}
