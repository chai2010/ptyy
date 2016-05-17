// Copyright 2016 <chaishushan{AT}gmail.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// 在其它语言中通过Id访问Go语言对象.
package goobj

import (
	"fmt"
	_ "runtime"
	"sync"
)

// Go对象id, 可以传递给其它语言
type Id int32

// 构造新的对象id
func NewId(new func() interface{}, deleter func(x interface{})) Id {
	return 0
}

// 是否为空id
func (id Id) IsNil() bool {
	return id == 0
}

// 获取原始对象(所有权属于Id)
func (id Id) Get() interface{} {
	return nil
}

// 释放对象
// 释放后将id更新为0
func (id *Id) Delete() {
	// TODO
}

// 对象信息
// 含有引用计数和析构函数
type countedObj struct {
	deleter func(x interface{})
	obj     interface{}
	cnt     int32
}

// 持有Go对象, 避免被GC回收, id可以传递给其它语言
var refs struct {
	sync.Mutex
	next int32 // next reference number to use for Go object, always negative
	refs map[interface{}]int32
	objs map[int32]countedObj
}

func init() {
	refs.Lock()
	refs.next = -24 // Go objects get negative reference numbers. Arbitrary starting point.
	refs.refs = make(map[interface{}]int32)
	refs.objs = make(map[int32]countedObj)
	refs.Unlock()
}

// also known to <goobjc.h>
const _NullRefNum = 41

// A Ref represents a Go object passed across the language boundary.
type _Ref struct {
	Bind_Num int32
}

type proxy interface {
	// Use a strange name and hope that user code does not implement it
	Bind_proxy_refnum__() int32
}

// ToRefNum increments the reference count for an object and
// returns its refnum.
func _ToRefNum(obj interface{}) int32 {
	// We don't track foreign objects, so if obj is a proxy
	// return its refnum.
	if r, ok := obj.(proxy); ok {
		refnum := r.Bind_proxy_refnum__()
		if refnum <= 0 {
			panic(fmt.Errorf("seq: proxy contained invalid Go refnum: %d", refnum))
		}
		return refnum
	}
	refs.Lock()
	num := refs.refs[obj]
	if num != 0 {
		s := refs.objs[num]
		refs.objs[num] = countedObj{obj: s.obj, cnt: s.cnt + 1}
	} else {
		num = refs.next
		refs.next--
		if refs.next > 0 {
			panic("refs.next underflow")
		}
		refs.refs[obj] = num
		refs.objs[num] = countedObj{obj: obj, cnt: 1}
	}
	refs.Unlock()

	return int32(num)
}

// FromRefNum returns the Ref for a refnum. If the refnum specifies a
// foreign object, a finalizer is set to track its lifetime.
func _FromRefNum(num int32) *_Ref {
	if num == _NullRefNum {
		return nil
	}
	ref := &_Ref{num}
	if num > 0 {
		// This is a foreign object reference.
		// Track its lifetime with a finalizer.
		//runtime.SetFinalizer(ref, FinalizeRef)
	}

	return ref
}

// Bind_IncNum increments the foreign reference count and
// return the refnum.
func (r *_Ref) Bind_IncNum() int32 {
	refnum := r.Bind_Num
	//IncForeignRef(refnum)
	return refnum
}

// Get returns the underlying object.
func (r *_Ref) Get() interface{} {
	refnum := r.Bind_Num
	refs.Lock()
	o, ok := refs.objs[refnum]
	refs.Unlock()
	if !ok {
		panic(fmt.Sprintf("unknown ref %d", refnum))
	}
	// This is a Go reference and its refnum was incremented
	// before crossing the language barrier.
	_Delete(refnum)
	return o.obj
}

// Delete decrements the reference count and removes the pinned object
// from the object map when the reference count becomes zero.
func _Delete(num int32) {
	refs.Lock()
	defer refs.Unlock()
	o, ok := refs.objs[num]
	if !ok {
		panic(fmt.Sprintf("goobj.Delete unknown refnum: %d", num))
	}
	if o.cnt <= 1 {
		delete(refs.objs, num)
		delete(refs.refs, o.obj)
	} else {
		refs.objs[num] = countedObj{obj: o.obj, cnt: o.cnt - 1}
	}
}
