package common

import (
	"fmt"
	"github.com/BurntSushi/toml"
	"path/filepath"
	"sync"
)

type yace struct {
	ObjectSize int64	`toml:"objectsize"`
	ObjectCount int64 `toml:"objectcount"`
	Parallel int64 `toml:"parallel"`
	Option string `toml:"option"`
	DataDir string `toml:"datadir"`
	LogFile string `toml:"logfile"`
}


type ceph_server struct {
	AccessKey string `toml:"access_key"`
	SecretKey string `toml:"secret_key"`
	Address string `toml:"address"`
	Port int `toml:"port"`
	Bucket string `toml:"bucket"`
}

type TomlConfig struct {
	Yace yace
	Ceph_server ceph_server
}

var (
	conf *TomlConfig
	once sync.Once
)
func Config(configFile string) *TomlConfig {
	// get abs dir of configFile
	file, err := filepath.Abs(configFile)
	if err != nil {
		fmt.Println("get abs dir of configFile is bad", err)
		return nil
	}
	// 使用单例模式读取configFile
	once.Do(func() {
		_, err1 := toml.DecodeFile(file, &conf)
		if err1 != nil {
			fmt.Println("tomlConfig from configFile is bad", err1)
			return
		}
	})
	return conf
}