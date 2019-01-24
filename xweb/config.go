package xweb

import (
	"github.com/hebl/gox/utils"

	log "github.com/sirupsen/logrus"
)

//CommonConfig 基础配置信息
type CommonConfig struct {
	Name         string    `json:"name"`
	Version      string    `json:"version"`
	Host         string    `json:"host"`
	Port         int       `json:"port"`
	Logging      bool      `json:"logging"`
	LogLevel     log.Level `json:"log_level"`
	LogFile      string    `json:"log_file"`
	CookieSecret string    `json:"cookie_secret"`
	Static       []struct {
		URI        string `json:"uri"`
		Filesystem string `json:"filesystem"`
	} `json:"static"`
	CertFile string `json:"cert_file"`
	KeyFile  string `json:"key_file"`
}

// LogLevel
/**
  0 PanicLevel
  1 FatalLevel
  2 ErrorLevel
  3 WarnLevel
  4 InfoLevel
  5 DebugLevel
*/

//DBConfig 数据库配置
type DBConfig struct {
	Database string `json:"database"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	PoolSize int    `json:"poolsize"`
}

//Config 系统配置信息
// Load Config : utils.ParseJSONFile(filename, &Config)
type Config struct {
	CommonConfig
	Database DBConfig `json:"database"`
}

// NewConfig 根据文件生成一个配置信息
func NewConfig(configFile string) *Config {
	var config Config
	utils.ParseJSONFile(configFile, &config)

	return &config
}
