========================
Slingshot Plugin Manager
========================

.. image:: https://img.shields.io/github/tag/klmitch/slingshot.svg
    :target: https://github.com/klmitch/slingshot/tags
.. image:: https://img.shields.io/hexpm/l/plug.svg
    :target: https://github.com/klmitch/slingshot/blob/master/LICENSE
.. image:: https://travis-ci.org/klmitch/slingshot.svg?branch=master
    :target: https://travis-ci.org/klmitch/slingshot
.. image:: https://coveralls.io/repos/github/klmitch/slingshot/badge.svg?branch=master
    :target: https://coveralls.io/github/klmitch/slingshot?branch=master
.. image:: https://godoc.org/github.com/klmitch/slingshot?status.svg
    :target: http://godoc.org/github.com/klmitch/slingshot
.. image:: https://img.shields.io/github/issues/klmitch/slingshot.svg
    :target: https://github.com/klmitch/slingshot/issues
.. image:: https://img.shields.io/github/issues-pr/klmitch/slingshot.svg
    :target: https://github.com/klmitch/slingshot/pulls
.. image:: https://goreportcard.com/badge/github.com/klmitch/slingshot
    :target: https://goreportcard.com/report/github.com/klmitch/slingshot

Package slingshot contains a plugin manager built on top of the
standard plugin package.  Plugins are loaded using the Load function,
which loads the specified plugin and invokes its SlingshotInit
function; this function will be passed a Slingshot object and any
parameters passed to the Load function, and is expected to use the
Register method of the Slingshot object to register one or more
plugins for use by the application (or the application's libraries).
