package db

import (
	"fmt"
	"msp-go/config"
)

func createUri(config config.Config) string {
	return fmt.Sprintf(
		config.MspCacheUri,
		config.MspCacheUser,
		config.MspCachePassword,
	)
}
