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

package command

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"

	"github.com/nlamirault/enigma/crypt"
	"github.com/nlamirault/enigma/logging"
	"github.com/nlamirault/enigma/store"
)

// generalOptionsUsage returns the usage documenation for commonly
// available options
func generalOptionsUsage() string {
	general := `
        --debug                       Debug mode enabled
	--bucket=name                 Bucket name
        --region=name                 Region name
`
	return strings.TrimSpace(general)
}

func checkArguments(args ...string) bool {
	for _, arg := range args {
		if len(arg) == 0 {
			return false
		}
	}
	return true
}

func getAWSConfig(region string, debug bool) *aws.Config {
	if debug {
		logging.SetLogging("DEBUG")
		return &aws.Config{
			Region:   aws.String(region),
			LogLevel: aws.LogLevel(aws.LogDebugWithHTTPBody),
		}
	}
	logging.SetLogging("INFO")
	return &aws.Config{
		Region: aws.String(region),
	}
}

// Client provides a keys manager and storage backend
type Client struct {
	Keys    crypt.KeyManager
	Storage store.StorageBackend
}

// NewClient creates a new instance of Client.
func NewClient(keysLabel string, storageLabel string) (*Client, error) {
	manager, err := crypt.New("gpg")
	if err != nil {
		return nil, err
	}
	storage, err := store.New("s3")
	if err != nil {
		return nil, err
	}
	return &Client{
		Keys:    manager,
		Storage: storage,
	}, nil
}
