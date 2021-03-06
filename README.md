signalbox
=========

An experimental Web-RTC signalling server written in Go. Designed to be compatible with the [signalling protocol](http://rtc.io/signalling-protocol.html#0) used with [rtc.io](http://rtc.io/).


[![Build Status](http://img.shields.io/travis/cfreeman/signalbox.svg?style=flat)](https://travis-ci.org/cfreeman/signalbox)&nbsp;
![alpha](https://img.shields.io/badge/stability-alpha-orange.svg?style=flat "Alpha")&nbsp;
![MIT License](http://img.shields.io/badge/license-MIT-blue.svg?style=flat "MIT License")

## Installation instructions (Ubuntu):

* [Download and Install Go](http://golang.org/doc/install)

* Get and compile Signalbox

		go get github.com/cfreeman/signalbox

* Install on your server

		sudo cp bin/signalbox /usr/sbin/signalbox

* Create signalbox configuration file

		vim /etc/signalbox.json
		{
		  "ListenAddress":":3000"
		}

* Create directory to hold signalbox logging output

		sudo mkdir /var/log/signalbox
		sudo chown /var/log/signalbox www-data
		sudo chgrp /var/log/signalbox www-data

* Create upstart configuration file

		vim /etc/init/signalbox.conf
		description "A webRTC signalling server."

		start on filesystem or runlevel [2345]
		stop on runlevel [!2345]

		exec start-stop

		exec start-stop-daemon --start --chuid www-data --exec /usr/sbin/signalbox /etc/signalbox.json 2>> /var/log/signalbox/signalbox.log

* Start signalbox

		sudo start signalbox

* Signalbox is now running on port 3000. You can open it up, or proxy pass from apache or nginx.



## License:

Copyright (c) 2014 Clinton Freeman

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.

