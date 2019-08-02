package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Bitmap (see gen/main.go). These are expressed
// as method signatures -- the Bitmap they operate on is an implicit
// receiver not shown in the signature.

import (
	"reflect"
)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionBitmap interface {
	BitmapOpType() OpType
}

type OpBitmapView func() (bool, int64, ReadOnlyBitmap)

func (OpBitmapView) BitmapOpType() OpType { return OpTypeView }

func LookupOpBitmapView(target ReadOnlyBitmap, name string) OpBitmapView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int64, ReadOnlyBitmap))
		return OpBitmapView(fn)
	}
	return nil
}

type OpBitmapUpdate func() (bool, int64, Bitmap)

func (OpBitmapUpdate) BitmapOpType() OpType { return OpTypeUpdate }

func LookupOpBitmapUpdate(target ReadOnlyBitmap, name string) OpBitmapUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int64, Bitmap))
		return OpBitmapUpdate(fn)
	}
	return nil
}

type OpBitmapViewRange func(uint64, uint64) (bool, int64, ReadOnlyBitmap)

func (OpBitmapViewRange) BitmapOpType() OpType { return OpTypeViewRange }

func LookupOpBitmapViewRange(target ReadOnlyBitmap, name string) OpBitmapViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) (bool, int64, ReadOnlyBitmap))
		return OpBitmapViewRange(fn)
	}
	return nil
}

type OpBitmapViewBit func(uint64) (bool, int64, ReadOnlyBitmap)

func (OpBitmapViewBit) BitmapOpType() OpType { return OpTypeViewBit }

func LookupOpBitmapViewBit(target ReadOnlyBitmap, name string) OpBitmapViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, int64, ReadOnlyBitmap))
		return OpBitmapViewBit(fn)
	}
	return nil
}

type OpBitmapUpdateBit func(uint64) (bool, int64, Bitmap)

func (OpBitmapUpdateBit) BitmapOpType() OpType { return OpTypeUpdateBit }

func LookupOpBitmapUpdateBit(target ReadOnlyBitmap, name string) OpBitmapUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, int64, Bitmap))
		return OpBitmapUpdateBit(fn)
	}
	return nil
}

// Add performs a default BitmapUpdateBit on a Bitmap.
type interfaceHasAdd interface {
	AddUpdateBit(uint64) (bool, int64, Bitmap)
}

func Add(target Bitmap, in1 uint64) (bool, int64, Bitmap) {
	if target, ok := target.(interfaceHasAdd); ok {
		return target.AddUpdateBit(in1)
	}
	return genericAdd(target, in1)
}

// Remove performs a default BitmapUpdateBit on a Bitmap.
type interfaceHasRemove interface {
	RemoveUpdateBit(uint64) (bool, int64, Bitmap)
}

func Remove(target Bitmap, in1 uint64) (bool, int64, Bitmap) {
	if target, ok := target.(interfaceHasRemove); ok {
		return target.RemoveUpdateBit(in1)
	}
	return genericRemove(target, in1)
}

type OpBitmapViewBitmap func(ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)

func (OpBitmapViewBitmap) BitmapOpType() OpType { return OpTypeViewOther }

func LookupOpBitmapViewBitmap(target ReadOnlyBitmap, name string) OpBitmapViewBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap))
		return OpBitmapViewBitmap(fn)
	}
	return nil
}

type OpBitmapUpdateBitmap func(ReadOnlyBitmap) (bool, int64, Bitmap)

func (OpBitmapUpdateBitmap) BitmapOpType() OpType { return OpTypeUpdateOther }

func LookupOpBitmapUpdateBitmap(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, int64, Bitmap))
		return OpBitmapUpdateBitmap(fn)
	}
	return nil
}

type OpBitmapViewBits func([]uint64) (bool, int64, ReadOnlyBitmap)

func (OpBitmapViewBits) BitmapOpType() OpType { return OpTypeViewBits }

func LookupOpBitmapViewBits(target ReadOnlyBitmap, name string) OpBitmapViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, int64, ReadOnlyBitmap))
		return OpBitmapViewBits(fn)
	}
	return nil
}

type OpBitmapUpdateBits func([]uint64) (bool, int64, Bitmap)

func (OpBitmapUpdateBits) BitmapOpType() OpType { return OpTypeUpdateBits }

func LookupOpBitmapUpdateBits(target ReadOnlyBitmap, name string) OpBitmapUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, int64, Bitmap))
		return OpBitmapUpdateBits(fn)
	}
	return nil
}

type OpBitmapViewBitmaps func([]ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)

func (OpBitmapViewBitmaps) BitmapOpType() OpType { return OpTypeViewOthers }

func LookupOpBitmapViewBitmaps(target ReadOnlyBitmap, name string) OpBitmapViewBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap))
		return OpBitmapViewBitmaps(fn)
	}
	return nil
}

type OpBitmapUpdateBitmaps func([]ReadOnlyBitmap) (bool, int64, Bitmap)

func (OpBitmapUpdateBitmaps) BitmapOpType() OpType { return OpTypeUpdateOthers }

func LookupOpBitmapUpdateBitmaps(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, int64, Bitmap))
		return OpBitmapUpdateBitmaps(fn)
	}
	return nil
}

type OpBitmapUpdateBytes func([]byte) (bool, int64, Bitmap)

func (OpBitmapUpdateBytes) BitmapOpType() OpType { return OpTypeUpdateBytes }

func LookupOpBitmapUpdateBytes(target ReadOnlyBitmap, name string) OpBitmapUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, int64, Bitmap))
		return OpBitmapUpdateBytes(fn)
	}
	return nil
}

// ImportRoaring performs a default BitmapUpdateBytes on a Bitmap.
type interfaceHasImportRoaring interface {
	ImportRoaringUpdateBytes([]byte) (bool, int64, Bitmap)
}

func ImportRoaring(target Bitmap, in1 []byte) (bool, int64, Bitmap) {
	if target, ok := target.(interfaceHasImportRoaring); ok {
		return target.ImportRoaringUpdateBytes(in1)
	}
	return genericImportRoaring(target, in1)
}

type OpBitmapViewGivesBytes func() []byte

func (OpBitmapViewGivesBytes) BitmapOpType() OpType { return OpTypeViewGivesBytes }

func LookupOpBitmapViewGivesBytes(target ReadOnlyBitmap, name string) OpBitmapViewGivesBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func() []byte)
		return OpBitmapViewGivesBytes(fn)
	}
	return nil
}

// ExportRoaring performs a default BitmapViewGivesBytes on a Bitmap.
type interfaceHasExportRoaring interface {
	ExportRoaringViewGivesBytes() []byte
}

func ExportRoaring(target ReadOnlyBitmap) []byte {
	if target, ok := target.(interfaceHasExportRoaring); ok {
		return target.ExportRoaringViewGivesBytes()
	}
	return genericExportRoaring(target)
}
