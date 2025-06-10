package config

type Config struct {
	Mysql  Mysql  `yaml:"mysql" json:"mysql"`
	System System `yaml:"system" json:"system"`
	Jwt    Jwt    `yaml:"jwt" json:"jwt"`
	Redis  Redis  `yaml:"redis" json:"redis"`
	Zap    Zap    `json:"zap" yaml:"zap"`
}
