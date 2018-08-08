package ats

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/RyanTKing/wombats/pkg/config"
	"github.com/RyanTKing/wombats/pkg/logging"
	log "github.com/sirupsen/logrus"
)

func checkDir(dir string, lastBuilt time.Time) bool {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Debug(err)
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

func getArgs(execFile string, cfg *config.Config) []string {
	args := append([]string{"-w"}, cfg.Package.PatsccArgs...)
	args = append(args, "-o", execFile, cfg.Package.EntryPoint)

	if len(cfg.Package.Clibs) > 0 {
		pkgCfgArgs := append(cfg.Package.Clibs, "--cflags", "--libs")
		cmd := exec.Command("pkg-config", pkgCfgArgs...)
		cflagsRaw, err := cmd.Output()
		if err != nil {
			log.Debug(err)
		}

		cflags := strings.Split(strings.TrimSpace(string(cflagsRaw)), " ")
		args = append(args, removeDuplicates(cflags)...)
	}

	return append(args, cfg.Package.GccArgs...)
}

func calcTime(d time.Duration) string {
	if d.Hours() >= 1.0 {
		return fmt.Sprintf("%0.2fh", d.Hours())
	} else if d.Minutes() >= 1.0 {
		return fmt.Sprintf("%0.2fm", d.Minutes())
	}

	return fmt.Sprintf("%0.2fs", d.Seconds())
}

func removeBuildFiles(small bool) error {
	if !small {
		if err := os.Chdir("./BUILD"); err != nil {
			return err
		}
	}

	files, err := ioutil.ReadDir("./")
	if err != nil {
		return err
	}

	r := regexp.MustCompile(".+_[d|s]ats\\.c")
	for _, f := range files {
		if r.MatchString(f.Name()) {
			err = os.Remove(f.Name())
			if err != nil {
				return err
			}
		}
	}

	if !small {
		return os.Chdir("../")
	}

	return nil
}

// Build compiles an ATS project
func Build(name string, cfg *config.Config) string {
	start := time.Now()
	execFile := fmt.Sprintf("./BUILD/%s", name)
	small := false
	if !strings.Contains(cfg.Package.EntryPoint, "./DATS/") {
		execFile = fmt.Sprintf("./%s", name)
		small = true
	}

	if !shouldBuild(execFile, small) {
		return execFile
	}

	log.Infof("Building '%s' project", name)
	patshomelocs := os.Getenv("PATSHOMELOCS")
	patshomelocs = fmt.Sprintf("./DEPS:%s", patshomelocs)
	os.Setenv("PATSHOMELOCS", patshomelocs)
	args := getArgs(execFile, cfg)
	out, err := ExecPatsccOutput(args...)
	if err != nil {
		logging.CheckErrors(strings.TrimSpace(out))
		log.Debug(out)
		os.Exit(1)
	}
	err = removeBuildFiles(small)
	if err != nil {
		log.Debug(err)
		log.Warnf("could not remove build files")
	}
	log.Infof("Finished building in %s", calcTime(time.Since(start)))

	return execFile
}
