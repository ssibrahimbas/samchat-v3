package config

import (
	"domain.com/shared/helper"
	"go.deanishe.net/env"
)

func LoadConfig(c interface{}) {
	err := env.Bind(c)
	helper.CheckErr(err)
}
