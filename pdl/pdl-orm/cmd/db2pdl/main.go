package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/kapablanka/pdl/pdl-orm/internal/db2pdl"
)

func main() {
	runFlag := flag.Bool("run", false, "execute the generator")
	configFlag := flag.String("config", "pdl.config.json", "path to configuration file")
	exitFlag := flag.Bool("exit", false, "exit process when complete")
	flag.Parse()
	if !*runFlag {
		fmt.Println("add --run to execute the generator")
		return
	}
	runner := db2pdl.Runner{
		ConfigPath:   *configFlag,
		ExitWhenDone: *exitFlag,
	}
	runErr := runner.Run(context.Background())
	if runErr != nil {
		fmt.Fprintf(os.Stderr, "%v\n", runErr)
		os.Exit(1)
	}
}
