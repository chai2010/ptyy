# 莆田系医院查询

[![GoDoc](https://godoc.org/github.com/chai2010/ptyy?status.svg)](https://godoc.org/github.com/chai2010/ptyy)

数据来源: https://github.com/open-power-workgroup/Hospital

iOS应用案例: https://appsto.re/cn/QH8ocb.i

![](./ios-app/yjyy/misc/appstore-yjyy.png)


## iOS应用案例截图

iOS版本的 [野鸡医院](./ios-app/yjyy) ( AppStore链接： https://appsto.re/cn/QH8ocb.i ), 底层使用了 `ptyy` 作为底层查询引擎:

![](./screenshots/yjyy-ios.png)


## 安装Go语言包

1. `go get github.com/chai2010/ptyy`
2. `go run hello.go`

## Go语言示例代码

```Go
package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/chai2010/ptyy"
)

var (
	flagLimits = flag.Int("limits", 0, "设置最大查询数目")
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage of %s: %s [query...]\n", filepath.Base(os.Args[0]), filepath.Base(os.Args[0]))
		flag.PrintDefaults()
	}
	flag.Parse()
	if flag.NArg() > 0 {
		for i := 0; i < flag.NArg(); i++ {
			for _, v := range ptyy.Search(flag.Arg(i), *flagLimits) {
				fmt.Printf("%s (%s)\n", v.Name, v.City)
			}
		}
	}
}
```

关键字查询:

```
go run hello.go -limits=0 武汉
武汉华仁医院 (武汉)
武汉华夏医院 (武汉)
武汉华美医院 (武汉)
武汉当代佳丽医院 (武汉)
武汉现代女子妇科医院 (武汉)
武汉阳光女子医院 (武汉)
湖北荣军医院 (武汉)
```

正则表达式查询:

```
go run hello.go "上海.*[男女]子"
上海九龙男子医院 (上海)
上海城市女子医院 (上海)
上海玛丽女子医院 (上海)
上海玫瑰女子医院 (上海)
上海真爱女子医院 (上海)
上海阿波罗男子医院 (上海)
```

支持拼音匹配:

```
go run hello.go wuhan
武汉华仁医院 (武汉)
武汉华夏医院 (武汉)
武汉华美医院 (武汉)
武汉当代佳丽医院 (武汉)
武汉现代女子妇科医院 (武汉)
武汉阳光女子医院 (武汉)
湖北荣军医院 (武汉)
```

## 报告问题

请联系 <chaishushan@gmail.com>.

谢谢!
