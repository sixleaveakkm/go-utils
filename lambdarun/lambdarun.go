// Package lambdarun starts a function as an AWS Lambda handler, or invokes it
// once against LOCAL_PAYLOAD when the LOCAL environment variable is "1".
//
// The handler signature is whatever github.com/aws/aws-lambda-go accepts: the
// local path reuses the same dispatch as the Lambda runtime.
package lambdarun

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
)

const (
	LocalEnv      = "LOCAL"
	LocalEnvValue = "1"
	PayloadEnv    = "LOCAL_PAYLOAD"
)

func Start(ctx context.Context, handler any) error {
	if os.Getenv(LocalEnv) != LocalEnvValue {
		lambda.StartWithOptions(handler, lambda.WithContext(ctx))
		return nil
	}

	out, err := lambda.NewHandler(handler).Invoke(ctx, localPayload())
	if err != nil {
		return err
	}

	_, err = fmt.Println(string(out))
	return err
}

func localPayload() []byte {
	payload := []byte(os.Getenv(PayloadEnv))
	if len(bytes.TrimSpace(payload)) == 0 {
		// Decodes to the zero value of the handler's input.
		return []byte("null")
	}
	return payload
}
