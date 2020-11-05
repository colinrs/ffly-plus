package config

import (
	"github.com/colinrs/pkgx/logger"
	"github.com/fsnotify/fsnotify"
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

var (
	// Conf ...
	Conf *Config
)

// Init init config
func Init(confPath string) error {
	err := initConfig(confPath)
	if err != nil {
		return err
	}
	return nil
}

// initConfig init config from conf file
func initConfig(confPath string) error {
	if confPath != "" {
		viper.SetConfigFile(confPath)
	} else {
		viper.AddConfigPath("conf")
		viper.SetConfigName("config.local")
	}
	viper.SetConfigType("json")
	viper.AutomaticEnv()
	if err := viper.ReadInConfig(); err != nil {
		return errors.WithStack(err)
	}

	// parse to config struct
	err := viper.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	logger.Info("config:(%#v)", Conf)
	watchConfig()

	return nil
}

// watchConfig ...
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logger.Info("Config file changed: %s", e.Name)
	})
}

// AppConfig ...
type AppConfig struct {
	Name      string `json:"name"`
	RunMode   string `json:"run_mode"`
	Addr      string `json:"addr"`
	JwtSecret string `json:"jwt_secret"`
	// JWTExpirationTime day
	JwtExpirationTime int `json:"jwt_expiration_time"`
}

// MySQLConfig ...
type MySQLConfig struct {
	Name            string `json:"name"`
	Addr            string `json:"addr"`
	DB              string `json:"db"`
	UserName        string `json:"username"`
	Password        string `json:"password"`
	MaxIdleConn     int    `json:"max_idel_conn"`
	MaxOpenConn     int    `json:"max_open_conn"`
	ConnMaxLifeTime int    `json:"conn_max_lifetime"`
}

// RedisConfig ...
type RedisConfig struct {
	Addr         string `json:"addr"`
	Password     string `json:"password"`
	DB           int    `json:"db"`
	DialTimeout  int    `json:"dial_timeout"`
	ReadTimeout  int    `json:"read_timeout"`
	WriteTimeout int    `json:"write_timeout"`
	PoolSize     int    `json:"pool_size"`
}

// CacheConfig ...
type CacheConfig struct {
	Driver string `json:"driver"`
	Prefix string `json:"prefix"`
}

// SentinelRuleConfig ...
type SentinelRuleConfig struct {
	Resource        string `json:"resource"`
	MetricType      string `json:"metric_type"`
	ControlBehavior string `json:"control_behavior"`
	Count           int64  `json:"count"`
}

// Config global config
// include common and biz config
type Config struct {
	// common
	App           AppConfig            `json:"app"`
	MySQL         MySQLConfig          `json:"mysql"`
	Redis         RedisConfig          `json:"redis"`
	Cache         CacheConfig          `json:"cache"`
	SentinelRules []SentinelRuleConfig `json:"sentinel_rules"`
}
