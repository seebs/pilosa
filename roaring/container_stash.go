// Copyright 2017 Pilosa Corp.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package roaring

import (
	"reflect"
	"runtime"
	"unsafe"
)

const (
	stashedArraySize = 5
	stashedRunSize   = (stashedArraySize / 2)
)

// Container represents a Container for uint16 integers.
//
// These are used for storing the low bits of numbers in larger sets of uint64.
// The high bits are stored in a Container's key which is tracked by a separate
// data structure. Integers in a Container can be encoded in one of three ways -
// the encoding used is usually whichever is most compact, though any Container
// type should be able to encode any set of integers safely. For containers with
// less than 4,096 values, an array is often used. Containers with long runs of
// integers would use run length encoding, and more random data usually uses
// bitmap encoding.
type Container struct {
	pointer  *uint16                  // the data pointer
	len, cap int32                    // length and cap
	n        int32                    // number of integers in container
	shared   bool                     // storage may be shared, copy-on-write
	typ      byte                     // array, bitmap, or run
	data     [stashedArraySize]uint16 // immediate data for small arrays or runs
}

// NewContainer returns a new instance of container. This trivial function
// may later become more interesting.
func NewContainer() *Container {
	statsHit("NewContainer")
	c := &Container{typ: containerArray, len: 0, cap: stashedArraySize}
	c.pointer = (*uint16)(unsafe.Pointer(&c.data[0]))
	return c
}

// NewContainerBitmap makes a bitmap container using the provided bitmap, or
// an empty one if provided bitmap is nil. If the provided bitmap is too short,
// it will be padded.
func NewContainerBitmap(n int32, bitmap []uint64) *Container {
	c := &Container{typ: containerBitmap, n: n, shared: bitmap != nil}
	if bitmap == nil {
		bitmap = make([]uint64, bitmapN)
	}
	// pad to required length
	if len(bitmap) < bitmapN {
		// No longer using the source data, so we're safe.
		c.shared = false
		bm2 := make([]uint64, bitmapN)
		copy(bm2, bitmap)
		bitmap = bm2
	}
	c.setBitmap(bitmap)
	return c
}

// NewContainerArray returns an array using the provided set of values. It's
// okay if the slice is nil; that's a length of zero.
func NewContainerArray(set []uint16) *Container {
	c := &Container{typ: containerArray, n: int32(len(set)), shared: set != nil}
	c.setArray(set)
	return c
}

// NewContainerRun creates a new run array using a provided (possibly nil)
// slice of intervals.
func NewContainerRun(set []interval16) *Container {
	c := &Container{typ: containerRun, shared: set != nil}
	c.setRuns(set)
	for _, run := range set {
		c.n += int32(run.last-run.start) + 1
	}
	return c
}

// N returns the internal n field.
func (c *Container) N() int32 {
	return c.n
}

// array yields the data viewed as a slice of uint16 values.
func (c *Container) array() []uint16 {
	if roaringParanoia {
		if c.typ != containerArray {
			panic("attempt to read non-array's array")
		}
	}
	return *(*[]uint16)(unsafe.Pointer(&reflect.SliceHeader{Data: uintptr(unsafe.Pointer(c.pointer)), Len: int(c.len), Cap: int(c.cap)}))
}

// setArray stores a set of uint16s as data.
func (c *Container) setArray(array []uint16) {
	if roaringParanoia {
		if c.typ != containerArray {
			panic("attempt to write non-array's array")
		}
	}
	// no array: start with our default 5-value array
	if array == nil {
		c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(&c.data[0])), 0, stashedArraySize
		return
	}
	h := (*reflect.SliceHeader)(unsafe.Pointer(&array))
	if h.Data == uintptr(unsafe.Pointer(c.pointer)) {
		// nothing to do but update length
		c.len = int32(h.Len)
		return
	}
	// array we can fit in data store:
	if len(array) <= stashedArraySize {
		copy(c.data[:stashedArraySize], array)
		c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(&c.data[0])), int32(len(array)), stashedArraySize
		// note: if this was previously actually-shared, we shouldn't have gotten
		// here. but it could have been set by NewContainerArray guessing that
		// a provided input array might be shared, in which case we don't care.
		c.shared = false
		return
	}
	c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(h.Data)), int32(h.Len), int32(h.Cap)
	runtime.KeepAlive(&array)
}

// bitmap yields the data viewed as a slice of uint64s holding bits.
func (c *Container) bitmap() []uint64 {
	if roaringParanoia {
		if c.typ != containerBitmap {
			panic("attempt to read non-bitmap's bitmap")
		}
	}
	return *(*[]uint64)(unsafe.Pointer(&reflect.SliceHeader{Data: uintptr(unsafe.Pointer(c.pointer)), Len: int(c.len), Cap: int(c.cap)}))
}

// setBitmap stores a set of uint64s as data.
func (c *Container) setBitmap(bitmap []uint64) {
	if roaringParanoia {
		if c.typ != containerBitmap {
			panic("attempt to write non-bitmap's bitmap")
		}
	}
	h := (*reflect.SliceHeader)(unsafe.Pointer(&bitmap))
	c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(h.Data)), int32(h.Len), int32(h.Cap)
	runtime.KeepAlive(&bitmap)
}

// runs yields the data viewed as a slice of intervals.
func (c *Container) runs() []interval16 {
	if roaringParanoia {
		if c.typ != containerRun {
			panic("attempt to read non-run's runs")
		}
	}
	return *(*[]interval16)(unsafe.Pointer(&reflect.SliceHeader{Data: uintptr(unsafe.Pointer(c.pointer)), Len: int(c.len), Cap: int(c.cap)}))
}

// setRuns stores a set of intervals as data.
func (c *Container) setRuns(runs []interval16) {
	if roaringParanoia {
		if c.typ != containerRun {
			panic("attempt to write non-run's runs")
		}
	}
	// no array: start with our default 2-value array
	if runs == nil {
		c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(&c.data[0])), 0, stashedRunSize
		return
	}
	h := (*reflect.SliceHeader)(unsafe.Pointer(&runs))
	if h.Data == uintptr(unsafe.Pointer(c.pointer)) {
		// nothing to do but update length
		c.len = int32(h.Len)
		return
	}

	// array we can fit in data store:
	if len(runs) <= stashedRunSize {
		newRuns := *(*[]interval16)(unsafe.Pointer(&reflect.SliceHeader{Data: uintptr(unsafe.Pointer(&c.data[0])), Len: stashedRunSize, Cap: stashedRunSize}))
		copy(newRuns, runs)
		c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(&c.data[0])), int32(len(runs)), stashedRunSize
		c.shared = false
		return
	}
	c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(h.Data)), int32(h.Len), int32(h.Cap)
	runtime.KeepAlive(&runs)
}

// Update updates the container
func (c *Container) Update(typ byte, n int32, shared bool) {
	c.typ = typ
	c.n = n
	c.shared = shared
	// we don't know that any existing slice is usable, so let's ditch it
	switch c.typ {
	case containerArray:
		c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(&c.data[0])), int32(0), stashedArraySize
	case containerRun:
		c.pointer, c.len, c.cap = (*uint16)(unsafe.Pointer(&c.data[0])), 0, stashedRunSize
	default:
		c.pointer, c.len, c.cap = nil, 0, 0
	}
}

// isArray returns true if the container is an array container.
func (c *Container) isArray() bool {
	return c.typ == containerArray
}

// isBitmap returns true if the container is a bitmap container.
func (c *Container) isBitmap() bool {
	return c.typ == containerBitmap
}

// isRun returns true if the container is a run-length-encoded container.
func (c *Container) isRun() bool {
	return c.typ == containerRun
}

// Mapped indicates whether the container might be mmapped, but is possibly
// no longer sane. Consider it deprecated.
func (c *Container) Mapped() bool {
	return c.shared
}

// unmapArray ensures that the container is not using mmapped storage.
func (c *Container) unmapArray() (oldC, newC *Container) {
	if !c.shared {
		return c, nil
	}
	newC = &Container{typ: containerArray, len: c.len, cap: c.cap, n: c.n}
	array := c.array()
	tmp := make([]uint16, c.len)
	copy(tmp, array)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&tmp))
	newC.pointer, newC.cap = (*uint16)(unsafe.Pointer(h.Data)), int32(h.Cap)
	runtime.KeepAlive(&tmp)
	newC.shared = false
	return newC, newC
}

// unmapBitmap ensures that the container is not using mmapped storage.
func (c *Container) unmapBitmap() (oldC, newC *Container) {
	if !c.shared {
		return c, nil
	}
	newC = &Container{typ: containerBitmap, len: c.len, n: c.n}
	bitmap := c.bitmap()
	tmp := make([]uint64, c.len)
	copy(tmp, bitmap)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&tmp))
	newC.pointer, newC.cap = (*uint16)(unsafe.Pointer(h.Data)), int32(h.Cap)
	runtime.KeepAlive(&tmp)
	newC.shared = false
	return newC, newC
}

// unmapRun ensures that the container is not using mmapped storage.
func (c *Container) unmapRun() (oldC, newC *Container) {
	if !c.shared {
		return c, nil
	}
	newC = &Container{typ: containerRun, len: c.len, n: c.n}
	runs := c.runs()
	tmp := make([]interval16, c.len)
	copy(tmp, runs)
	h := (*reflect.SliceHeader)(unsafe.Pointer(&tmp))
	newC.pointer, newC.cap = (*uint16)(unsafe.Pointer(h.Data)), int32(h.Cap)
	runtime.KeepAlive(&tmp)
	newC.shared = false
	return newC, newC
}
