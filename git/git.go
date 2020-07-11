package git

import (
	"fmt"
	"os/exec"
	"regexp"
	"strings"
)

var execCommand = exec.Command

func getBranch() *[]string {
	out, err := execCommand("git", "branch", "--merged").Output()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	items := strings.Split(string(out), "\n")
	items = items[:len(items)-1]
	return &items
}

func deleteMergedBranch(branches *[]string) {
	for _, branch := range *branches {
		if !isCurrentBranch(branch) {
			deleteBranch(branch)
		}
	}
}

func deleteBranch(branch string) {
	formattedBranchName := strings.Replace(branch, " ", "", -1)
	err := execCommand("git", "branch", "-d", formattedBranchName).Run()
	if err != nil {
		fmt.Errorf(err.Error())
	}
}

func isCurrentBranch(branch string) bool {
	currentBranchPattern := regexp.MustCompile(`\*\ .+`)
	return currentBranchPattern.MatchString(branch)
}

func Delete() {
	items := getBranch()
	deleteMergedBranch(items)
}
