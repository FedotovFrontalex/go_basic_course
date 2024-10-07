package api

import (
	"jsonBin/config"
	"jsonBin/print"
)

func Init(config *config.Config) {
	print.Success(config)
}
