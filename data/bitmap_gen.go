package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Bitmap (see gen/main.go). These are expressed
// as method signatures -- the Bitmap they operate on is an implicit
// receiver not shown in the signature.

import (
	"io"
	"reflect"
)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionBitmap interface {
	BitmapOpType() OpType
}

type OpBitmapView func() (bool, uint64, ReadOnlyBitmap)

func (OpBitmapView) BitmapOpType() OpType { return OpTypeView }

func LookupOpBitmapView(target ReadOnlyBitmap, name string) OpBitmapView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, uint64, ReadOnlyBitmap))
		return OpBitmapView(fn)
	}
	return nil
}

type OpBitmapViewGivesBool func() bool

func (OpBitmapViewGivesBool) BitmapOpType() OpType { return OpTypeViewGivesBool }

func LookupOpBitmapViewGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func() bool)
		return OpBitmapViewGivesBool(fn)
	}
	return nil
}

// Any performs a default BitmapViewGivesBool on a Bitmap.
type interfaceHasAny interface {
	AnyViewGivesBool() bool
}

func Any(target ReadOnlyBitmap) bool {
	if target, ok := target.(interfaceHasAny); ok {
		return target.AnyViewGivesBool()
	}
	return genericAny(target)
}

type OpBitmapViewGivesBit func() uint64

func (OpBitmapViewGivesBit) BitmapOpType() OpType { return OpTypeViewGivesBit }

func LookupOpBitmapViewGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func() uint64)
		return OpBitmapViewGivesBit(fn)
	}
	return nil
}

// Count performs a default BitmapViewGivesBit on a Bitmap.
type interfaceHasCount interface {
	CountViewGivesBit() uint64
}

func Count(target ReadOnlyBitmap) uint64 {
	if target, ok := target.(interfaceHasCount); ok {
		return target.CountViewGivesBit()
	}
	return genericCount(target)
}

type OpBitmapViewRangeGivesBool func(uint64, uint64) bool

func (OpBitmapViewRangeGivesBool) BitmapOpType() OpType { return OpTypeViewRangeGivesBool }

func LookupOpBitmapViewRangeGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) bool)
		return OpBitmapViewRangeGivesBool(fn)
	}
	return nil
}

// AnyRange performs a default BitmapViewRangeGivesBool on a Bitmap.
type interfaceHasAnyRange interface {
	AnyRangeViewRangeGivesBool(uint64, uint64) bool
}

func AnyRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) bool {
	if target, ok := target.(interfaceHasAnyRange); ok {
		return target.AnyRangeViewRangeGivesBool(in1, in2)
	}
	return genericAnyRange(target, in1, in2)
}

type OpBitmapViewRangeGivesBit func(uint64, uint64) uint64

func (OpBitmapViewRangeGivesBit) BitmapOpType() OpType { return OpTypeViewRangeGivesBit }

func LookupOpBitmapViewRangeGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) uint64)
		return OpBitmapViewRangeGivesBit(fn)
	}
	return nil
}

// CountRange performs a default BitmapViewRangeGivesBit on a Bitmap.
type interfaceHasCountRange interface {
	CountRangeViewRangeGivesBit(uint64, uint64) uint64
}

func CountRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) uint64 {
	if target, ok := target.(interfaceHasCountRange); ok {
		return target.CountRangeViewRangeGivesBit(in1, in2)
	}
	return genericCountRange(target, in1, in2)
}

type OpBitmapViewRangeGivesBitmap func(uint64, uint64) ReadOnlyBitmap

func (OpBitmapViewRangeGivesBitmap) BitmapOpType() OpType { return OpTypeViewRangeGivesOther }

func LookupOpBitmapViewRangeGivesBitmap(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) ReadOnlyBitmap)
		return OpBitmapViewRangeGivesBitmap(fn)
	}
	return nil
}

// OffsetRange performs a default BitmapViewRangeGivesBitmap on a Bitmap.
type interfaceHasOffsetRange interface {
	OffsetRangeViewRangeGivesBitmap(uint64, uint64) ReadOnlyBitmap
}

func OffsetRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) ReadOnlyBitmap {
	if target, ok := target.(interfaceHasOffsetRange); ok {
		return target.OffsetRangeViewRangeGivesBitmap(in1, in2)
	}
	return genericOffsetRange(target, in1, in2)
}

type OpBitmapViewRangeGivesBitsBool func(uint64, uint64) ([]uint64, bool)

func (OpBitmapViewRangeGivesBitsBool) BitmapOpType() OpType { return OpTypeViewRangeGivesBitsBool }

func LookupOpBitmapViewRangeGivesBitsBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitsBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitsBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) ([]uint64, bool))
		return OpBitmapViewRangeGivesBitsBool(fn)
	}
	return nil
}

// SliceRange performs a default BitmapViewRangeGivesBitsBool on a Bitmap.
type interfaceHasSliceRange interface {
	SliceRangeViewRangeGivesBitsBool(uint64, uint64) ([]uint64, bool)
}

func SliceRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) ([]uint64, bool) {
	if target, ok := target.(interfaceHasSliceRange); ok {
		return target.SliceRangeViewRangeGivesBitsBool(in1, in2)
	}
	return genericSliceRange(target, in1, in2)
}

type OpBitmapUpdate func() (bool, uint64, Bitmap)

func (OpBitmapUpdate) BitmapOpType() OpType { return OpTypeUpdate }

func LookupOpBitmapUpdate(target ReadOnlyBitmap, name string) OpBitmapUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, uint64, Bitmap))
		return OpBitmapUpdate(fn)
	}
	return nil
}

type OpBitmapViewRange func(uint64, uint64) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewRange) BitmapOpType() OpType { return OpTypeViewRange }

func LookupOpBitmapViewRange(target ReadOnlyBitmap, name string) OpBitmapViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewRange(fn)
	}
	return nil
}

type OpBitmapViewBit func(uint64) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBit) BitmapOpType() OpType { return OpTypeViewBit }

func LookupOpBitmapViewBit(target ReadOnlyBitmap, name string) OpBitmapViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBit(fn)
	}
	return nil
}

type OpBitmapUpdateBit func(uint64) (bool, uint64, Bitmap)

func (OpBitmapUpdateBit) BitmapOpType() OpType { return OpTypeUpdateBit }

func LookupOpBitmapUpdateBit(target ReadOnlyBitmap, name string) OpBitmapUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, uint64, Bitmap))
		return OpBitmapUpdateBit(fn)
	}
	return nil
}

// Add performs a default BitmapUpdateBit on a Bitmap.
type interfaceHasAdd interface {
	AddUpdateBit(uint64) (bool, uint64, Bitmap)
}

func Add(target Bitmap, in1 uint64) (bool, uint64, Bitmap) {
	if target, ok := target.(interfaceHasAdd); ok {
		return target.AddUpdateBit(in1)
	}
	return genericAdd(target, in1)
}

// Remove performs a default BitmapUpdateBit on a Bitmap.
type interfaceHasRemove interface {
	RemoveUpdateBit(uint64) (bool, uint64, Bitmap)
}

func Remove(target Bitmap, in1 uint64) (bool, uint64, Bitmap) {
	if target, ok := target.(interfaceHasRemove); ok {
		return target.RemoveUpdateBit(in1)
	}
	return genericRemove(target, in1)
}

type OpBitmapViewBitmap func(ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBitmap) BitmapOpType() OpType { return OpTypeViewOther }

func LookupOpBitmapViewBitmap(target ReadOnlyBitmap, name string) OpBitmapViewBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBitmap(fn)
	}
	return nil
}

type OpBitmapUpdateBitmap func(ReadOnlyBitmap) (bool, uint64, Bitmap)

func (OpBitmapUpdateBitmap) BitmapOpType() OpType { return OpTypeUpdateOther }

func LookupOpBitmapUpdateBitmap(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, uint64, Bitmap))
		return OpBitmapUpdateBitmap(fn)
	}
	return nil
}

type OpBitmapViewBits func([]uint64) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBits) BitmapOpType() OpType { return OpTypeViewBits }

func LookupOpBitmapViewBits(target ReadOnlyBitmap, name string) OpBitmapViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBits(fn)
	}
	return nil
}

type OpBitmapUpdateBits func([]uint64) (bool, uint64, Bitmap)

func (OpBitmapUpdateBits) BitmapOpType() OpType { return OpTypeUpdateBits }

func LookupOpBitmapUpdateBits(target ReadOnlyBitmap, name string) OpBitmapUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, uint64, Bitmap))
		return OpBitmapUpdateBits(fn)
	}
	return nil
}

type OpBitmapViewBitmaps func([]ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBitmaps) BitmapOpType() OpType { return OpTypeViewOthers }

func LookupOpBitmapViewBitmaps(target ReadOnlyBitmap, name string) OpBitmapViewBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBitmaps(fn)
	}
	return nil
}

type OpBitmapUpdateBitmaps func([]ReadOnlyBitmap) (bool, uint64, Bitmap)

func (OpBitmapUpdateBitmaps) BitmapOpType() OpType { return OpTypeUpdateOthers }

func LookupOpBitmapUpdateBitmaps(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, uint64, Bitmap))
		return OpBitmapUpdateBitmaps(fn)
	}
	return nil
}

type OpBitmapUpdateBytes func([]byte) (bool, uint64, Bitmap)

func (OpBitmapUpdateBytes) BitmapOpType() OpType { return OpTypeUpdateBytes }

func LookupOpBitmapUpdateBytes(target ReadOnlyBitmap, name string) OpBitmapUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, uint64, Bitmap))
		return OpBitmapUpdateBytes(fn)
	}
	return nil
}

// ImportRoaring performs a default BitmapUpdateBytes on a Bitmap.
type interfaceHasImportRoaring interface {
	ImportRoaringUpdateBytes([]byte) (bool, uint64, Bitmap)
}

func ImportRoaring(target Bitmap, in1 []byte) (bool, uint64, Bitmap) {
	if target, ok := target.(interfaceHasImportRoaring); ok {
		return target.ImportRoaringUpdateBytes(in1)
	}
	return genericImportRoaring(target, in1)
}

type OpBitmapViewWriterGivesError func(io.Writer) error

func (OpBitmapViewWriterGivesError) BitmapOpType() OpType { return OpTypeViewWriterGivesError }

func LookupOpBitmapViewWriterGivesError(target ReadOnlyBitmap, name string) OpBitmapViewWriterGivesError {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewWriterGivesError")
	if method.IsValid() {
		fn, _ := method.Interface().(func(io.Writer) error)
		return OpBitmapViewWriterGivesError(fn)
	}
	return nil
}

// ExportRoaring performs a default BitmapViewWriterGivesError on a Bitmap.
type interfaceHasExportRoaring interface {
	ExportRoaringViewWriterGivesError(io.Writer) error
}

func ExportRoaring(target ReadOnlyBitmap, in1 io.Writer) error {
	if target, ok := target.(interfaceHasExportRoaring); ok {
		return target.ExportRoaringViewWriterGivesError(in1)
	}
	return genericExportRoaring(target, in1)
}
