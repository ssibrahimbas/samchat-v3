package config

import (
	"github.com/ssibrahimbas/samchat-v3.shared/pkg/helper"
	"go.deanishe.net/env"
)

func LoadConfig(c interface{}) {
	err := env.Bind(c)
	helper.CheckErr(err)
}
