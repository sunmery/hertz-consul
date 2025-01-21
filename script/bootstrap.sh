#! /usr/bin/env bash
CURDIR=$(cd $(dirname $0); pwd)
echo "$CURDIR/bin/hertz"
exec "$CURDIR/bin/hertz"
