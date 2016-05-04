// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy_test

import (
	"fmt"
	"log"

	"github.com/chai2010/ptyy"
)

func ExampleSearch_wuhan() {
	for _, v := range ptyy.Search("武汉", 0) {
		fmt.Printf("%s (%s)\n", v.Name, v.City)
	}

	// Output:
	// 武汉华仁医院 (武汉)
	// 武汉华夏医院 (武汉)
	// 武汉华美医院 (武汉)
	// 武汉当代佳丽医院 (武汉)
	// 武汉现代女子妇科医院 (武汉)
	// 武汉阳光女子医院 (武汉)
	// 湖北荣军医院 (武汉)
}

func ExampleSearch_wuhanLimit3() {
	for _, v := range ptyy.Search("武汉", 3) {
		fmt.Printf("%s (%s)\n", v.Name, v.City)
	}

	// Output:
	// 武汉华仁医院 (武汉)
	// 武汉华夏医院 (武汉)
	// 武汉华美医院 (武汉)
}

func ExampleSearchByRegexp() {
	results, err := ptyy.SearchByRegexp("上海.*[男女]子", 0)
	if err != nil {
		log.Fatal(err)
	}

	for _, v := range results {
		fmt.Printf("%s (%s)\n", v.Name, v.City)
	}

	// Output:
	// 上海九龙男子医院 (上海)
	// 上海城市女子医院 (上海)
	// 上海玛丽女子医院 (上海)
	// 上海玫瑰女子医院 (上海)
	// 上海真爱女子医院 (上海)
	// 上海阿波罗男子医院 (上海)
}
