# regx - Regular Expression Generator
[![build status](https://travis-ci.org/thepeterstone/regx.svg?branch=master)](https://travis-ci.org/thepeterstone/regx)

`regx` reads lines of input and attempts to create a compact regular expression that
will match all similar lines. 

For instance, given the input

> 1

> 2

> 3

> 4

`regx` will interpret this as a `[[:digit:]]` and match those lines as well as `5`, `6`, etc.

## Status

**`regx` is in development and is in no way ready to be used for anything
serious.**
