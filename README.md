# LIP, the Linter Issue Printer

The Linter Issue Printer is a project to shorten the time for creating BOSH packages and releases. <p>
BOSH is an Open Source system to define and deploy software. <p>
The linter performs a quick static analysis of all the code, configuration and scripts that make up a BOSH release. It checks for common mistakes, enforce a set of agreed-upon formatting rules, and thus give a reasonable outlook whether the current state of the code has a good chance to succeed being deployed.

## Setup
The project is configured to use `direnv`. Please install and setup `direnv` accordingly, if you want to use it
```
$ brew install direnv
```
then run `direnv allow`

```
$ direnv allow
```
You will also need the Go testing framework Ginkgo and Golang. Please install Ginkgo and Gomega if you haven't done it already.

```
$ go get github.com/onsi/ginkgo/ginkgo
$ go get github.com/onsi/gomega/...
```

## Usage
Now you can change in the directory `executable` and run the test, if the file is executable.
```
$ cd executable
$ ginkgo
```
Open `executable_test.go` to change the file which should be tested. In the function `IsExecutable` you can write the filename as a String input/parameter. With `ginkgo` you can go on with testing.
```
Expect(IsExecutable("filename")
$ ginkgo
```
But please consider that the tested file have to be in the `executable` directory.