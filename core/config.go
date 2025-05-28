package core

import (
	"github.com/Demonx24/Vblog-backend/config"
	"github.com/Demonx24/Vblog-backend/utils"
	"gopkg.in/yaml.v3"
	"log"
)

func InitConfig() *config.Config {
	c := &config.Config{}
	yamlConf, err := utils.LoadConfig()
	if err != nil {
		log.Fatalf("无法加载配置文件: %v", err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		log.Fatalf("无法解析配置文件:%v", err)
	}
	return c
}
