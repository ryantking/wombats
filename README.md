# Wombats
[![Build Status](https://travis-ci.org/RyanTKing/wombats.svg?branch=master)](https://travis-ci.org/RyanTKing/wombats) [![codecov](https://codecov.io/gh/RyanTKing/wombats/branch/master/graph/badge.svg)](https://codecov.io/gh/RyanTKing/wombats) [![Go Report Card](https://goreportcard.com/badge/github.com/RyanTKing/wombats)](https://goreportcard.com/report/github.com/RyanTKing/wombats)

<img src="logo.png" alt="Wombats logo" title="The simple Wombat" align="right" />

> Oh how the family affections combat<br/>
> Within this heart, and each hour flings a bomb at<br/>
> My burning soul! Neither from owl nor from bat<br/>
> Can peace be gained until I clasp my wombat.<br/>
> - from [The Wombat](https://www.poemhunter.com/poem/the-wombat-2/) by Dante Gabriel Rosetti

Wombats is a package manager and build system for ATS in order to automate
building and deploying packages.

Inspired by the likes of [Go](https://github.com/golang/go), [Cargo](https://github.com/rust-lang/cargo), and [Leiningen](https://github.com/technomancy/leiningen).

**Note:** There is currently no server setup to manage the language packages,
so this is not ready for use out side of experimental use.

# Installation

Currently, Wombats is not on any package managers, but can easily be installed
using go:

    $ go get github.com/RyanTKing/wombats/cmd/wom

ATS is also required, and instructions for that can be found on [my blog](http://ryanking.com/blog/joy-of-ats-1-installing-ats/).


# Usage

A project can be initialized and ran as follows:

    $ wom new [DIR] -n [NAME]
    $ wom run

The following commands are available:

    $ wom new [DIR] # Create a new project
    $ wom run # Run project and build if necessary (works for basic project)
    $ wom build # Build the current project (unimplemented)
    $ wom install # Install the current project (unimplemented)
    $ wom version # Show ATS version information
    $ wom fetch # Fetch all specified dependencies (unimplemented)
    $ wom fetch [PKG] # Fetches a package (unimplemented)

# License

Source Copyright &copy; 2018 Ryan King, distributed under the MIT License.
