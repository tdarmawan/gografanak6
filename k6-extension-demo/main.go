// main.go
package demo

import "go.k6.io/k6/js/modules"

func init() {
	modules.Register("k6/x/demo", new(RootModule))
}

type RootModule struct{}

func (r *RootModule) NewModuleInstance() modules.Instance { return &ModuleInstance{} }

type ModuleInstance struct{}

func (m *ModuleInstance) Exports() modules.Exports {
	return modules.Exports{Named: map[string]interface{}{"test": func() string { return "OK" }}}
}
