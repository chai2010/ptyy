// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

class MyRootController: UITableViewController, UISearchBarDelegate {

	var results:[String] = [ "123", "abc" ]
	
	override func viewDidLoad() {
		super.viewDidLoad()
		
		self.title = "野鸡医院查询"
		self.view.backgroundColor = UIColor.whiteColor()

		let searchBar = UISearchBar(frame: CGRectMake(0, 0, tableView.frame.size.width, 0))
		searchBar.placeholder = "名字 或 拼音 或 正则"
		searchBar.showsCancelButton = false
		searchBar.delegate = self
		searchBar.sizeToFit()

		self.tableView.tableHeaderView = searchBar
		self.tableView.dataSource = self
		self.tableView.delegate = self
		
		// 注册TableViewCell
		self.tableView.registerClass(UITableViewCell.self, forCellReuseIdentifier: "cell")
	}

	override func didReceiveMemoryWarning() {
		super.didReceiveMemoryWarning()
	}
	
	func tableView(tableView tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
		return 2 //self.results.count
	}
	
	func tableView(tableView tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
		let identify:String = "cell"
		let cell = tableView.dequeueReusableCellWithIdentifier(identify, forIndexPath: indexPath) as UITableViewCell
		//cell.accessoryType = UITableViewCellAccessoryType.DisclosureIndicator
		cell.textLabel?.text = "22" //self.results[indexPath.row]
		print("aaaa")
		return cell
	}
	
	func searchBarSearchButtonClicked(searchBar searchBar: UISearchBar) {
		searchBar.resignFirstResponder()
	}

}
