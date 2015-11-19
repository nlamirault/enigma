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
	"flag"
	"fmt"
	//"log"
	//"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	// "github.com/aws/aws-sdk-go/aws/awsutil"
	// "github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/cli"

	//"github.com/nlamirault/enigma/crypt"
	//"github.com/nlamirault/enigma/store"
)

type SecretCommand struct {
	UI cli.Ui
}

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

func (c *SecretCommand) Synopsis() string {
	return "Manage secrets from Amazon S3"
}

func (c *SecretCommand) Run(args []string) int {
	var debug bool
	var bucket, region, key, text, keysManager string
	f := flag.NewFlagSet("bucket", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }

	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&bucket, "bucket", "", "Glacier vault's bucket")
	f.StringVar(&region, "region", "eu-west-1", "AWS region name")
	f.StringVar(&key, "key", "", "Key for store data")
	f.StringVar(&text, "text", "", "Text to store")
	f.StringVar(&keysManager, "keys-manager", "kms", "Keys Manager")
	//f.StringVar(&action, "action", "", "Action to perform")

	if err := f.Parse(args); err != nil {
		return 1
	}
	args = f.Args()
	if len(args) != 1 {
		f.Usage()
		return 1
	}

	client, err := NewClient("kms", "s3")
	if err != nil {
		c.UI.Error(err.Error())
		return 1
	}

	config := getAWSConfig(region, debug)
	action := args[0]
	//fmt.Printf("Action: %s\n", action)

	switch action {
	case "list":
		valid := checkArguments(bucket, region)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket and region."))
			return 1
		}
		c.doList(client, config, bucket)
	case "delete":
		valid := checkArguments(bucket, region, key)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket, region and key."))
			return 1
		}
		c.doDelete(client, config, bucket, key)
	case "put":
		valid := checkArguments(bucket, region, key, text)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket, region key and text."))
			return 1
		}
		c.doPutText(client, config, bucket, key, text)
	case "get":
		valid := checkArguments(bucket, region, key)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket, region and key."))
			return 1
		}
		c.doGetText(client, config, bucket, key)
	default:
		f.Usage()
	}
	return 0
}

func (c *SecretCommand) doGetText(client *Client, config *aws.Config, bucket string, key string) {
	c.UI.Info(fmt.Sprintf("Retrive secret text for key : %s", key))
	blob, err := client.Storage.Get(bucket, []byte(key))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	// var ev crypt.Envelope
	// err = crypt.UnmarshalJSON(blob, &ev)
	// if err != nil {
	// 	c.UI.Error(err.Error())
	// 	return
	// }
	keyID := getKeyID()
	// decrypted, err := client.Keys.Decrypt(keyID, &ev)
	decrypted, err := client.Keys.Decrypt(keyID, blob)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Decrypted: %s", string(decrypted)))
}

func (c *SecretCommand) doPutText(client *Client, config *aws.Config, bucket string, key string, text string) {
	c.UI.Info(fmt.Sprintf("Store secret text %s with key %s", text, key))

	keyID := getKeyID()
	// ev, err := client.Keys.Encrypt(keyID, []byte(text))
	// if err != nil {
	// 	c.UI.Error(err.Error())
	// 	return
	// }
	// log.Printf("[DEBUG] Encrypted: %v", ev)
	// output, err := crypt.MarshalJSON(ev)

	output, err := client.Keys.Encrypt(keyID, []byte(text))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}

	err = client.Storage.Put(bucket, []byte(key), output)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	c.UI.Output(fmt.Sprintf("Successfully uploaded data with key %s", key))
}

// func (c *SecretCommand) doPutFile(config *aws.Config, bucket string, key string, path string) {
// c.UI.Info(fmt.Sprintf("Store secret : %s %s %s", bucket, key, path))
// file, err := os.Open(path)
// if err != nil {
// 	c.UI.Error(err.Error())
// 	return
// }
// s3Client := eaws.GetS3Client(config)
// result, err := s3Client.PutObject(&s3.PutObjectInput{
// 	Bucket: &bucket,
// 	Key:    aws.String(key),
// 	Body:   file,
// })
// if err != nil {
// 	c.UI.Error(err.Error())
// 	return
// }
// log.Printf("[DEBUG] %s", awsutil.Prettify(result))
// c.UI.Output(fmt.Sprintf("Uploaded: %s : %s", path, result))
// }

func (c *SecretCommand) doDelete(client *Client, config *aws.Config, bucket string, key string) {
	c.UI.Info(fmt.Sprintf("Delete secret with key %s", key))
	client.Storage.Delete(bucket, []byte(key))
	c.UI.Output("Deleted")
}

func (c *SecretCommand) doList(client *Client, config *aws.Config, bucket string) {
	c.UI.Info(fmt.Sprintf("List secrets :"))
	secrets, err := client.Storage.List(bucket)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	for _, key := range secrets {
		c.UI.Output(fmt.Sprintf("- %s", key))
	}
}
