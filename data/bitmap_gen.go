package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Bitmap (see gen/main.go). These are expressed
// as method signatures -- the Bitmap they operate on is an implicit
// receiver not shown in the signature.

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionBitmap interface {
	BitmapOpType() OpType
}

type OpBitmapView func() (bool, int64, ReadOnlyBitmap)
func (OpBitmapView) BitmapOpType() OpType { return OpTypeView }
var zeroOpBitmapView OpBitmapView

type OpBitmapUpdate func() (bool, int64, Bitmap)
func (OpBitmapUpdate) BitmapOpType() OpType { return OpTypeUpdate }
var zeroOpBitmapUpdate OpBitmapUpdate

type OpBitmapViewRange func(uint64, uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewRange) BitmapOpType() OpType { return OpTypeViewRange }
var zeroOpBitmapViewRange OpBitmapViewRange

type OpBitmapViewBit func(uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBit) BitmapOpType() OpType { return OpTypeViewBit }
var zeroOpBitmapViewBit OpBitmapViewBit

type OpBitmapUpdateBit func(uint64) (bool, int64, Bitmap)
func (OpBitmapUpdateBit) BitmapOpType() OpType { return OpTypeUpdateBit }
var zeroOpBitmapUpdateBit OpBitmapUpdateBit

type OpBitmapViewBitmap func(ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBitmap) BitmapOpType() OpType { return OpTypeViewOther }
var zeroOpBitmapViewBitmap OpBitmapViewBitmap

type OpBitmapUpdateBitmap func(ReadOnlyBitmap) (bool, int64, Bitmap)
func (OpBitmapUpdateBitmap) BitmapOpType() OpType { return OpTypeUpdateOther }
var zeroOpBitmapUpdateBitmap OpBitmapUpdateBitmap

type OpBitmapViewBits func([]uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBits) BitmapOpType() OpType { return OpTypeViewBits }
var zeroOpBitmapViewBits OpBitmapViewBits

type OpBitmapUpdateBits func([]uint64) (bool, int64, Bitmap)
func (OpBitmapUpdateBits) BitmapOpType() OpType { return OpTypeUpdateBits }
var zeroOpBitmapUpdateBits OpBitmapUpdateBits

type OpBitmapViewBitmaps func([]ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBitmaps) BitmapOpType() OpType { return OpTypeViewOthers }
var zeroOpBitmapViewBitmaps OpBitmapViewBitmaps

type OpBitmapUpdateBitmaps func([]ReadOnlyBitmap) (bool, int64, Bitmap)
func (OpBitmapUpdateBitmaps) BitmapOpType() OpType { return OpTypeUpdateOthers }
var zeroOpBitmapUpdateBitmaps OpBitmapUpdateBitmaps

type OpBitmapUpdateBytes func([]byte) (bool, int64, Bitmap)
func (OpBitmapUpdateBytes) BitmapOpType() OpType { return OpTypeUpdateBytes }
var zeroOpBitmapUpdateBytes OpBitmapUpdateBytes

type OpBitmapViewGivesBytes func() ([]byte)
func (OpBitmapViewGivesBytes) BitmapOpType() OpType { return OpTypeViewGivesBytes }
var zeroOpBitmapViewGivesBytes OpBitmapViewGivesBytes

// OpType to reflect.Type lookup table
var lookupBitmapFunctionTypes = [OpTypeMax]OpFunctionBitmap {
	OpTypeView: zeroOpBitmapView,
	OpTypeUpdate: zeroOpBitmapUpdate,
	OpTypeViewRange: zeroOpBitmapViewRange,
	OpTypeViewBit: zeroOpBitmapViewBit,
	OpTypeUpdateBit: zeroOpBitmapUpdateBit,
	OpTypeViewOther: zeroOpBitmapViewBitmap,
	OpTypeUpdateOther: zeroOpBitmapUpdateBitmap,
	OpTypeViewBits: zeroOpBitmapViewBits,
	OpTypeUpdateBits: zeroOpBitmapUpdateBits,
	OpTypeViewOthers: zeroOpBitmapViewBitmaps,
	OpTypeUpdateOthers: zeroOpBitmapUpdateBitmaps,
	OpTypeUpdateBytes: zeroOpBitmapUpdateBytes,
	OpTypeViewGivesBytes: zeroOpBitmapViewGivesBytes,
}
