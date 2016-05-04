// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy

import (
	"testing"
)

func TestSearch_onlyOne(t *testing.T) {
	infoList := Search("万州阳光眼科", 0)
	tAssert(t, len(infoList) == 1, infoList)
	tAssertEQ(t, infoList[0].Name, "万州阳光眼科医院")
	tAssertEQ(t, infoList[0].City, "重庆")
}

func TestSearchByRegexp(t *testing.T) {
	infoList, err := SearchByRegexp("^454", 0)
	tAssertNil(t, err)
	tAssert(t, len(infoList) == 1, infoList)

	infoList, err = SearchByRegexp("上海.*[男女]子", 0)
	tAssertNil(t, err)
	tAssert(t, len(infoList) > 1, infoList)
}
