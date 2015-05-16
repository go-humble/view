// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package view

import (
	"honnef.co/go/js/dom"
)

// DefaultView provides an implementation of the Element method
// of the View interface, and is meant to be embedded. When Element
// is called on DefaultView (or a struct that embeds DefaultView),
// it will create a new div element if an element has not yet been
// assigned to the view.
type DefaultView struct {
	el dom.Element
}

// Element satisfies View.Element. If the view does not already have
// an element assigned, it creates a new div element (but does not
// insert it into the DOM). If the view already has an element assigned,
// it returns the element.
func (v *DefaultView) Element() dom.Element {
	if v.el == nil {
		v.el = document.CreateElement("div")
	}
	return v.el
}

// SetElement can be used to manually set the element for a view.
func (v *DefaultView) SetElement(el dom.Element) {
	v.el = el
}
