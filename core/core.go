package core

// References:
//   https://www.cnblogs.com/zyndev/p/14454891.html
//   https://h1z3y3.me/posts/simple-and-powerful-reverse-proxy-in-golang/
//   https://segmentfault.com/a/1190000039778241

import (
	"github.com/go-zoox/zoox"
	"github.com/go-zoox/zoox/defaults"
)

type Core interface {
	Version() string
	Run() error
}

type core struct {
	app *zoox.Application

	version string
	cfg     *Config
}

func New(version string, cfg *Config) (Core, error) {
	c := &core{
		app: defaults.Default(),
		//
		version: version,
		//
		cfg: cfg,
	}

	return c, nil
}

func (c *core) Version() string {
	return c.version
}
