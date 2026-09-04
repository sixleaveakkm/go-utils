# lambdarun

Starts a function as an AWS Lambda handler, or invokes it once against
`LOCAL_PAYLOAD` when the `LOCAL` environment variable is `1`.

The handler signature is anything `github.com/aws/aws-lambda-go` accepts; both
modes use the same dispatch, so a local run behaves like a real invocation.

```go
func handle(ctx context.Context, in Request) (Response, error) { ... }

func main() {
	if err := lambdarun.Start(context.Background(), handle); err != nil {
		log.Fatalln(err)
	}
}
```

```console
$ LOCAL=1 LOCAL_PAYLOAD='{"name":"sixleaf"}' go run .
{"greeting":"hello sixleaf"}
```

Without `LOCAL_PAYLOAD` the handler is invoked with `null`, which decodes to the
zero value of its input. The response is written to stdout as a single JSON
line.
