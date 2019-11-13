package spec

type PublicKey struct {
}

func (pk *PublicKey) Verify(s Signature) bool {
	panic("implement me!")
}

type Signature struct {
}
