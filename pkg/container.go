package pkg

import (
	"reflect"

	"github.com/sirupsen/logrus"
)

var AppContainer *Container

func init() {
	AppContainer = &Container{Services: make(map[string]any)}
}

type Container struct {
	Services map[string]any
}

func (c *Container) Register(name string, service any) {
	if c.Has(name) || c.HasType(name, reflect.TypeOf(service).Elem()) {
		t := reflect.TypeOf(service).Elem()
		logrus.Fatalf("Service %s of type %s already exists", name, t.Name())
	}
	c.Services[name] = service
}

func (c *Container) Get(name string) any {
	if !c.Has(name) {
		logrus.Fatalf("Service %s not found", name)
	}
	return c.Services[name]
}

func (c *Container) Remove(name string) {
	delete(c.Services, name)
}

func (c *Container) Has(name string) bool {
	_, ok := c.Services[name]
	return ok
}

func (c *Container) HasType(name string, t reflect.Type) bool {
	if !c.Has(name) {
		return false
	}
	service := c.Get(name)
	return reflect.TypeOf(service).Elem() == t
}
