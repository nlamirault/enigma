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

package aws

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetS3Client return S3 service client
func GetS3Client(cfg *aws.Config) *s3.S3 {
	c := s3.New(cfg)
	return c
}

func ListObjects(s3Client *s3.S3, bucket string) ([]string, error) {
	var l []string
	resp, err := s3Client.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(bucket),
	})
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

func StoreText(s3Client *s3.S3, bucket string, key string, text string) (*s3.PutObjectOutput, error) {
	return s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(text),
	})
}

func RetrieveText(s3Client *s3.S3, bucket string, key string) ([]byte, error) {
	resp, err := s3Client.GetObject(&s3.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return nil, err
	}
	blob := make([]byte, *resp.ContentLength)
	_, err = resp.Body.Read(blob)
	return blob, nil
}
