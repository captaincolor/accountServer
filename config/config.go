package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		cfg,
	}

	// init config
	if err := c.initConfig(); err != nil {
		return err
	}

	// 监控配置文件，热加载程序(不重启api进程,使api加载最新配置项的值)
	c.watchConfig()

	return nil
}

func (c *Config) initConfig() error {
	if c.Name != "" {
		viper.SetConfigFile(c.Name) // 解析指定配置文件
	} else {
		viper.AddConfigPath("conf") // 解析默认配置文件
		viper.SetConfigName("config")
	}

	viper.SetConfigType("yaml")         // 设置配置文件格式
	viper.AutomaticEnv()                // 读取匹配的环境变量
	viper.SetEnvPrefix("ACCOUNTSERVER") // 读取环境变量前缀
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Printf("Config file changed: %s", in.Name)
	})
}
