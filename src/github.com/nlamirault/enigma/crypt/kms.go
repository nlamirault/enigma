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
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/kms"
	"golang.org/x/crypto/nacl/secretbox"
)

const (
	kmsLabel    = "kms"
	keyLength   = 32
	nonceLength = 24
)

func init() {
	registry[kmsLabel] = NewKms
}

// Kms is a KeyManager for AWS KMS.
type Kms struct {
	client *kms.KMS
}

// NewKms returns a new Kms.
func NewKms() KeyManager {
	return &Kms{
		client: kms.New(session.New(&aws.Config{
			Region: aws.String("eu-west-1")})),
	}
}

// Name returns kmsLabel
func (k *Kms) Name() string {
	return kmsLabel
}

// Decrypt decrypts the encrypted key.
func (k *Kms) Decrypt(blob []byte) ([]byte, error) {
	var ev Envelope
	err := unmarshalJSON(blob, &ev)
	if err != nil {
		return nil, err
	}
	res, err := k.client.Decrypt(&kms.DecryptInput{
		CiphertextBlob: ev.EncryptedKey,
	})
	log.Printf("[DEBUG] %s", awsutil.Prettify(res))
	if err != nil {
		return nil, err
	}

	var key [keyLength]byte
	copy(key[:], res.Plaintext[0:keyLength])

	var nonce [nonceLength]byte
	copy(nonce[:], ev.Nonce[0:nonceLength])

	var dec []byte
	dec, ok := secretbox.Open(dec, ev.Ciphertext, &nonce, &key)
	if !ok {
		return nil, fmt.Errorf("Can't decrypt data")
	}
	return dec, nil
}

// Encrypt encrypt the text using a plaintext key
func (k *Kms) Encrypt(plaintext []byte) ([]byte, error) {
	encKey, err := k.generateEnvelopKey(getKey())
	var key [keyLength]byte
	copy(key[:], encKey.Plaintext[0:keyLength])

	rand, err := k.generateNonce()
	if err != nil {
		return nil, err
	}
	var nonce [nonceLength]byte
	copy(nonce[:], rand[0:nonceLength])

	var enc []byte
	enc = secretbox.Seal(enc, plaintext, &nonce, &key)

	ev := &Envelope{
		Ciphertext:   enc,
		EncryptedKey: encKey.CiphertextBlob,
		Nonce:        nonce[:],
	}
	output, err := marshalJSON(ev)
	if err != nil {
		return nil, err
	}
	return output, nil
}

// Generate generates an EnvelopeKey under a specific KeyID.
func (k *Kms) generateEnvelopKey(keyID string) (*kms.GenerateDataKeyOutput, error) {
	resp, err := k.client.GenerateDataKey(&kms.GenerateDataKeyInput{
		KeyId:         aws.String(keyID),
		NumberOfBytes: aws.Int64(keyLength),
	})
	log.Printf("[DEBUG] %s", awsutil.Prettify(resp))
	if err != nil {
		return nil, err
	}
	//return &EnvelopeKey{dk.Plaintext, dk.CiphertextBlob}, nil
	return resp, nil
}

func (k *Kms) generateNonce() ([]byte, error) {
	resp, err := k.client.GenerateRandom(
		&kms.GenerateRandomInput{
			NumberOfBytes: aws.Int64(nonceLength),
		},
	)
	log.Printf("[DEBUG] %s", awsutil.Prettify(resp))
	if err != nil {
		return nil, err
	}
	return resp.Plaintext, nil
}

func marshalJSON(ev *Envelope) ([]byte, error) {
	return json.Marshal(ev)
}

func unmarshalJSON(data []byte, ev *Envelope) error {
	return json.Unmarshal(data, ev)
}

func getKey() string {
	return os.Getenv("ENIGMA_KEYID")
}

// GetKmsClient returns KMS service client
// func GetKmsClient(cfg *aws.Config) *kms.KMS {
// 	c := kms.New(cfg)
// 	return c
// }

// func Decrypt(kmsClient *kms.KMS, ciphertext *[]byte) ([]byte, error) {
// 	resp, err := kmsClient.Decrypt(&kms.DecryptInput{
// 		CiphertextBlob: *ciphertext,
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.Plaintext, nil
// }

// func Encrypt(kmsClient *kms.KMS, keyID string, plaintext []byte) ([]byte, error) {
// 	resp, err := kmsClient.Encrypt(&kms.EncryptInput{
// 		Plaintext: plaintext,
// 		KeyId:     aws.String(keyID),
// 	})
// 	if err != nil {
// 		return nil, err
// 	}
// 	return resp.CiphertextBlob, nil
// }
