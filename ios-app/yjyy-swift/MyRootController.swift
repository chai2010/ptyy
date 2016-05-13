// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

class MyRootController: UITableViewController, UISearchBarDelegate {
	override func viewDidLoad() {
		super.viewDidLoad()

		let searchBar = UISearchBar(frame: CGRectMake(0, 0, tableView.frame.size.width, 0))
		searchBar.delegate = self
		searchBar.showsCancelButton = true
		searchBar.sizeToFit()
		searchBar.placeholder = "名字 或 拼音 或 正则"

		view.backgroundColor = UIColor.whiteColor()
		tableView.tableHeaderView = searchBar
		tableView.dataSource = self
		tableView.delegate = self
	}

	override func didReceiveMemoryWarning() {
		super.didReceiveMemoryWarning()
	}
}
