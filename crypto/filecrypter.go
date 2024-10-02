package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"io"
	"os"

	"golang.org/x/crypto/pbkdf2"
)

func FileEncrypter(src string, dst string, key []byte) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	plaintext, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	nonce := make([]byte, 12)
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return err
	}

	block, err := aes.NewCipher(pbkdf2.Key(key, nonce, 4096, 32, sha1.New))
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	ciphertext = append(ciphertext, nonce...)

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err = io.Copy(f, bytes.NewReader(ciphertext)); err != nil {
		return err
	}

	return nil
}

func FileDecrypter(src string, dst string, key []byte) error {
	if _, err := os.Stat(src); os.IsNotExist(err) {
		return err
	}

	ciphertext, err := os.ReadFile(src)
	if err != nil {
		return err
	}

	nonce, err := hex.DecodeString(hex.EncodeToString(ciphertext[len(ciphertext)-12:]))
	if err != nil {
		return err
	}

	block, err := aes.NewCipher(pbkdf2.Key(key, nonce, 4096, 32, sha1.New))
	if err != nil {
		return err
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		return err
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext[:len(ciphertext)-12], nil)
	if err != nil {
		return err
	}

	f, err := os.Create(dst)
	if err != nil {
		return err
	}
	if _, err = io.Copy(f, bytes.NewReader(plaintext)); err != nil {
		return err
	}

	return nil
}
