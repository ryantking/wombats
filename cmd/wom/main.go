package main

import (
	"os"

	"github.com/RyanTKing/wombats/pkg/womapp"
)

func main() {
	app := womapp.New()
	app.Run(os.Args)
}
