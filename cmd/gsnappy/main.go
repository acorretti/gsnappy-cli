package main

import (
	"bufio"
	"io"
	"log"
	"os"

	"github.com/golang/snappy"
	"github.com/urfave/cli/v2"
)

const (
	inputFlag  string = "i"
	outputFlag string = "o"
)

var ioFlagNames = []string{
	inputFlag,
	outputFlag,
}

// global flags values override subcommand ones to allow specifying the command first
// yet still print useful global usage help
func copyLineageFlags(c *cli.Context) {
	for _, ctxt := range c.Lineage() {
		for _, flag := range ioFlagNames {
			if ctxt.IsSet(flag) {
				c.Set(flag, ctxt.String(flag))
			}
		}
	}
}

func main() {
	var inputFile string
	var outputFile string

	ioFlags := []cli.Flag{
		&cli.StringFlag{
			Name:        inputFlag,
			Aliases:     []string{"input"},
			Usage:       "read input data from `FILE`",
			DefaultText: "stdin",
		},
		&cli.StringFlag{
			Name:        outputFlag,
			Aliases:     []string{"output"},
			Usage:       "write output data to `FILE`",
			DefaultText: "stdout",
		},
	}

	app := &cli.App{
		Name:    "gsnappy",
		Version: "1.0",
		Usage:   "Snappy compression utility",
		Flags:   ioFlags,
		Action: func(c *cli.Context) error {
			// run 'compress' command by default
			c.App.Command("c").Run(c)
			return nil
		},
		Commands: []*cli.Command{
			{
				Name:    "c",
				Aliases: []string{"compress"},
				Usage:   "compress data",
				Flags:   ioFlags,
				Action: func(c *cli.Context) error {
					copyLineageFlags(c)
					inputFile = c.String(inputFlag)
					outputFile = c.String(outputFlag)
					compress(inputFile, outputFile)
					return nil
				},
			},
			{
				Name:    "d",
				Aliases: []string{"decompress"},
				Usage:   "decompress data",
				Flags:   ioFlags,
				Action: func(c *cli.Context) error {
					copyLineageFlags(c)
					inputFile = c.String(inputFlag)
					outputFile = c.String(outputFlag)
					decompress(inputFile, outputFile)
					return nil
				},
			},
		},
	}
	checkForPanic(
		app.Run(os.Args),
	)
}

func compress(inputFile string, outputFile string) {
	writeOutputBytes(outputFile, snappy.Encode(nil, readInputBytes(inputFile)))
}

func decompress(inputFile string, outputFile string) {
	decoded, err := snappy.Decode(nil, readInputBytes(inputFile))
	checkForPanic(err)
	writeOutputBytes(outputFile, decoded)
}

func readInputBytes(inputFile string) []byte {
	reader := os.Stdin
	if inputFile != "" {
		file, err := os.Open(inputFile)
		checkForPanic(err)
		defer reader.Close()
		reader = file
	}

	in := bufio.NewReader(reader)
	out, err := io.ReadAll(in)
	checkForPanic(err)
	return out
}

func writeOutputBytes(outputFile string, data []byte) {
	if outputFile == "" {
		os.Stdout.Write(data)
	} else {
		file, err := os.Create(outputFile)
		checkForPanic(err)
		defer file.Close()
		_, err = file.Write(data)
		checkForPanic(err)
	}
}

func checkForPanic(e error) {
	if e != nil {
		log.Fatal(e)
		os.Exit(1)
	}
}