// Copyright ©2019-2020 by Richard A. Wilkes. All rights reserved.
//
// This Source Code Form is subject to the terms of the Mozilla Public
// License, version 2.0. If a copy of the MPL was not distributed with
// this file, You can obtain one at http://mozilla.org/MPL/2.0/.
//
// This Source Code Form is "Incompatible With Secondary Licenses", as
// defined by the Mozilla Public License, version 2.0.

package keys

// Known key codes.
var (
	A              = &Key{Code: 0x26, Name: "A"}
	B              = &Key{Code: 0x38, Name: "B"}
	C              = &Key{Code: 0x36, Name: "C"}
	D              = &Key{Code: 0x28, Name: "D"}
	E              = &Key{Code: 0x1a, Name: "E"}
	F              = &Key{Code: 0x29, Name: "F"}
	G              = &Key{Code: 0x2a, Name: "G"}
	H              = &Key{Code: 0x2b, Name: "H"}
	I              = &Key{Code: 0x1f, Name: "I"}
	J              = &Key{Code: 0x2c, Name: "J"}
	K              = &Key{Code: 0x2d, Name: "K"}
	L              = &Key{Code: 0x2e, Name: "L"}
	M              = &Key{Code: 0x3a, Name: "M"}
	N              = &Key{Code: 0x39, Name: "N"}
	O              = &Key{Code: 0x20, Name: "O"}
	P              = &Key{Code: 0x21, Name: "P"}
	Q              = &Key{Code: 0x18, Name: "Q"}
	R              = &Key{Code: 0x1b, Name: "R"}
	S              = &Key{Code: 0x27, Name: "S"}
	T              = &Key{Code: 0x1c, Name: "T"}
	U              = &Key{Code: 0x1e, Name: "U"}
	V              = &Key{Code: 0x37, Name: "V"}
	W              = &Key{Code: 0x19, Name: "W"}
	X              = &Key{Code: 0x35, Name: "X"}
	Y              = &Key{Code: 0x1d, Name: "Y"}
	Z              = &Key{Code: 0x34, Name: "Z"}
	One            = &Key{Code: 0x0a, Name: "1"}
	Two            = &Key{Code: 0x0b, Name: "2"}
	Three          = &Key{Code: 0x0c, Name: "3"}
	Four           = &Key{Code: 0x0d, Name: "4"}
	Five           = &Key{Code: 0x0e, Name: "5"}
	Six            = &Key{Code: 0x0f, Name: "6"}
	Seven          = &Key{Code: 0x10, Name: "7"}
	Eight          = &Key{Code: 0x11, Name: "8"}
	Nine           = &Key{Code: 0x12, Name: "9"}
	Zero           = &Key{Code: 0x13, Name: "0"}
	Return         = &Key{Code: 0x24, Name: "Return"}
	Escape         = &Key{Code: 0x09, Name: "Escape"}
	Backspace      = &Key{Code: 0x16, Name: "Backspace"}
	Tab            = &Key{Code: 0x17, Name: "Tab"}
	Space          = &Key{Code: 0x41, Name: "Space"}
	Minus          = &Key{Code: 0x14, Name: "Minus"}
	Equal          = &Key{Code: 0x15, Name: "="}
	LeftBracket    = &Key{Code: 0x22, Name: "["}
	RightBracket   = &Key{Code: 0x23, Name: "]"}
	Backslash      = &Key{Code: 0x33, Name: `\`}
	Semicolon      = &Key{Code: 0x2f, Name: ";"}
	Quote          = &Key{Code: 0x30, Name: "'"}
	Backquote      = &Key{Code: 0x31, Name: "`"}
	Comma          = &Key{Code: 0x3b, Name: ","}
	Period         = &Key{Code: 0x3c, Name: "."}
	Slash          = &Key{Code: 0x3d, Name: "/"}
	F1             = &Key{Code: 0x43, Name: "F1"}
	F2             = &Key{Code: 0x44, Name: "F2"}
	F3             = &Key{Code: 0x45, Name: "F3"}
	F4             = &Key{Code: 0x46, Name: "F4"}
	F5             = &Key{Code: 0x47, Name: "F5"}
	F6             = &Key{Code: 0x48, Name: "F6"}
	F7             = &Key{Code: 0x49, Name: "F7"}
	F8             = &Key{Code: 0x4a, Name: "F8"}
	F9             = &Key{Code: 0x4b, Name: "F9"}
	F10            = &Key{Code: 0x4c, Name: "F10"}
	F11            = &Key{Code: 0x5f, Name: "F11"}
	F12            = &Key{Code: 0x60, Name: "F12"}
	F13            = &Key{Code: 0xbf, Name: "F13"}
	F14            = &Key{Code: 0xc0, Name: "F14"}
	F15            = &Key{Code: 0xc1, Name: "F15"}
	Delete         = &Key{Code: 0x77, Name: "Delete"}
	Home           = &Key{Code: 0x6e, Name: "Home"}
	End            = &Key{Code: 0x73, Name: "End"}
	PageUp         = &Key{Code: 0x70, Name: "PageUp"}
	PageDown       = &Key{Code: 0x75, Name: "PageDown"}
	Left           = &Key{Code: 0x71, Name: "Left"}
	Up             = &Key{Code: 0x6f, Name: "Up"}
	Right          = &Key{Code: 0x72, Name: "Right"}
	Down           = &Key{Code: 0x74, Name: "Down"}
	Clear          = &Key{Code: 0x4d, Name: "Clear"}
	NumpadDivide   = &Key{Code: 0x6a, Name: "/"}
	NumpadMultiply = &Key{Code: 0x3f, Name: "*"}
	NumpadAdd      = &Key{Code: 0x56, Name: "+"}
	NumpadSubtract = &Key{Code: 0x52, Name: "-"}
	NumpadDecimal  = &Key{Code: 0x5b, Name: "."}
	NumpadEnter    = &Key{Code: 0x68, Name: "Enter"}
	Numpad1        = &Key{Code: 0x57, Name: "1"}
	Numpad2        = &Key{Code: 0x58, Name: "2"}
	Numpad3        = &Key{Code: 0x59, Name: "3"}
	Numpad4        = &Key{Code: 0x53, Name: "4"}
	Numpad5        = &Key{Code: 0x54, Name: "5"}
	Numpad6        = &Key{Code: 0x55, Name: "6"}
	Numpad7        = &Key{Code: 0x4f, Name: "7"}
	Numpad8        = &Key{Code: 0x50, Name: "8"}
	Numpad9        = &Key{Code: 0x51, Name: "9"}
	Numpad0        = &Key{Code: 0x5a, Name: "0"}
)
