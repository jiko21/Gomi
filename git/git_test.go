package git

import (
	"fmt"
	"os"
	"os/exec"
	"reflect"
	"testing"
)

// test helper
// https://gist.github.com/hichihara/46f6b278f3b6a1a9901666f27bcaa61b
var testCase string

func fakeExecCommand(command string, args ...string) *exec.Cmd {
	cs := []string{"-test.run=TestHelperProcess", "--", command}
	cs = append(cs, args...)
	cmd := exec.Command(os.Args[0], cs...)
	cmd.Env = []string{"GO_WANT_HELPER_PROCESS=1", "TEST_CASE=" + testCase}
	return cmd
}

func TestHelperProcess(t *testing.T) {
	if os.Getenv("GO_WANT_HELPER_PROCESS") != "1" {
		return
	}
	switch os.Getenv("TEST_CASE") {
	case "branch":
		fmt.Fprintf(os.Stdout, branches)
	}
	os.Exit(0)
}

// test helper ended

var branches = `  a-branch
* master
  x-branch
`

func Test_getBranch(t *testing.T) {
	execCommand = fakeExecCommand
	testCase = "branch"
	defer func() { execCommand = exec.Command }()

	expected := []string{
		"  a-branch",
		"* master",
		"  x-branch",
	}
	tests := []struct {
		name string
		want *[]string
	}{
		{
			"should getBranch correctly returns branch array",
			&expected,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getBranch(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_isCurrentBranch(t *testing.T) {
	type args struct {
		branch string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			"should branch name `* master` be current branch",
			args{
				"* master",
			},
			true,
		},
		{
			"should branch name `  not-current-branch` not be current branch",
			args{
				"  not-current-branch",
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isCurrentBranch(tt.args.branch); got != tt.want {
				t.Errorf("isCurrentBranch() = %v, want %v", got, tt.want)
			}
		})
	}
}
