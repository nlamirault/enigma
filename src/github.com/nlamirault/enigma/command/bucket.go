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
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/mitchellh/cli"

	eaws "github.com/nlamirault/enigma/providers/aws"
)

type BucketCommand struct {
	UI cli.Ui
}

func (c *BucketCommand) Help() string {
	helpText := `
Usage: enigma bucket [options] action
	Manage buckets from Amazon S3

Options:
	` + generalOptionsUsage() + `

Action :
        list                          Display bucket's content
        create                        Create a bucket
        delete                        Delete a bucket
`
	return strings.TrimSpace(helpText)
}

func (c *BucketCommand) Synopsis() string {
	return "Manage buckets from Amazon S3"
}

func (c *BucketCommand) Run(args []string) int {
	var debug bool
	var bucket, region string
	f := flag.NewFlagSet("bucket", flag.ContinueOnError)
	f.Usage = func() { c.UI.Error(c.Help()) }

	f.BoolVar(&debug, "debug", false, "Debug mode enabled")
	f.StringVar(&bucket, "bucket", "", "Glacier vault's name")
	f.StringVar(&region, "region", "eu-west-1", "AWS region name")
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
	switch action {
	case "list":
		valid := checkArguments(bucket, region)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket and region."))
			return 1
		}
		c.doListBucket(config, bucket)
	case "create":
		valid := checkArguments(bucket, region)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket and region."))
			return 1
		}
		c.doCreateBucket(config, bucket)
	case "delete":
		valid := checkArguments(bucket, region)
		if !valid {
			f.Usage()
			c.UI.Error(fmt.Sprintf(
				"\nSecret expects arguments: bucket and region."))
			return 1
		}
		c.doDeleteBucket(config, bucket)
	default:
		f.Usage()
	}
	return 0
}

func (c *BucketCommand) doListBucket(config *aws.Config, bucket string) {
	c.UI.Info(fmt.Sprintf("List bucket secrets : %s", bucket))
	s3Client := eaws.GetS3Client(config)
	list, err := eaws.ListObjects(s3Client, bucket)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Println("[DEBUG] Size:", len(list))
	for _, key := range list {
		c.UI.Output(fmt.Sprintf("- %s", key))
	}
}

func (c *BucketCommand) doCreateBucket(config *aws.Config, bucket string) {
	c.UI.Info(fmt.Sprintf("Create bucket : %s", bucket))
	s3Client := eaws.GetS3Client(config)
	result, err := s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
	c.UI.Output(fmt.Sprintf("Created: %s", *result.Location))
}

func (c *BucketCommand) doDeleteBucket(config *aws.Config, bucket string) {
	c.UI.Info(fmt.Sprintf("Delete bucket %s", bucket))
	s3Client := eaws.GetS3Client(config)
	list, err := eaws.ListObjects(s3Client, bucket)
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] Delete bucket objects")
	for _, key := range list {
		res, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: &bucket,
			Key:    aws.String(key),
		})
		if err != nil {
			c.UI.Error(err.Error())
			return
		}
		log.Printf("[DEBUG] %s", awsutil.Prettify(res))
	}
	log.Printf("[DEBUG] Delete bucket")
	result, err := s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		c.UI.Error(err.Error())
		return
	}
	log.Printf("[DEBUG] %s", awsutil.Prettify(result))
	c.UI.Output("Deleted")
}
