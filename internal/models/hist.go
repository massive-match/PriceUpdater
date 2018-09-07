package models

import (
	"net/http"
	"github.com/nicolasvasquez/workerPrice/internal/utils"
)

type HistProducts struct {
	Price string
}

func (hist *HistProducts) UpdatePrice() error {
	if _, err := http.PostForm(utils.Config.Connections.Base + utils.Config.Connections.Update, hist); err != nil {
		return err
	} else {
		return nil
	}
}
