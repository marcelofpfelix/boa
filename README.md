# boa

An cli argument parser, using spf13/cobra, for shell scripts.

### Example

![boa](docs/example.gif)


Similar to https://github.com/charmbracelet/gum, this binary file takes use of go parsing and validation.

As seen in the [example](docs/example.gif), the cli can be defined in a bash script like so:

```sh
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
```
