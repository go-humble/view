Humble/View
=============

[![GoDoc](https://godoc.org/github.com/go-humble/view?status.svg)](https://godoc.org/github.com/go-humble/view)

Version 0.2.0

A small library for organizing view-related code written in pure go which
compiles to javascript via [gopherjs](https://github.com/gopherjs/gopherjs).
View includes a View interface and some helper functions for operating on views
(e.g. Append, Replace, Remove, AddEventListener, etc.). View works great as a
stand-alone package or in combination with other
[Humble](https://github.com/go-humble) packages.

View is written in pure go. It feels like go, follows go idioms when possible,
and compiles with the go tools. But it is meant to be compiled to javascript and
run in the browser.


Browser Support
---------------

View works with IE9+ (with a
[polyfill for typed arrays](https://github.com/inexorabletash/polyfill/blob/master/typedarray.js))
and all other modern browsers. View compiles to javascript via [gopherjs](https://github.com/gopherjs/gopherjs)
and this is a gopherjs limitation.

View is regularly tested with the latest versions of Firefox, Chrome, and Safari on Mac OS.
Each major or minor release is tested with IE9+ and the latest versions of Firefox and Chrome
on Windows.


Installation
------------

Install view like you would any other go package:

```bash
go get github.com/go-humble/view
```

You will also need to install gopherjs if you don't already have it. The latest version is
recommended. Install gopherjs with:

```
go get -u github.com/gopherjs/gopherjs
```


Quickstart Guide
----------------

### What is a View?

The `View` interface consists of only two methods:

```
type View interface {
	Render() error
	Element() dom.Element
}
```

Element expects an element object from the gopherjs
[dom bindings](http://dominik.honnef.co/go/js/dom) package.

### The Default View

If you want, you can embed `DefaultView` to satisfy the Element method.
When you call `Element` on `DefaultView`, if there is not already an element
assigned to the view, it creates a `<div>` element for you. `DefaultView` also
provides a `SetElement` method which you can call to set the element manually.
`DefaultView` is pretty simple, so here's the full implementation:

```go
type DefaultView struct {
	el dom.Element
}

func (v *DefaultView) Element() dom.Element {
	if v.el == nil {
		// Create an element if there is not one already
		v.el = document.CreateElement("div")
	}
	return v.el
}

func (v *DefaultView) SetElement(el dom.Element) {
	v.el = el
}

```

If you are using an embedded `DefaultView`, you will often need to work with a
pointer to your `View` type (e.g. `*TodoView` instead of `TodoView`). This is
because the `Element` method of `DefaultView` requires a pointer receiver.

### Defining the Render Method

DefaultView doesn't define a Render method, so if you are embedding you still
need to define one. Here's an example of a `TodoView`:

```go
type TodoView struct {
	title       string
	isCompleted bool
	view.DefaultView
}

func (todo TodoView) Render() {
	todo.Element().SetInnerHTML(`"<div class="todo-item" <span class="title"></span>`)
}
```

The `Render` method is really flexible and the view package is templating language agnostic.
For simpler views, it is totally fine to use string literals or `fmt.Sprintf`. However,
for more comlicated views, we highly recommend [temple](https://github.com/go-humble/temple),
a small wrapper around go's builtin template/html package that is designed to work with
gopherjs and humble.

### Using the Helper Functions

The view package provides a number of helper functions for inserting and removing views
from the DOM. The goal is to eliminate the need to use `github.com/gopherjs/gopherjs/js`
or the [dom package](http://dominik.honnef.co/go/js/dom) directly, at least for the most
common use cases. Here's a list of all the helper functions:

- [Append](http://godoc.org/github.com/go-humble/view#Append)
- [AppendToEl](http://godoc.org/github.com/go-humble/view#AppendToEl)
- [Replace](http://godoc.org/github.com/go-humble/view#Replace)
- [ReplaceEl](http://godoc.org/github.com/go-humble/view#ReplaceEl)
- [Remove](http://godoc.org/github.com/go-humble/view#Remove)
- [Hide](http://godoc.org/github.com/go-humble/view#Hide)
- [Show](http://godoc.org/github.com/go-humble/view#Show)

You can also view all the [documentation on godoc.org](http://godoc.org/github.com/go-humble/view).


Testing
-------

View uses the [karma test runner](http://karma-runner.github.io/0.12/index.html) to test
the code running in actual browsers.

The tests require the following additional dependencies:

- [node.js](http://nodejs.org/) (If you didn't already install it above)
- [karma](http://karma-runner.github.io/0.12/index.html)
- [karma-qunit](https://github.com/karma-runner/karma-qunit)

Don't forget to also install the karma command line tools with `npm install -g karma-cli`.

You will also need to install a launcher for each browser you want to test with, as well as the
browsers themselves. Typically you install a karma launcher with `npm install -g karma-chrome-launcher`.
You can edit the config file at `karma/test-mac.conf.js` or create a new one (e.g. `karma/test-windows.conf.js`)
if you want to change the browsers that are tested on.

Once you have installed all the dependencies, start karma with `karma start karma/test-mac.conf.js` (or
your customized config file, if applicable). Once karma is running, you can keep it running in between tests.

Next you need to compile the test.go file to javascript so it can run in the browsers:

```
gopherjs build karma/go/view_test.go -o karma/js/view_test.js
```

Finally run the tests with `karma run karma/test-mac.conf.js` (changing the name of the config file if needed).

If you are on a unix-like operating system, you can recompile and run the tests in one go by running
the provided bash script: `./karma/test.sh`.


Contributing
------------

See [CONTRIBUTING.md](https://github.com/go-humble/view/blob/master/CONTRIBUTING.md)


License
-------

View is licensed under the MIT License. See the [LICENSE](https://github.com/go-humble/view/blob/master/LICENSE)
file for more information.
