package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/alecthomas/kong"
	"github.com/winebarrel/codeowners"
)

var version string

func init() {
	log.SetFlags(0)
}

func parseArgs() *codeowners.Options {
	var CLI struct {
		codeowners.Options
		Version kong.VersionFlag
	}

	parser := kong.Must(&CLI, kong.Vars{"version": version})
	parser.Model.HelpFlag.Help = "Show help."
	_, err := parser.Parse(os.Args[1:])
	parser.FatalIfErrorf(err)

	return &CLI.Options
}

func main() {
	options := parseArgs()
	ctx := context.Background()
	cos, err := codeowners.List(ctx, options)

	if err != nil {
		log.Fatal(err)
	}

	raw, _ := json.MarshalIndent(cos, "", "  ")
	fmt.Println(string(raw))
}
