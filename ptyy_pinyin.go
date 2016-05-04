// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy

// 名字和城市的拼音表格
// 1. 全部拼音组合(忽略音调)
// 2. 首字母组合
var _NamePinyinLongMap = make(map[string]string)
var _NamePinyinShortMap = make(map[string]string)

// 初始化拼音表
func init() {
	for _, info := range All {
		if _NamePinyinLongMap[info.Name] == "" || _NamePinyinShortMap[info.Name] == "" {
			pyLong, pyShort := makePinyin(info.Name)
			_NamePinyinLongMap[info.Name] = pyLong
			_NamePinyinShortMap[info.Name] = pyShort
		}
		if _NamePinyinLongMap[info.City] == "" || _NamePinyinShortMap[info.City] == "" {
			pyLong, pyShort := makePinyin(info.City)
			_NamePinyinLongMap[info.City] = pyLong
			_NamePinyinShortMap[info.City] = pyShort
		}
	}
}

// 生成单词的拼音列表
func makePinyin(name string) (pyLong, pyShort string) {
	for _, r := range name {
		py, _ := _RunePinyinTable[r]
		if py == "" {
			return "", ""
		}
		py = goodRunePinyin(py)
		pyLong += py      // 完整的拼音, 不含声调
		pyShort += py[:1] // 第一个拼音字母
	}
	return
}

// 规范汉子的拼音(去掉声调)
func goodRunePinyin(py string) string {
	lastRune := py[len(py)-1]
	if lastRune >= '0' && lastRune <= '9' {
		py = py[:len(py)-1]
	}
	return py
}
