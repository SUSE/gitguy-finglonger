# github-project-mgr

[![Release](https://img.shields.io/github/release/chentex/github-project-mgr.svg)](https://github.com/chentex/github-project-mgr/releases/latest)
[![Build Status](https://img.shields.io/travis/chentex/github-project-mgr/master.svg)](https://travis-ci.org/chentex/github-project-mgr)
[![codecov](https://codecov.io/gh/chentex/github-project-mgr/branch/master/graph/badge.svg)](https://codecov.io/gh/chentex/github-project-mgr)

![License: Apache 2.0](https://img.shields.io/github/license/chentex/github-project-mgr.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/chentex/github-project-mgr)](https://goreportcard.com/report/github.com/chentex/github-project-mgr)

Manage your github project automatically, reacting to actions on your issues and labeling actions

## Install

Download the latest version from [releases](https://github.com/chentex/github-project-mgr/releases/latest).

## Test

Clone repository into your $GOPATH. You can also use go get:

`go get github.com/chentex/github-project-mgr`

### Dependencies

* `go >= 1.11`

Note:
We use golang modules but you still need to work inside your $GOPATH for developing `github-project-mgr`.
Working outside GOPATH is currently **not supported**

### Running tests

To run test on this package simply run:

`make test`

#### Testing with Docker

`make test.unit`

## Code Coverage

Run first the tests. Then use `make coverage` for visualizing coverage.

Feel free to read more about this on: https://blog.golang.org/cover.

## Building

Be sure you have all prerequisites.

A simple `make` should be enough. This should compile [the main
function](main.go) and generate a `github-project-mgr` binary.

Your binary will be stored under `bin` folder

## Generating Releases

Run first the tests. Then use `make release` to generate the release assets.

They will be created in the [release](release/) folder.

## Community

