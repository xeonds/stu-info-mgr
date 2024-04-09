package config

import "stu-info-mgr/lib"

type Config struct {
	Server struct {
		Port int `yaml:"port"`
	} `yaml:"server"`
	lib.DatabaseConfig
}
