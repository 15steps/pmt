package cmd

import (
	"errors"
	"fmt"
	"github.com/urfave/cli"
	"log"
	"os"
)

// flag names
const  (
	Pattern = "pattern"
	MaxEdit = "edit"
	Algorithm = "algorithm"
	Count = "count"
)

// algorithms
const (
	ShiftOr = "shiftor"
)

func BuildCli() {
	app := cli.NewApp()

	// general info
	app.Name = "pmt"
	app.Usage = "fast text search"
	app.Author = "Wellington Felix - wfmf@cin.ufpe.br"
	app.Version = "0.0.1"

	app.UsageText = "pmt [options] pattern textfile [textfile...]"

	// flags
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name: getFlagName(Pattern),
			Usage: "file containing a pattern per line",
		},
		cli.Int64Flag{
			Name: getFlagName(MaxEdit),
			Usage: "max error for approximate matching",
		},
		cli.StringFlag{
			Name: getFlagName(Algorithm),
			Usage: "algorithm to use during search",
			Value: "shiftor",
		},
		cli.BoolFlag{
			Name: getFlagName(Count),
			Usage: "if present show the total number of occurrences",
		},
	}

	app.Action = func(c *cli.Context) error {
		pattern := c.Args().First()
		files := c.Args().Tail()
		if algorithm := c.String(Algorithm); algorithm == "" {
			algorithm = ShiftOr
		}
        if len(pattern) == 0 {
        	return errors.New("a pattern is required")
		}
		if len(files) == 0 {
			return errors.New("a least one file is required")
		}
		fmt.Printf("searching for %s in %v, %s\n", pattern, files, c.String(Algorithm))
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func getFlagName(flag string) string {
	return fmt.Sprintf("%s, %c", flag, flag[0])
}