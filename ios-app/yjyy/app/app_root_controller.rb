# Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
# Use of this source code is governed by a BSD-style
# license that can be found in the LICENSE file.

class AppRootController < UITableViewController
  def viewDidLoad
    super

    @searchWasCancelled = false
    @runeList = DataEngin.Search('', 100)

    searchBar = UISearchBar.alloc.initWithFrame(CGRectMake(0, 0, tableView.frame.size.width, 0))
    searchBar.delegate = self
    searchBar.showsCancelButton = false
    searchBar.sizeToFit
    searchBar.placeholder = "名字 或 拼音 或 正则"
    #searchBar.keyboardType = UIKeyboardTypeASCIICapable

    view.backgroundColor = UIColor.whiteColor
    view.tableHeaderView = searchBar
    view.dataSource = self
    view.delegate = self

    searchBarSearchButtonClicked(searchBar)
  end

  # 禁止旋转
  def shouldAutorotate
    true
  end

  def supportedInterfaceOrientations
    UIInterfaceOrientationMaskPortrait
  end

  def preferredInterfaceOrientationForPresentation
    UIInterfaceOrientationMaskPortrait
  end

  # 检索栏状态
  def searchBarTextDidBeginEditing(searchBar)
    @searchWasCancelled = false
    searchBar.showsCancelButton = true
  end

  def searchBarTextDidEndEditing(searchBar)
    searchBar.text = '' if @searchWasCancelled
  end

  def searchBar(searchBar, textDidChange:searchText)
    @searchWasCancelled = true
    searchBar.showsCancelButton = true

    # 获取查询条件
    query = searchBar.text

    # 根据查询条件查询结果
    @runeList = DataEngin.Search(query, 100)

    # 更新列表视图
    view.reloadData

    # 更新检索栏状态
    # searchBar.resignFirstResponder
  end

  def searchBarCancelButtonClicked(searchBar)
    @searchWasCancelled = true
    searchBar.showsCancelButton = false
    searchBar.text = ''

    # 根据查询条件查询结果
    @runeList = DataEngin.Search('', 100)

    # 更新列表视图
    view.reloadData

    # 更新检索栏状态
    searchBar.resignFirstResponder
  end

  def searchBarSearchButtonClicked(searchBar)
    @searchWasCancelled = false

    # 隐藏取消按钮
    searchBar.showsCancelButton = false

    # 获取查询条件
    query = searchBar.text

    # 根据查询条件查询结果
    @runeList = DataEngin.Search(query, 100)

    # 更新列表视图
    view.reloadData

    # 更新检索栏状态
    searchBar.resignFirstResponder
  end

  # 取消检索操作
  def searchBarCancelButtonClicked(searchBar)
    # 隐藏取消按钮
    searchBar.showsCancelButton = false

    # 根据查询条件查询结果(没有查询条件)
    @runeList = DataEngin.Search('', 100)

    # 更新列表视图
    view.reloadData

    # 更新检索栏状态
    searchBar.resignFirstResponder
  end

  def tableView(_tableView, numberOfRowsInSection:section)
    @runeList.count
  end

  def tableView(tableView, cellForRowAtIndexPath:indexPath)
    @reuseIdentifier ||= 'CELL_IDENTIFIER'
    cell = tableView.dequeueReusableCellWithIdentifier(@reuseIdentifier) || begin
      UITableViewCell.alloc.initWithStyle(UITableViewCellStyleDefault, reuseIdentifier: @reuseIdentifier)
    end
    cell.textLabel.text = @runeList[indexPath.row]
    cell
  end

  # 3.10 stuffs
  # if yes/true then iOS will display the context menu for the table view cell whose index
  # gets passed to you through the shouldShowMenueForRowAtIndexPath param
  def tableView(_tableView, shouldShowMenuForRowAtIndexPath:indexPath)
    true
  end

  def tableView(_tableView, canPerformAction:action, forRowAtIndexPath:indexPath, withSender:sender)
    p NSStringFromSelector(action)

    # filter actions, return true to enable and false to disable
    # But for now, always return true (enable all)
    action == :"copy:"
  end

  def tableView(tableView, performAction:action, forRowAtIndexPath:indexPath, withSender:sender)
    # Example implementation of how to do copy, should that have been selected
    if action == :"copy:"
      cell = tableView.cellForRowAtIndexPath(indexPath)
      paste_board = UIPasteboard.generalPasteboard
      # take current string and put it in paste
      paste_board.setString(cell.textLabel.text)
    end
  end
end
