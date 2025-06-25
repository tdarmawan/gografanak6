package demo

import (
	"go.k6.io/k6/js/modules"
)

func init() {
	modules.Register("k6/x/minimal", new(Module))
}

type Module struct{}

func (m *Module) NewModuleInstance() modules.Instance {
	return &Instance{}
}

type Instance struct{}

func (i *Instance) Exports() modules.Exports {
	return modules.Exports{
		Named: map[string]interface{}{
			"test": func() string { return "OK" },
		},
	}
}
