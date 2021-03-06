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
	//"bufio"
	"bytes"
	//"encoding/base64"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/docker/docker/pkg/homedir"
	"golang.org/x/crypto/openpgp"
	"golang.org/x/crypto/openpgp/armor"
	"golang.org/x/crypto/ssh/terminal"

	"github.com/nlamirault/enigma/config"
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
	PublicKeyring  string
	PrivateKeyring string
	Email          string
}

// NewGpg returns a new Gpg instance
func NewGpg(conf *config.Configuration) (KeyManager, error) {
	home := homedir.Get()
	publicKeyring := filepath.Join(home, defaultGPGPath, "pubring.gpg")
	privateKeyring := filepath.Join(home, defaultGPGPath, "secring.gpg")
	return &Gpg{
		PublicKeyring:  publicKeyring,
		PrivateKeyring: privateKeyring,
		Email:          conf.Gpg.Email,
	}, nil
}

// Name return the GPG name provider
func (g *Gpg) Name() string {
	return gpgLabel
}

// Encrypt encrypts a message
func (g *Gpg) Encrypt(b []byte) ([]byte, error) {
	log.Printf("[DEBUG] GPG Open public keyring %s", g.PublicKeyring)
	publicRingBuffer, err := os.Open(g.PublicKeyring)
	if err != nil {
		return nil, fmt.Errorf(
			"opening public key %s failed: %v", g.PublicKeyring, err)
	}
	defer publicRingBuffer.Close()
	log.Printf("[DEBUG] GPG Read public keyring")
	publicRing, err := openpgp.ReadKeyRing(publicRingBuffer)
	if err != nil {
		return nil, err
	}
	publicKey := getKeyByEmail(publicRing, g.Email)
	if publicKey == nil {
		return nil, fmt.Errorf("Can't find GPG public key")
	}

	var buffer = &bytes.Buffer{}
	armoredWriter, err := armor.Encode(buffer, "PGP MESSAGE", nil)
	if err != nil {
		return nil, err
	}
	cipheredWriter, err := openpgp.Encrypt(
		armoredWriter, []*openpgp.Entity{publicKey}, nil, nil, nil)
	if err != nil {
		return nil, err
	}
	_, err = cipheredWriter.Write(b)
	if err != nil {
		return nil, err
	}

	cipheredWriter.Close()
	armoredWriter.Close()

	return buffer.Bytes(), nil

}

// Decrypt decrypt an encrypted message
func (g *Gpg) Decrypt(blob []byte) ([]byte, error) {

	// Open the private key file
	privateRingBuffer, err := os.Open(g.PrivateKeyring)
	if err != nil {
		return nil, err
	}
	defer privateRingBuffer.Close()

	privateRing, err := openpgp.ReadKeyRing(privateRingBuffer)
	if err != nil {
		return nil, err
	}
	privateKey := getKeyByEmail(privateRing, g.Email)

	fmt.Print("GPG Passphrase: ")
	passphrase, err := terminal.ReadPassword(0)
	fmt.Println("")

	log.Printf("[DEBUG] GPG Decrypting private key using passphrase")
	//passphraseByte := []byte(passphrase)
	privateKey.PrivateKey.Decrypt(passphrase)
	for _, subkey := range privateKey.Subkeys {
		subkey.PrivateKey.Decrypt(passphrase)
	}
	log.Printf("[DEBUG] GPG Finished decrypting private key using passphrase")

	armoredBlock, err := armor.Decode(bytes.NewReader(blob))
	if err != nil {
		return nil, err
	}

	// Decrypt it with the contents of the private key
	md, err := openpgp.ReadMessage(armoredBlock.Body, privateRing, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("GPG Read message failed: %v", err)
	}
	plain, err := ioutil.ReadAll(md.UnverifiedBody)
	if err != nil {
		return nil, err
	}
	return plain, nil
}

func getKeyByEmail(keyring openpgp.EntityList, email string) *openpgp.Entity {
	log.Printf("[DEBUG] GPG Search key into keyring using %s", email)
	for _, entity := range keyring {
		for _, ident := range entity.Identities {
			if ident.UserId.Email == email {
				return entity
			}
		}
	}
	return nil
}
