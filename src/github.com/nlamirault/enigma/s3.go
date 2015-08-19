// Copyright (C) 2015 Nicolas Lamirault <nicolas.lamirault@gmail.com>

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

package main

import (
	"strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func getS3Client return S3 service client
func getS3Client(cfg *aws.Config) *s3.S3 {
	c := s3.New(cfg)
	return c
}

func listObjects(s3Client *s3.S3, bucket string) ([]string, error) {
	l := make([]string, 0)
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

func storeText(s3Client *s3.S3, bucket string, key string, text string) (*s3.PutObjectOutput, error) {
	return s3Client.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   strings.NewReader(text),
	})
}

func retrieveText(s3Client *s3.S3, bucket string, key string) ([]byte, error) {
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
