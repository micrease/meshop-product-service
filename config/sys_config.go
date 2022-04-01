package config

import (
	"github.com/micrease/micrease-core/config"
	"os"
	"sync"
)

//继承CommonConfig此基础上扩展
type SysConfig struct {
	config.CommonConfig `yaml:",inline"`
}

var once sync.Once
var sysConfig *SysConfig

func InitSysConfig() *SysConfig {
	once.Do(func() {
		sysConfig = new(SysConfig)
		err := config.LoadConfigTo(sysConfig)
		if err != nil {
			os.Exit(1)
		}
	})
	return sysConfig
}

func Get() *SysConfig {
	return sysConfig
}
