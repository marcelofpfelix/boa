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
commands:
  - use: list
    short: example command 1
    aliases:
      - l
    long: |
      This is an example command
      for multiple commands
    commands:
      - use: all
        short: example subcommand 1
        long: |
          This is an example command
          for multiple commands
  - use: add
    short: example command 2
    long: |
      This is an example command
      for multiple commands
    flags:
      - use: nitrogen
        short: It has symbol N and atomic number 7.
        alias: n
        type: int
      - use: oxygen
        short: It has symbol O and atomic number 8.
        alias: o
"""

# parse the cli arguments
result=$(echo "$cli" | ../bin/boa "$@")
readarray -t args <<<"$result"
if [[ ${args[-1]} != "false" ]]; then
    # print the help message
    echo "$result"
    exit 1
else
    cmd=${args[0]}
    carbon=${args[1]}
    nitrogen=${args[2]}
    oxygen=${args[3]}
    echo "$result"
    # print the result
    echo "Running:" "${cmd} -c ${carbon}" "-n ${nitrogen}" "-o ${oxygen}"
fi
