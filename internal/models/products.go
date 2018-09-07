package models

import (
	"github.com/gocolly/colly"
	"github.com/nicolasvasquez/workerPrice/internal/utils"
	"net/http"
	"encoding/json"
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

func GetProducts() ([]Products, error) {
	var products []Products

	if r, err := http.Get(utils.Config.Connections.Base + utils.Config.Connections.Get); err != nil {
		return products, err
	} else {
		if err := json.NewDecoder(r.Body).Decode(&products); err != nil {
			return products, err
		} else {
			return products, nil
		}
	}
}
