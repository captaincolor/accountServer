package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"strings"
)

type Config struct {
	Name string
}

func Init(cfg string) error {
	c := Config{
		Name: cfg,
	}

	// init config
	if err := c.initConfig(); err != nil {
		return err
	}

	// init log, after initiation of config
	// 需要读取与日志相关的配置
	c.initLog()

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
	viper.SetEnvPrefix("ACCOUNTSERVER") // 读取环境变量前缀, 利用viper从环境变量读取配置时用到
	replacer := strings.NewReplacer(".", "_")
	viper.SetEnvKeyReplacer(replacer)

	// viper解析配置文件
	if err := viper.ReadInConfig(); err != nil {
		return err
	}

	return nil
}

func (c *Config) initLog() {
	passLagerCfg := log.PassLagerCfg{
		Writers:        viper.GetString("log.writers"),
		LoggerLevel:    viper.GetString("log.logger_level"),
		LoggerFile:     viper.GetString("log.logger_file"),
		LogFormatText:  viper.GetBool("log.log_format_text"),
		RollingPolicy:  viper.GetString("log.rollingPolicy"),
		LogRotateDate:  viper.GetInt("log.log_rotate_date"),
		LogRotateSize:  viper.GetInt("log.log_rotate_size"),
		LogBackupCount: viper.GetInt("log.log_backup_count"),
	}

	_ = log.InitWithConfig(&passLagerCfg)
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		log.Infof("Config file changed: %s", in.Name)
	})
}
