// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy

import (
	"testing"
)

func TestSearch_onlyOne(t *testing.T) {
	infoList := Search("454", 0)
	tAssert(t, len(infoList) == 1, infoList)
	tAssert(t, infoList[0].Name, "454医院植发科")
	tAssert(t, infoList[0].City, "南京")
}
