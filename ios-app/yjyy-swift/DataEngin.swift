// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

// 导入Go语言函数
@_silgen_name("YjyySearch")
func GoYjyySearch(query:UnsafePointer<CChar>, limits:Int32) -> UnsafeMutablePointer<CChar>

class DataEngin {
	let cache = LRUCache<String, [String]>(capacity: 16)

	func Search(query:String) -> [String] {
		if let result = self.cache[query] {
			return result
		}
		
		let p = GoYjyySearch(query, limits: 0)
		let s = String.fromCString(p)!
		p.dealloc(1)
		
		let result = s.componentsSeparatedByCharactersInSet(NSCharacterSet.newlineCharacterSet())
		self.cache[query] = result
		
		return result
	}
}
