# rat
[![ci](https://github.com/Yakiyo/rat/actions/workflows/ci.yml/badge.svg?branch=main)](https://github.com/Yakiyo/rat/actions/workflows/ci.yml) ![GitHub go.mod Go version (subdirectory of monorepo)](https://img.shields.io/github/go-mod/go-version/Yakiyo/rat) ![GitHub tag (with filter)](https://img.shields.io/github/v/tag/Yakiyo/rat?label=version)

Writing a cat (and [bat](https://github.com/sharkdp/bat)) clone in Go for learning purposes

## Installation
Installation of it requires the [Go](https://go.dev) toolchain.

```bash
$ go install github.com/Yakiyo/rat
```
or the repository can be locally cloned and installed too.

```bash
$ git clone https://github.com/Yakiyo/rat

$ cd rat

$ go build .
```

## Usage
It works similar to cat
```shell
# pretty print a js file
$ rat index.js

# print multiple files of different language
$ rat index.js main.go

# implicitly mention language to use
$ rat --language typescript index.js

# choose a style
$ rat --style githubdark main.cpp

# read from stdin
$ echo "console.log('hello')" | rat

# first print README, then use stdin and then index.js
$ rat README.md - index.js

# disable pretty printing (print plain text)
$ rat --style none app.py

```

For syntax highlighting, the [chroma](https://github.com/alecthomas/chroma) library is used. A list of supported languages can be found [here](https://github.com/alecthomas/chroma#supported-languages). All available chroma styles can be viewed at the [styles gallery](https://xyproto.github.io/splash/docs/).