// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package searcher

import (
	"testing"
)

var tDbList = []map[string]string{
	{
		"123": "123",
		"234": "234",
		"200": "200",
	},
	{
		"abc": "abc",
		"bcd": "bcd",
		"x34": "x34",
	},
}

func TestSearcher_SearchByKeyAll(t *testing.T) {
	p := New(tDbList...)

	ss := p.SearchByKeyAll("2")
	tAssertTrue(t, len(ss) == 3)

	tAssertContain(t, ss, "123")
	tAssertContain(t, ss, "234")
	tAssertContain(t, ss, "200")

	ss = p.SearchByKeyAll("23")
	tAssertTrue(t, len(ss) == 2)

	tAssertContain(t, ss, "123")
	tAssertContain(t, ss, "234")

	ss = p.SearchByKeyAll("3")
	tAssertTrue(t, len(ss) == 3)

	tAssertContain(t, ss, "123")
	tAssertContain(t, ss, "234")
	tAssertContain(t, ss, "x34")
}

func TestSearcher_SearchByRegexpAll(t *testing.T) {
	//
}

func tAssertTrue(t *testing.T, v bool, arg ...interface{}) {
	if !v {
		t.Fatal(arg...)
	}
}

func tAssertContain(t *testing.T, ss []string, s string, arg ...interface{}) {
	for _, v := range ss {
		if v == s {
			return
		}
	}
	t.Fatal(arg...)
}
