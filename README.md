[![GoDoc](https://godoc.org/github.com/mschneider82/milter?status.svg)](https://godoc.org/github.com/mschneider82/milter)

# milter

This is a fork of github.com/phalaaxx/milter, added the following pull requests:

* https://github.com/phalaaxx/milter/pull/16
* https://github.com/phalaaxx/milter/pull/14
* https://github.com/phalaaxx/milter/pull/11

and also _test cases using my [milterclient](https://github.com/mschneider82/milterclient) library.


### Added the functions in the interface:

* Init() which is called before the first mail and after the end-of-body, and also on RSET (abort command), so you can cleanup and init.

* Disconnect() which is called when the client disconnects (if you have a concurrent session counter you can decrease the counter there, this was not possible before)
