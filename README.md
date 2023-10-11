# GoProxy - A CLI for [GoProxy](https://goproxy.cn/)

[![PkgGoDev](https://pkg.go.dev/badge/github.com/go-zoox/goproxy)](https://pkg.go.github.com/go-zoox/goproxy)
[![Build Status](https://github.com/go-zoox/goproxy/actions/workflows/release.yml/badge.svg?branch=master)](httpgithub.com/go-zoox/goproxy/actions/workflows/release.yml)
[![Go Report Card](https://goreportcard.com/badge/github.com/go-zoox/goproxy)](https://goreportcard.com/repgithub.com/go-zoox/goproxy)
[![Coverage Status](https://coveralls.io/repos/github/go-zoox/goproxy/badge.svg?branch=master)](https://coveralls.io/github/go-zoox/goproxy?branch=master)
[![GitHub issues](https://img.shields.io/github/issues/go-zoox/goproxy.svg)](https://github.com/go-zoox/goproxy/issues)
[![Release](https://img.shields.io/github/tag/go-zoox/goproxy.svg?label=Release)](https://github.com/go-zoox/goproxy/tags)


## Installation
To install the package, run:
```bash
go install github.com/go-zoox/goproxy/cmd/goproxy@latest
```

## Quick Start

```bash
# start goproxy, cached in memory, default udp port: 8080
goproxy

# start goproxy with config (see conf/goproxy.yml for more options)
goproxy -c goproxy.yml
```

## Configuration
See the [configuration file](conf/goproxy.yml).

## License
GoZoox is released under the [MIT License](./LICENSE).
