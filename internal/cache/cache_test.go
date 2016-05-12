// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cache

import (
	"strconv"
	"testing"
)

const (
	tCacheSize = 1000
)

type TCache struct {
	*Cache
}

func tNewTCache(capacity int) *TCache {
	return &TCache{Cache: New(capacity)}
}

func (p *TCache) Insert(key, value int) {
	p.Cache.Set(strconv.Itoa(key), value)
}

func (p *TCache) Lookup(key int) int {
	if v, ok := p.Cache.Get(strconv.Itoa(key)); ok {
		return v.(int)
	}
	return -1
}

func TestCache_hitAndMiss(t *testing.T) {
	c := tNewTCache(tCacheSize)

	tAssertEQ(t, -1, c.Lookup(100))

	c.Insert(100, 101)
	tAssertEQ(t, 101, c.Lookup(100))
	tAssertEQ(t, -1, c.Lookup(200))
	tAssertEQ(t, -1, c.Lookup(200))

	c.Insert(200, 201)
	tAssertEQ(t, 101, c.Lookup(100))
	tAssertEQ(t, 201, c.Lookup(200))
	tAssertEQ(t, -1, c.Lookup(300))

	c.Insert(100, 102)
	tAssertEQ(t, 102, c.Lookup(100))
	tAssertEQ(t, 201, c.Lookup(200))
	tAssertEQ(t, -1, c.Lookup(300))
}

func TestCache_evictionPolicy(t *testing.T) {
	c := tNewTCache(tCacheSize)

	c.Insert(100, 101)

	// Frequently used entry must be kept around
	for i := 0; i < tCacheSize+100; i++ {
		if i < tCacheSize {
			tAssertEQ(t, 101, c.Lookup(100))
		} else {
			tAssertEQ(t, -1, c.Lookup(100))
		}

		c.Insert(1000+i, 2000+i)
		tAssertEQ(t, 2000+i, c.Lookup(1000+i))
	}
	tAssertEQ(t, -1, c.Lookup(100))
	tAssertEQ(t, -1, c.Lookup(200))
}

func tAtoi(s string, defaultV int) int {
	if v, err := strconv.Atoi(s); err == nil {
		return v
	}
	return defaultV
}
