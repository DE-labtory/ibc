package ibc

import "github.com/DE-labtory/ibc/spec"

// Connection is a temporarily implementation ConnectionEnd.
type Connection struct {
	state                           spec.ConnectionState
	counterPartyConnectionIdentifer spec.Identifier
	counterPartyPrefix              spec.CommitmentPrefix
	clientIdentifier                spec.Identifier
	counterPartyClientIdentifer     spec.Identifier
	version                         []string
}

func (c *Connection) getState() spec.ConnectionState {
	return c.state
}

func (c *Connection) getCounterPartyPrefix() spec.CommitmentPrefix {
	return c.counterPartyPrefix
}

func (c *Connection) getCounterPartyConnectionIdentifer() spec.Identifier {
	return c.counterPartyConnectionIdentifer
}

func (c *Connection) getCounterPartyClientIdentifer() spec.Identifier {
	return c.counterPartyClientIdentifer
}

func (c *Connection) getClientIdentifer() spec.Identifier {
	return c.clientIdentifier
}

func (c *Connection) getVersion() []string {
	return c.version
}
