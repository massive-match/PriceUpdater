package crawler

import (
	"fmt"
	"github.com/nicolasvasquez/workerPrice/internal/models"
)

func Init() {
	var p models.Products

	p.Url = "https://articulo.mercadolibre.cl/MLC-463653695-pack-10-foco-led-panel-placa-18w-embutido-71504-fernapet-_JM"

	p.Crawl()

	fmt.Printf("%v\n", p)
	fmt.Println("TERMINADO\n")
}
