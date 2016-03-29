// Copyright (C) 2015, 2016 Nicolas Lamirault <nicolas.lamirault@gmail.com>

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
	"errors"
	"sort"

	"github.com/nlamirault/enigma/config"
)

var (
	registry                 = make(map[string]func(*config.Configuration) (KeyManager, error))
	errUnsupportedKeyManager = errors.New("Unsupported key manager")
)

// KeyManager represents a service that can generate envelope keys and provide decryption
// keys.
type KeyManager interface {

	// Name identify the key manager
	Name() string

	// Encrypt encrypt the data
	Encrypt(plaintext []byte) ([]byte, error)

	// Decrypt the secret
	Decrypt(ciphertext []byte) ([]byte, error)
}

// New returns a KeyManager
func New(conf *config.Configuration) (KeyManager, error) {
	if constructor, present := registry[conf.Encryption]; present {
		return constructor(conf)
	}
	return nil, errUnsupportedKeyManager
}

// GetKeyManagers returns a list of registered key managers.
func GetKeyManagers() []string {
	var collector []string
	for k := range registry {
		collector = append(collector, k)
	}
	sort.Strings(collector)
	return collector
}

// Envelope represents the structure used in envelope encryption.
type Envelope struct {
	Ciphertext   []byte
	EncryptedKey []byte
	Nonce        []byte
}
