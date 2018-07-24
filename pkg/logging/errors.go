package logging

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"
)

const (
	lineRegex = `^([\/\w+]+\.[d|s|h]ats): \d+\(line=(\d+), offs=(\d+)\) -- ` +
		`\d+\(line=(\d+), offs=(\d+)\): error\(\d+\): (.+).$`
	typeRegex  = `^The (actual|needed) term is: (.+)$`
	countRegex = `^patsopt\(\w+\): there are \[(\d+)\] errors in total.$`
)

// CheckErrors checks output for errors and prints them
func CheckErrors(output string) {
	lines := strings.Split(output, "\n")

	atsErrors := ParseErrors(lines)
	for _, e := range atsErrors {
		if e.Type() == ErrorCount {
			fmt.Print(" ")
		}
		e.Print()
	}
}

// ParseErrors parses a series of output lines and returns the errors
func ParseErrors(lines []string) []ATSErrorLine {
	atsErrors := []ATSErrorLine{}

	for i := 0; i < len(lines); i++ {
		atsError := parseError(lines[i])
		if atsError != nil {
			atsErrors = append(atsErrors, atsError)
		}
	}

	return atsErrors
}

func parseError(line string) ATSErrorLine {
	if atsErr := parseLineError(line); atsErr != nil {
		return atsErr
	}

	if atsErr := parseTypeError(line); atsErr != nil {
		return atsErr
	}

	if atsErr := parseCountError(line); atsErr != nil {
		return atsErr
	}

	return nil
}

func parseLineError(line string) ATSErrorLine {
	r := regexp.MustCompile(lineRegex)
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

	return &ATSLineError{
		Fname:  fname,
		StartL: startL - 1,
		EndL:   endL - 1,
		StartO: startO - 1,
		EndO:   endO - 1,
		Error:  matches[6],
	}
}

func parseTypeError(line string) ATSErrorLine {
	r := regexp.MustCompile(typeRegex)
	matches := r.FindStringSubmatch(line)
	if len(matches) != 3 {
		return nil
	}

	return &ATSTypeError{
		Status:  matches[1],
		ATSType: matches[2],
	}
}

func parseCountError(line string) ATSErrorLine {
	r := regexp.MustCompile(countRegex)
	matches := r.FindStringSubmatch(line)
	if len(matches) != 2 {
		return nil
	}

	count, err := strconv.Atoi(matches[1])
	if err != nil {
		return nil
	}

	return &ATSCountError{Count: count}
}

func printLineNum(n, max int) {
	printSpaces(false)
	nStr := strconv.Itoa(n)
	for i := 0; i < max-len(nStr); i++ {
		fmt.Print(" ")
	}
	fmt.Printf("%d | ", n)
}

// Print ...
func (e ATSLineError) Print() {
	printSpaces(false)
	errfmt.Printf("error: ")
	fmt.Printf("%s\n", e.Error)

	f, err := os.Open(e.Fname)
	if err != nil {
		return
	}
	defer f.Close()

	var lines []string
	s := bufio.NewScanner(f)
	for s.Scan() {
		lines = append(lines, s.Text())
	}

	max := len(strconv.Itoa(e.EndL))
	printLineNum(e.StartL, max)
	fmt.Print(lines[e.StartL][:e.StartO])
	if e.StartL == e.EndL {
		errfmt.Print(lines[e.StartL][e.StartO:e.EndO])
	} else {
		errfmt.Printf("%s\n", lines[e.StartL][e.StartO:])
		for i, badL := range lines[e.StartL+1 : e.EndL-1] {
			printLineNum(i+e.StartL+1, max)
			errfmt.Printf("%s\n", badL)
		}
		printLineNum(e.EndL, max)
		errfmt.Print(lines[e.EndL][:e.EndO])
	}
	fmt.Printf("%s\n", lines[e.EndL][e.EndO:])
}

// Print ...
func (e ATSTypeError) Print() {
	printSpaces(false)
	boldfmt.Printf("%s: ", e.Status)
	fmt.Printf("%s\n", e.ATSType)
}

// Print ...
func (e ATSCountError) Print() {
	printSpaces(true)
	errfmt.Printf("Found ")
	errors := "errors"
	if e.Count == 1 {
		errors = "error"
	}
	fmt.Printf("%d %s\n", e.Count, errors)
}
