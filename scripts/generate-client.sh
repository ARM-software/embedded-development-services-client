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

    openapi-generator generate -g go -o "$OUTPUT" --additional-properties packageName=client,isGoSubmodule=true --git-user-id="Arm-Debug" --git-repo-id="solar-services-client" -i "$INPUT"

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
