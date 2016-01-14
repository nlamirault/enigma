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


## Configuration


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
