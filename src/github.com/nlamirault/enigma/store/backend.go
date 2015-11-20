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

package store

import (
	"errors"
	"sort"
)

var (
	backends                     = make(map[string]func() (StorageBackend, error))
	errUnsupportedStorageBackend = errors.New("Unsupported storage backend")
)

// StorageBackend represents a storage backend
type StorageBackend interface {

	// Name identify the key manager
	Name() string

	// Put a value at the specified key
	Put(bucket string, key []byte, value []byte) error

	// Get a value given its key
	Get(bucket string, key []byte) ([]byte, error)

	// Delete a value given its key
	Delete(bucket string, key []byte) error

	// List values
	List(bucket string) ([]string, error)
}

// New returns a new storage backend using the label
func New(label string) (StorageBackend, error) {
	if constructor, present := backends[label]; present {
		return constructor()
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
