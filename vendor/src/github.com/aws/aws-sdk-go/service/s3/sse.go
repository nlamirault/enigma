package s3

import (
	"crypto/md5"
	"encoding/base64"

	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/aws/service"
)

var errSSERequiresSSL = awserr.New("ConfigError", "cannot send SSE keys over HTTP.", nil)

func validateSSERequiresSSL(r *service.Request) {
	if r.HTTPRequest.URL.Scheme != "https" {
		p := awsutil.ValuesAtPath(r.Params, "SSECustomerKey||CopySourceSSECustomerKey")
		if len(p) > 0 {
			r.Error = errSSERequiresSSL
		}
	}
}

func computeSSEKeys(r *service.Request) {
	headers := []string{
		"x-amz-server-side-encryption-customer-key",
		"x-amz-copy-source-server-side-encryption-customer-key",
	}

	for _, h := range headers {
		md5h := h + "-md5"
		if key := r.HTTPRequest.Header.Get(h); key != "" {
			// Base64-encode the value
			b64v := base64.StdEncoding.EncodeToString([]byte(key))
			r.HTTPRequest.Header.Set(h, b64v)

			// Add MD5 if it wasn't computed
			if r.HTTPRequest.Header.Get(md5h) == "" {
				sum := md5.Sum([]byte(key))
				b64sum := base64.StdEncoding.EncodeToString(sum[:])
				r.HTTPRequest.Header.Set(md5h, b64sum)
			}
		}
	}
}
