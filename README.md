# GOMI
Branch delete tool made with Golang
## About
GOMI is a tool for deleting branches that have been merged into the local git repository.


When working with git or GitHub flow,
you often encounter many branches that have already been merged into the master or develop branches.

You can delete branches that have been merged into your current branch simply by using the following command in your project:

```bash
gomi
```

## How to install
GOMI can be installed via Homebrew:

```bash
brew tap jiko21/gomi
```

then,Then, install it with:

```bash
brew install gomi
```

## Features

### branch delete block
You can specify branches that you do not want to delete.

This can be done using a .gomiignore file in your project's root directory.

If you don't use a .gomiignore file,
the default rule applies: master, main, develop, and release branches cannot be deleted by GOMI.

An example of a .gomiignore file is shown below:

```
master
do-not-delete-branch
master-xxx
# this is commented out
# lines like this are ignored in the .gomiignore file
```

### Auto-Execution After Merging a Branch
You can configure GOMI to run automatically after merging a branch.

This feature can be enabled during the initialization of GOMI.

### Initialization
You can generate a .gomiignore file with the command shown below.

Afterward, Git hooks will be created for actions following a merge commit.
```
$ gomi init
```
