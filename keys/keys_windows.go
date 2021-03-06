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
	A              = &Key{Code: 0x41, Name: "A"}
	B              = &Key{Code: 0x42, Name: "B"}
	C              = &Key{Code: 0x43, Name: "C"}
	D              = &Key{Code: 0x44, Name: "D"}
	E              = &Key{Code: 0x45, Name: "E"}
	F              = &Key{Code: 0x46, Name: "F"}
	G              = &Key{Code: 0x47, Name: "G"}
	H              = &Key{Code: 0x48, Name: "H"}
	I              = &Key{Code: 0x49, Name: "I"}
	J              = &Key{Code: 0x4a, Name: "J"}
	K              = &Key{Code: 0x4b, Name: "K"}
	L              = &Key{Code: 0x4c, Name: "L"}
	M              = &Key{Code: 0x4d, Name: "M"}
	N              = &Key{Code: 0x4e, Name: "N"}
	O              = &Key{Code: 0x4f, Name: "O"}
	P              = &Key{Code: 0x50, Name: "P"}
	Q              = &Key{Code: 0x51, Name: "Q"}
	R              = &Key{Code: 0x52, Name: "R"}
	S              = &Key{Code: 0x53, Name: "S"}
	T              = &Key{Code: 0x54, Name: "T"}
	U              = &Key{Code: 0x55, Name: "U"}
	V              = &Key{Code: 0x56, Name: "V"}
	W              = &Key{Code: 0x57, Name: "W"}
	X              = &Key{Code: 0x58, Name: "X"}
	Y              = &Key{Code: 0x59, Name: "Y"}
	Z              = &Key{Code: 0x5a, Name: "Z"}
	One            = &Key{Code: 0x31, Name: "1"}
	Two            = &Key{Code: 0x32, Name: "2"}
	Three          = &Key{Code: 0x33, Name: "3"}
	Four           = &Key{Code: 0x34, Name: "4"}
	Five           = &Key{Code: 0x35, Name: "5"}
	Six            = &Key{Code: 0x36, Name: "6"}
	Seven          = &Key{Code: 0x37, Name: "7"}
	Eight          = &Key{Code: 0x38, Name: "8"}
	Nine           = &Key{Code: 0x39, Name: "9"}
	Zero           = &Key{Code: 0x30, Name: "0"}
	Return         = &Key{Code: 0x0d, Name: "Return"}
	Escape         = &Key{Code: 0x1b, Name: "Escape"}
	Backspace      = &Key{Code: 0x08, Name: "Backspace"}
	Tab            = &Key{Code: 0x09, Name: "Tab"}
	Space          = &Key{Code: 0x20, Name: "Space"}
	Minus          = &Key{Code: 0xbd, Name: "Minus"}
	Equal          = &Key{Code: 0xbb, Name: "="}
	LeftBracket    = &Key{Code: 0xdb, Name: "["}
	RightBracket   = &Key{Code: 0xdd, Name: "]"}
	Backslash      = &Key{Code: 0xde, Name: `\`}
	Semicolon      = &Key{Code: 0xba, Name: ";"}
	Quote          = &Key{Code: 0xc0, Name: "'"}
	Backquote      = &Key{Code: 0xdc, Name: "`"}
	Comma          = &Key{Code: 0xbc, Name: ","}
	Period         = &Key{Code: 0xbe, Name: "."}
	Slash          = &Key{Code: 0xbf, Name: "/"}
	F1             = &Key{Code: 0x70, Name: "F1"}
	F2             = &Key{Code: 0x71, Name: "F2"}
	F3             = &Key{Code: 0x72, Name: "F3"}
	F4             = &Key{Code: 0x73, Name: "F4"}
	F5             = &Key{Code: 0x74, Name: "F5"}
	F6             = &Key{Code: 0x75, Name: "F6"}
	F7             = &Key{Code: 0x76, Name: "F7"}
	F8             = &Key{Code: 0x77, Name: "F8"}
	F9             = &Key{Code: 0x78, Name: "F9"}
	F10            = &Key{Code: 0x79, Name: "F10"}
	F11            = &Key{Code: 0x7a, Name: "F11"}
	F12            = &Key{Code: 0x7b, Name: "F12"}
	F13            = &Key{Code: 0x7c, Name: "F13"}
	F14            = &Key{Code: 0x7d, Name: "F14"}
	F15            = &Key{Code: 0x7e, Name: "F15"}
	Delete         = &Key{Code: 0x2e, Name: "Delete"}
	Home           = &Key{Code: 0x24, Name: "Home"}
	End            = &Key{Code: 0x23, Name: "End"}
	PageUp         = &Key{Code: 0x21, Name: "PageUp"}
	PageDown       = &Key{Code: 0x22, Name: "PageDown"}
	Left           = &Key{Code: 0x25, Name: "Left"}
	Up             = &Key{Code: 0x26, Name: "Up"}
	Right          = &Key{Code: 0x27, Name: "Right"}
	Down           = &Key{Code: 0x28, Name: "Down"}
	Clear          = &Key{Code: 0x90, Name: "Clear"}
	NumpadDivide   = &Key{Code: 0x6f, Name: "/"}
	NumpadMultiply = &Key{Code: 0x6a, Name: "*"}
	NumpadAdd      = &Key{Code: 0x6b, Name: "+"}
	NumpadSubtract = &Key{Code: 0x6c, Name: "-"}
	NumpadDecimal  = &Key{Code: 0x6d, Name: "."}
	NumpadEnter    = &Key{Code: 0x0d, Name: "Enter"}
	Numpad1        = &Key{Code: 0x61, Name: "1"}
	Numpad2        = &Key{Code: 0x62, Name: "2"}
	Numpad3        = &Key{Code: 0x63, Name: "3"}
	Numpad4        = &Key{Code: 0x64, Name: "4"}
	Numpad5        = &Key{Code: 0x65, Name: "5"}
	Numpad6        = &Key{Code: 0x66, Name: "6"}
	Numpad7        = &Key{Code: 0x67, Name: "7"}
	Numpad8        = &Key{Code: 0x68, Name: "8"}
	Numpad9        = &Key{Code: 0x69, Name: "9"}
	Numpad0        = &Key{Code: 0x60, Name: "0"}
)
