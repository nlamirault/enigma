# Enigma

[![Circle CI](https://circleci.com/gh/nlamirault/enigma.svg?style=svg)](https://circleci.com/gh/nlamirault/enigma)

This tool is a personal safe using [Amazon S3][] and [Amazon KMS][]

## Usage

Setup your AWS credentials :

    $ export AWS_ACCESS_KEY_ID='AKID'
    $ export AWS_SECRET_ACCESS_KEY='SECRET'

Creates a KMS key via the AWS Console and store its ID (a UUID) in an environment variable:

    $ export ENIGMA_KEYID = "be3338c8-4e4e-2384-644f-1ac0af044fe4"

Initialize your bucket into S3 :

    $ enigma create -bucket enigma-ex
    2015/08/19 01:15:00 Create bucket : enigma-ex
    2015/08/19 01:15:01 {
        Location: "http://enigma-ex.s3.amazonaws.com/"
    }

Store a text :

    $ bin/enigma -bucket enigma-ex -put-text -text foo -key test
    2015/08/19 01:15:33 Encrypt text:  foo
    2015/08/19 01:15:33 Encrypted:  [10 32 186 66 75 45 140 19 56 122 39 234 232 28 64 15 92 9 186 15 190 53 46 133 167 224 203 218 32 83 220 212 15 214 18 138 1 1 1 2 0 120 186 66 75 45 140 19 56 122 39 234 232 28 64 15 92 9 186 15 190 53 46 133 167 224 203 218 32 83 220 212 15 214 0 0 0 97 48 95 6 9 42 134 72 134 247 13 1 7 6 160 82 48 80 2 1 0 48 75 6 9 42 134 72 134 247 13 1 7 1 48 30 6 9 96 134 72 1 101 3 4 1 46 48 17 4 12 249 162 128 88 135 114 55 206 77 142 214 84 2 1 16 128 30 39 241 164 53 219 182 228 33 128 39 85 36 255 61 183 121 158 165 245 80 188 197 243 78 3 142 21 17 163 229]
    2015/08/19 01:15:34 Successfully  uploaded data with key test

Check your enigma :

    $ bin/enigma -bucket enigma-ex -list
    2015/08/19 01:15:49 Files:
    2015/08/19 01:15:49 Size: 1
    2015/08/19 01:15:49 Object:  test

Retrieve your data :

    $ bin/enigma -bucket enigma-ex -get-text -key test
    2015/08/19 09:58:39 Retrive text for key : test
    2015/08/19 09:58:40 Successfully decrypted: foo


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



[Amazon S3]:https://aws.amazon.com/s3/
[Amazon KMS]: https://aws.amazon.com/kms/
