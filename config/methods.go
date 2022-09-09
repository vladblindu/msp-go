package config

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"path"
)

func (cfg *Config) getConfig() {
	configPath := path.Join(rootDir(), "config.yaml")
	yamlFile, err := os.ReadFile(configPath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}
	err = yaml.Unmarshal(yamlFile, cfg)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}
}

func (cfg *Config) GetMspCacheCredentials() *Config {

	mspCachePath := path.Join(userDir(), "msp.cache")
	yamlFile, err := os.ReadFile(mspCachePath)
	if err != nil {
		log.Printf("yamlFile.Get err   #%v ", err)
	}

	mspData := MspCacheCredentials{}
	err = yaml.Unmarshal(yamlFile, &mspData)
	if err != nil {
		log.Fatalf("Unmarshal: %v", err)
	}

	cfg.MspCacheUser = mspData.MspCacheUser
	cfg.MspCachePassword = mspData.MspCachePassword
	cfg.MspCacheUri = mspData.MspCacheUri

	return cfg
}
