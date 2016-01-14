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

package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

func TestGetConfiguration(t *testing.T) {
	templateFile, err := ioutil.TempFile("", "configuration")
	if err != nil {
		t.Fatal(err)
	}
	defer os.RemoveAll(templateFile.Name())
	data := []byte(`# Enigma configuration file

# Encryption provider
encryption = "gpg"

# Storage backend
backend = "boltdb"

[gpg]
email = "foo.bar@gmail.com"

[kms]
region = "eu-west-1"
keyID = "0123456789-abc-xyz"

[s3]
region = "eu-west-1"
bucket = "enigma"

[boltdb]
file = "/tmp/ut.db"
bucket = "enigma"`)
	err = ioutil.WriteFile(templateFile.Name(), data, 0700)
	if err != nil {
		t.Fatal(err)
	}
	configuration, err := LoadFileConfig(templateFile.Name())
	if err != nil {
		t.Fatalf("Error with configuration: %v", err)
	}
	fmt.Printf("Configuration : %#v\n", configuration)
	if configuration.Backend != "boltdb" {
		t.Fatalf("Configuration backend failed")
	}
	if configuration.Encryption != "gpg" {
		t.Fatalf("Configuration encryption failed")
	}

	// Storage
	if configuration.BoltDB.Bucket != "enigma" ||
		configuration.BoltDB.File != "/tmp/ut.db" {
		t.Fatalf("Configuration BoldDB failed")
	}
	if configuration.S3.Bucket != "enigma" ||
		configuration.S3.Region != "eu-west-1" {
		t.Fatalf("Configuration S3 failed")
	}

	// Encryption
	if configuration.Kms.KeyID != "0123456789-abc-xyz" ||
		configuration.S3.Region != "eu-west-1" {
		t.Fatalf("Configuration KMS failed")
	}
	if configuration.Gpg.Email != "foo.bar@gmail.com" {
		t.Fatalf("Configuration GPG failed")
	}
}
