// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 简化版LRU缓存实现
package cache

import (
	"sync"
)

// 缓存对象
type Cache struct {
	sync.Mutex
	keys    []string
	results map[string]interface{}
}

// 构造缓存
func New(capacity int) (p *Cache) {
	return &Cache{
		keys:    make([]string, 0, capacity),
		results: make(map[string]interface{}),
	}
}

func (p *Cache) Value(key string, defaultValue ...interface{}) interface{} {
	p.Lock()
	defer p.Unlock()

	if value, ok := p.results[key]; ok {
		return value
	}
	if len(defaultValue) > 0 {
		return defaultValue[0]
	}
	return nil
}

func (p *Cache) Get(key string) (value interface{}, ok bool) {
	p.Lock()
	defer p.Unlock()

	value, ok = p.results[key]
	return
}

func (p *Cache) Set(key string, value interface{}) {
	p.Lock()
	defer p.Unlock()

	// 已经存在
	if _, ok := p.results[key]; ok {
		// 将key移动到开头
		for idx, v := range p.keys {
			if v == key {
				for i := idx; i > 0; i-- {
					p.keys[i] = p.keys[i-1]
				}
				p.keys[0] = key
				break
			}
		}
		p.results[key] = value
		return
	}

	// 缓存已经满了, 需要清理一个空间
	if len(p.keys) == cap(p.keys) {
		delete(p.results, p.keys[len(p.keys)-1])
		p.keys = p.keys[:len(p.keys)-1]
	}

	// 添加新的数据key到开头(保持keys内存布局不变)
	p.keys = p.keys[:len(p.keys)+1]
	for i := len(p.keys) - 1; i > 0; i-- {
		p.keys[i] = p.keys[i-1]
	}
	p.keys[0] = key

	// 添加value
	p.results[key] = value
	return
}

// 清理缓存
func (p *Cache) Clear() {
	p.Lock()
	defer p.Unlock()

	p.keys = p.keys[:0]
	p.results = make(map[string]interface{})
}
