#!/bin/bash

cd $(dirname $0)

cli="""
use: boash
short: a boa example
long: |
  This is an cli example from sh
  of how to use boa in a shell script
flags:
  - use: carbon
    short: It has symbol C and atomic number 6.
    alias: c
    type: string
  - use: nitrogen
    short: It has symbol N and atomic number 7.
    alias: n
    type: int
  - use: oxygen
    short: It has symbol O and atomic number 8.
    alias: o
"""

# parse the cli arguments
result=$(echo "$cli" | boa $@)
readarray -t args <<<"$result"
if [[ ${args[-1]} != "false" ]]; then
    # print the help message
    echo "$result"
    exit 1
else
    carbon=${args[0]}
    nitrogen=${args[1]}
    oxygen=${args[2]}

    # print the result
    echo "Running cli:" "boash" "-c ${carbon}" "-n ${nitrogen}" "-o ${oxygen}"
fi
