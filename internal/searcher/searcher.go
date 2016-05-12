// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 简易检索器
package searcher

import (
	"regexp"
	"strings"
)

type Searcher struct {
	db []map[string]string // 开头的优先权更高
}

func New(db ...map[string]string) *Searcher {
	return &Searcher{db: db}
}

// 关键字查询: 返回列表
func (p *Searcher) SearchByKeyAll(query string) []string {
	var ss []string
	for s := range p.SearchByKey(query) {
		ss = append(ss, s)
	}
	return ss
}

// 关键字查询: 返回管道
func (p *Searcher) SearchByKey(query string) <-chan string {
	// 构造一个带缓存的管道, 用于返回结果
	ch := make(chan string, 32)

	// 在另一个goroutine中执行查询操作
	go func() {
		defer close(ch)

		// 空值认为是没有匹配结果
		if query == "" {
			return
		}

		// 构建一个map，避免重复的查询
		foundMap := make(map[string]bool)

		// 根据优先级顺序模糊查询
		// 因为map无序, 最终的结果可能是无序的
		for _, db := range p.db {
			// 前缀匹配优先级较高
			for k, v := range db {
				if strings.HasPrefix(k, query) {
					if !foundMap[v] {
						foundMap[v] = true
						ch <- v
					}
				}
			}

			// 根据包含关系查询
			for k, v := range db {
				if !foundMap[v] {
					if strings.Contains(k, query) {
						foundMap[v] = true
						ch <- v
					}
				}
			}
		}
	}()

	return ch
}

// 正则查询: 返回列表
func (p *Searcher) SearchByRegexpAll(query *regexp.Regexp) []string {
	var ss []string
	for s := range p.SearchByRegexp(query) {
		ss = append(ss, s)
	}
	return ss
}

// 正则查询: 返回管道
func (p *Searcher) SearchByRegexp(query *regexp.Regexp) <-chan string {
	// 构造一个带缓存的管道, 用于返回结果
	ch := make(chan string, 32)

	// 在另一个goroutine中执行查询操作
	go func() {
		defer close(ch)

		// 构建一个map，避免重复的查询
		foundMap := make(map[string]bool)

		// 根据优先级顺序模糊查询
		// 因为map无序, 最终的结果可能是无序的
		for _, db := range p.db {
			// 前缀匹配优先级较高
			for k, v := range db {
				if !foundMap[v] {
					if reFindFirstIndex(k, query) == 0 {
						foundMap[v] = true
						ch <- v
					}
				}
			}

			// 根据包含关系查询
			for k, v := range db {
				if !foundMap[v] {
					if reFindFirstIndex(k, query) > 0 {
						foundMap[v] = true
						ch <- v
					}
				}
			}
		}
	}()

	return ch
}

// 正则第一个匹配的位置
func reFindFirstIndex(value string, query *regexp.Regexp) int {
	if idx := query.FindStringIndex(value); len(idx) > 0 {
		return idx[0]
	}
	return -1
}
