package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/kapablanka/pdl/pdl/internal/pipeline"
)

func main() {
	configPath := flag.String("config", "pdl.config.json", "path to project configuration file")
	cleanStage := flag.Bool("clean-stage", false, "clean staged directories before copying")
	flag.Parse()

	action := pipeline.ActionBuild
	if len(flag.Args()) > 0 {
		switch strings.ToLower(flag.Args()[0]) {
		case "db2pdl":
			action = pipeline.ActionDb2Pdl
		case "rebuild":
			action = pipeline.ActionRebuild
		case "rebuild-all":
			action = pipeline.ActionRebuildAll
		case "clean":
			action = pipeline.ActionClean
		case "build":
			action = pipeline.ActionBuild
		default:
			fmt.Fprintf(os.Stderr, "unknown command: %s\n", flag.Args()[0])
			os.Exit(1)
		}
	}

	opts := pipeline.Options{
		ConfigPath: *configPath,
		CleanStage: *cleanStage,
		Action:     action,
	}
	if err := pipeline.Run(context.Background(), opts); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
