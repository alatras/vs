package main

import (
	"bitbucket.verifone.com/validation-service/cmd"
	"github.com/jessevdk/go-flags"
	"log"
	"os"
)

var version = "unknown"
var appName = "Validation Service"

// Opts with all cli commands and flags
type Opts struct {
	ServerCmd cmd.ServerCommand `command:"server"`
	cmd.CommonOpts
}

func main() {
	var opts Opts
	cmd.AppName = appName
	cmd.Version = version

	p := flags.NewParser(&opts, flags.Default)

	if len(os.Args) == 1 {
		os.Args = append(os.Args, "server")
	}

	p.CommandHandler = func(command flags.Commander, args []string) error {
		commonOpts := cmd.CommonOpts{Log: opts.Log}

		c := command.(cmd.CommonOptionsCommander)

		c.SetCommon(commonOpts)

		err := c.Execute(args)

		if err != nil {
			log.Printf("[ERROR] failed with %+v", err)
		}

		return err
	}

	if _, err := p.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}

}
