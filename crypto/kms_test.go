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

package crypto

import (
	"fmt"
	"os"
	"testing"

	"github.com/nlamirault/enigma/config"
)

func Test_EnigmaKmsManager(t *testing.T) {
	plaintext := "In tartiflette we trust !"
	keyID := os.Getenv("ENIGMA_KEYID")
	if len(keyID) == 0 {
		t.Fatalf("ENIGMA_KEYID not found")
	}
	manager, err := NewKms(&config.Configuration{
		Encryption: "kms",
		Kms: &config.KmsConfiguration{
			Region: "eu-west-1",
			KeyID:  keyID,
		},
	})
	if err != nil {
		t.Fatalf("Can't create KMS manager : %v", err)
	}

	// Encrypt plaintext
	ev, err := manager.Encrypt([]byte(plaintext))
	if err != nil {
		t.Fatalf("Can't encrypt : %v", err)
	}
	fmt.Println("Encrypted: ", ev)

	// Decrypt ciphertext
	decrypted, err := manager.Decrypt(ev)
	if err != nil {
		t.Fatalf("Can't decrypt : %v", err)
	}
	fmt.Println("Decrypted: ", decrypted)

	if plaintext != string(decrypted) {
		t.Fatalf("Enigma failed %v", err)
	}
}
