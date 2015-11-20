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

package store

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

const (
	s3Label = "s3"
)

func init() {
	backends[s3Label] = NewS3
}

// S3 is the AWS S3 backend.
type S3 struct {
	Client *s3.S3
}

// NewS3 returns a new Kms.
func NewS3() (StorageBackend, error) {
	return &S3{
		Client: s3.New(session.New(&aws.Config{
			Region: aws.String("eu-west-1")})),
	}, nil
}

// Name returns S3 label
func (s *S3) Name() string {
	return s3Label
}

func (s *S3) List(bucket string) ([]string, error) {
	var l []string
	log.Printf("[DEBUG] S3 list secrets %s", bucket)
	resp, err := s.Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})
	log.Printf("[DEBUG] %s", awsutil.Prettify(resp))
	if err != nil {
		return l, err
	}
	for _, k := range resp.Contents {
		//if *k.Size <= 5000 {
		l = append(l, *k.Key)
		//}
	}
	return l, nil
}

func (s *S3) Put(bucket string, key []byte, value []byte) error {
	log.Printf("[DEBUG] S3 Put : %s %v %v", bucket, string(key), string(value))
	resp, err := s.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(string(key)),
		Body:   strings.NewReader(string(value)),
	})
	log.Printf("[DEBUG] %s", awsutil.Prettify(resp))
	return err
}

func (s *S3) Get(bucket string, key []byte) ([]byte, error) {
	log.Printf("[DEBUG] S3 Get : %s %v", bucket, string(key))
	resp, err := s.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(string(key)),
	})
	log.Printf("[DEBUG] %s", awsutil.Prettify(resp))
	if err != nil {
		return nil, err
	}
	blob := make([]byte, *resp.ContentLength)
	_, err = resp.Body.Read(blob)
	return blob, nil
}

func (s *S3) Delete(bucket string, key []byte) error {
	log.Printf("[DEBUG] S3 Delete : %s %v", bucket, string(key))
	resp, err := s.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &bucket,
		Key:    aws.String(string(key)),
	})
	log.Printf("[DEBUG] %s", awsutil.Prettify(resp))
	return err
}
