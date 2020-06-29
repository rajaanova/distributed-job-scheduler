#!/bin/bash

# Variable Definitions
RED="\033[0;31m"             # ANSI color code Red - Used to color encode error text
NOCOLOR="\033[0m"            # ANSI color code No color - Used to unset previously set color encoding
workdir=.cover               # directory to store all coverage output into (include intermediate files)
profile="$workdir/cover.out" # file to store aggregated coverate output
COVERAGE_PATTERN="./..."     # Namespace pattern used by "go list" to include relevant packages

go test ./... -coverpkg ${COVERAGE_PATTERN} -coverprofile ${profile}

PERCENTTEST=`go tool cover -func ${profile}|grep total |awk '{ print substr($3,1,length($3)-1) }'`

if [ 1 -eq `echo "${PERCENTTEST} < 80" | bc` ]
then
 echo -e "${RED}very less test coverage $PERCENTTEST $NOCOLOR";
 exit  1
fi

