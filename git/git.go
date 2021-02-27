package git

import (
	"bufio"
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

func (g *Git) Delete() error {
	items, err := getMergedBranch()
	if err != nil {
		return err
	}
	return g.deleteMergedBranch(&items)
}

func (g *Git) deleteMergedBranch(branches *[]string) error {
	for _, branch := range *branches {
		if !isCurrentBranch(branch) && g.isBrachDeletable(branch) {
			err := deleteBranch(branch)
			if err != nil {
				return err
			}
		}
	}
	return nil
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

func deleteBranch(branch string) error {
	formattedBranchName := strings.Replace(branch, " ", "", -1)
	return execCommand("git", "branch", "-d", formattedBranchName).Run()
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
			"main",
			"develop",
			"release",
		}
	}
	fileScanner := bufio.NewScanner(f)
	fileScanner.Split(bufio.ScanLines)
	ignoreBranches := make([]string, 0)

	for fileScanner.Scan() {
		if fileScanner.Text()[0] != '#' {
			ignoreBranches = append(ignoreBranches, fileScanner.Text())
		}
	}
	return ignoreBranches
}

func getMergedBranch() ([]string, error) {
	out, err := execCommand("git", "branch", "--merged").Output()
	if err != nil {
		return nil, err
	}
	items := strings.Split(string(out), "\n")
	items = items[:len(items)-1]
	return items, nil
}

func GetBranch() ([]string, error) {
	out, err := execCommand("git", "branch").Output()
	if err != nil {
		return nil, err
	}
	items := strings.Split(string(out), "\n")
	items = items[:len(items)-1]
	branches := []string{}
	for _, item := range items {
		branches = append(branches, item[2:])
	}
	return branches, nil
}
