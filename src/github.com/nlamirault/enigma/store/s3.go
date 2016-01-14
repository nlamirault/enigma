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

package store

import (
	"log"
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"

	"github.com/nlamirault/enigma/config"
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
	Bucket string
}

// NewS3 returns a new Kms.
func NewS3(conf *config.Configuration) (StorageBackend, error) {
	return &S3{
		Client: s3.New(session.New(&aws.Config{
			Region: aws.String(conf.S3.Region)})),
		Bucket: conf.S3.Bucket,
	}, nil
}

// Name returns S3 label
func (s *S3) Name() string {
	return s3Label
}

// List returns all secrets
func (s *S3) List() ([]string, error) {
	var l []string
	log.Printf("[DEBUG] Amazon S3 list secrets %s", s.Bucket)
	resp, err := s.Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(s.Bucket),
	})
	log.Printf("[DEBUG] Amazon S3 %s", awsutil.Prettify(resp))
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

// Put a value at the specified key
func (s *S3) Put(key []byte, value []byte) error {
	log.Printf("[DEBUG] Amazon S3 Put : %s %v %v", s.Bucket, string(key), string(value))
	resp, err := s.Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(string(key)),
		Body:   strings.NewReader(string(value)),
	})
	log.Printf("[DEBUG] Amazon S3 %s", awsutil.Prettify(resp))
	return err
}

// Get a value given its key
func (s *S3) Get(key []byte) ([]byte, error) {
	log.Printf("[DEBUG] Amazon S3 Get : %s %v", s.Bucket, string(key))
	resp, err := s.Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(s.Bucket),
		Key:    aws.String(string(key)),
	})
	log.Printf("[DEBUG] Amazon S3 %s", awsutil.Prettify(resp))
	if err != nil {
		return nil, err
	}
	blob := make([]byte, *resp.ContentLength)
	_, err = resp.Body.Read(blob)
	return blob, nil
}

// Delete the value at the specified key
func (s *S3) Delete(key []byte) error {
	log.Printf("[DEBUG] Amazon S3 Delete : %s %v", s.Bucket, string(key))
	resp, err := s.Client.DeleteObject(&s3.DeleteObjectInput{
		Bucket: &s.Bucket,
		Key:    aws.String(string(key)),
	})
	log.Printf("[DEBUG] Amazon S3 %s", awsutil.Prettify(resp))
	return err
}
