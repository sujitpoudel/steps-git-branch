package main

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-steputils/stepconf"
	"github.com/bitrise-io/go-utils/log"
	"github.com/bitrise-steplib/steps-git-branch/gitbranch"
)

func failf(format string, v ...interface{}) {
	log.Errorf(format, v...)
	os.Exit(1)
}

func main() {
	//var cfg gitbranch.Config
	cfg := gitbranch.Config{
		Base:   "master",
		Branch: "release-4",
	}
	if err := stepconf.Parse(&cfg); err != nil {
		failf("Error: %s\n", err)
	}
	stepconf.Print(cfg)

	if err := gitbranch.Execute(cfg); err != nil {
		failf("Error: %v", err)
	}

	fmt.Println()
	log.Donef("Success")
}
