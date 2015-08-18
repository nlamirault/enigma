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
	"fmt"
	"os"
	"testing"

	"github.com/aws/aws-sdk-go/aws"
)

const (
	testregion = "eu-west-1"
	plaintext  = "In tartiflette we trust !"
)

func Test_Enigma(t *testing.T) {

	var cfg *aws.Config
	cfg = &aws.Config{Region: aws.String(testregion)}

	kmsClient := getKmsClient(cfg)
	keyID := os.Getenv("ENIGMA_KEYID")
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
