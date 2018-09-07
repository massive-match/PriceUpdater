package app

import (
	"github.com/nicolasvasquez/workerPrice/internal/utils"
	"github.com/nicolasvasquez/workerPrice/internal/models"
	"fmt"
)

func Init(routeConfig string) {
	utils.LoadConfig(routeConfig)
	if products, err := models.GetProducts(); err != nil {
		fmt.Println(err)
	} else {
		for _, product := range products {
			product.Crawl()
			length := len(product.Hist)
			if product.Hist[length - 1].Price != product.Hist[length - 2].Price {
				if err := product.Hist[length - 1].UpdatePrice(); err != nil {
					fmt.Println(err)
				}
			}
		}
	}
	fmt.Println("TERMINADO")
}
