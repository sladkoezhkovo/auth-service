package configs

import (
	"errors"
	"github.com/ilyakaznacheev/cleanenv"
)

func SetupConfig(path string, config interface{}) error {
	if err := cleanenv.ReadConfig(path, config); err != nil {
		return errors.New("cannot read config")
	}
	return nil
}
