// Copyright 2018 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package quote collects pithy sayings.
package quote // import "rsc.io/quote"

// Hello returns a greeting.
func Hello() string {
	return "Hello, world."
}

// Glass returns a useful phrase for world travelers.
func Glass() string {
	// See http://www.oocities.org/nodotus/hbglass.html.
	return "I can eat glass and it doesn't hurt me."
}

// Go returns a Go proverb.
func Go() string {
	return "Don't communicate by sharing memory, share memory by communicating."
}
