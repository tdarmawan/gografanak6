package demo

import (
	"context"
	"errors"
	"fmt"

	"go.k6.io/k6/js/modules"
)

// Register the module
func init() {
	modules.Register("k6/x/demo", new(Demo))
}

type Demo struct{}

type Greeter struct {
	Prefix string
}

func (d *Demo) NewGreeter(prefix string) *Greeter {
	return &Greeter{Prefix: prefix}
}

func (g *Greeter) Greet(name string) string {
	return fmt.Sprintf("%s, %s!", g.Prefix, name)
}

func (d *Demo) XAdd(ctx context.Context, a, b int) (int, error) {
	if a < 0 || b < 0 {
		return 0, errors.New("negative values not allowed")
	}
	return a + b, nil
}
