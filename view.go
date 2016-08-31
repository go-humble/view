// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package view

import (
	"strings"

	"honnef.co/go/js/dom"
)

var (
	document = dom.GetWindow().Document()
)

// View is an interface that must be satisfied by all views.
// You can embed DefaultView in a struct to satisfy the Element
// method automatically.
type View interface {
	Render() error
	Element() dom.Element
}

// Append appends child to a parent View. More specifically, it
// appends child.Element() to parent.Element() using the appendChild
// method from the DOM API.
func Append(parent View, child View) {
	parent.Element().AppendChild(child.Element())
}

// AppendToEl appends child to a parent element. More specifically, it
// appends child.Element() to parent using the appendChild method
// from the DOM API.
func AppendToEl(parent dom.Element, child View) {
	parent.AppendChild(child.Element())
}

// InsertBefore inserts v directly before before. More specifically, it
// inserts v.Element() before before.Element() using the insertBefore
// method from the DOM API.
func InsertBefore(v View, before View) {
	before.Element().ParentNode().InsertBefore(v.Element(), before.Element())
}

// InsertBeforeEl inserts v directly before before. More specifically, it
// inserts v.Element() using the insertBefore method from the DOM API.
func InsertBeforeEl(v View, before dom.Element) {
	before.ParentNode().InsertBefore(v.Element(), before)
}

// Replace replaces an old View with new. More specifically, it replaces
// old.Element() with new.Element() using the replaceChild method
// from the DOM API.
func Replace(new View, old View) {
	old.Element().ParentElement().ReplaceChild(new.Element(), old.Element())
}

// ReplaceEl replaces an old element with new. More specifically, it
// replaces old.Element() with new.Element() using the replaceChild
// method from the DOM API.
func ReplaceEl(new View, old dom.Element) {
	old.ParentElement().ReplaceChild(new.Element(), old)
}

// Remove removes the view from the DOM entirely. It does not destory the
// Element propery of the view.
func Remove(v View) {
	v.Element().ParentElement().RemoveChild(v.Element())
}

// Hide hides the view from the DOM by adding the inline style "display:none".
// Hide is safe to use even if you have other attributes and inline styles. It
// has no effect if the view is already hidden.
func Hide(v View) {
	oldStyles := v.Element().GetAttribute("style")
	newStyles := ""
	switch {
	case oldStyles == "":
		// There was no style attribute. We can safely set
		// the style attribute directly.
		newStyles = "display:none"
	case strings.Contains(oldStyles, "display:none"):
		// The element is already hidden. We should do
		// nothing.
		return
	case oldStyles[len(oldStyles)] == ';':
		// There was a style attribute and it ended in a semicolon,
		// We can safely append the new styles to the old.
		newStyles = oldStyles + "display:none;"
	default:
		// There was a style attribute and it didn't end in a semicolon,
		// in this case we should add our own semicolon.
		newStyles = oldStyles + ";display:none;"
	}
	v.Element().SetAttribute("style", newStyles)
}

// Show shows a previously hidden view by removing the inline style
// "display:none". Show is safe to use even if you have other attributes and
// inline styles. It has no effect if the view is already visible.
func Show(v View) {
	oldStyles := v.Element().GetAttribute("style")
	// Try removing the with a semicolon version first.
	// If there is not a semicolon, this will have no effect.
	newStyles := strings.Replace(oldStyles, "display:none;", "", 1)
	// Then try removing the without a semicolon version.
	// If there was a semicolon, this will have no effect.
	newStyles = strings.Replace(newStyles, "display:none", "", 1)
	v.Element().SetAttribute("style", newStyles)
}
