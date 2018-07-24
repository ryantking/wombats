package logging

import (
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

const errorRegex = `^([\/\w+]+\.[d|s|h]ats): \d+\(line=(\d+), offs=(\d+)\) ` +
	`-- \d+\(line=(\d+), offs=(\d+)\): error(\d+): (.+)$`

func parseErrors(lines []string) []*ATSError {
	atsErrors := []*ATSError{}
	var prevError *ATSError

	for i := 0; i < len(lines); i++ {
		atsError := parseError(lines[i])

	}

	return atsErrors
}

func (e *ATSError) equal(e2 *ATSError) bool {
	return e.Fname == e2.Fname &&
		e.StartL == e2.StartL &&
		e.EndL == e2.EndL &&
		e.StartO == e2.StartO &&
		e.EndO == e2.EndO
}

func parseError(line string) *ATSError {
	r, err := regexp.Compile(errorRegex)
	if err != nil {
		return nil
	}

	matches := r.FindStringSubmatch(line)
	if len(matches) != 7 {
		return nil
	}

	wd, err := os.Getwd()
	if err != nil {
		return nil
	}

	fname, err := filepath.Rel(wd, matches[1])
	if err != nil {
		return nil
	}

	startL, err := strconv.Atoi(matches[2])
	if err != nil {
		return nil
	}
	startO, err := strconv.Atoi(matches[3])
	if err != nil {
		return nil
	}
	endL, err := strconv.Atoi(matches[4])
	if err != nil {
		return nil
	}
	endO, err := strconv.Atoi(matches[5])
	if err != nil {
		return nil
	}

	return &ATSError{
		Fname:  fname,
		StartL: startL,
		EndL:   endL,
		StartO: startO,
		EndO:   endO,
		Errors: []string{matches[6]},
	}
}
