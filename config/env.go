package config

type Env struct {
	ApplicationPort string `env:"PORT"`
	DataBaseHost string `env:"DB_HOST"`
	DataBasePort int `env:"DB_PORT"`
	DataBaseName string `env:"DB_NAME"`
	DataBasePassword string `env:"DB_PWD"`
	DataBaseUser string `env:"DB_USER"`
	DataBaseSchema string `env:"DB_SCHEMA"`
}