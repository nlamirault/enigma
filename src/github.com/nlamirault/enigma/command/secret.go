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
	"flag"
	"fmt"
	"strings"

	"github.com/mitchellh/cli"
)

// SecretCommand defines the CLI command to manage secrets
type SecretCommand struct {
	UI cli.Ui
}

// Help display help message about the command
func (c *SecretCommand) Help() string {
	helpText := `
Usage: enigma secret [options] action
	Manage secrets

Options:
        ` + generalOptionsUsage() + `

Secret options:
        --text=text                   Text to encrypt
        --key=key                     Key for text

Action :
        list                     List all secrets
        put                      Put a secret
        get                      Retrieve a secret
        delete                   Delete a secret
`
	return strings.TrimSpace(helpText)
}

// Synopsis return the command message
func (c *SecretCommand) Synopsis() string {
	return "Manage secrets from Amazon S3"
}

// Run launch the command
func (c *SecretCommand) Run(args []string) int {
	var debug bool
	var key, text, config string
	f := flag.NewFlagSet("secrets", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }

	defaultConfigFile := getConfigurationFile()

	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&config, "config", defaultConfigFile, "Configuration filename")
	f.StringVar(&key, "key", "", "Key for store data")
	f.StringVar(&text, "text", "", "Text to store")

	if err := f.Parse(args); err != nil {
		return 1
	}
	args = f.Args()
	if len(args) != 1 {
		f.Usage()
		return 1
	}
	setLogging(debug)

	client, err := NewClient(config)
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	action := args[0]
	//fmt.Printf("Action: %s\n", action)

	switch action {
	case "list":
		c.doList(client)
	case "delete":
		valid := checkArguments(key)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: key."))
			return 1
		}
		c.doDelete(client, key)
	case "put":
		valid := checkArguments(key, text)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: key and text."))
			return 1
		}
		c.doPutText(client, key, text)
	case "get":
		valid := checkArguments(key)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: key."))
			return 1
		}
		c.doGetText(client, key)
	default:
		f.Usage()
	}
	return 0
}

func (c *SecretCommand) doGetText(client *Client, key string) {
	c.UI.Info(fmt.Sprintf("Retrive secret text for key : %s", key))
	blob, err := client.Storage.Get([]byte(key))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	if len(blob) == 0 {
		c.UI.Output(fmt.Sprintf("No value with key %s", key))
		return
	}
	decrypted, err := client.Keys.Decrypt(blob)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Decrypted: %s", string(decrypted)))
}

func (c *SecretCommand) doPutText(client *Client, key string, text string) {
	c.UI.Info(fmt.Sprintf("Store secret text %s with key %s", text, key))
	output, err := client.Keys.Encrypt([]byte(text))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}

	err = client.Storage.Put([]byte(key), output)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Successfully uploaded data with key %s", key))
}

func (c *SecretCommand) doDelete(client *Client, key string) {
	c.UI.Info(fmt.Sprintf("Delete secret with key %s", key))
	client.Storage.Delete([]byte(key))
	c.UI.Output("Deleted")
}

func (c *SecretCommand) doList(client *Client) {
	c.UI.Info(fmt.Sprintf("List secrets :"))
	secrets, err := client.Storage.List()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	for _, key := range secrets {
		c.UI.Output(fmt.Sprintf("- %s", key))
	}
}
