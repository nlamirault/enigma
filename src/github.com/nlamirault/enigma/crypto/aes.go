// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"

	"github.com/nlamirault/enigma/config"
)

const (
	aesLabel     = "aes"
	aesKeyLength = 32
)

func init() {
	registry[aesLabel] = NewAes
}

// Aes is a KeyManager using AES encrypt
type Aes struct {
	Key string
}

// NewAes returns a new Aes.
func NewAes(conf *config.Configuration) (KeyManager, error) {
	return &Aes{
		Key: conf.Aes.Key,
	}, nil
}

// Name returns kmsLabel
func (a *Aes) Name() string {
	return kmsLabel
}

// Decrypt decrypts the encrypted key.
func (a *Aes) Decrypt(text []byte) ([]byte, error) {
	ciphertext := []byte(text)
	key := []byte(a.Key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, fmt.Errorf("Text is too short")
	}

	// Get the 16 byte IV
	iv := ciphertext[:aes.BlockSize]

	// Remove the IV from the ciphertext
	ciphertext = ciphertext[aes.BlockSize:]

	// Return a decrypted stream
	stream := cipher.NewCFBDecrypter(block, iv)

	// Decrypt bytes from ciphertext
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}

// Encrypt encrypt the text using a plaintext key
func (a *Aes) Encrypt(text []byte) ([]byte, error) {
	plaintext := []byte(text)
	key := []byte(a.Key)

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))

	iv := ciphertext[:aes.BlockSize]

	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	// Return an encrypted stream
	stream := cipher.NewCFBEncrypter(block, iv)

	// Encrypt bytes from plaintext to ciphertext
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}
