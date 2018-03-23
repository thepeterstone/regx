# regx - Regular Expression Generator
[![build status][https://travis-ci.org/thepeterstone/regx.svg?branch=master][https://travis-ci.org/thepeterstone/regx]]

`regx` reads lines of input and attempts to create a compact regular expression that
will match all similar lines. 

For instance, given the input

> 1

> 2

> 3

> 4

`regx` will interpret this as `\d` and match those lines as well as `5`, `6`, etc.

## Status

**`regx` is in development and is in no way ready to be used for anything
serious.**

## License

This work is [free software](http://www.fsf.org/about/); you can redistribute it and/or modify it under the
terms of the [GNU General Public License](http://www.gnu.org/licenses/gpl.html) 
as published by the Free Software Foundation; either version 3 of the License, or
any later version. This work is distributed in the hope that it will be useful, **but
without any warranty**; without even the implied warranty of **merchantability** or
**fitness for a particular purpose**.

