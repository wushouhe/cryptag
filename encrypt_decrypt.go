// Steve Phillips / elimisteve
// 2015.08.05

package cryptag

import (
	"crypto/rand"
	"fmt"

	"golang.org/x/crypto/nacl/secretbox"
)

const (
	validKeyLength   = 32
	validNonceLength = 24
)

var (
	ErrDecrypt      = fmt.Errorf("Error decrypting ciphertext")
	ErrDecryptEmpty = fmt.Errorf("Error decrypting empty ciphertext")
	ErrInvalidKey   = fmt.Errorf("Invalid key")
	ErrNilKey       = fmt.Errorf("Nil key")
	ErrNilNonce     = fmt.Errorf("Nil nonce")
	ErrInvalidNonce = fmt.Errorf("Invalid nonce")
)

func Encrypt(plain []byte, nonce *[24]byte, key *[32]byte) ([]byte, error) {
	if nonce == nil {
		return nil, ErrNilNonce
	}
	if key == nil {
		return nil, ErrNilKey
	}

	cipher := secretbox.Seal(nil, plain, nonce, key)
	return cipher, nil
}

func Decrypt(cipher []byte, nonce *[24]byte, key *[32]byte) ([]byte, error) {
	if nonce == nil {
		return nil, ErrNilNonce
	}
	if key == nil {
		return nil, ErrNilKey
	}
	if len(cipher) == 0 {
		return nil, ErrDecryptEmpty
	}

	plain, ok := secretbox.Open(nil, cipher, nonce, key)
	if !ok {
		return nil, ErrDecrypt
	}
	return plain, nil
}

func ConvertKey(key []byte) (goodKey *[32]byte, err error) {
	if len(key) != validKeyLength {
		return nil, fmt.Errorf("Invalid key; must be of length %d, has length %d",
			validKeyLength, len(key))
	}

	// []byte -> *[32]byte
	var good [validKeyLength]byte
	copy(good[:], key)

	return &good, nil
}

func UnconvertKey(goodKey *[32]byte) ([]byte, error) {
	if goodKey == nil {
		return nil, ErrNilKey
	}
	return (*goodKey)[:], nil
}

func ConvertNonce(nonce []byte) (goodNonce *[24]byte, err error) {
	if len(nonce) != validNonceLength {
		return nil, ErrInvalidNonce
	}
	var b [24]byte
	copy(b[:], nonce[:])
	return &b, nil
}

func RandomNonce() (*[24]byte, error) {
	var b [24]byte
	_, err := rand.Reader.Read(b[:])
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func RandomKey() (*[32]byte, error) {
	var b [32]byte
	_, err := rand.Reader.Read(b[:])
	if err != nil {
		return nil, err
	}
	return &b, nil
}

func RandomKeySlice() ([]byte, error) {
	k, err := RandomKey()
	if err != nil {
		return nil, err
	}
	return (*k)[:], nil
}
