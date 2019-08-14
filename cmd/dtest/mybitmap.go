package main

import (
	"fmt"
	"io"

	"github.com/pilosa/pilosa/data"
)

type MyBitmap struct {
	_ int
}

func (b *MyBitmap) Any() bool {
	return false
}

func (b *MyBitmap) AnyRange(first, last uint64) bool {
	return false
}

func (b *MyBitmap) Count() uint64 {
	return 0
}

func (b *MyBitmap) CountRange(first, last uint64) uint64 {
	return 0
}

func (b *MyBitmap) Slice() ([]uint64, bool) {
	return nil, false
}

func (b *MyBitmap) SliceRange(first, last uint64) ([]uint64, bool) {
	return nil, false
}

func (b *MyBitmap) GetContainer(uint64) data.Container {
	return nil
}

func (b *MyBitmap) OffsetRange(offset, first, last uint64) data.Bitmap {
	return nil
}

func (b *MyBitmap) ViewContainers(fn func(key uint64, c data.Container) (done bool, err error)) error {
	return nil
}

func (b *MyBitmap) ViewContainersRange(first, last uint64, fn func(key uint64, c data.Container) (done bool, err error)) error {
	return nil
}

func (b *MyBitmap) ExportRoaring() []byte {
	return nil
}

func (b *MyBitmap) WriteRoaringTo(io.Writer) error {
	return nil
}

func (b *MyBitmap) UnionViewBitmap(other data.ReadOnlyBitmap) (bool, int64, data.ReadOnlyBitmap) {
	fmt.Printf("self %p, other %p\n", b, other)
	return false, 0, nil
}

func (b *MyBitmap) ExportRoaringViewGivesBytes() []byte {
	return nil
}
