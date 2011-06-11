gocurses
========

go bindings for curses library

fetching
--------

on the shell run::

	git clone https://github.com/dforsyth/gocurse.git


building
--------

to build the bindings you will need to:

 * set the GOROOT environment variable

   * if not set when you installed go use: export GOROOT=/path/to/go
   * you should add this to your .bashrc file

 * have the development files for the curses library

   * on fedora: sudo yum install ncurses-devel

after this run::

	make
	make install

running the example
-------------------

::

	8g sample.go
	8l -o sample sample.8
	./sample

