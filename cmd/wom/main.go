package main

import (
	"fmt"
	"os"
)

func build(args []string) error {
	routine := "default"
	if len(args) > 0 {
		routine = args[0]
	}

	_, err := os.Stat("build.hats")
	if os.IsNotExist(err) {
		return fmt.Errorf("no build file defined")
	}
	fmt.Printf("Building: %s\n", routine)

	return nil
}

func usage() {
	fmt.Printf("Usage of %s:\n", os.Args[0])
	fmt.Printf("    %s build <build_routine>\n", os.Args[0])
}

func main() {
	if len(os.Args) == 1 {
		usage()
		os.Exit(0)
	}

	command := os.Args[1]
	var err error

	if command == "build" {
		err = build(os.Args[2:])
	}

	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
