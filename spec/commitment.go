package spec

type CommitmentState interface{}
type CommitmentRoot interface {
	VerifyMemebership(path Path, value Value, proof CommitmentProof) bool
	VerifyNonMembership(path Path, proof CommitmentProof) bool
}
type CommitmentPath interface{}
type CommitmentPrefix interface{}
type CommitmentProof interface{}
