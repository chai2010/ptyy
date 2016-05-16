// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 改进思路:
// 1. 缓存最近的查询
// 2. 尽量缓存全部失败的查询
// 3. objc环境返回的数据采用NSData类型

package main

//#include "./yjyy.h"
//#include <stdint.h>
import "C"

import (
	"bytes"
	"fmt"

	"github.com/chai2010/ptyy"
)

//export YjyySearch
func YjyySearch(query *C.char, limits C.int32_t) *C.char {
	var buf bytes.Buffer
	var list = ptyy.Search(C.GoString(query), int(limits))
	for i, v := range list {
		if i < len(list)-1 {
			fmt.Fprintln(&buf, v.Name)
		} else {
			fmt.Fprint(&buf, v.Name)
		}
	}
	return C.CString(string(buf.Bytes()))
}

//export YjyyAdd
func YjyyAdd(a, b C.int32_t) C.int32_t {
	return a + b
}
