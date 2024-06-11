package ssmenvs

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-sdk-go-v2/aws"
	awsconfig "github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/ssm"
	"github.com/pkg/errors"
)

type SSMEnvs struct {
	ssm    *ssm.Client
	config Config
}

type Config struct {
	Prefix   string // Prefix for SSM parameters. Default is "SSM_" if empty.
	Override bool   // Override existing env vars with same name if set to true.
}

var DefaultConfig = Config{
	Prefix:   "SSM_",
	Override: false,
}

func New(region string, config Config) SSMEnvs {
	awsCfg, err := awsconfig.LoadDefaultConfig(context.TODO(), awsconfig.WithRegion(region))
	if err != nil {
		panic(errors.WithStack(err))
	}
	return SSMEnvs{
		ssm:    ssm.NewFromConfig(awsCfg),
		config: config,
	}
}

func (se SSMEnvs) Prefix() string {
	if se.config.Prefix == "" {
		return "SSM_"
	}
	return se.config.Prefix
}

func (se SSMEnvs) Load() map[string]string {
	newEnvs := make(map[string]string)
	for _, s := range os.Environ() {
		if strings.HasPrefix(s, se.Prefix()) {
			a := strings.TrimPrefix(s, se.Prefix())
			pair := strings.SplitN(a, "=", 2)
			key, value := pair[0], pair[1]

			out, err := se.ssm.GetParameter(context.TODO(), &ssm.GetParameterInput{
				Name:           aws.String(value),
				WithDecryption: aws.Bool(true),
			})
			if err != nil {
				panic(errors.WithStack(err))
			}
			if out != nil && out.Parameter != nil && out.Parameter.Value != nil {
				newEnvs[key] = *out.Parameter.Value
			}
		}
	}
	return newEnvs
}

func (se SSMEnvs) LoadAndSet() {
	newEnvs := se.Load()
	for k, v := range newEnvs {
		if v != "" && (se.config.Override || os.Getenv(k) == "") {
			_ = os.Setenv(k, v)
		}
	}
}
