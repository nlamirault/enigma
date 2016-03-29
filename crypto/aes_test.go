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
	"testing"

	"github.com/nlamirault/enigma/config"
)

var (
	IV = []byte{0x00, 0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0a, 0x0b, 0x0c, 0x0d, 0x0e, 0x0f}
)

func Test_EnigmaAesManager(t *testing.T) {
	plaintext := "In tartiflette we trust !"
	manager, err := NewAes(&config.Configuration{
		Encryption: "aes",
		Aes: &config.AesConfiguration{
			Key: "averyveryveryveryveryverylongkey",
		},
	})
	if err != nil {
		t.Fatalf("Can't create AES manager : %v", err)
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
