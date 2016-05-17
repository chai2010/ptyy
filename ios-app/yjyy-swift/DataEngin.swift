// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

extension String {
	subscript (i: Int) -> Character {
		return self[self.startIndex.advancedBy(i)]
	}
	
	subscript (i: Int) -> String {
		return String(self[i] as Character)
	}
	
	subscript (r: Range<Int>) -> String {
		let start = startIndex.advancedBy(r.startIndex)
		let end = start.advancedBy(r.endIndex - r.startIndex)
		return self[Range(start ..< end)]
	}
}

// 导入Go语言函数
@_silgen_name("YjyySearch")
func GoYjyySearch(query:UnsafePointer<CChar>, limits:Int32) -> UnsafeMutablePointer<CChar>

class DataEngin {
	let cache = LRUCache<String, [String]>(capacity: 16)
	let cacheV2 = LRUCache<String, [[String]]>(capacity: 16)

	func SearchV2(query:String) -> [[String]] {
		if let result = self.cacheV2[query] {
			return result
		}
		
		var resultMap = [String: [String]]()
		let letters : NSString = "ABCDEFGHIJKLMNOPQRSTUVWXYZ#"
		
		for i in 0..<letters.length {
			let c = letters.characterAtIndex(i)
			resultMap[String(UnicodeScalar(c))] = []
		}
		
		for s in self.Search(query) {
			if !s.isEmpty {
				let key = self.getFistLetter(s)
				resultMap[key]!.append(s)
			}
		}
		
		var resultList = [[String]]()
		for i in 0..<letters.length {
			let c = letters.characterAtIndex(i)
			resultList.append(resultMap[String(UnicodeScalar(c))]!)
		}

		return resultList
	}

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
	
	// 拼音首字母, [A-Z]#
	func getFistLetter(str: String)-> String {
		if str.isEmpty {
			return "#"
		}
		
		// 多音字: 长沙/涡阳/厦门
		if str[0] == "长" {
			return "C"
		}
		if str[0] == "涡" {
			return "G"
		}
		if str[0] == "厦" {
			return "X"
		}


		let mutStr = NSMutableString(string: str) as CFMutableString
		
		// 取得带音调拼音
		if CFStringTransform(mutStr, nil,kCFStringTransformMandarinLatin, false) {
			// 取得不带音调拼音
			if CFStringTransform(mutStr, nil, kCFStringTransformStripDiacritics, false) {
				let pinyin = (mutStr as String).uppercaseString
				if pinyin.characters.count > 0 {
					if pinyin[0] >= "A" && pinyin[0] <= "Z"{
						return pinyin[0]
					}
				}
			}
		}
		
		return "#"
	}
}
