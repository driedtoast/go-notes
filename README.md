
This is just a mess around space for looking at the best ways to manage a go project.

Goal is to just have a simple app that just adds a blob to a postgres db, eventually.

*Status*

* Created a `Makefile` for simple building
* Use a `.rvmrc` file for setting pathing 
* Add a line for each dependency I have in the `dependencies` file (this is loaded in make file)
* Simple `hello` package just to test out the make file / dependencies

*External Dependencies:*

* github.com/lib/pq
* github.com/nu7hatch/gouuid