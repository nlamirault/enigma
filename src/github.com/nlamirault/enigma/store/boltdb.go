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
	"fmt"
	"log"
	"os"
	"path"

	"github.com/boltdb/bolt"

	"github.com/nlamirault/enigma/config"
)

const (
	// Bucket defines the Enigma bucket
	// Bucket      = "enigma"
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
func NewBoltDB(conf *config.Configuration) (StorageBackend, error) {
	//f := filepath.Join(conf.BoltDB.Directory, "enigma.db")
	f := conf.BoltDB.File
	directory := path.Dir(f)
	if _, err := os.Stat(directory); os.IsNotExist(err) {
		log.Printf("Create %s", directory)
		err = os.Mkdir(directory, 0755)
		if err != nil {
			return nil, err
		}
	}
	db, err := bolt.Open(f, 0600, nil)
	if err != nil {
		return nil, err
	}

	return &BoltDB{
		DB:         db,
		Path:       conf.BoltDB.File,
		BucketName: conf.BoltDB.Bucket,
	}, nil
}

// Name returns BoltDB label
func (db *BoltDB) Name() string {
	return boltdbLabel
}

// Create intialize the storage backend
func (db *BoltDB) Create() error {
	return db.Update(func(tx *bolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(db.BucketName))
		if err != nil {
			return fmt.Errorf("BoltDB Can't create bucket: %s", err)
		}
		return nil
	})
}

// Destroy remove the storage backend bucket
func (db *BoltDB) Destroy() error {
	directory := path.Dir(db.BucketName)
	log.Printf("Create %s", directory)
	return nil
}

// List returns all secrets
func (db *BoltDB) List() ([]string, error) {
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
func (db *BoltDB) Get(key []byte) ([]byte, error) {
	log.Printf("[DEBUG] BoltDB Search entry with key : %v", string(key))
	var value []byte
	db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.ForEach(func(k, v []byte) error {
			//log.Printf("[BoltDB] Entry : %s %s", string(k), string(v))
			if string(k) == string(key) {
				//log.Printf("[DEBUG] Find : %s", string(v))
				value = v
				return nil
			}
			return nil
		})
		return nil
	})
	return value, nil
}

// Put a value at the specified key
func (db *BoltDB) Put(key []byte, value []byte) error {
	log.Printf("[DEBUG] BoltDB Put : %v %v", string(key), string(value))
	return db.DB.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket([]byte(db.BucketName))
		b.Put(key, value)
		return nil
	})
}

// Delete the value at the specified key
func (db *BoltDB) Delete(key []byte) error {
	log.Printf("[DEBUG] BoltDB Delete : %v", string(key))
	return fmt.Errorf("Not implemented")
}
