package ibc

import "github.com/DE-labtory/ibc/spec"

// Client is temporarily implementation of Client interface.
type Client struct {
	identifier     spec.Identifier
	clientType     spec.ClientType
	consensusState spec.ConsensusState
}

func (c *Client) getIdentifier() spec.Identifier {
	return c.identifier
}

func (c *Client) getClientType() spec.ClientType {
	return c.clientType
}

func (c *Client) getConsensusState() spec.ConsensusState {
	return c.consensusState
}
