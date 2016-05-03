// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build ingore

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"unicode/utf8"
)

type HospitalInfo struct {
	Name    string   // 名称
	City    string   // 城市
	Owner   []string // 投资者
	Comment []string // 注释
}

func main() {
	// 读取列表文件
	data, err := ioutil.ReadFile("./list.txt")
	if err != nil {
		log.Fatal(err)
	}

	// 分析行信息
	var (
		lines           = strings.Split(string(data), "\n")
		allInfoMap      = make(map[string]*HospitalInfo)
		curInfo         HospitalInfo
		curInfo_hasElem = false
	)

	for i := 0; i < len(lines); i++ {
		curLine := strings.TrimSpace(lines[i])

		// 跳过忽略的行
		if isIngoreLine(curLine) {
			continue
		}

		// 城市名/公司名/注释
		if curLine[0] == '#' {
			if curInfo_hasElem {
				curInfo = HospitalInfo{}
				curInfo_hasElem = false
			}

			s := strings.TrimLeft(curLine, "# \t")
			switch {
			case isCityName(s):
				curInfo.City = s
			case isCompanyName(s):
				curInfo.Owner = append(curInfo.Owner, s)
			default:
				curInfo.Comment = append(curInfo.Comment, s)
			}
		}

		// 是否为医院名
		if curLine[0] == '-' {
			curInfo_hasElem = true

			s := strings.TrimLeft(curLine, "- \t")
			if isHospitalName(s) {
				curInfo.Name = s
			}
		}

		// 打印信息
		if curInfo.Name != "" {
			pInfo := allInfoMap[curInfo.Name]
			if pInfo == nil {
				info := curInfo
				allInfoMap[curInfo.Name] = &info
				pInfo = &info
			}
			pInfo.Owner = append(pInfo.Owner, curInfo.Owner...)
			pInfo.Comment = append(pInfo.Comment, curInfo.Comment...)

		}
	}

	// 输出
	for _, pInfo := range allInfoMap {
		fmt.Printf("%#v\n", pInfo)
	}
}

// 是否为忽略的行
func isIngoreLine(line string) bool {
	if strings.TrimSpace(line) == "" {
		return true
	}
	for _, key := range g_ingoreLineKeywordsList {
		if strings.Contains(line, key) {
			return true
		}
	}
	return false
}

// 是否是城市名
func isCityName(name string) bool {
	name = strings.TrimSpace(name)
	for s, _ := range g_CityMap {
		if s == name {
			return true
		}
	}
	return false
}

// 公司名
func isCompanyName(name string) bool {
	if utf8.RuneCountInString(name) >= 20 {
		return false
	}
	for _, key := range g_CompanyNameKeywordsList {
		if strings.Contains(name, key) {
			return true
		}
	}
	return false
}

// 是否为医院名(`-`开头)
func isHospitalName(name string) bool {
	for _, key := range g_HospitalKeywordsList {
		if strings.Contains(name, key) {
			return true
		}
	}
	return false
}

// 忽略行的关键字
var g_ingoreLineKeywordsList = []string{
	"BlackheartedHospital",
	"最新补充",
	"欢迎更新",
	"版本",
	"- 1.",
	"- 2.",
	"- 3.",
	"- 4.",
	"- 5.",
	"- 6.",
	"- 7.",
	"- 8.",
	"- 9.",
}

// 医院名关键字
var g_HospitalKeywordsList = []string{
	"医院",
	"医疗中心",
	"中医",
	"妇科",
	"门诊部",
	"美容诊所",
	"五官中心",
	"北京天院",
	"新医科",
	"眼科中心",
	"产科中心",
	"体检中心",
	"前列腺专科",
	"长征院",
	"长征医院",
	"心理院",
	"保健中心院",
}

// 公司名关键字
var g_CompanyNameKeywordsList = []string{
	"公司",
	"生态园",
	"研究所",
	"全资机构",
	"整形网",
	"不育网",
	"肿瘤网",
}

// 城市名列表
var g_CityMap = map[string]bool{
	"上海":   true,
	"北京":   true,
	"苏州":   true,
	"天津":   true,
	"广州":   true,
	"珠海":   true,
	"惠州":   true,
	"中山":   true,
	"汕头":   true,
	"东莞":   true,
	"江门":   true,
	"肇庆":   true,
	"佛山":   true,
	"深圳":   true,
	"昆明":   true,
	"玉溪":   true,
	"曲靖":   true,
	"重庆":   true,
	"成都":   true,
	"雅安":   true,
	"遵义":   true,
	"凉山":   true,
	"南充":   true,
	"乐山":   true,
	"福州":   true,
	"舟山":   true,
	"厦门":   true,
	"莆田":   true,
	"宁波":   true,
	"杭州":   true,
	"湖州":   true,
	"泉州":   true,
	"金华":   true,
	"嘉兴":   true,
	"台州":   true,
	"温州":   true,
	"龙岩":   true,
	"济南":   true,
	"潍坊":   true,
	"青岛":   true,
	"德州":   true,
	"威海":   true,
	"聊城":   true,
	"淄博":   true,
	"哈尔滨":  true,
	"长春":   true,
	"四平":   true,
	"延边":   true,
	"沈阳":   true,
	"大连":   true,
	"无锡":   true,
	"南京":   true,
	"张家港":  true,
	"泰州":   true,
	"盐城":   true,
	"宿迁":   true,
	"淮安":   true,
	"南通":   true,
	"武汉":   true,
	"荆州":   true,
	"黄冈":   true,
	"黄石":   true,
	"襄阳":   true,
	"乌海":   true,
	"呼和浩特": true,
	"贵阳":   true,
	"铜仁":   true,
	"安顺":   true,
	"毕节":   true,
	"长沙":   true,
	"郴州":   true,
	"湘潭":   true,
	"娄底":   true,
	"南昌":   true,
	"九江":   true,
	"吉安":   true,
	"萍乡":   true,
	"赣州":   true,
	"上饶":   true,
	"太原":   true,
	"临汾":   true,
	"阳泉":   true,
	"长治":   true,
	"大同":   true,
	"晋城":   true,
	"晋中":   true,
	"运城":   true,
	"西安":   true,
	"包头":   true,
	"蚌埠":   true,
	"亳州":   true,
	"芜湖":   true,
	"巢湖":   true,
	"淮北":   true,
	"合肥":   true,
	"安阳":   true,
	"郑州":   true,
	"许昌":   true,
	"廊坊":   true,
	"保定":   true,
	"唐山":   true,
	"洛阳":   true,
	"信阳":   true,
	"平顶山":  true,
	"漯河":   true,
	"石家庄":  true,
	"邯郸":   true,
	"拉萨":   true,
	"银川":   true,
	"兰州":   true,
	"桂林":   true,
	"柳州":   true,
	"伊犁":   true,
	"伊宁":   true,
	"乌鲁木齐": true,
	"海口":   true,
}
