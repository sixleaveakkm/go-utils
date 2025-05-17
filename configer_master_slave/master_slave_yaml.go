package configer_master_slave

import (
	"github.com/mitchellh/mapstructure"
	"github.com/pkg/errors"
	"gopkg.in/yaml.v3"
	"reflect"
)

type YamlLoader interface {
	Load(path string) ([]byte, error)
}

func LoadConfig[T any](loader YamlLoader, path string, cfg *T) error {
	v := reflect.ValueOf(cfg).Elem()
	t := v.Type()
	if t.Kind() != reflect.Struct {
		return errors.Errorf("config must be a struct type, got %v", t.Kind())
	}

	masterData := make(map[string]string)
	err := loadYaml(loader, path, &masterData)
	if err != nil {
		return err
	}

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		if !field.IsExported() {
			continue
		}
		fn := field.Name

		fp, ok := masterData[fn]
		if !ok {
			return errors.Errorf("'%s' not found in master data", fn)
		}

		fv := v.Field(i).Addr().Interface()
		/// fv (field value) is a interface{} from the field type.
		/// Directly use it to yaml result fv becomes a map[string]interface{}
		/// Then original field type is not changed.
		/// Using mapstructure to convert map[string]interface{} to field type.

		m := make(map[string]any)
		err = loadYaml(loader, fp, &m)
		if err != nil {
			return errors.Errorf("failed to load '%s' from ssm, path: %s", fn, fp)
		}
		err = mapstructure.Decode(m, fv)
		if err != nil {
			return errors.Errorf("failed to map data '%s', path: %s", fn, fp)
		}

	}

	return nil
}

func loadYaml(loader YamlLoader, path string, obj any) error {
	bytes, err := loader.Load(path)
	if err != nil {
		return errors.WithStack(err)
	}
	err = yaml.Unmarshal(bytes, obj)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}
