# Clara


[![Go Reference](https://pkg.go.dev/badge/github.com/gobuffalo/clara.svg)](https://pkg.go.dev/github.com/gobuffalo/clara)
[![Actions Status](https://github.com/gobuffalo/clara/workflows/Tests/badge.svg)](https://github.com/gobuffalo/clara/actions)


Clara is a tool for checking your environment for [Go](https://golang.org) and [Buffalo](https://gobuffalo.io) development. It will check things like versions, pathing, etc... and will let you know what needs to be fixed, setup, or updated, in order to have a system that hums along while you develop.

Note that Clara officially supports the last two versions of Go, which at the moment are:
* 1.17
* 1.18

Even though it may (or may not) work on the older versions, we encourage you to upgrade your Go development environment.

## Installation

### Binary Releases

Pre-built binaries for most platforms can be found at [https://github.com/gobuffalo/clara/releases](https://github.com/gobuffalo/clara/releases).

### Using Go Install

For go version 1.17 or higher,

```console
$ go install github.com/gobuffalo/clara/v2@latest
```

### From Source

```console
$ make install
```

## Usage

```console
$ clara
```

### Example Output

```console
$ clara

-> Go: Checking installation
✓ The `go` executable was found on your system at: /usr/local/go/bin/go

-> Go: Checking minimum version requirements
✓ Your version of Go, 1.12.1, meets the minimum requirements.

-> Go: Checking GOPATH
✓ You appear to be operating inside of your GOPATH.

-> Go: Checking Package Management
⚠ You do not appear to be using a package management system.

It is strongly suggested that you use one of the following package management systems:

* Go Modules (Recommended) - https://gobuffalo.io/en/docs/gomods
* Dep - https://github.com/golang/dep

For help setting up your Go environment please follow the instructions for you platform at:

https://www.gopherguides.com/courses/preparing-your-environment-for-go-development

-> Go: Checking PATH
✓ Your PATH contains /Users/markbates/Dropbox/go/bin.

-> Buffalo: Checking installation
✘ The `buffalo` executable could not be found on your system.
For help setting up your Buffalo environment please follow the instructions for you platform at:

https://gobuffalo.io/en/docs/installation

For help setting up your Go environment please follow the instructions for you platform at:

https://www.gopherguides.com/courses/preparing-your-environment-for-go-development

-> Buffalo: Checking minimum version requirements
✘ : exec: "buffalo": executable file not found in $PATH

-> Node: Checking installation
✓ The `node` executable was found on your system at: /usr/local/bin/node

-> Node: Checking minimum version requirements
✓ Your version of Node, v10.11.0, meets the minimum requirements.

-> NPM: Checking installation
✓ The `npm` executable was found on your system at: /usr/local/bin/npm

-> NPM: Checking minimum version requirements
✓ Your version of NPM, 6.4.1, meets the minimum requirements.

-> Yarn: Checking installation
✓ The `yarnpkg` executable was found on your system at: /usr/local/bin/yarnpkg

-> Yarn: Checking minimum version requirements
✘ You version of Yarn, 1.10.1, does not meet the minimum requirements.

Minimum versions of Yarn are:

* >=1.12

For help setting up your Yarn environment please follow the instructions for you platform at:

https://yarnpkg.com/en/docs/install

```

## Why the Name Clara?

This project was named after [Clara Barton](https://en.wikipedia.org/wiki/Clara_Barton). Clarissa "Clara" Harlowe Barton (December 25, 1821 – April 12, 1912) was a pioneering nurse who founded the American Red Cross. She was a hospital nurse in the American Civil War, a teacher, and patent clerk. Nursing education was not very formalized at that time and she did not attend nursing school, so she provided self-taught nursing care.[1] Barton is noteworthy for doing humanitarian work at a time when relatively few women worked outside the home.[2] She was inducted into the National Women's Hall of Fame in 1973.[3]
