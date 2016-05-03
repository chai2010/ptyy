// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

//go:generate go run gen_helper.go

package ptyy

import (
	"regexp"
	"strings"
)

// 医院信息
type HospitalInfo struct {
	Name    string   // 名称
	City    string   // 城市
	Owner   []string // 投资者
	Comment []string // 注释
}

var (
	All             []HospitalInfo // 医院列表
	AllHospitalList []HospitalInfo // 医院列表
	AllCityList     []string       // 城市列表
	AllOwnerList    []string       // 单位列表
)

// 查询列表
func Search(query string) []HospitalInfo {
	// 规范化: 删除前后空白
	query = strings.TrimSpace(query)

	// 如果为空的话, 返回全部
	if query == "" {
		return All
	}

	// 根据关键字查询
	if results := searchByKeywords(query); len(results) > 0 {
		return results
	}

	// 如果没有匹配的, 则尝试正则查询
	if re, err := regexp.Compile(query); err == nil {
		if results := searchByRegexp(re); len(results) > 0 {
			return results
		}
	}

	// 没有匹配
	return nil
}

// 根据关键字查询
// TODO: 以后可扩展为多个关键字, 采用非字符字符分隔
func searchByKeywords(key string) []HospitalInfo {
	return nil
}

// 根据正则表达式查询
func searchByRegexp(re *regexp.Regexp) []HospitalInfo {
	// TODO
	return nil
}
