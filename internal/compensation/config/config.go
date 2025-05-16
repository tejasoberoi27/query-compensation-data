package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigtoml"
)

type Postgres struct {
	Host     string `toml:"host" json:"host" default:"127.0.0.1" env:"HOST"`
	Database string `toml:"database" json:"database" default:"postgres" env:"DATABASE"`
	Port     int    `toml:"port" json:"port" default:"5432" env:"PORT"`
	User     string `toml:"user" json:"user" default:"admin" env:"USER"`
	Password string `toml:"password" json:"password" default:"root" env:"PASSWORD"`
}

type Config struct {
	Postgres Postgres `toml:"postgres" env:"POSTGRES"`
}

func New() (Config, error) {
	var cfg Config

	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		AllowUnknownEnvs: true,
		SkipFlags:        true,
		EnvPrefix:        "COMPENSATION",
		Files:            []string{"./config/compensation/config.toml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".toml": aconfigtoml.New(),
		},
	})

	if err := loader.Load(); err != nil {
		return Config{}, err
	}

	return cfg, nil
}
