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

package store

import (
	"errors"
	"sort"

	"github.com/nlamirault/enigma/config"
)

var (
	backends                     = make(map[string]func(*config.Configuration) (StorageBackend, error))
	errUnsupportedStorageBackend = errors.New("Unsupported storage backend")
)

// StorageBackend represents a storage backend
type StorageBackend interface {

	// Name identify the key manager
	Name() string

	// Put a value at the specified key
	Put(key []byte, value []byte) error

	// Get a value given its key
	Get(key []byte) ([]byte, error)

	// Delete a value given its key
	Delete(key []byte) error

	// List values
	List() ([]string, error)
}

// New returns a new storage backend using the label
func New(conf *config.Configuration) (StorageBackend, error) {
	if constructor, present := backends[conf.Backend]; present {
		return constructor(conf)
	}
	return nil, errUnsupportedStorageBackend
}

// GetBackends returns a list of registered storage backends
func GetBackends() []string {
	var collector []string
	for k := range backends {
		collector = append(collector, k)
	}
	sort.Strings(collector)
	return collector
}
