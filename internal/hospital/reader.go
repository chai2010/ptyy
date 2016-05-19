// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package hospital

import (
	"encoding/json"
	"io"
	"os"
)

// 全部医院信息(map[医院名]信息)
type HospitalDB map[string]HospitalInfo

// 医院信息
type HospitalInfo struct {
	Name      string   // 名字
	City      string   // 城市
	Keywords  []string // 关键字
	Addr      []string // 地址
	WebSite   []string // 网站
	WeiXin    []string // 微信
	WeiXinPub []string // 微信公众号
	Tel       []string // 电话
	Comment   []string // 注释
	Feedback  []string // 用户反馈
}

// 医院信息, 用于内部解析json
type _HospitalInfo struct {
	Addr      []string `json:"地址"`
	WebSite   []string `json:"网址"`
	WeiXin    []string `json:"微信"`
	WeiXinPub []string `json:"微信公众号"`
	Tel       []string `json:"电话"`
	Comment   []string `json:"注释"`
	Feedback  []string `json:"用户反馈"`
}

// 读取Json文件
func LoadJsonFile(filename string) (HospitalDB, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	return readJsonFrom(f)
}

// 从输入流读取
func ReadJsonFrom(r io.Reader) (HospitalDB, error) {
	return readJsonFrom(r)
}

// 从输入流读取
func readJsonFrom(r io.Reader) (HospitalDB, error) {
	rawDb, err := parseHospitalJsonDB(r)
	if err != nil {
		return nil, err
	}

	db := make(HospitalDB)
	for keyword, hospitals := range rawDb {
		if keyword == "网站" {
			continue
		}
		for name, info := range hospitals {
			v, _ := db[name]
			v.Name = name
			if !isNotCity(keyword) {
				v.City = keyword
			}
			v.Keywords = append(v.Keywords, keyword)
			v.Addr = append(v.Addr, info.Addr...)
			v.WebSite = append(v.WebSite, info.WebSite...)
			v.Tel = append(v.Tel, info.Tel...)
			db[name] = v
		}
	}

	return db, nil
}

// 解析数据
func parseHospitalJsonDB(r io.Reader) (map[string]map[string]_HospitalInfo, error) {
	result := make(map[string]map[string]_HospitalInfo)
	if err := json.NewDecoder(r).Decode(&result); err != nil {
		return nil, err
	}
	return result, nil
}

func isNotCity(name string) bool {
	for _, s := range g_notCityList {
		if s == name {
			return true
		}
	}
	return false
}

var g_notCityList = []string{
	"网站",
	"林氏家族",
	"黄氏家族",
	"陈氏家族",
	"妇科/产科",
	"三甲医院外包科室名单如下",
	"整形科",
	"不孕症",
}
