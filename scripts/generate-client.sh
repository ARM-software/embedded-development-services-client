#!/bin/bash
# Copyright (C) 2020 Arm Ltd. All rights reserved.
#
# The aim of this script is to generate an client from an OpenAPi definition
# Requires:
# - https://github.com/OpenAPITools/openapi-generator

set -e

usage() {
    cat <<"EOF"
USAGE:
   generate-client.sh [arguments]
ARGUMENTS:
   -i input                               REQUIRED: The input OpenAPI definition
   -o output                              REQUIRED: Where the client will be generated
   -d                                     Include to disable the automatic formatting of the client
   --help, -h                             Show help
EOF
}

run() {
    while getopts i:o:d option; do
        case "${option}"
        in
            i) INPUT=${OPTARG};;
            o) OUTPUT=${OPTARG};;
            d) DISABLE_FMT='true';;
            *) usage
              exit 1 ;;
        esac
    done

    [[ -z "$INPUT" ]] && { echo "Must provide an input file" ; exit 1; }
    [[ -z "$OUTPUT" ]] && { echo "Must provide a output directory" ; exit 1; }

    openapi-generator generate -g go -o "$OUTPUT" --additional-properties packageName=client,isGoSubmodule=true --git-user-id="ARM-software" --git-repo-id="embedded-development-services-client" -i "$INPUT"
    # FIXME the following is a stop gap to overcome a bug in the generator https://github.com/OpenAPITools/openapi-generator/issues/12810
    sed -i -e 's/**os.File/*os.File/g' ./client/api_build_jobs.go
    sed -i -e 's/**os.File/*os.File/g' ./client/api_intellisense_jobs.go
    # FIXME remove the following line when the import in tests has been fixed in the generator
    rm -rf "$OUTPUT"/test
    rm -rf "$OUTPUT"/api
    # FIXME the following is a stop gap to overcome the fact urls are already escaped in the services so we don't the client to escape them again as go believes that the resulting string is an f-string and complains because the argument is missing. remove if an option to not escape URLs is added
    find ./client -type f -name "api_*.go" -print0 | xargs -0 sed -i -e 's@url.PathEscape(\([^)]*\))@\1@g'

    if [[ -z "$DISABLE_FMT"  ]]; then
      ./fmt-go.sh -i "$OUTPUT"
    fi
}

main() {
    if [[ -z "${1}" || "${1}" == '-h' || "${1}" == '--help' ]]; then
        usage
    else
        run "$@"
    fi
}

main "$@"
