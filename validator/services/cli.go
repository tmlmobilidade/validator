package services

import (
	"flag"
	"fmt"
)

type CliOptions struct {
	InputPath  string
	OutputPath string
	LogLevel   string
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
	flag.StringVar(&c.Options.OutputPath, "output", "", "Path to the output file")
	flag.StringVar(&c.Options.LogLevel, "log", "info", "Log level (debug, info, error)")

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
	c.Validate()
}

var AppCLI = NewCLI("GTFS Validator")
