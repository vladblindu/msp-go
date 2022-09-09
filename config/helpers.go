package config

import (
	"log"
	"os/user"
	"path"
	"path/filepath"
	"runtime"
)

func rootDir() string {
	_, b, _, _ := runtime.Caller(0)
	d := path.Join(path.Dir(b))
	return filepath.Dir(d)
}

func userDir() string {
	u, err := user.Current()
	if err != nil {
		log.Fatalf(err.Error())
	}
	return u.HomeDir
}
