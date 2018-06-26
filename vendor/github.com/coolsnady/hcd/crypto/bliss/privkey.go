package bliss

import (
	dcrcrypto "github.com/coolsnady/hcd/crypto"
	"github.com/coolsnady/bliss"
)

type PrivateKey struct {
	dcrcrypto.PrivateKeyAdapter
	bliss.PrivateKey
}

// Public returns the PublicKey corresponding to this private key.
func (p PrivateKey) PublicKey() dcrcrypto.PublicKey {
	blissPkp := p.PrivateKey.PublicKey()
	pk := &PublicKey{
		PublicKey: *blissPkp,
	}
	return pk
}

// GetType satisfies the bliss PrivateKey interface.
func (p PrivateKey) GetType() int {
	return pqcTypeBliss
}

func (p PrivateKey) Serialize() []byte {
	return p.PrivateKey.Serialize()
}
