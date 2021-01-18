# GOMI
Branch delete tool made by Golang
## About
GOMI is a tool for deleting merged branches in local git repository.

When you write code with git or github flow,
there are many branches that are already merged into master / develop branch.

You can delete branches merged into your current branch only with

```bash
gomi
```

at your project.
## How to install
You can install via homebrew

```bash
brew tap jiko21/gomi
```

then,

```bash
brew install gomi
```

## Features

### branch delete block
You can specify the branches that you don't want to delete.

You can specify them with `.gomiignore`, on your project root.

If you don't use `.gomiignore`, then default rule will be adopted; `master`, `main`, `develop`, `release` cannot be deleted by gomi.

The example of `.gomiignore` is shown below.

```
master
do-not-delete-branch
master-xxx
```
