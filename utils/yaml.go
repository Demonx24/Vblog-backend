package utils

import "os"

const configFile = "E:\\go代码\\Vblog-backend\\config.yaml"

func LoadConfig() ([]byte, error) {
	return os.ReadFile(configFile)
}
