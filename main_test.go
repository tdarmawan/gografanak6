package main

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGreeter(t *testing.T) {
	d := &Demo{}
	g := d.NewGreeter("Hi")
	assert.Equal(t, "Hi, Test!", g.Greet("Test"))
}

func TestXAdd(t *testing.T) {
	d := &Demo{}
	res, err := d.XAdd(context.Background(), 2, 3)
	assert.NoError(t, err)
	assert.Equal(t, 5, res)

	_, err = d.XAdd(context.Background(), -1, 5)
	assert.Error(t, err)
}
