// Copyright 2021 alfonso corretti. All rights reserved.
// Use of this source code is governed by The MIT License that can be found in
// the LICENSE file.

/*
CLI utility for golang/snappy compression & decompression of Snappy block
format. File and stdio capable.

Compatible with Google C++, Xerial Java and snappyjs JavaScript implementations.

Usage

Print a usage overview with `gsnappy help`

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

*/
package gsnappy_cli
