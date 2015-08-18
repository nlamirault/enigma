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
