[![GoDoc](https://godoc.org/github.com/cwedgwood/milter?status.svg)](https://godoc.org/github.com/cwedgwood/milter)

This is an incompatible fork of
https://github.com/mschneider82/milter, which is an excellent update
to github.com/phalaaxx/milter.

# API / Interface

The present API is not guaranteed to be stable though any changes will
be minimized.

There were quite a few iterations as code using this evolved to get it
to it's present state and now various projects depend on the interface
not breaking.


# Fork-Specific Changes

* The interface has been altered from mschneider82's code such that
  it's now easier to determine the boundaries between messages.

* The session and message IDs have been removed, they can be
  implemented in the specific Milter code if needed.

* A small race in .Close() has been fixed.

* Macros accumulate/mutate throughout the session and are now usable.

* Testing; updated for new API. `Milter` is no longer embedded in the
  `TestMilter` as it's not necessary or desirable.

* At least one chance after this was forked has been ignored so far
  (IPv6 addresses are working as expected for me, if a change is made
  to trim the IPv6 prefix is should probably only be done after
  checking it exists).

# Rough Guide to the Lifecycle of Calls

`NewSession` is called at the start of milter session.  `EndSession`
is called when it's closed.

`Connect` is called once per session, it contains the string (name) of
the remote host and it's address information.

`Helo` is called one (or more) times per session in response to `HELO`
or `EHLO`.

`Reset` can be called at almost anytime to reset message-specific
state, this corresponds to SMTPs `RSET` verb.

`NewMessage` is called before an email message is to be delivered.
Usually you will see one message per session, but it's possible for
there to be zero or two or more.  Some high-performance MTAs will
reuse SMTP connections so that you will see multiple messages from
difference sources to different addresses in a given session.

`MailFrom` is called for `MAIL FROM`, one of which per message.

`RcptTo` is called for `RCPT TO`, one or more of which per message.

`Header` is called for head header passed in to the `BODY` stage of SMTP transaction.

`Headers` is called upon completion of all headers.

`BodyChunk` is called (possibly many times) with parts (chunks) of the
message `BODY`.

`Body` is called upon completion of the entire `BODY`.


## Example Sessions

# Macros

Macros are available from some function at `milter.Modifiers.Macros`,
these accumulate and mutate throughout the session.  Some macro values
are valid for the duration of the session, some are only valid for the
current or previous message.  The lifecycle of macros is not well
defined, this usually isn't a problem for session level state but for
things that change per-message use with care.

<!--  LocalWords:  GoDoc mschneider Milter TestMilter Lifecycle Helo
 -->
<!--  LocalWords:  NewSession milter EndSession HELO EHLO SMTPs RSET
 -->
<!--  LocalWords:  NewMessage MTAs SMTP MailFrom RcptTo BodyChunk IPv
 -->
<!--  LocalWords:  lifecycle
 -->
