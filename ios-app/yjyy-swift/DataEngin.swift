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

	// 失败的查询(对于非正则只要包含, 则必然是失败的)
	// 正则的特例: .*?|
	let failedQueryCache = LRUCache<String, Bool>(capacity: 64)

	// 判断是非是失败的查询
	private func assertFailedQuery(query:String) -> Bool {
		// 1. 如果完全包含, 则必然是失败的查询
		if self.failedQueryCache[query] != nil {
			return true
		}

		// 2. 查看是否是失败集合的子集
		for (key, _) in self.failedQueryCache.Map() {
			if query.rangeOfString(key) != nil{
				return true
			}
		}
		
		// 不确定是否失败
		return false
	}

	// 注册失败的查询
	private func addFailedQuery(query:String) {
		// 正则的特例: .*?|
		for c in query.characters {
			for x in ".*?|。＊？｜".characters {
				if c == x {
					return
				}
			}
		}
		if self.assertFailedQuery(query) {
			return
		}
		self.failedQueryCache[query] = true
	}

	func SearchV2(rawQuery:String) -> [[String]] {
		// 删除空白
		let query = rawQuery.stringByTrimmingCharactersInSet(
			NSCharacterSet.whitespaceAndNewlineCharacterSet()
		)

		// 已经缓存成功的查询
		if let result = self.cacheV2[query] {
			return result
		}

		var resultMap = [String: [String]]()
		let letters : NSString = "ABCDEFGHIJKLMNOPQRSTUVWXYZ#"

		for i in 0..<letters.length {
			let c = letters.characterAtIndex(i)
			resultMap[String(UnicodeScalar(c))] = []
		}
		
		// 执行查询
		var resultTemp = [String]()

		// 如果不确定是失败的查询, 则进行查询操作
		if !self.assertFailedQuery(query) {
			resultTemp = self.Search(query)
		}

		// 处理查询结果
		for s in resultTemp {
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

		// 缓存成功和失败的查询
		if resultTemp.isEmpty {
			self.addFailedQuery(query)
		} else {
			self.cacheV2[query] = resultList
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
