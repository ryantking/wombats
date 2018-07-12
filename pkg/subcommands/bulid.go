package subcommands

import (
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli"
)

// Build compiles the current project into a binary.
func Build(c *cli.Context) error {
	fmt.Println("Building...")
	info, err := os.Stat("./BUILD")
	if os.IsNotExist(err) {
		os.Mkdir("./BUILD", os.ModeDir)
	} else if err != nil {
		log.Fatalln("err")
	} else if !info.Mode().IsDir() {
		log.Fatalln("BUILD is a regular file.")
	}

	os.Chdir("./BUILD")

	return nil
}
