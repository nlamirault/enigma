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
	"flag"
	"log"
	"os"
	// "strings"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/s3"
)

var (
	bucket             string
	region             string
	printVersion       bool
	printEnigmaSecrets bool
	doCreateBucket     bool
	doDeleteBucket     bool
	doPutText          bool
	doGetText          bool
	text               string
	key                string
	file               string
)

func init() {
	flag.StringVar(&bucket, "bucket", "enigma", "s3 bucket name")
	flag.StringVar(&region, "region", "eu-west-1", "aws region")
	flag.BoolVar(&doCreateBucket, "create", false, "create bucket")
	flag.BoolVar(&doDeleteBucket, "delete", false, "delete bucket")
	flag.BoolVar(&doPutText, "put-text", false, "store text")
	flag.BoolVar(&doGetText, "get-text", false, "retrieve text")
	flag.StringVar(&text, "text", "", "text to store")
	flag.StringVar(&key, "key", "", "key for store data")
	flag.BoolVar(&printEnigmaSecrets, "list", false, "print files")
	flag.BoolVar(&printVersion, "version", false, "print version and exit")
}

func main() {
	flag.Parse()
	if printVersion {
		log.Println("Version", Version)
		os.Exit(0)
	}
	if printEnigmaSecrets {
		// checkArgument(region, "S3 region")
		// checkArgument(bucket, "S3 bucket")
		listEnigmaSecrets()
		os.Exit(0)
	}
	if doCreateBucket {
		// checkArgument(region, "S3 region")
		// checkArgument(bucket, "S3 bucket")
		createBucket()
	}
	if doDeleteBucket {
		// checkArgument(region, "S3 region")
		// checkArgument(bucket, "S3 bucket")
		deleteBucket()
	}
	if doPutText {
		checkArgument(text, "an input text to store")
		checkArgument(key, "a key name for data")
		putText()
	}
	if doGetText {
		checkArgument(key, "a key name for data")
		getText()
	}
}

func getKeyID() string {
	return os.Getenv("ENIGMA_KEYID")
}

func checkArgument(key string, value string) {
	if key == "" {
		log.Printf("Please specify %s. Exiting.\n", value)
		os.Exit(1)
	}
}

func getAWSConfig(region string) *aws.Config {
	return &aws.Config{Region: aws.String(region)}

}

func createBucket() {
	log.Printf("Create bucket : %s\n", bucket)
	s3Client := getS3Client(getAWSConfig(region))
	result, err := s3Client.CreateBucket(&s3.CreateBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		log.Println(err)
		return
	}
	// log.Println("Successfully created bucket", bucket, "in", *result.Location)
	log.Println(awsutil.Prettify(result))
}

func deleteBucket() {
	log.Println("Delete bucket objects")
	s3Client := getS3Client(getAWSConfig(region))
	list, err := listObjects(s3Client, bucket)
	if err != nil {
		log.Println(err)
		return
	}
	for _, key := range list {
		_, err := s3Client.DeleteObject(&s3.DeleteObjectInput{
			Bucket: &bucket,
			Key:    aws.String(key),
		})
		if err != nil {
			log.Println(err)
			return
		}
	}
	log.Println("Delete bucket")
	result, err := s3Client.DeleteBucket(&s3.DeleteBucketInput{
		Bucket: &bucket,
	})
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(awsutil.Prettify(result))
}

func listEnigmaSecrets() {
	log.Println("Files:")
	s3Client := getS3Client(getAWSConfig(region))
	list, err := listObjects(s3Client, bucket)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("Size:", len(list))
	for _, key := range list {
		log.Println("Object: ", key)
	}
}

func putText() {
	log.Printf("Encrypt text %s with key %s\n", text, key)
	cfg := getAWSConfig(region)
	keyID := getKeyID()

	encrypted, err := encrypt(getKmsClient(cfg), keyID, []byte(text))
	if err != nil {
		log.Printf("Can't encrypt : %v\n", err)
		return
	}
	log.Println("Encrypted: ", encrypted)

	_, err = storeText(getS3Client(cfg), bucket, key, string(encrypted))
	if err != nil {
		log.Printf("Failed to upload data to %s %v\n,",
			bucket, err)
		return
	}
	// log.Println(awsutil.Prettify(result))
	log.Printf("Successfully uploaded data with key %s\n", key)
}

func getText() {
	log.Printf("Retrive text for key : %s\n", key)
	cfg := getAWSConfig(region)
	blob, err := retrieveText(getS3Client(cfg), bucket, key)
	if err != nil {
		log.Printf("Failed to upload data to %s %v\n,",
			bucket, err)
		return
	}
	decrypted, err := decrypt(getKmsClient(cfg), &blob)
	if err != nil {
		log.Printf("Can't decrypt data: %v\n", err)
		return
	}
	log.Printf("Successfully decrypted: %s\n", decrypted)
}

// func putFile {
// 	s3Client := getS3Client(cfg)
// 	file, err := os.Open(path)
// 	if err != nil {
// 		log.Println("Failed opening file", path, err)
// 		continue
// 	}
// 	result, err := s3Client.PutObject(&s3.PutObjectInput{
// 		Bucket: &bucket,
// 		Key:    aws.String(rel),
// 		Body:   file,
// 	})
// 	if err != nil {
// 		log.Fatalln("Failed to upload data to %s/%s,",
// 			bucket, path, err)
// 		return
// 	}
// 	log.Printf("Uploaded: %s : %s\n", path, result)
// }
