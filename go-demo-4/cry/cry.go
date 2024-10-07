package cry

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

type Crypto struct {
	Key string
}

func NewCrypto(key string) *Crypto {
	return &Crypto{
		Key: key,
	}
}

func (cry *Crypto) Encrypt(str []byte) []byte {
	block, err := aes.NewCipher([]byte(cry.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonce := make([]byte, aesGSM.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		panic(err.Error())
	}
	return aesGSM.Seal(nonce, nonce, str, nil)
}

func (cry *Crypto) Decrypt(encStr []byte) []byte {
	block, err := aes.NewCipher([]byte(cry.Key))
	if err != nil {
		panic(err.Error())
	}
	aesGSM, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}
	nonceSize := aesGSM.NonceSize()
	nonce, cipherText := encStr[:nonceSize], encStr[nonceSize:]
	str, err := aesGSM.Open(nil, nonce, cipherText, nil)
	if err != nil {
		panic(err.Error())
	}
	return str
}
