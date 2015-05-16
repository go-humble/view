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

type ContentView struct {
	content string
	view.DefaultView
}

func (v *ContentView) Render() error {
	v.Element().SetInnerHTML(v.content)
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
		inner := &ContentView{
			content: "foo",
		}
		inner.SetElement(document.CreateElement("li"))
		_ = inner.Render()
		view.Append(outer, inner)
		assert.Equal(container.InnerHTML(), "<ul><li>foo</li></ul>", "inner view was not appended to outer view")
	})

	qunit.Test("AppendToEl", func(assert qunit.QUnitAssert) {
		defer reset()
		// Create the ul element
		list := document.CreateElement("ul")
		container.AppendChild(list)
		// Append an inner view
		inner := &ContentView{
			content: "foo",
		}
		inner.SetElement(document.CreateElement("li"))
		_ = inner.Render()
		view.AppendToEl(list, inner)
		assert.Equal(container.InnerHTML(), "<ul><li>foo</li></ul>", "inner view was not appended to outer view")
	})

	qunit.Test("Replace", func(assert qunit.QUnitAssert) {
		defer reset()
		// Create the first view
		fooView := &ContentView{
			content: "foo",
		}
		_ = fooView.Render()
		view.AppendToEl(container, fooView)
		// Create the view which will replace it
		barView := &ContentView{
			content: "bar",
		}
		_ = barView.Render()
		view.Replace(barView, fooView)
		assert.Equal(container.InnerHTML(), "<div>bar</div>", "inner view was not appended to outer view")
	})

	qunit.Test("ReplaceEl", func(assert qunit.QUnitAssert) {
		defer reset()
		// Create the element to be replaced
		fooEl := document.CreateElement("div")
		fooEl.SetInnerHTML("foo")
		container.AppendChild(fooEl)
		// Create the view which will replace it
		barView := &ContentView{
			content: "bar",
		}
		_ = barView.Render()
		view.ReplaceEl(barView, fooEl)
		assert.Equal(container.InnerHTML(), "<div>bar</div>", "inner view was not appended to outer view")
	})

	qunit.Test("Remove", func(assert qunit.QUnitAssert) {
		defer reset()
		removeMe := &ContentView{
			content: "removeMe",
		}
		_ = removeMe.Render()
		view.AppendToEl(container, removeMe)
		view.Remove(removeMe)
		assert.Equal(container.InnerHTML(), "", "inner view was not appended to outer view")
	})

	qunit.Test("Hide", func(assert qunit.QUnitAssert) {
		defer reset()
		hideMe := &ContentView{
			content: "hideMe",
		}
		view.AppendToEl(container, hideMe)
		// Add some additional attributes to make sure they are not messed up.
		hideMe.Element().SetAttribute("data-power-level", "9001")
		hideMe.Element().SetAttribute("style", `color:#ff0000`)
		view.Hide(hideMe)
		assert.Ok(hideMe.Element().HasAttribute("data-power-level"), "data-power-level attribute was removed")
		assert.Ok(hideMe.Element().HasAttribute("style"), "style attribute was removed")
		assert.Equal(hideMe.Element().GetAttribute("style"), `color:#ff0000;display:none;`, "attributes were not set correctly")
	})

	qunit.Test("Show", func(assert qunit.QUnitAssert) {
		defer reset()
		showMe := &ContentView{
			content: "showMe",
		}
		view.AppendToEl(container, showMe)
		// Add some additional attributes to make sure they are not messed up.
		showMe.Element().SetAttribute("data-answer-to-everything", "42")
		showMe.Element().SetAttribute("style", `color:#ff0000`)
		view.Hide(showMe)
		view.Show(showMe)
		assert.Ok(showMe.Element().HasAttribute("data-answer-to-everything"), "data-answer-to-everything attribute was removed. Maybe it will appear again in  7.5 million years?")
		assert.Ok(showMe.Element().HasAttribute("style"), "style attribute was removed")
		assert.Equal(showMe.Element().GetAttribute("style"), `color:#ff0000;`, "attributes were not set correctly")
	})
}

func reset() {
	container.SetInnerHTML("")
}
