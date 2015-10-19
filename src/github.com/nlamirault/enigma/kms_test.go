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
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

const (
	testregion = "eu-west-1"
	plaintext  = "In tartiflette we trust !"
)

func Test_EnigmaKms(t *testing.T) {

	var cfg *aws.Config
	cfg = &aws.Config{Region: aws.String(testregion)}

	kmsClient := getKmsClient(cfg)
	keyID := os.Getenv("ENIGMA_KEYID")
	if len(keyID) == 0 {
		t.Fatalf("ENIGMA_KEYID not found")
	}

	// Encrypt plaintext
	encrypted, err := encrypt(kmsClient, keyID, []byte(plaintext))
	if err != nil {
		t.Fatalf("Can't encrypt : %v", err)
	}
	fmt.Println("Encrypted: ", encrypted)

	// Decrypt ciphertext
	decrypted, err := decrypt(kmsClient, &encrypted)
	if err != nil {
		t.Fatalf("Can't decrypt : %v", err)
	}
	fmt.Println("Decrypted: ", decrypted)

	if plaintext != string(decrypted) {
		t.Fatalf("Enigma failed %v", err)
	}

}
