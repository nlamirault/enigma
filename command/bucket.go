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

// BucketCommand defines the CLI command to manage buckets
type BucketCommand struct {
	UI cli.Ui
}

// Help display help message about the command
func (c *BucketCommand) Help() string {
	helpText := `
Usage: enigma bucket [options] action
	Manage buckets

Options:
	` + generalOptionsUsage() + `

Action :
        create                        Create a bucket
        delete                        Delete a bucket
`
	return strings.TrimSpace(helpText)
}

// Synopsis return the command message
func (c *BucketCommand) Synopsis() string {
	return "Manage buckets."
}

// Run launch the command
func (c *BucketCommand) Run(args []string) int {
	var debug bool
	var config string
	f := flag.NewFlagSet("bucket", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }

	defaultConfigFile := getConfigurationFile()

	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&config, "config", defaultConfigFile, "Configuration filename")

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
	case "create":
		c.doCreateBucket(client)
	case "delete":
		c.doDeleteBucket(client)
	default:
		f.Usage()
	}
	return 0
}

func (c *BucketCommand) doCreateBucket(client *Client) {
	c.UI.Info(fmt.Sprintf("Create bucket"))
	err := client.Storage.Create()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Bucket successfully created"))
}

func (c *BucketCommand) doDeleteBucket(client *Client) {
	c.UI.Info(fmt.Sprintf("Delete bucket"))
	err := client.Storage.Destroy()
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Bucket successfully deleted"))
}
