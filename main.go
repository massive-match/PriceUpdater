package main

import (
	"github.com/nicolasvasquez/workerPrice/internal/app"
	"os"
)

func main() {
	arg := ""
	if len(os.Args) > 1 {
		arg = os.Args[1]
	}
	app.Init(arg)
}
