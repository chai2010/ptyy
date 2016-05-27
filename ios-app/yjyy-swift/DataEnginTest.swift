// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import XCTest

@testable import yjyy

class DataEnginTest: XCTestCase {

    override func setUp() {
        super.setUp()
    }

    override func tearDown() {
        super.tearDown()
    }

	// 查询的性能
	func testBenchSearch() {
		self.measureBlock {
			let db = DataEngin()
			let result = db.SearchV2("")
			XCTAssert(!result.isEmpty)
		}
	}

}
