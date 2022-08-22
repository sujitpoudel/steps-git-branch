package gitbranch

import (
	"testing"

	"github.com/bitrise-io/go-utils/command/git"
	"github.com/stretchr/testify/assert"
)

const rawCmdError = "dummy_cmd_error"

// SubmoduleUpdate
var branchOffBaseTestCases = [...]struct {
	name     string
	cfg      Config
	wantCmds []string
}{
	{
		name: "Given submodule update depth is 1 when the submodules are updated then expect the --depth=1 flag on the command",
		cfg: Config{
			Branch: "release-1.1.1",
			Base:   "master",
		},
		wantCmds: []string{
			`git "checkout" "-b" "release-1.1.1" "master"`,
		},
	},
}

func Test_BranchFromBase(t *testing.T) {
	for _, tt := range branchOffBaseTestCases {
		t.Run(tt.name, func(t *testing.T) {
			// Given
			mockRunner := givenMockRunnerSucceeds()
			runner = mockRunner

			// When
			actualErr := createBranch(git.Git{}, tt.cfg)

			// Then
			assert.NoError(t, actualErr)
			assert.Equal(t, tt.wantCmds, mockRunner.Cmds())
		})
	}
}

// SetupSparseCechkout
//var sparseCheckoutTestCases = [...]struct {
//	name              string
//	sparseDirectories []string
//	wantCmds          []string
//}{
//	{
//		name:              "Sparse-checkout single directory",
//		sparseDirectories: []string{"client/android"},
//		wantCmds: []string{
//			`git "sparse-checkout" "init" "--cone"`,
//			`git "sparse-checkout" "set" "client/android"`,
//			`git "config" "extensions.partialClone" "origin" "--local"`,
//		},
//	},
//	{
//		name:              "Sparse-checkout multiple directory",
//		sparseDirectories: []string{"client/android", "client/ios"},
//		wantCmds: []string{
//			`git "sparse-checkout" "init" "--cone"`,
//			`git "sparse-checkout" "set" "client/android" "client/ios"`,
//			`git "config" "extensions.partialClone" "origin" "--local"`,
//		},
//	},
//}

//func Test_SetupSparseCheckout(t *testing.T) {
//	for _, tt := range sparseCheckoutTestCases {
//		t.Run(tt.name, func(t *testing.T) {
//			// Given
//			mockRunner := givenMockRunnerSucceeds()
//			runner = mockRunner
//
//			// When
//			actualErr := setupSparseCheckout(git.Git{}, tt.sparseDirectories)
//
//			// Then
//			assert.NoError(t, actualErr)
//			assert.Equal(t, tt.wantCmds, mockRunner.Cmds())
//		})
//	}
//}

// Mocks
func givenMockRunner() *MockRunner {
	mockRunner := new(MockRunner)
	mockRunner.GivenRunForOutputSucceeds()
	return mockRunner
}

func givenMockRunnerSucceeds() *MockRunner {
	return givenMockRunnerSucceedsAfter(0)
}

func givenMockRunnerSucceedsAfter(times int) *MockRunner {
	return givenMockRunner().
		GivenRunWithRetrySucceedsAfter(times).
		GivenRunSucceeds()
}

type MockPatchSource struct {
	diffFilePath string
	err          error
}

func (m MockPatchSource) getDiffPath(_, _ string) (string, error) {
	return m.diffFilePath, m.err
}
