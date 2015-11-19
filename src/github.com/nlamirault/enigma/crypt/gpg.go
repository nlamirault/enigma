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

package crypt

import (
	//"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/homedir"
	"golang.org/x/crypto/openpgp"
)

const (
	gpgLabel              = "gpg"
	defaultGPGPath string = ".gnupg/"
)

func init() {
	registry[gpgLabel] = NewGpg
}

// Gpg is a KeyManager using GPG.
type Gpg struct {
	PublicKeyring string
	SecretKeyring string
}

// NewGpg returns a new Gpg instance
func NewGpg() KeyManager {
	home := homedir.Get()
	publicKeyring := filepath.Join(home, defaultGPGPath, "pubring.gpg")
	secretKeyring := filepath.Join(home, defaultGPGPath, "secring.gpg")
	return &Gpg{
		PublicKeyring: publicKeyring,
		SecretKeyring: secretKeyring,
	}
}

func (g *Gpg) Name() string {
	return gpgLabel
}

func (g *Gpg) Encrypt(publicKeyring string, b []byte) ([]byte, error) {
	// Read in public key
	keyringFileBuffer, err := os.Open(publicKeyring)
	if err != nil {
		return nil, fmt.Errorf(
			"opening public key %s failed: %v", publicKeyring, err)
	}
	defer keyringFileBuffer.Close()

	entityList, err := openpgp.ReadKeyRing(keyringFileBuffer)
	if err != nil {
		return nil, err
	}

	// encrypt string
	buf := new(bytes.Buffer)
	w, err := openpgp.Encrypt(buf, entityList, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	defer w.Close()

	// write the byte to our encrypt writer
	if _, err = w.Write(b); err != nil {
		return nil, err
	}

	// close the writer
	if err = w.Close(); err != nil {
		return nil, err
	}

	bytes, err := ioutil.ReadAll(buf)
	if err != nil {
		return nil, err
	}

	// base64 encode
	encStr := base64.StdEncoding.EncodeToString(bytes)
	return []byte(encStr), nil
}

// func readInput(in io.Reader, out io.Writer) []byte {
// 	reader := bufio.NewReader(in)
// 	line, _, err := reader.ReadLine()
// 	if err != nil {
// 		fmt.Fprintln(out, err.Error())
// 		os.Exit(1)
// 	}
// 	return line
// }

// // Decrypt a io.Reader with the given secretKeyring.
// // You can optionally pass a defaultGPGKey to use for the
// // decryption, otherwise it will use the first entity.
// func Decrypt(f io.Reader, secretKeyring, defaultGPGKey string) (io.Reader, error) {
func (g *Gpg) Decrypt(secretKeyring string, blob []byte) ([]byte, error) {

	// Open the private key file
	keyringFileBuffer, err := os.Open(secretKeyring)
	if err != nil {
		return nil, err
	}
	defer keyringFileBuffer.Close()

	entityList, err := openpgp.ReadKeyRing(keyringFileBuffer)
	if err != nil {
		return nil, err
	}
	entity := entityList[0]

	// 	var entity *openpgp.Entity
	// 	if defaultGPGKey != "" {

	// 		// loop through their keys until we find the one they want
	// 		var foundKey bool
	// 		for _, e := range entityList {
	// 			// we can match on the fingerprint or the keyid because
	// 			// why not? I bet no one knows the difference
	// 			if e.PrimaryKey.KeyIdString() == defaultGPGKey ||
	// 				e.PrimaryKey.KeyIdShortString() == defaultGPGKey ||
	// 				fmt.Sprintf("%X", e.PrimaryKey.Fingerprint) == defaultGPGKey {
	// 				foundKey = true
	// 				entity = e
	// 				break
	// 			}
	// 		}

	// 		if !foundKey {
	// 			// we didn't find the key they specified
	// 			return nil, fmt.Errorf("Could not find private GPG Key with id: %s", defaultGPGKey)
	// 		}

	// 	} else {
	// 		// they didn't set a default key
	// 		// so let's hope it is the first one :/
	// 		// TODO(jfrazelle): maybe prompt here if they have
	// 		// more than one private key
	// 		entity = entityList[0]
	// 	}

	// var identityString string
	// for _, identity := range entity.Identities {
	// 	identityString = fmt.Sprintf(" %s [%s]",
	// 		identity.Name, entity.PrimaryKey.KeyIdString())
	// 	break
	// }

	fmt.Print("GPG Passphrase: ")
	var passphrase string
	fmt.Scanln(&passphrase)

	log.Printf("[DEBUG] Decrypting private key using passphrase")

	passphraseByte := []byte(passphrase)
	entity.PrivateKey.Decrypt(passphraseByte)
	for _, subkey := range entity.Subkeys {
		subkey.PrivateKey.Decrypt(passphraseByte)
	}
	log.Printf("[DEBUG] Finished decrypting private key using passphrase")

	dec, err := base64.StdEncoding.DecodeString(string(blob))
	if err != nil {
		return nil, err
	}

	// Decrypt it with the contents of the private key
	md, err := openpgp.ReadMessage(bytes.NewBuffer(dec), entityList, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GPG Read message failed: %v", err)
	}
	bytes, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}
	return bytes, nil
}
