// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

import UIKit

@UIApplicationMain
class MyAppDelegate: UIResponder, UIApplicationDelegate {

	var window: UIWindow?

	func application(application: UIApplication, didFinishLaunchingWithOptions launchOptions: [NSObject: AnyObject]?) -> Bool {

		// 主界面控制器
		let myRootController = MyRootController()
		myRootController.title = "野鸡医院查询"
		myRootController.view.backgroundColor = UIColor.whiteColor()

		// 导航栏
		let navigationController = UINavigationController(rootViewController: myRootController)

		// 创建主窗口
		window = UIWindow(frame: UIScreen.mainScreen().bounds)
		window!.rootViewController = navigationController
		window!.makeKeyAndVisible()

		// 延长启动界面显示时间
		NSThread.sleepForTimeInterval(2.0)
		return true
	}

	func applicationWillResignActive(application: UIApplication) {
		//
	}

	func applicationDidEnterBackground(application: UIApplication) {
		//
	}

	func applicationWillEnterForeground(application: UIApplication) {
		//
	}

	func applicationDidBecomeActive(application: UIApplication) {
		//
	}

	func applicationWillTerminate(application: UIApplication) {
		//
	}

}
