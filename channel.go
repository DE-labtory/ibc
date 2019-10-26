package ibc

import "github.com/DE-labtory/ibc/spec"

// Channel is a data structure on one chain storing channel metadata. (temporarily)
type Channel struct {
	state                         spec.ChannelState
	ordering                      spec.ChannelOrder
	counterPartyPortIdentifier    spec.Identifier
	counterPartyChannelIdentifier spec.Identifier
	connectionHops                []spec.Identifier
	version                       string
}

func (c *Channel) getState() spec.ChannelState {
	return c.state
}

func (c *Channel) getOrdering() spec.ChannelOrder {
	return c.ordering
}

func (c *Channel) getCounterPartyPortIdentifier() spec.Identifier {
	return c.counterPartyPortIdentifier
}

func (c *Channel) getCounterPartyChannelIdentifer() spec.Identifier {
	return c.counterPartyChannelIdentifier
}

func (c *Channel) getConnectionHops() []spec.Identifier {
	return c.connectionHops
}

func (c *Channel) getVersion() string {
	return c.version
}
