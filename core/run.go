package core

import (
	"fmt"
	"os"

	"github.com/go-zoox/zoox"
	"github.com/goproxy/goproxy"
)

func (c *core) Run() error {
	c.app.Use(func(ctx *zoox.Context) {
		zoox.WrapH(&goproxy.Goproxy{
			GoBinEnv: append(
				os.Environ(),
				fmt.Sprintf("HTTPS_PROXY=%s", c.cfg.Proxy),
				fmt.Sprintf("GOPROXY=%s", c.cfg.Upstream),
				fmt.Sprintf("GOPRIVATE=%s", c.cfg.Private), // Solve the problem of pulling private modules
			),
			ProxiedSUMDBs: []string{
				// Proxy the default checksum database
				// "sum.golang.org https://goproxy.cn/sumdb/sum.golang.org",
				c.cfg.SumDB,
			},
		})(ctx)
	})

	return c.app.Run(fmt.Sprintf(":%d", c.cfg.Port))
}
