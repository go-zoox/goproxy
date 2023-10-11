package core

type Config struct {
	Port int64 `config:"port,default=8080"`
	//
	Upstream string `config:"upstream,default=https://proxy.golang.org,direct"`
	//
	SumDB string `config:"sumdb,default=sum.golang.org"`
	//
	Private string `config:"private,default=*.corp.example.com|*.cn"`
	//
	Proxy string `config:"proxy,default=direct"`
}
