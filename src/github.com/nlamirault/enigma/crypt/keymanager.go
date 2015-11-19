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

package crypt

import (
	"errors"
	"sort"
)

var (
	registry                 = make(map[string]func() KeyManager)
	errUnsupportedKeyManager = errors.New("Unsupported key manager")
)

// KeyManager represents a service that can generate envelope keys and provide decryption
// keys.
type KeyManager interface {

	// Name identify the key manager
	Name() string

	// Encrypt encrypt the data
	Encrypt(key string, text []byte) ([]byte, error)

	// Decrypt the secret
	Decrypt(key string, enc []byte) ([]byte, error)
}

// New returns a KeyManager
func New(label string) (KeyManager, error) {
	if constructor, present := registry[label]; present {
		return constructor(), nil
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
