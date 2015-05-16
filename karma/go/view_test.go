// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package main

import (
	"github.com/go-humble/view"
	"github.com/rusco/qunit"
	"honnef.co/go/js/dom"
)

var (
	document  = dom.GetWindow().Document()
	body      = document.QuerySelector("body")
	container dom.Element
)

func init() {
	container = document.CreateElement("div")
	container.SetID("container")
	body.AppendChild(container)
}

type TestView struct {
	view.DefaultView
}

func (v *TestView) Render() error {
	v.Element().SetInnerHTML("foo")
	return nil
}

type NoOpView struct {
	view.DefaultView
}

func (v *NoOpView) Render() error {
	// A no-op
	return nil
}

func main() {
	qunit.Test("Append", func(assert qunit.QUnitAssert) {
		defer reset()
		// Create the ul view wrapper
		outer := &NoOpView{}
		list := document.CreateElement("ul")
		container.AppendChild(list)
		outer.SetElement(list)
		// Append an inner view
		inner := &TestView{}
		inner.SetElement(document.CreateElement("li"))
		view.Append(outer, inner)
		_ = inner.Render()
		assert.Equal(container.InnerHTML(), "<ul><li>foo</li></ul>", "inner view was not appended to outer view")
	})
}

func reset() {
	container.SetInnerHTML("")
}
