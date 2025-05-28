package config

import (
	"gorm.io/gorm/logger"
	"strconv"
	"strings"
)

type Mysql struct {
	Host         string `yaml:"host" json:"host"`
	Port         int    `yaml:"port" json:"port"`
	Config       string `yaml:"config" json:"config"`
	DBName       string `yaml:"db_name" json:"db_name"`
	Username     string `yaml:"username" json:"username"`
	Password     string `yaml:"password" json:"password"`
	MaxIdleConns int    `yaml:"max_idle_conns" json:"max_idle_conns"`
	MaxOpenConns int    `yaml:"max_open_conns" json:"max_open_conns"`
	LogMode      string `yaml:"log_mode" json:"log_mode"`
}

func (m Mysql) Dsn() string {
	return m.Username + ":" + m.Password + "@tcp(" +
		m.Host + ":" + strconv.Itoa(m.Port) + ")/" +
		m.DBName + "?" + m.Config
}

func (m Mysql) LogLevel() logger.LogLevel {
	switch strings.ToLower(m.LogMode) {
	case "silent", "Silent":
		return logger.Silent
	case "error", "Error":
		return logger.Error
	case "warn", "Warn":
		return logger.Warn
	case "info", "Info":
		return logger.Info
	}
	return logger.Info
}
