package gitbranch

import (
	"fmt"
	"os"

	"github.com/bitrise-io/go-utils/command/git"
)

// Config is the git clone step configuration
type Config struct {
	Base   string `env:"base"`
	Branch string `env:"branch"`
}

const (
	originRemoteName        = "origin"
	forkRemoteName          = "fork"
	updateSubmodelFailedTag = "update_submodule_failed"
	createBranchTag         = "create_branch"
)

func createBranch(gitCmd git.Git, cfg Config) error {
	var opts []string
	opts = append(opts, "")
	if err := runner.Run(gitCmd.NewBranchFrom(cfg.Branch, cfg.Base)); err != nil {
		return newStepError(
			createBranchTag,
			fmt.Errorf("unable to create branch/tag %v", err),
			"Creating new branch from base branch failed.",
		)
	}
	return nil
}

// Execute is the entry point of the git clone process
func Execute(cfg Config) error {
	fmt.Println(os.Environ())
	return nil
	//return createBranch(git.Git{}, cfg)
}
