base
# Base

Based on "Tutorial: Get started with Go - The Go Programming Language" at https://go.dev/doc/tutorial/getting-started

Based on "Tutorial: Create a Go module" at https://go.dev/doc/tutorial/create-module

Based on "Go in Gitpod" at https://www.gitpod.io/docs/introduction/languages/go

Based on "Threagility" at https://github.com/Threagile/threagile

Based on "How To Deploy a Go Web Application with Docker" at https://semaphoreci.com/community/tutorials/how-to-deploy-a-go-web-application-with-docker

This repository contains the **base** of the 'products' that will be released as part of Creations Global:

- **Frame**: see https://github.com/creations-global/frame
- ...

Base is mostly written in the **Go** programming language and heavily inspired by [Threagile](https://github.com/Threagile/threagile).

## Usage

Run the code as follows from e.g. GitPod:

```
$ cd internal/base
$ go mod tidy
```

then,

```
$ cd internal/base
$ go run .
```

## Execution via Docker Container

The easiest way to execute Base on the commandline is via its Docker container:

```
docker run --rm -it creations/base --help


```

To reduce the text to be typed, one can equally make use of the shell script (```base.sh```), as follows:

```
$ ./base.sh --help
```