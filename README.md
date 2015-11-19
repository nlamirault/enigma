# Enigma

[![License Apache 2][badge-license]](LICENSE)

Master :
* [![Circle CI](https://circleci.com/gh/nlamirault/enigma/tree/master.svg?style=svg)](https://circleci.com/gh/nlamirault/enigma/tree/master)

Develop :
* [![Circle CI](https://circleci.com/gh/nlamirault/enigma/tree/develop.svg?style=svg)](https://circleci.com/gh/nlamirault/enigma/tree/develop)


This tool is a personal safe.

* Storage backend :
- [Amazon S3][] to store secrets encrypted


* Secret provider :
- [Amazon KMS][] to manage encryption keys which encrypt/decrypt your secrets


## Usage

### Amazon

Setup your AWS credentials :

    $ export AWS_ACCESS_KEY_ID='AKID'
    $ export AWS_SECRET_ACCESS_KEY='SECRET'

Creates a KMS key via the AWS Console and store its ID (a UUID) in an environment variable:

    $ export ENIGMA_KEYID = "be3338c8-4e4e-2384-644f-1ac0af044fe4"

Initialize your bucket into S3 :

    $ enigma bucket --bucket=my-enigma-bucket create
    Create bucket : my-enigma-bucket
    Created: http://my-enigma-bucket.s3.amazonaws.com/

Store a text :

    $ enigma secret --bucket=my-enigma-bucket --key=foo --text="A text for my enigma" put-text
    Store secret text A text for my enigma with key foo
    Successfully uploaded data with key foo

Check your enigma :

    $ enigma bucket --bucket=my-enigma-bucket list
    List bucket secrets : my-enigma-bucket
    - foo

Retrieve your data :

    $ enigma secret --bucket=my-enigma-bucket --key=foo get-text
    Retrive secret text for key : foo
    Decrypted: A text for my enigma

Delete your bucket :

    $ enigma bucket --bucket=my-enigma-bucket delete
    Delete bucket my-enigma-bucket
    Deleted


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

[Amazon S3]:https://aws.amazon.com/s3/
[Amazon KMS]: https://aws.amazon.com/kms/
