package cmd

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/RyanTKing/wombats/pkg/ats"
	"github.com/RyanTKing/wombats/pkg/config"
	"github.com/fatih/color"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// ATSError holds info about a line error
type ATSError struct {
	fname                      string
	startL, endL, startO, endO int
	errors                     []string
}

// checkCmd represents the build command
var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Typecheck the current project",
	Long: `Use patscc to typecheck the project, but not compile it and report
any errors.

All arguments passed to the command will be added as arguments to the patscc
compiler command.`,
	Run: runCheck,
}

func init() {
	rootCmd.AddCommand(checkCmd)
}

func runCheck(cmd *cobra.Command, args []string) {
	config, err := config.Read()
	if err != nil {
		log.Debugf("error reading config: %s", err)
		log.Fatalf(
			"could not find '%s' in this directory or any parent directory",
			"Wombats.toml",
		)
	}

	log.Infof("Typechecking '%s' project", config.Package.Name)
	output, err := ats.ExecPatsccOutput("-tcats", "./**/*.dats", "./**/*.sats")
	if err == nil {
		log.Info("Found no typechecking errors")
		return
	}

	atsErrors := ATSError{}
	prevError := ATSError{}
	lines := strings.Split(strings.TrimSpace(output), "\n")
	for i := 0; i < len(lines); i++ {
		atsError := parseError(lines[i])
	}
	for _, line := range strings.Split(strings.TrimSpace(output), "\n") {
		if err := formatLine(line); err != nil {
			log.Errorf("could not parse line: %s", line)
			continue
		}
	}

}

func sameLine(err1, err2 ATSError) bool {
	return err1.startL == err2.startL &&
		err1.endL == err2.endL &&
		err1.startO == err2.startO &&
		err1.endO == err2.endO
}

func parseError(line string) ATSError {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	r, err := regexp.Compile(`(^[\/\w+]+.[d|s|h]ats): \d+\(line=(\d+), ` +
		`offs=(\d+)\) -- \d+\(line=(\d+), offs=(\d+)\): error(\d+): (.+)$`)
	if err != nil {
		return err
	}

	matches := r.FindStringSubmatch(line)
	if len(matches) == 6 {
		fname, err := filepath.Rel(wd, matches[1])
		if err != nil {
			return err
		}

		startL, err := strconv.Atoi(matches[2])
		if err != nil {
			return err
		}
		startO, err := strconv.Atoi(matches[3])
		if err != nil {
			return err
		}
		endL, err := strconv.Atoi(matches[4])
		if err != nil {
			return err
		}
		endO, err := strconv.Atoi(matches[5])
		if err != nil {
			return err
		}

		err = formatFileline(fname, startL-1, endL-1, startO-1, endO-1)
		if err != nil {
			return err
		}
	}

	return nil
}

func printLineNum(n, max int) {
	nStr := strconv.Itoa(n)
	for i := 0; i < max-len(nStr); i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d ", n)
}

func formatFileline(fname string, startL, endL, startO, endO int) error {
	f, err := os.Open(fname)
	if err != nil {
		return err
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	max := len(strconv.Itoa(endL))
	errfmt := color.New(color.FgRed).Add(color.Bold)
	printLineNum(startL, max)
	fmt.Print(lines[startL][:startO])
	if endL == startL {
		errfmt.Printf("%s", lines[startL][startO:endO])
	} else {
		errfmt.Printf("%s\n", lines[startL][startO:])
		for i, badL := range lines[startL+1 : endL-1] {
			printLineNum(i+startL+1, max)
			errfmt.Printf("%s\n", badL)
		}
		printLineNum(endL, max)
		errfmt.Print(lines[endL][:endO])
	}
	fmt.Printf("%s\n", lines[endL][endO:])

	return nil
}

func formatLine(line string) error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	r, err := regexp.Compile(`(^[\/\w+]+.[d|s|h]ats): \d+\(line=(\d+), ` +
		`offs=(\d+)\) -- \d+\(line=(\d+), offs=(\d+)\)`)
	if err != nil {
		return err
	}

	matches := r.FindStringSubmatch(line)
	if len(matches) == 6 {
		fname, err := filepath.Rel(wd, matches[1])
		if err != nil {
			return err
		}

		startL, err := strconv.Atoi(matches[2])
		if err != nil {
			return err
		}
		startO, err := strconv.Atoi(matches[3])
		if err != nil {
			return err
		}
		endL, err := strconv.Atoi(matches[4])
		if err != nil {
			return err
		}
		endO, err := strconv.Atoi(matches[5])
		if err != nil {
			return err
		}

		err = formatFileline(fname, startL-1, endL-1, startO-1, endO-1)
		if err != nil {
			return err
		}
	}

	return nil
}
