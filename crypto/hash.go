package crypto

import "crypto/sha256"

type Hashable interface {
	GetHashData() map[string]interface{}
}

func SHA256(data []byte) [32]byte {
	return sha256.Sum256(data)
}
