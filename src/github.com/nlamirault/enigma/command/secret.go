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
	//"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/cli"

	eaws "github.com/nlamirault/enigma/providers/aws"
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
        put-text                      Put a secret text
        get-text                      Retrieve a secret text
`
	return strings.TrimSpace(helpText)
}

func (c *SecretCommand) Synopsis() string {
	return "Manage secrets from Amazon S3"
}

func (c *SecretCommand) Run(args []string) int {
	var debug bool
	var bucket, region, key, text string
	f := flag.NewFlagSet("bucket", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }

	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&bucket, "bucket", "", "Glacier vault's bucket")
	f.StringVar(&region, "region", "eu-west-1", "AWS region name")
	f.StringVar(&key, "key", "", "Key for store data")
	f.StringVar(&text, "text", "", "Text to store")
	//f.StringVar(&action, "action", "", "Action to perform")

	if err := f.Parse(args); err != nil {
		return 1
	}
	args = f.Args()
	if len(args) != 1 {
		f.Usage()
		return 1
	}
	config := getAWSConfig(region, debug)
	action := args[0]
	//fmt.Printf("Action: %s\n", action)
	switch action {
	case "put-text":
		valid := checkArguments(bucket, region, key, text)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket, region key and text."))
			return 1
		}
		c.doPutText(config, bucket, key, text)
	case "get-text":
		valid := checkArguments(bucket, region, key)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket, region and key."))
			return 1
		}
		c.doGetText(config, bucket, key)
	default:
		f.Usage()
	}
	return 0
}

func (c *SecretCommand) doGetText(config *aws.Config, bucket string, key string) {
	c.UI.Info(fmt.Sprintf("Retrive text for key : %s", key))
	blob, err := eaws.RetrieveText(eaws.GetS3Client(config), bucket, key)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	decrypted, err := eaws.Decrypt(eaws.GetKmsClient(config), &blob)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(decrypted))
	c.UI.Output(fmt.Sprintf("Decrypted: %s", decrypted))
}

func (c *SecretCommand) doPutText(config *aws.Config, bucket string, key string, text string) {
	c.UI.Info(fmt.Sprintf("Encrypt text %s with key %s", text, key))
	keyID := getKeyID()
	encrypted, err := eaws.Encrypt(eaws.GetKmsClient(config), keyID, []byte(text))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] Encrypted: %v", encrypted)
	result, err := eaws.StoreText(
		eaws.GetS3Client(config), bucket, key, string(encrypted))
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
	c.UI.Output(fmt.Sprintf("Successfully uploaded data with key %s", key))
}

func (c *SecretCommand) doPutFile(config *aws.Config, bucket string, key string, path string) {
	c.UI.Info(fmt.Sprintf("Store secret : %s %s %s", bucket, key, path))
	file, err := os.Open(path)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	s3Client := eaws.GetS3Client(config)
	result, err := s3Client.PutObject(&s3.PutObjectInput{
		Bucket: &bucket,
		Key:    aws.String(key),
		Body:   file,
	})
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
	c.UI.Output(fmt.Sprintf("Uploaded: %s : %s", path, result))
}
