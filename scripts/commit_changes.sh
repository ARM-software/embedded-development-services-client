#!/bin/bash
# Copyright (C) 2021 Arm Ltd. All rights reserved.
#
# The aim of this script is to commit changes from a specific branch to github

# Getting the branch name from arguments
branch_name=$1

git config user.email "monty-bot@arm.com"
git config user.name "monty-bot"

# Modifying the git remote URL and protocol
git remote set-url origin https://github.com/$(git remote get-url origin | sed 's/https:\/\/github.com\///' | sed 's/git@github.com://')

# Getting the latest changes from the remote repository in case some happened.
git branch --set-upstream-to origin/${branch_name}
git pull
git status

echo "Committing changes"
for object in "$2" ; do
    git add $object
done

git status

# To ensure the next commit does not trigger a build.
git commit -m ":sparkles: Automatic changes -> $3"

git push