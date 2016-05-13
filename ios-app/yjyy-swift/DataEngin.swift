// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

@_silgen_name("YjyyAdd")
func MyAdd(a:Int32, b:Int32) -> Int32

class DataEngin {
	func Search(query:String) -> [String] {
		let sum = MyAdd(3, b: 20)
		print(Int(sum))
		
		if query == "" {
			return [ "123", "abc" ]
		}
		return ["111","222"]
	}
}
