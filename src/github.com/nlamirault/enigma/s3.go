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
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3"
)

// func getS3Client return S3 service client
func getS3Client(cfg *aws.Config) *s3.S3 {
	c := s3.New(cfg)
	return c
}

func listObjects(c *s3.S3, b string) ([]string, error) {
	l := make([]string, 0)
	resp, err := c.ListObjects(&s3.ListObjectsInput{
		Bucket: aws.String(b),
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
