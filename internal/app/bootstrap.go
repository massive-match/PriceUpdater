package app

import (
	"github.com/nicolasvasquez/workerPrice/internal/utils"
	"github.com/nicolasvasquez/workerPrice/internal/crawler"
)

func Init(routeConfig string) {
	utils.LoadConfig(routeConfig)
	crawler.Init()
}
