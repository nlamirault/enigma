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

package command

import (
	"path/filepath"
	"strings"

	"github.com/docker/docker/pkg/homedir"

	"github.com/nlamirault/enigma/config"
	"github.com/nlamirault/enigma/crypto"
	"github.com/nlamirault/enigma/logging"
	"github.com/nlamirault/enigma/store"
)

// generalOptionsUsage returns the usage documenation for commonly
// available options
func generalOptionsUsage() string {
	general := `
        --debug                       Debug mode enabled
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

func setLogging(debug bool) {
	if debug {
		logging.SetLogging("DEBUG")
	} else {
		logging.SetLogging("INFO")
	}
}

func getConfigurationFile() string {
	home := homedir.Get()
	return filepath.Join(home, ".config/enigma/enigma.toml")
}

// Client provides a keys manager and storage backend
type Client struct {
	Keys    crypto.KeyManager
	Storage store.StorageBackend
}

// NewClient creates a new instance of Client.
func NewClient(filename string) (*Client, error) {
	conf, err := config.LoadFileConfig(filename)
	if err != nil {
		return nil, err
	}
	manager, err := crypto.New(conf)
	if err != nil {
		return nil, err
	}
	storage, err := store.New(conf)
	if err != nil {
		return nil, err
	}
	return &Client{
		Keys:    manager,
		Storage: storage,
	}, nil
}
