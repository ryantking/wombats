package ats

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"github.com/RyanTKing/wombats/pkg/logging"
	log "github.com/sirupsen/logrus"
)

func checkDir(dir string, lastBuilt time.Time) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Debugf("error reading directory: %s: %s", dir, err)
		log.Fatalf("could not read directory: %s", dir)
	}

	for _, f := range files {
		if f.Mode().IsDir() {
			newDir := fmt.Sprintf("%s%s/", dir, f.Name())
			if changed := checkDir(newDir, lastBuilt); changed {
				return true
			}
			continue
		}

		if strings.Contains(filepath.Ext(f.Name()), "ats") &&
			f.ModTime().After(lastBuilt) {

			return true
		}
	}

	return false
}

func shouldBuild(execFile string, small bool) bool {
	execInfo, err := os.Stat(execFile)
	if err != nil {
		return true
	}

	return checkDir("./", execInfo.ModTime())
}

func removeDuplicates(args []string) []string {
	seen := map[string]bool{}
	i := 0
	for _, arg := range args {
		if _, ok := seen[arg]; ok {
			continue
		}
		seen[arg] = true
		args[i] = arg
		i++
	}

	return args[:i]
}

func getArgs(execFile, entryPoint string, clibs []string) []string {
	args := []string{"-w", "-o", execFile, entryPoint, "-DATS_MEMALLOC_LIBC"}
	if len(clibs) == 0 {
		return args
	}

	pkgCfgArgs := append(clibs, "--cflags", "--libs")
	cmd := exec.Command("pkg-config", pkgCfgArgs...)
	cflagsRaw, err := cmd.Output()
	if err != nil {
		log.Debugf("pkg-config error: %s", err)
	}

	cflags := strings.Split(strings.TrimSpace(string(cflagsRaw)), " ")
	return append(args, removeDuplicates(cflags)...)
}

func calcTime(d time.Duration) string {
	if d.Hours() >= 1.0 {
		return fmt.Sprintf("%0.2fh", d.Hours())
	} else if d.Minutes() >= 1.0 {
		return fmt.Sprintf("%0.2fm", d.Minutes())
	}

	return fmt.Sprintf("%0.2fs", d.Seconds())
}

// Build compiles an ATS project
func Build(name, entryPoint string, clibs []string) string {
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
	args := getArgs(execFile, entryPoint, clibs)
	out, err := ExecPatsccOutput(args...)
	if err != nil {
		logging.CheckErrors(strings.TrimSpace(out))
		log.Fatalln("build failed")
	}
	log.Infof("Finished building in %s", calcTime(time.Since(start)))

	return execFile
}
