package models

import (
	"github.com/gocolly/colly"
	"github.com/nicolasvasquez/workerPrice/internal/utils"
)

type Products struct {
	Url            string
	MarketplacesId int
	Hist           []HistProducts
}

func (product *Products) Crawl() {
	c := colly.NewCollector()

	c.OnHTML(utils.Config.Selectors.MercadoLibre.Price, func(e *colly.HTMLElement) {
		product.Hist = append(product.Hist, HistProducts{Price: e.Text})
	})

	c.Visit(product.Url)
}
