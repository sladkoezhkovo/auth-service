package configs

import "github.com/ilyakaznacheev/cleanenv"

type AppConfig struct {
	Name  string      `yaml:"name"`
	Env   string      `yaml:"env"`
	Redis RedisConfig `yaml:"redis"`
	Pg    SqlConfig   `yaml:"pg"`
	Jwt   JwtConfig   `yaml:"jwt"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   int    `yaml:"db" env-default:"0"`
}

type SqlConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   string `yaml:"db"`
	TLS  string `yaml:"tls" env-default:"disabled"`
}

type JwtConfig struct {
	AccessTTL  int `yaml:"accessTTL"`
	RefreshTTL int `yaml:"refreshTTL"`
}

func SetupConfig(path string) *AppConfig {
	var config AppConfig

	if err := cleanenv.ReadConfig(path, &config); err != nil {
		panic("cannot read config")
	}

	return &config
}
