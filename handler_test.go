package ibc

import (
	"github.com/DE-labtory/ibc/spec"
	"testing"
)

var handler = Handler{}

func TestClientStatePath(t *testing.T) {
	statePath := handler.ClientStatePath(spec.Identifier("testId"))
	if statePath != "clients/testId/state" {
		t.Fatalf("test failed")
	}
}

func TestClientTypePath(t *testing.T) {
	typePath := handler.ClientTypePath(spec.Identifier("testId"))
	if typePath != "clients/testId/type" {
		t.Fatalf("test failed")
	}
}

func TestConsensusStatePath(t *testing.T) {
	consensusPath := handler.ConsensusStatePath(spec.Identifier("testId"))
	if consensusPath != "clients/testId/consensusState" {
		t.Fatalf("test failed")
	}
}
