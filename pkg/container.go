package pkg

var AppContainer *Container

func init() {
	AppContainer = &Container{Services: make(map[string]any)}
}

type Container struct {
	Services map[string]any
}

func (c *Container) Register(name string, service any) {
	c.Services[name] = service
}

func (c *Container) Get(name string) any {
	return c.Services[name]
}

func (c *Container) Remove(name string) {
	delete(c.Services, name)
}

func (c *Container) Has(name string) bool {
	_, ok := c.Services[name]
	return ok
}
