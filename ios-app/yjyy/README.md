# 莆田系等 野鸡医院 查询

*关键字: 莆田, 魏则西, 莆田医院, 野鸡医院, 百度*

用于查询中国大陆比较常见的非公有或私人承包的野鸡医院。

## 运行效果

AppStore: https://appsto.re/cn/QH8ocb.i

![](./misc/screenshots/yjyy-ios.png)

扫码安装:

![](./misc/appstore-yjyy.png)

## 构建步骤

**安装环境**

- 安装Go语言
- 安装Ruby环境
- 安装RubyMotion(下载免费版): http://www.rubymotion.com/
- 安装 `ptyy` Go语言包: `go get github.com/chai2010/ptyy`
- 安装 Xcode 7.3 (参考RubyMotion具体要求)
- （可选）安装 `ios-deploy` 工具: `npm install ios-deploy`


**模拟器运行**

构建并在模拟器运行: `rake gopkg && rake`

**真机运行**

真机运行需要配置用于签名的账号和 provisioning profile 配置文件。

如果已经加入Apple的付费账号，将相关信息填入 Rakefile 文件开头的配置参数中。
然后usb链接上设备，然后运行命令: `rake gopkg && rake device`

如果是免费的账号，需要自己是用Xcode生成对应的provisioning profile 配置文件，
然后运行 `rake build:device` 命令构建目标程序，然后运行 `rake mydevide` 安装到设备（依赖ios-deploy工具）。


## 报告问题

请联系 <chaishushan@gmail.com>.

谢谢!
