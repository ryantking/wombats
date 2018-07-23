package ats

import (
	"os"
	"os/exec"
	"strings"

	log "github.com/sirupsen/logrus"
)

var (
	// Patshome points to the PATSHOME environment variable
	Patshome string

	// Patscc points to the patscc binary
	Patscc string

	// Patsopt points to the patsopt binary
	Patsopt string
)

func init() {
	Patshome = os.Getenv("PATSHOME")
	if Patshome == "" {
		log.Fatalf("could not find PATSHOME environment variable")
	}

	// Check if patscc and patsopt are in the path
	for _, loc := range strings.Split(os.Getenv("PATH"), ":") {
		if _, err := os.Stat(loc + "/patscc"); err == nil {
			Patscc = "patscc"
		}

		if _, err := os.Stat(loc + "/patsopt"); err == nil {
			Patsopt = "patsopt"
		}

		if Patscc != "" && Patsopt != "" {
			break
		}
	}

	if Patscc == "" {
		log.Warnf("could not find patscc in PATH")

		Patscc = Patshome + "/bin/patscc"
		if _, err := os.Stat(Patscc); err != nil {
			log.Fatalf("could not find patscc binary")
		}
	}

	if Patsopt == "" {
		log.Warnf("could not find patsopt in PATH")

		Patsopt = Patshome + "/bin/patsopt"
		if _, err := os.Stat(Patsopt); err != nil {
			log.Fatalf("could not find patsopt binary")
		}
	}
}

// ExecPatscc executes a patscc command
func ExecPatscc(args ...string) error {
	cmd := exec.Command(Patscc, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Debugf("patscc error: %s", err)
		return err
	}
	return nil
}

// ExecPatsopt executes a patsopt command
func ExecPatsopt(args ...string) error {
	cmd := exec.Command(Patsopt, args...)
	cmd.Env = os.Environ()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		log.Debugf("patsopt error: %s", err)
		return err
	}
	return nil
}
