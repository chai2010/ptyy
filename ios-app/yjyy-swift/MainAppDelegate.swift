
// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

enum MenuAction:String{
	case Copy = "copy:"
	case MyCustomSearch = "mySearchAction:"
	case MyCustomNavigate = "myNavigateAction:"

	static func supportedAction(action: Selector) -> Bool {
		return action == MenuAction.MyCustomSearch.selector() || action == MenuAction.MyCustomNavigate.selector()
	}


	// We need this awkward little conversion because «enum»'s can only have literals as raw value types.
	// And «Selector»s aren't literal.
	func selector()->Selector{
		return Selector(self.rawValue)
	}
}

class MyCustomMenuCell : UITableViewCell{
	static let ReuseIdentifier = "MyCustomMenuCell"

	// Bing搜索
	func mySearchAction(sender:AnyObject?) {
		if self.textLabel!.text == nil {
			return
		}

		let urlStr = "http://cn.bing.com/?q=\(self.textLabel!.text!)"
		let bingUrl = NSURL(string: urlStr.stringByAddingPercentEncodingWithAllowedCharacters(
			NSCharacterSet.URLQueryAllowedCharacterSet())!)

		if let url = bingUrl {
			if UIApplication.sharedApplication().canOpenURL(url) {
				UIApplication.sharedApplication().openURL(url)
			} else {
				print("can't open url: \(url)")
			}
		} else {
			print("cant open url: \(bingUrl)")
		}
	}
	
	// 地图导航
	func myNavigateAction(sender:AnyObject?) {
		if self.textLabel!.text == nil {
			return
		}
		
		let urlStr = "http://maps.apple.com/?q=\(self.textLabel!.text!)"
		let bingUrl = NSURL(string: urlStr.stringByAddingPercentEncodingWithAllowedCharacters(
			NSCharacterSet.URLQueryAllowedCharacterSet())!)
		
		if let url = bingUrl {
			if UIApplication.sharedApplication().canOpenURL(url) {
				UIApplication.sharedApplication().openURL(url)
			} else {
				print("can't open url: \(url)")
			}
		} else {
			print("cant open url: \(bingUrl)")
		}
	}
}

@UIApplicationMain
class MainAppDelegate: UIResponder, UIApplicationDelegate {

	var window: UIWindow?

	func application(application: UIApplication, didFinishLaunchingWithOptions launchOptions: [NSObject: AnyObject]?) -> Bool {

		// 添加自定义菜单
		addCustomMenuItems()

		// 创建主窗口(包含导航栏)
		window = UIWindow(frame: UIScreen.mainScreen().bounds)
		window!.rootViewController = MainViewController(nibName: "MainViewController", bundle: nil)
		window!.makeKeyAndVisible()

		// 延长启动界面显示时间
		if #available(iOS 8.0, *) {
			NSThread.sleepForTimeInterval(0.5)
		}

		return true
	}

 	private func addCustomMenuItems() {
		// 添加自定义的搜索按钮
		let mySearchItem = UIMenuItem(title: "搜索", action: MenuAction.MyCustomSearch.selector())
		let myNavigateItem = UIMenuItem(title: "地图", action: MenuAction.MyCustomNavigate.selector())

		let menuController = UIMenuController.sharedMenuController()
		var newItems = menuController.menuItems ?? [UIMenuItem]()
		
		newItems.append(mySearchItem)
		newItems.append(myNavigateItem)

		menuController.menuItems = newItems
	}
}
