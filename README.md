# gsnappy-cli
[![Go Reference](https://pkg.go.dev/badge/github.com/acorretti/gsnappy-cli.svg)](https://pkg.go.dev/github.com/acorretti/gsnappy-cli)
[![Go Report Card](https://goreportcard.com/badge/github.com/acorretti/gsnappy-cli)](https://goreportcard.com/report/github.com/acorretti/gsnappy-cli)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![GitHub release](https://img.shields.io/github/release/acorretti/gsnappy-cli.svg)](https://github.com/acorretti/gsnappy-cli/releases/)

Go CLI utility for [Golang's Snappy](github.com/golang/snappy) compression/decompression **block** format. File and *stdio* capable.

Compatible with [Google](https://github.com/google/snappy) C++, [Xerial](https://github.com/xerial/snappy-java/) Java and [snappyjs](https://github.com/zhipeng-jia/snappyjs) JavaScript implementations.

### Install

```
go install github.com/acorretti/gsnappy-cli/cmd/gsnappy@latest
```

### Usage

```
NAME:
   gsnappy - Snappy compression utility

USAGE:
   gsnappy [global options] command [command options] [arguments...]

VERSION:
   1.x

COMMANDS:
   c, compress    compress data
   d, decompress  decompress data
   help, h        Shows a list of commands or help for one command

GLOBAL OPTIONS:
   -i FILE, --input FILE   read input data from FILE (default: stdin)
   -o FILE, --output FILE  write output data to FILE (default: stdout)
   --help, -h              show help (default: false)
   --version, -v           print the version (default: false)
```
Compress from `stdin`, write to `stdout`:
```
gsnappy
```
(press `Ctrl+D` to stop input and run)

Compress from `stdin`, write to a file:
```
cat /tmp/input | gsnappy c -o /tmp/output
```

Decompress from `stdin`, write to `stdout` redirected to a file:
```
gsnappy d > /tmp/out
```

Decompress a file, write to another file:
```
gsnappy d -i /tmp/input -o /tmp/output
```
