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

package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/kms"
)

// Function getKmsClient returns KMS service client
func getKmsClient(cfg *aws.Config) *kms.KMS {
	c := kms.New(cfg)
	return c
}

func decrypt(kmsClient *kms.KMS, ciphertext *[]byte) ([]byte, error) {
	resp, err := kmsClient.Decrypt(&kms.DecryptInput{
		CiphertextBlob: *ciphertext,
	})
	if err != nil {
		return nil, err
	}
	return resp.Plaintext, nil
}

func encrypt(kmsClient *kms.KMS, keyID string, plaintext []byte) ([]byte, error) {
	resp, err := kmsClient.Encrypt(&kms.EncryptInput{
		Plaintext: plaintext,
		KeyId:     aws.String(keyID),
	})
	if err != nil {
		return nil, err
	}
	return resp.CiphertextBlob, nil
}
