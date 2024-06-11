package configer

import (
	"errors"
	"github.com/imdario/mergo"
)

type ConfigLoader[T any] interface {
	Parse(*T) error
}

// For is a function to initialize a config factory for the given type.
func For[T any](config T) *Configer[T] {
	return &Configer[T]{
		configStruct: &config,
	}
}

// Parse is a helper function for simple case that takes a pointer to a struct as default config and load environment variables into it.
func Parse[T any](config *T) error {
	if config == nil {
		return errors.New("input parameter is nil")
	}

	c, err := For(*config).LoadEnv().Parse()
	if err != nil {
		return err
	}
	*config = c
	return nil
}

type Configer[T any] struct {
	configStruct *T
	loaders      []ConfigLoader[T]
}

func (c *Configer[T]) AddLoader(loader ConfigLoader[T]) *Configer[T] {
	c.loaders = append(c.loaders, loader)
	return c
}

func (c *Configer[T]) LoadYamlFile(path string) *Configer[T] {
	return c.AddLoader(YamlLoaderFromFile[T](path))
}

func (c *Configer[T]) LoadYamlBytes(bs []byte) *Configer[T] {
	return c.AddLoader(YamlLoaderFromBytes[T](bs))
}

func (c *Configer[T]) LoadEnv() *Configer[T] {
	return c.AddLoader(NewEnvLoader[T]())
}

func (c *Configer[T]) Parse() (T, error) {
	cfg := *c.configStruct

	for i := 0; i < len(c.loaders); i++ {
		var t T
		err := c.loaders[i].Parse(&t)
		if err != nil {
			return cfg, err
		}
		err = mergo.Merge(&cfg, t, mergo.WithOverride)
		if err != nil {
			return cfg, err
		}
	}
	return cfg, nil
}
