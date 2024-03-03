package configs

type Config struct {
	App   AppConfig   `yaml:"app"`
	Redis RedisConfig `yaml:"redis"`
	Pg    SqlConfig   `yaml:"pg"`
	Jwt   JwtConfig   `yaml:"jwt"`
}

type AppConfig struct {
	Name string `yaml:"name"`
	Env  string `yaml:"env"`
	Port int    `yaml:"port"`
}

type RedisConfig struct {
	Host string `yaml:"host"`
	Port string `yaml:"port"`
	Db   int    `yaml:"db" env-default:"0"`
}

type SqlConfig struct {
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
	Db   string `yaml:"db"`
	TLS  string `yaml:"tls" env-default:"disabled"`
}

type JwtConfig struct {
	AccessTTL  int `yaml:"accessTTL"`
	RefreshTTL int `yaml:"refreshTTL"`
}
