# Enigma

This tool is a personal safe using [Amazon S3][] and [Amazon KMS][]

## Usage

Creates a KMS key via the AWS Console and store its ID (a UUID) in an environment variable:

    $ export ENIGMA_KEYID = "be3338c8-4e4e-2384-644f-1ac0af044fe4"

Initialize your bucket into S3 :

    $ enigma create


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

See [LICENSE][] for the complete license.


## Changelog

A [changelog](ChangeLog.md) is available


## Contact

Nicolas Lamirault <nicolas.lamirault@gmail.com>



[Amazon S3]:https://aws.amazon.com/s3/
[Amazon KMS]: https://aws.amazon.com/kms/
