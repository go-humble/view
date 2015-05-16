// Copyright 2015 Alex Browne and Soroush Pour.
// Allrights reserved. Use of this source code is
// governed by the MIT license, which can be found
// in the LICENSE file.

package view

import (
	"honnef.co/go/js/dom"
)

type DefaultView struct {
	el dom.Element
}

func (v *DefaultView) Element() dom.Element {
	if v.el == nil {
		v.el = document.CreateElement("div")
	}
	return v.el
}

func (v *DefaultView) SetElement(el dom.Element) {
	v.el = el
}
