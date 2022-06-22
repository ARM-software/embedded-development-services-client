#!/usr/bin/env bash

# Dependabot will only update the Dockerfiles however we need to share the version of Go in multiple places
# So this script will fetch the version of Go from one of the Dockerfiles and then use this to update the:
# CircleCI config
# The README
# The go.mod files
# The changes will be committed using the commit_changes.sh script

function update() {
  DOCKERFILE="Dockerfile.acceptancetest"

  GO_FULL_VERSION=`grep -o 'golang:[0-9.]*' ${DOCKERFILE} | cut -d ":" -f2`
  GO_SHORT_VERSION=`echo ${GO_FULL_VERSION} | sed 's/\.[^.]*$//'`

  sed -i -e "s/golang:[0-9.]*$/golang:${GO_FULL_VERSION}/g" .circleci/config.yml
  sed -i -e "s/go-v[0-9.]*/go-v${GO_FULL_VERSION}/g" README.md
  sed -i -e "s/^go [0-9.]*$/go ${GO_SHORT_VERSION}/g" acceptancetests/go.mod

  git status
}

update
