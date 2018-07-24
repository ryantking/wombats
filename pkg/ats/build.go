package ats

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

func shouldBuild(execFile string, small bool) bool {
	execInfo, err := os.Stat(execFile)
	if err != nil {
		return true
	}

	wd, err := os.Getwd()
	if err != nil {
		log.Fatalf("could not get current directory")
	}

	files, err := filepath.Glob(filepath.Join(wd, "/*.[dshc]ats"))
	if err != nil {
		log.Fatalf("could not find files in directory")
	}
	if !small {
		files2, err := filepath.Glob(filepath.Join(wd, "/**/*.[dshc]ats"))
		if err != nil {
			log.Fatalf("could not find files in directory")
		}
		files = append(files, files2...)
	}

	for _, file := range files {
		info, err := os.Stat(file)
		if err != nil {
			log.Fatalf("error getting stats on '%s'", file)
		}

		if info.ModTime().After(execInfo.ModTime()) {
			return true
		}
	}

	return false
}

// Build compiles an ATS project
func Build(name, entryPoint string) string {
	start := time.Now()
	execFile := fmt.Sprintf("./BUILD/%s", name)
	small := false
	if !strings.Contains(entryPoint, "./DATS/") {
		execFile = fmt.Sprintf("./%s", name)
		small = true
	}

	if !shouldBuild(execFile, small) {
		return execFile
	}

	log.Infof("Building '%s' project", name)
	if err := ExecPatscc("-o", execFile, entryPoint); err != nil {
		log.Debugf("build error: %s", err)
		log.Fatalln("could not build")
	}

	dur := time.Since(start)
	unit := "s"
	t := dur.Seconds()
	if t >= 60 && t < 3600 {
		unit = "m"
		t = dur.Minutes()
	} else if t >= 36000 {
		unit = "h"
		t = dur.Hours()
	}
	log.Infof("Finished building in %.2f%s", t, unit)

	return execFile
}
