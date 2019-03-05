# gitguy-finglonger

[![Release](https://img.shields.io/github/release/SUSE/gitguy-finglonger.svg)](https://github.com/SUSE/gitguy-finglonger/releases/latest)
[![Build Status](https://img.shields.io/travis/SUSE/gitguy-finglonger/master.svg)](https://travis-ci.org/SUSE/gitguy-finglonger)
[![codecov](https://codecov.io/gh/SUSE/gitguy-finglonger/branch/master/graph/badge.svg)](https://codecov.io/gh/SUSE/gitguy-finglonger)

![License: Apache 2.0](https://img.shields.io/github/license/SUSE/gitguy-finglonger.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/SUSE/gitguy-finglonger)](https://goreportcard.com/report/github.com/SUSE/gitguy-finglonger)

Manage your github project automatically, reacting to actions on your issues and labeling actions

## Install

Download the latest version from [releases](https://github.com/SUSE/gitguy-finglonger/releases/latest).

## Test

Clone repository into your $GOPATH. You can also use go get:

`go get github.com/SUSE/gitguy-finglonger`

### Dependencies

* `go >= 1.11`

Note:
We use golang modules but you still need to work inside your $GOPATH for developing `gitguy-finglonger`.
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
function](main.go) and generate a `gitguy-finglonger` binary.

Your binary will be stored under `bin` folder

## Generating Releases

Run first the tests. Then use `make release` to generate the release assets.

They will be created in the [release](release/) folder.

## Community

