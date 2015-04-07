#!/bin/bash
set -e

# Inspired on https://github.com/docker-library/postgres script

case "$1" in
    "check")
        shift
        OPTIONS="$@"
        exec check ${OPTIONS}
    ;;
esac
exec "$@"
