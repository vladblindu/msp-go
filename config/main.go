package config

func New(noDb bool) *Config {
	cfg := new(Config)
	cfg.getConfig()
	if !noDb {
		cfg.GetMspCacheCredentials()
	}
	return cfg
}
