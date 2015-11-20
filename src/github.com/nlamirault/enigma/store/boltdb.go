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
	"fmt"
	"log"

	"github.com/boltdb/bolt"
)

const (
	// Bucket defines the Enigma bucket
	Bucket      = "enigma"
	boltdbLabel = "boltdb"
)

func init() {
	backends[boltdbLabel] = NewBoltDB
}

// BoltDB is the Boltdb backend.
type BoltDB struct {
	*bolt.DB
	BucketName string
	Path       string
}

// NewBoltDB opens a new BoltDB connection to the specified path and bucket
func NewBoltDB() (StorageBackend, error) {
	path := fmt.Sprintf("/tmp/%s", Bucket)
	log.Printf("[DEBUG] Init BoltDB storage : %v", path)
	db, err := bolt.Open(path, 0600, nil)
	if err != nil {
		return nil, err
	}
	db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(Bucket))
		if err != nil {
			return fmt.Errorf("Can't create BoltDB bucket: %s", err)
		}
		return nil
	})
	return &BoltDB{
		DB:         db,
		Path:       path,
		BucketName: Bucket,
	}, nil
}

// Name returns BoltDB label
func (db *BoltDB) Name() string {
	return boltdbLabel
}

// List returns all secrets
func (db *BoltDB) List(bucket string) ([]string, error) {
	log.Printf("[DEBUG] BoltDB list secrets")
	var l []string
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.ForEach(func(key, value []byte) error {
			//log.Println(string(key), string(value))
			l = append(l, string(key))
			return nil
		})
		return nil
	})
	return l, nil
}

// Get a value given its key
func (db *BoltDB) Get(bucket string, key []byte) ([]byte, error) {
	log.Printf("[DEBUG] Search entry with key : %v", string(key))
	var value []byte
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.ForEach(func(k, v []byte) error {
			// log.Printf("[BoltDB] Entry : %s %s", string(k), string(v))
			if string(k) == string(key) {
				//log.Printf("[DEBUG] Find : %s", string(v))
				value = v
			}
			return nil
		})
		return nil
	})
	return value, nil
}

// Put a value at the specified key
func (db *BoltDB) Put(bucket string, key []byte, value []byte) error {
	log.Printf("[DEBUG] Put : %v %v", string(key), string(value))
	return db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.Put(key, value)
		return nil
	})
}

// Delete the value at the specified key
func (db *BoltDB) Delete(bucket string, key []byte) error {
	log.Printf("[DEBUG] Delete : %v", string(key))
	return fmt.Errorf("Not implemented")
}