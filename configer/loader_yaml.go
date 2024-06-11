package configer

import (
	"github.com/sixleaveakkm/go-utils/ptr"
	"gopkg.in/yaml.v3"
)

type YamlLoader[T any] struct {
	file  *string
	bytes []byte
}

func YamlLoaderFromFile[T any](path string) ConfigLoader[T] {
	return &YamlLoader[T]{file: ptr.String(path)}
}

func YamlLoaderFromBytes[T any](bs []byte) ConfigLoader[T] {
	return &YamlLoader[T]{bytes: bs}
}

func (y *YamlLoader[T]) Parse(c *T) error {
	return yaml.Unmarshal(y.bytes, c)
}
