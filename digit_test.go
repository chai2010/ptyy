// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy

import (
	"testing"
)

func TestAdjustDigitString(t *testing.T) {
	tests := []struct {
		s      string
		s0, zh string
		ok     bool
	}{
		{"0", "0", "零", true},
		{"零", "0", "零", true},

		{"37", "37", "三七", true},
		{"3七", "37", "三七", true},
		{"三7", "37", "三七", true},
		{"三七", "37", "三七", true},
		{"三拾七", "37", "三七", true},

		{"3零7", "307", "三零七", true},

		{"0a", "", "", false},
	}
	for i, v := range tests {
		s0, zh, ok := adjustDigitString(v.s)
		if ok != v.ok {
			t.Fatalf("%d: faild, %v", i, v)
		}
		if s0 != v.s0 {
			t.Fatalf("%d: faild, %v", i, v)
		}
		if zh != v.zh {
			t.Fatalf("%d: faild, %v", i, v)
		}
	}
}
