package common

import (
	"os"
	"sync"
)

func env(key, defaultVal string) string {
	val, exists := os.LookupEnv(key)
	if exists {
		return val
	}
	return defaultVal
}

type settings struct {
	config map[string]string
}

var s *settings
var once sync.Once

func GetSettings() *settings {
	once.Do(func() {
		s = &settings{}
		config := make(map[string]string)
		//settings
		config["INFLUXDB_HOST"] = env("INFLUXDB_HOST", "http://127.0.0.1")
		config["INFLUXDB_WRITE_PORT"] = env("INFLUXDB_WRITE_PORT", "18086")
		config["INFLUXDB_READ_PORT"] = env("INFLUXDB_READ_PORT", "18086")
		config["INFLUXDB_DATABASE"] = env("INFLUXDB_DATABASE", "test")
		s.config = config

	})
	return s
}

func (s *settings) Get(key string) (string, bool) {
	val, exists := s.config[key]
	return val, exists
}

func (s *settings) Getv(key string) string {
	return s.config[key]
}
