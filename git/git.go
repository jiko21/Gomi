package git

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

var execCommand = exec.Command

type Git struct {
	blockBranches []string
}

func ConstructGit(path string) *Git {
	ignoreBranches := loadGomiIgnore(path)
	git := Git{ignoreBranches}
	return &git
}

func (g *Git) Delete() {
	items := getBranch()
	g.deleteMergedBranch(items)
}

func (g *Git) deleteMergedBranch(branches *[]string) {
	for _, branch := range *branches {
		if !isCurrentBranch(branch) && g.isBrachDeletable(branch) {
			deleteBranch(branch)
		}
	}
}

func (g *Git) isBrachDeletable(branch string) bool {
	formattedBranchName := strings.Replace(branch, " ", "", -1)
	for _, branchName := range g.blockBranches {
		if formattedBranchName == branchName {
			return false
		}
	}
	return true
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

func loadGomiIgnore(path string) []string {
	f, err := os.Open(path)
	defer f.Close()
	if err != nil {
		return []string{
			"master",
			"develop",
			"release",
		}
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	ignoreBranches := make([]string, 0)

	for fileScanner.Scan() {
		ignoreBranches = append(ignoreBranches, fileScanner.Text())
	}
	return ignoreBranches
}

func getBranch() *[]string {
	out, err := execCommand("git", "branch", "--merged").Output()
	if err != nil {
		fmt.Errorf(err.Error())
	}
	items := strings.Split(string(out), "\n")
	items = items[:len(items)-1]
	return &items
}
