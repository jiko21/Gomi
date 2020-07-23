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

func Test_loadGomiIgnore(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			"should load gomiigore",
			args{
				".gomiignore.test",
			},
			[]string{
				"master",
				"hoge",
			},
		},
		{
			"should return default file when gomiignore not found",
			args{
				".gomiignore",
			},
			[]string{
				"master",
				"develop",
				"release",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := loadGomiIgnore(tt.args.path); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("loadGomiIgnore() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGit_isBrachDeletable(t *testing.T) {
	type fields struct {
		blockBranches []string
	}
	type args struct {
		branch string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   bool
	}{
		{
			"should master is not deletable",
			fields{
				[]string{
					"master",
					"develop",
				},
			},
			args{
				"master",
			},
			false,
		},
		{
			"should hogehoge not deletable",
			fields{
				[]string{
					"master",
					"develop",
				},
			},
			args{
				"hogehoge",
			},
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Git{
				blockBranches: tt.fields.blockBranches,
			}
			if got := g.isBrachDeletable(tt.args.branch); got != tt.want {
				t.Errorf("Git.isBrachDeletable() = %v, want %v", got, tt.want)
			}
		})
	}
}
