// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

// 导入Go语言函数
@_silgen_name("YjyySearch")
func GoYjyySearch(query:UnsafePointer<CChar>, limits:Int32) -> UnsafeMutablePointer<CChar>

class DataEngin {
	func Search(query:String) -> [String] {
		let p = GoYjyySearch(query, limits: 0)
		let s = String.fromCString(p)!
		p.dealloc(1)
		
		return s.componentsSeparatedByCharactersInSet(NSCharacterSet.newlineCharacterSet())
	}
}
