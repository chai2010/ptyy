// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

class DataEngin {
	func Search(query:String) -> [String] {
		if query == "" {
			return [ "123", "abc" ]
		}
		return ["111","222"]
	}
}
