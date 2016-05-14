// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

class MyRootController: UITableViewController, UISearchBarDelegate {

	let CellReuseIdentifier = "yjyy.result.cell"
	var searchWasCancelled = false

	var db:DataEngin = DataEngin()
	var results:[String] = []

	override func viewDidLoad() {
		super.viewDidLoad()

		self.title = "野鸡医院查询"
		self.view.backgroundColor = UIColor.whiteColor()

		let aboutBtn = UIBarButtonItem(title: "关于", style: UIBarButtonItemStyle.Plain, target: self, action:#selector(showAbout))
		self.navigationItem.rightBarButtonItem = aboutBtn

		let searchBar = UISearchBar(frame: CGRectMake(0, 0, tableView.frame.size.width, 0))
		searchBar.placeholder = "名字 或 拼音 或 正则"
		searchBar.showsCancelButton = false
		searchBar.delegate = self
		searchBar.sizeToFit()

		self.tableView.tableHeaderView = searchBar
		self.tableView.dataSource = self
		self.tableView.delegate = self

		// 注册TableViewCell
		self.tableView.registerClass(UITableViewCell.self, forCellReuseIdentifier: CellReuseIdentifier)

		// 生成初始列表
		self.searchBarSearchButtonClicked(searchBar)
	}
	
	// 关于
	func showAbout(b:UIBarButtonItem) {
		let message = "" +
		"用于查询中国大陆较常见的非公有或私人承包的野鸡医院或科室(以莆田系为主)，支持 拼音/汉字/正则 查询。\n" +
		"\n" +
		"原始数据来源于GitHub网站：http://github.com/open-power-workgroup/Hospital\n" +
		"\n" +
		"查询结果只是作为就医前的一个参考，本应用并不保证结果的真实性和准确性，请用户自行判断真伪。\n" +
		"\n" +
		"http://github.com/chai2010/ptyy\n" +
		"版权 2016"
		
		UIAlertView(
			title: "关于 野鸡医院",
			message: message,
			delegate: nil,
			cancelButtonTitle: "确定"
		).show()
	}

	// 表格单元数目
	override func tableView(tableView: UITableView, numberOfRowsInSection section: Int) -> Int {
		return self.results.count
	}

	// 表格单元
	override func tableView(tableView: UITableView, cellForRowAtIndexPath indexPath: NSIndexPath) -> UITableViewCell {
		let cell = tableView.dequeueReusableCellWithIdentifier(CellReuseIdentifier, forIndexPath: indexPath) as UITableViewCell
		cell.textLabel?.text = self.results[indexPath.row]
		return cell
	}

	// 点击搜索按钮
	func searchBarSearchButtonClicked(searchBar: UISearchBar) {
		self.searchWasCancelled = false
		searchBar.showsCancelButton = false

		// 根据查询条件查询结果
		self.results = self.db.Search(searchBar.text!)

		// 更新列表视图
		self.tableView.reloadData()

		searchBar.resignFirstResponder()
	}

	// 检索词发生变化
	func searchBar(searchBar: UISearchBar, textDidChange searchText: String) {
		self.searchWasCancelled = true
		searchBar.showsCancelButton = true

		// 根据查询条件查询结果
		self.results = self.db.Search(searchBar.text!)

		// 更新列表视图
		self.tableView.reloadData()
	}


	// 取消搜索
	func searchBarCancelButtonClicked(searchBar: UISearchBar) {
		// 隐藏取消按钮
		searchBar.showsCancelButton = false
		searchBar.text = ""

		// 根据查询条件查询结果(没有查询条件)
		self.results = self.db.Search("")

		// 更新列表视图
		self.tableView.reloadData()

		// 更新检索栏状态
		searchBar.resignFirstResponder()
	}

	func searchBarTextDidBeginEditing(searchBar: UISearchBar) {
		self.searchWasCancelled = false
		searchBar.showsCancelButton = true
	}

	override func tableView(tableView: UITableView, shouldShowMenuForRowAtIndexPath indexPath: NSIndexPath) -> Bool {
		return true
	}

	override func tableView(tableView: UITableView, canPerformAction action: Selector, forRowAtIndexPath indexPath: NSIndexPath, withSender sender: AnyObject?) -> Bool {
		return action == #selector(NSObject.copy(_:))
	}

	override func tableView(tableView: UITableView, performAction action: Selector, forRowAtIndexPath indexPath: NSIndexPath, withSender sender: AnyObject?) {
		if action == #selector(NSObject.copy(_:)) {
			let cell = tableView.cellForRowAtIndexPath(indexPath)
			UIPasteboard.generalPasteboard().string = cell!.textLabel!.text
		}
	}

	// 选择列时隐藏搜索键盘
	override func tableView(tableView: UITableView, didSelectRowAtIndexPath indexPath: NSIndexPath) {
		self.searchWasCancelled = true
		let searchBar = self.tableView.tableHeaderView as! UISearchBar
		searchBar.showsCancelButton = false

		// 更新检索栏状态
		searchBar.resignFirstResponder()

		// 已经选择的话, 则取消选择
		if indexPath == tableView.indexPathForSelectedRow {
			tableView.deselectRowAtIndexPath(indexPath, animated: true)
		}
	}

	// 内存报警
	override func didReceiveMemoryWarning() {
		super.didReceiveMemoryWarning()
	}

}
