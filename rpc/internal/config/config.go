package config

import (
	"github.com/neumathe/simple-admin-common/plugins/casbin"
	"github.com/zeromicro/go-zero/zrpc"

	"github.com/neumathe/simple-admin-common/config"
)

type Config struct {
	zrpc.RpcServerConf
	DatabaseConf config.DatabaseConf
	CasbinConf   casbin.CasbinConf
	RedisConf    config.RedisConf
}
