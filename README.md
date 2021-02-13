## Requirements

- git 
- sed

## Install

1. Clone [`scripts/pre-commit`](https://github.com/aomerk/go-modules-githook/blob/main/scripts/pre-commit)  and [`scripts/post-commit`](https://github.com/aomerk/go-modules-githook/blob/main/scripts/post-commit) to your computer

2. Make `pre-commit` and `post-commit` executable

3. Copy `pre-commit` and `post-commit` scripts to:
    - your repositories' hooks (`.git/hooks`) 
    - or your global hooks. See this [page](https://coderwall.com/p/jp7d5q/create-a-global-git-commit-hook) for a tutorial about adding global git hooks
