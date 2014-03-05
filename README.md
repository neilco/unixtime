# unixtime

__A command-line utility for converting a date or date & time into a UNIX timestamp__

### Usage

`unixtime [datetime]`

`datetime`: The date or date & time to parse. Can be one of the following formats:

	  now
	  YYYY-MM-DD
	  YYYY-MM-DD HH:MM
	  YYYY-MM-DD HH:MM:SS
	  YYYY-MM-DD HH:MM:SS.SSS
	  YYYY-MM-DDTHH:MMZ
	  YYYY-MM-DDTHH:MM:SSZ
	  YYYY-MM-DDTHH:MM:SS.SSSZ

If no argument is supplied, `now` is assumed.

### Build

`go build -o unixtime main.go`

### Contact

[Neil Cowburn](http://github.com/neilco)  
[@neilco](https://twitter.com/neilco)

## License

[MIT license](http://neil.mit-license.org)

Copyright (c) 2014 Neil Cowburn (http://github.com/neilco/)

MIT License

Permission is hereby granted, free of charge, to any person obtaining
a copy of this software and associated documentation files (the
"Software"), to deal in the Software without restriction, including
without limitation the rights to use, copy, modify, merge, publish,
distribute, sublicense, and/or sell copies of the Software, and to
permit persons to whom the Software is furnished to do so, subject to
the following conditions:

The above copyright notice and this permission notice shall be
included in all copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND,
EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND
NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE
LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION
OF CONTRACT, TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION
WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
