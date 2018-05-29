package registry

import (
	"github.com/pkg/errors"
	"reflect"
)

const (
	repositoryName = "go-struct-registry"
)

var (
	registeredStructs = make(map[string]reflect.Type)
)

func RegisterStruct(name string, st interface{}) error {
	if _, ok := registeredStructs[name]; ok {
		return errors.Errorf(repositoryName+": struct already registered: %s", name)
	}

	registeredStructs[name] = reflect.TypeOf(st)

	return nil
}

func NewStruct(name string) (interface{}, error) {
	if _, ok := registeredStructs[name]; !ok {
		return nil, errors.Errorf(repositoryName+": struct not found: %s", name)
	}

	return reflect.New(registeredStructs[name]).Elem().Interface(), nil
}
