package subcommands

import (
	"os"
)

var (
	patshome string
	patscc   string
)

func init() {
	patshome = os.Getenv("PATSHOME")
	patscc = patshome + "/bin/patscc"
}
