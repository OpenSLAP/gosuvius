package main

import (
	"github.com/urfave/cli"
	"os"
)

func FileExists(name string) bool {
	if _, err := os.Stat(name); err == nil {
		return true
	} else {
		return false
	}
}

func HandleArgs(c *cli.Context) error {
	target := c.Args().Get(0)
	if len(target) == 0 {
		return cli.NewExitError("Must specify a target", 1)
	}
	if !FileExists(target) {
		return cli.NewExitError("Target specified does not exist", 2)
	}
	return nil
}

func HandleLint(c *cli.Context) error {
	target := c.Args().Get(0)
	f, err := os.Open(target)
	if err != nil {
		return err
	}
	defer f.Close()

	p := NewParser(f)
	for {
		l, _ := p.Parse()
		if l == nil {
			break
		}
		/*
			fmt.Printf("%s ", l.RawCode)
			for _, param := range l.Params {
				fmt.Printf("%s ", param.RawParam)
			}
			fmt.Printf("\n")
		*/
	}
	return nil
}

func main() {
	// Set up the CLI
	app := cli.NewApp()
	app.Name = "gosuvius"
	app.Usage = "G-code linter/sanitizer for vesuvius"
	app.Authors = []cli.Author{
		{Name: "Charles Neill"},
		{Name: "Zach Marcantel"},
	}
	app.Commands = []cli.Command{
		{
			Name:      "lint",
			Usage:     "lint a gcode file for errors",
			Action:    HandleLint,
			Before:    HandleArgs,
			ArgsUsage: "TARGET",
		},
		{
			Name:  "sanitize",
			Usage: "sanitize a gcode file and output vesuvius-compatible gcode",
			Action: func(c *cli.Context) error {
				return cli.NewExitError("NOT IMPLEMENTED", 3)
			},
		},
		{
			Name:  "stats",
			Usage: "get stats about a gcode file",
			Action: func(c *cli.Context) error {
				return cli.NewExitError("NOT IMPLEMENTED", 3)
			},
		},
	}
	app.Run(os.Args)
}
