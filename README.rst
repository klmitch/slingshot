========================
Slingshot Plugin Manager
========================

.. image:: https://travis-ci.org/klmitch/slingshot.svg?branch=master
    :target: https://travis-ci.org/klmitch/slingshot
.. image:: https://coveralls.io/repos/github/klmitch/slingshot/badge.svg?branch=master
    :target: https://coveralls.io/github/klmitch/slingshot?branch=master
.. image:: https://godoc.org/github.com/klmitch/slingshot?status.svg
    :target: http://godoc.org/github.com/klmitch/slingshot

Package slingshot contains a plugin manager built on top of the
standard plugin package.  Plugins are loaded using the Load function,
which loads the specified plugin and invokes its SlingshotInit
function; this function will be passed a Slingshot object and any
parameters passed to the Load function, and is expected to use the
Register method of the Slingshot object to register one or more
plugins for use by the application (or the application's libraries).
