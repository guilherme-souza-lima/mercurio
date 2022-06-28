package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"ssMercurio/entities"
)

type CryptoPassword struct {
	Key string
}

func NewCryptoPassword(Key string) CryptoPassword {
	return CryptoPassword{Key}
}

func (c CryptoPassword) Encrypt(password string) (string, error) {
	key := []byte(entities.PASSWORD_CRYPTO)
	plaintext := []byte(password)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	nonce := []byte(entities.NONCE_CRYPTO)
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	encodedString := hex.EncodeToString(ciphertext)
	return encodedString, nil
}

func (c CryptoPassword) Decrypt(crypt string) (string, error) {
	key := []byte(entities.PASSWORD_CRYPTO)
	ciphertext, _ := hex.DecodeString(crypt)
	nonce := []byte(entities.NONCE_CRYPTO)
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}
	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	password := string(plaintext)
	return password, nil
}
