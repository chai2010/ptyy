// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ptyy

// 是否为数字
func isDigitRune(c rune) bool {
	if c >= '0' && c <= '9' {
		return true
	}
	for _, v := range "零一二三四五六七八九十" {
		if c == v {
			return true
		}
	}
	return false
}

// 是否为数字字符串
func isDigitString(s string) bool {
	for _, c := range s {
		if !isDigitRune(c) {
			return false
		}
	}
	return true
}

// 规范化数字
// 返回: 阿拉伯数字, 中文数字
// 3七 -> 37, 三七
// 10 -> 10, 一零
func adjustDigitString(s string) (s0, zh string, ok bool) {
	runes := make([]rune, 0, len(s))

Loop:
	for _, c := range s {
		// 阿拉伯数字
		if c >= '0' && c <= '9' {
			runes = append(runes, c)
			continue Loop
		}

		// 忽略中文的 ［十/拾］
		if c == '十' || c == '拾' || c == '百' || c == '千' {
			continue Loop // 忽略
		}

		// 将中文数字转为阿拉伯数字
		for i, v := range []rune("零一二三四五六七八九") {
			if c == v {
				runes = append(runes, rune('0'+i))
				continue Loop
			}
		}
		for i, v := range []rune("零壹贰叁肆伍陆柒捌玖") {
			if c == v {
				runes = append(runes, rune('0'+i))
				continue Loop
			}
		}

		// 不是数字, 失败
		return "", "", false
	}

	runesZh := make([]rune, len(runes))
	runesZhTable := []rune("零一二三四五六七八九")

	for i, c := range runes {
		runesZh[i] = runesZhTable[int(c-'0')]
	}

	s0 = string(runes)
	zh = string(runesZh)
	ok = true
	return
}
