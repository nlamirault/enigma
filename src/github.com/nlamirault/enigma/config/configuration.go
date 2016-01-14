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
	//"log"

	"github.com/BurntSushi/toml"
)

// Configuration holds configuration for Enigma.
type Configuration struct {
	Encryption string
	Backend    string

	S3     *S3Configuration
	BoltDB *BoltDBConfiguration

	Kms *KmsConfiguration
	Gpg *GpgConfiguration
	Aes *AesConfiguration
}

// New returns a Configuration with default values
func New() *Configuration {
	return &Configuration{
		Backend:    "boltdb",
		Encryption: "gpg",
		Gpg:        &GpgConfiguration{},
		BoltDB:     &BoltDBConfiguration{},
		Aes:        &AesConfiguration{},
	}
}

// LoadFileConfig returns a Configuration from reading the specified file (a toml file).
func LoadFileConfig(file string) (*Configuration, error) {
	configuration := New()
	if _, err := toml.DecodeFile(file, configuration); err != nil {
		return nil, err
	}
	return configuration, nil
}

// GpgConfiguration defines the configuration for Gpg provider
type GpgConfiguration struct {
	Email string
}

// KmsConfiguration defines the configuration for AWS KMS provider
type KmsConfiguration struct {
	Region string
	KeyID  string
}

// AesConfiguration defines the configuration for AES provider
type AesConfiguration struct {
	Key string
}

// BoltDBConfiguration defines the configuration for BoltDB storage backend
type BoltDBConfiguration struct {
	Bucket string
	File   string
}

// S3Configuration defines the configuration for S3 storage backend
type S3Configuration struct {
	Region string
	Bucket string
}
