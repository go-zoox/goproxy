package main

import (
	"github.com/go-zoox/cli"
	"github.com/go-zoox/config"
	"github.com/go-zoox/core-utils/fmt"
	"github.com/go-zoox/fs"
	api "github.com/go-zoox/goproxy"
	"github.com/go-zoox/goproxy/core"
	"github.com/go-zoox/logger"
)

func main() {
	app := cli.NewSingleProgram(&cli.SingleProgramConfig{
		Name:    "goproxy",
		Usage:   "An Easy Self Hosted API Gateway",
		Version: api.Version,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name: "config",
				// Value:   "conf/goproxy.yaml",
				Usage:   "The path to the configuration file",
				Aliases: []string{"c"},
				// Required: true,
			},
			&cli.Int64Flag{
				Name:    "port",
				Usage:   "The port to listen on",
				Aliases: []string{"p"},
				// Value:   "8080",
			},
			&cli.StringFlag{
				Name:    "upstream",
				Usage:   "Upstream proxy",
				Aliases: []string{"u"},
				EnvVars: []string{"UPSTREAM", "GOPROXY"},
				Value:   "https://proxy.golang.org,direct",
			},
			&cli.StringFlag{
				Name:    "sumdb",
				Usage:   "Upstream sumdb",
				Aliases: []string{"s"},
				EnvVars: []string{"SUMDB", "GOSUMDB"},
				Value:   "sum.golang.org",
			},
			&cli.StringFlag{
				Name:    "private",
				Usage:   "Private modules",
				Aliases: []string{"r"},
				EnvVars: []string{"PRIVATE", "GOPRIVATE"},
			},
			&cli.StringFlag{
				Name:    "proxy",
				Usage:   "The proxy to use when fetching modules.",
				EnvVars: []string{"PROXY", "HTTPS_PROXY", "ALL_PROXY"},
			},
		},
	})

	app.Command(func(c *cli.Context) error {
		configFilePath := c.String("config")
		if configFilePath == "" {
			configFilePath = "/etc/goproxy/config.yaml"
		}

		var cfg core.Config

		if configFilePath != "" {
			if !fs.IsExist(configFilePath) {
				// return fmt.Errorf("config file(%s) not found", configFilePath)
			} else {
				if err := config.Load(&cfg, &config.LoadOptions{
					FilePath: configFilePath,
				}); err != nil {
					return fmt.Errorf("failed to read config file: %s", err)
				}
			}
		}

		if cfg.Port == 0 {
			cfg.Port = 8080
		}

		if c.Int64("port") != 0 {
			cfg.Port = c.Int64("port")
		}

		if c.String("upstream") != "" {
			cfg.Upstream = c.String("upstream")
		}

		if c.String("sumdb") != "" {
			cfg.SumDB = c.String("sumdb")
		}

		if c.String("private") != "" {
			cfg.Private = c.String("private")
		}

		if c.String("proxy") != "" {
			cfg.Proxy = c.String("proxy")
		}

		// @TODO
		if logger.IsDebugLevel() {
			// logger.Debugf("config: %v", cfg)
			fmt.PrintJSON("config:", cfg)
		}

		app, err := core.New(api.Version, &cfg)
		if err != nil {
			return fmt.Errorf("failed to create core: %s", err)
		}

		return app.Run()
	})

	if err := app.RunWithError(); err != nil {
		logger.Fatal("%s", err.Error())
	}
}
