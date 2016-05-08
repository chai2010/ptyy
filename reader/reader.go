// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Json数据解析
package ptyy_reader

import (
	"encoding/json"
	"io"
	"os"
)

// 全部医院信息(map[医院名]信息)
type HospitalDB map[string]HospitalInfo

// 医院信息
type HospitalInfo struct {
	Name    string   // 名字
	City    string   // 城市
	Addr    []string // 地址
	WebSite []string // 网站
	Tel     []string // 电话
}

// 医院信息, 用于内部解析json
type _HospitalInfo struct {
	Addr    []string `json:"地址"`
	WebSite []string `json:"网址"`
	Tel     []string `json:"电话"`
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
		city := keyword
		for name, info := range hospitals {
			v, _ := db[name]
			v.Name = name
			v.City = city
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
