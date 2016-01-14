# Enigma

[![License Apache 2][badge-license]](LICENSE)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/enigma/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/enigma/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/enigma/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/enigma/tree/develop)


This tool is a personal safe.

## Storage backend

- [BoltDB][]
- [Amazon S3][]
- [Google Cloud Storage][]


## Secret provider
- [Amazon KMS][]
- [GPG][]
- [AES][]


## Installation

You can download the binaries :

* Architecture i386 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/enigma_linux_386) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/enigma_darwin_386) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_freebsd_386) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_netbsd_386) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_openbsd_386) / [windows](https://bintray.com/artifact/download/nlamirault/oss/enigma_windows_386.exe) ]
* Architecture amd64 [ [linux](https://bintray.com/artifact/download/nlamirault/oss/enigma_linux_amd64) / [darwin](https://bintray.com/artifact/download/nlamirault/oss/enigma_darwin_amd64) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_freebsd_amd64) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_netbsd_amd64) / [openbsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_openbsd_amd64) / [windows](https://bintray.com/artifact/download/nlamirault/oss/enigma_windows_amd64.exe) ]
* Architecture arm [ [linux](https://bintray.com/artifact/download/nlamirault/oss/enigma_linux_arm) / [freebsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_freebsd_arm) / [netbsd](https://bintray.com/artifact/download/nlamirault/oss/enigma_netbsd_arm) ]



## Configuration

Enigma configuration use [toml][] format. File is located into `$HOME/.config/enigma/enigma.toml`.

### KMS

To use the Amazon KMS, :

* Creates a KMS key via the AWS Console and store its ID (a UUID)
* Setup the AWS region

* Setup into the configuration file :

        [kms]
        region = "eu-west-1"
        keyID = "xxxx-xxxx-xxxx"

### S3

* Initialize your bucket into S3 :

        $ enigma bucket --bucket=my-enigma-bucket create
        Create bucket : my-enigma-bucket
        Created: http://my-enigma-bucket.s3.amazonaws.com/

* Setup into the configuration file :

        [s3]
        region = "eu-west-1"
        bucket = "my-enigma-bucket"


### GPG

Specify the email to use with your public key:

        [gpg]
        email = "foo.bar@gmail.com"

### BoltDB

You must specify where database file will be saved and the bucket name :

        [boltdb]
        file = "/tmp/enigma.db"
        bucket = "enigma"



### Example

```toml
# enigma.toml

# Encryption provider
backend = "gpg"

# Storage backend
storage = "boltdb"

[gpg]
email = "foo.bar@gmail.com"

[kms]
region = "eu-west-1"
keyID = "xxxx-xxxx-xxxx"

[aes]
key = "abcdefghijklmnop"

[s3]
region = "eu-west-1"
bucket = "enigma"

[boltdb]
file = "/tmp/enigma.db"
bucket = "enigma"

```

## Usage

### KMS / BoltDB

* List all secrets:

        $ enigma secret list
        List secrets :

* Store a new secret :

        $ enigma secret --key="mysecret" --text="mypassword" put
        Store secret text mypassword with key mysecret
        Successfully uploaded data with key mysecret

        $ enigma secret list
        List secrets :
        - mysecret

* Retrieve a secret :

        $ enigma secret --key="mysecret" get
        Retrive secret text for key : mysecret
        Decrypted: mypassword


### GPG / BoltDB

* Store a new secret :

        $ enigma secret --debug --key="nicolas" --text="mypassword" put
        2016/01/14 23:08:04 [DEBUG] Init BoltDB storage : /tmp/enigma.db
        Store secret text mypassword with key nicolas
        2016/01/14 23:08:04 [DEBUG] GPG Open public keyring /home/nlamirault/.gnupg/pubring.gpg
        2016/01/14 23:08:04 [DEBUG] GPG Read public keyring
        2016/01/14 23:08:04 [DEBUG] GPG Search key into keyring using nicolas.lamirault@gmail.com
        2016/01/14 23:08:04 [DEBUG] Put : nicolas -----BEGIN PGP MESSAGE-----
        [...]
        4AHkPJd4QQaimnFACYR8pTeEUuEgOODO4Arhwt/gDOKYMAIv4ILjI5qsqqWR+qjg
        zOF8/+Dp5GSbF7vp19ilGb8OubCpgHTiL/fIquGi8AA=
        =9agp
        -----END PGP MESSAGE-----
        Successfully uploaded data with key nicolas

* Retrieve a secret :

        $ bin/enigma secret --debug --key="nicolas" get
        2016/01/14 23:10:06 [DEBUG] Init BoltDB storage : /tmp/enigma.db
        Retrive secret text for key : nicolas
        2016/01/14 23:10:06 [DEBUG] Search entry with key : nicolas
        2016/01/14 23:10:06 [DEBUG] GPG Search key into keyring using nicolas.lamirault@gmail.com
        GPG Passphrase:
        2016/01/14 23:10:11 [DEBUG] GPG Decrypting private key using passphrase
        2016/01/14 23:10:11 [DEBUG] GPG Finished decrypting private key using passphrase
        Decrypted: mypassword



## Development

* Initialize environment

        $ make init

* Build tool :

        $ make build

* Launch unit tests :

        $ make test

## Contributing

See [CONTRIBUTING](CONTRIBUTING.md).


## License

See [LICENSE](LICENSE) for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>


[badge-license]: https://img.shields.io/badge/license-Apache2-green.svg?style=flat

[BoltDB]: https://github.com/boltdb/bolt

[Amazon S3]:https://aws.amazon.com/s3/
[Google Cloud Storage]: https://cloud.google.com/storage/

[Amazon KMS]: https://aws.amazon.com/kms/
[GPG]: https://www.gnupg.org/
[AES]: https://en.wikipedia.org/wiki/Advanced_Encryption_Standard


[toml]: https://github.com/toml-lang/toml
