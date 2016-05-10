// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

//#include "./libgo.h"
import "C"

import (
	"bytes"
	"fmt"

	"github.com/chai2010/ptyy"
)

//export GopkgYjyySearch
func GopkgYjyySearch(query *C.char, limits C.int) *C.char {
	var buf bytes.Buffer
	for _, v := range ptyy.Search(C.GoString(query), int(limits)) {
		fmt.Fprintln(&buf, v.Name)
	}
	return C.CString(string(buf.Bytes()))
}
