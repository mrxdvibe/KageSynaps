package crypto

import (
	"crypto/ecdh"
	"crypto/rand"
	"crypto/sha256"
)

func GenerateEphemeralKeyPair() (*ecdh.PrivateKey, []byte, error) {
	curve := ecdh.P256()
	privKey, err := curve.GenerateKey(rand.Reader)
	if err != nil {
		return nil, nil, err
	}
	return privKey, privKey.PublicKey().Bytes(), nil
}

func DeriveSharedSecret(priv *ecdh.PrivateKey, pubBytes []byte) ([]byte, error) {
	curve := ecdh.P256()
	pubKey, err := curve.NewPublicKey(pubBytes)
	if err != nil {
		return nil, err
	}

	secret, err := priv.ECDH(pubKey)
	if err != nil {
		return nil, err
	}

	hash := sha256.Sum256(secret)
	return hash[:], nil
}
