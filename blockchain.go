package ibc

import "github.com/DE-labtory/ibc/spec"

type Blockchain struct {
	genesis   spec.ConsensusState
	consensus spec.Consensus
}
