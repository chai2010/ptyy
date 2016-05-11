// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"sync"
)

// 最大的缓存数目
const maxCachedSize = 16

// 查询缓冲
var cache struct {
	sync.Mutex
	keys    []string
	results map[string]string
}

func init() {
	cache.Lock()
	defer cache.Unlock()

	cache.keys = make([]string, 0, maxCachedSize)
	cache.results = make(map[string]string)
}

// 从缓存查询
func getFromCache(key string) (value string, ok bool) {
	cache.Lock()
	defer cache.Unlock()

	value, ok = cache.results[key]
	return
}

// 更新缓存
func setToCache(key, value string) {
	cache.Lock()
	defer cache.Unlock()

	// 已经存在
	if _, ok := cache.results[key]; ok {
		// 将key移动到开头
		for idx, v := range cache.keys {
			if v == key {
				for i := idx; i > 0; i-- {
					cache.keys[i] = cache.keys[i-1]
				}
				cache.keys[0] = key
				break
			}
		}
		cache.results[key] = value
		return
	}

	// 缓存已经满了, 需要清理一个空间
	if len(cache.keys) == cap(cache.keys) {
		delete(cache.results, cache.keys[len(cache.keys)-1])
		cache.keys = cache.keys[:len(cache.keys)-1]
	}

	// 添加新的数据
	cache.keys = append([]string{key}, cache.keys...)
	cache.results[key] = value
	return
}

// 清理缓存
func clearCache() {
	cache.Lock()
	defer cache.Unlock()

	cache.keys = cache.keys[:0]
	cache.results = make(map[string]string)
}
