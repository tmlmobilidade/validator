package services

import (
	"flag"
	"fmt"
	"os"
)

type CliOptions struct {
	InputPath  string // Path to the GTFS zip file
	OutputPath string // Path to the output file
	LogLevel   string // Log level (debug, info, error)
	RulesPath  string // Path to the rules file
	Version    bool   // Show version
}

type CLI struct {
	title   string
	Options CliOptions
}

func NewCLI(title string) *CLI {
	return &CLI{
		title: title,
		Options: CliOptions{
			OutputPath: "",
			LogLevel:   "info",
		},
	}
}

func (c *CLI) Parse() {
	flag.StringVar(&c.Options.InputPath, "input", "", "Path to the GTFS zip file")
	flag.StringVar(&c.Options.OutputPath, "out", "", "Path to the output file")
	flag.StringVar(&c.Options.OutputPath, "o", "", "Path to the output file")
	flag.StringVar(&c.Options.LogLevel, "log", "info", "Log level (debug, info, error)")
	flag.StringVar(&c.Options.RulesPath, "rules", "", "Path to the rules file")
	flag.BoolVar(&c.Options.Version, "v", false, "Show version")
	flag.BoolVar(&c.Options.Version, "version", false, "Show help")

	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s:\n", os.Args[0])
		flag.PrintDefaults()
	}

	flag.Parse()
}

func (c *CLI) Validate() error {
	if c.Options.InputPath == "" {
		return fmt.Errorf("input path is required")
	}

	return nil
}

func (c *CLI) Run() {
	c.Parse()

	if c.Options.Version {
		const version = "0.0.0"
		fmt.Printf("GTFS Validator v%s\n", version)
		os.Exit(0)
	}

	c.Validate()
}

var AppCLI = NewCLI("GTFS Validator")
