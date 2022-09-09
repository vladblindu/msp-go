package config

type Config struct {
	MspCacheCredentials
	MspCacheCredentialsFileName   string `yaml:"mspCacheFileName"`
	MspCacheCredentialsDatabase   string `yaml:"mspCacheDatabase"`
	MspCacheCredentialsCollection string `yaml:"mspCacheCollection"`
}

type MspCacheCredentials struct {
	MspCacheUser     string `yaml:"cacheUser"`
	MspCachePassword string `yaml:"cachePassword"`
	MspCacheUri      string `yaml:"cacheUri"`
}
