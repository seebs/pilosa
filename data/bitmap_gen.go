package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Bitmap (see gen/main.go)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionBitmap interface {
	Type(Bitmap) OpType
}

// OpBitmapView is a View operation on a ReadOnlyBitmap with no other parameters.
type OpBitmapView func() (bool, int64, ReadOnlyBitmap)
func (OpBitmapView) Type(Bitmap) OpType { return OpTypeView }
var zeroOpBitmapView OpBitmapView

// OpBitmapUpdate is an Update operation on a Bitmap with no other parameters.
type OpBitmapUpdate func() (bool, int64, Bitmap)
func (OpBitmapUpdate) Type(Bitmap) OpType { return OpTypeUpdate }
var zeroOpBitmapUpdate OpBitmapUpdate

// OpBitmapViewBit is a View operation on a ReadOnlyBitmap and one Bit.
type OpBitmapViewBit func(uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBit) Type(Bitmap) OpType { return OpTypeViewBit }
var zeroOpBitmapViewBit OpBitmapViewBit

// OpBitmapUpdateBit is an Update operation on a Bitmap and one Bit.
type OpBitmapUpdateBit func(uint64) (bool, int64, Bitmap)
func (OpBitmapUpdateBit) Type(Bitmap) OpType { return OpTypeUpdateBit }
var zeroOpBitmapUpdateBit OpBitmapUpdateBit

// OpBitmapViewBitmap is a View operation on a ReadOnlyBitmap and one other Bitmap.
type OpBitmapViewBitmap func(ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBitmap) Type(Bitmap) OpType { return OpTypeViewOther }
var zeroOpBitmapViewBitmap OpBitmapViewBitmap

// OpBitmapUpdateBitmap is an Update operation on a Bitmap and one other Bitmap.
type OpBitmapUpdateBitmap func(ReadOnlyBitmap) (bool, int64, Bitmap)
func (OpBitmapUpdateBitmap) Type(Bitmap) OpType { return OpTypeUpdateOther }
var zeroOpBitmapUpdateBitmap OpBitmapUpdateBitmap

// OpBitmapViewBits is a View operation on a ReadOnlyBitmap and one or more Bits.
type OpBitmapViewBits func(...uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBits) Type(Bitmap) OpType { return OpTypeViewBits }
var zeroOpBitmapViewBits OpBitmapViewBits

// OpBitmapUpdateBits is an Update operation on a Bitmap and one or more Bits.
type OpBitmapUpdateBits func(...uint64) (bool, int64, Bitmap)
func (OpBitmapUpdateBits) Type(Bitmap) OpType { return OpTypeUpdateBits }
var zeroOpBitmapUpdateBits OpBitmapUpdateBits

// OpBitmapViewBitmaps is a View operation on a ReadOnlyBitmap and one or more other Bitmaps.
type OpBitmapViewBitmaps func(...ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBitmaps) Type(Bitmap) OpType { return OpTypeViewOthers }
var zeroOpBitmapViewBitmaps OpBitmapViewBitmaps

// OpBitmapUpdateBitmaps is an Update operation on a Bitmap and one or more other Bitmaps.
type OpBitmapUpdateBitmaps func(...ReadOnlyBitmap) (bool, int64, Bitmap)
func (OpBitmapUpdateBitmaps) Type(Bitmap) OpType { return OpTypeUpdateOthers }
var zeroOpBitmapUpdateBitmaps OpBitmapUpdateBitmaps

// OpBitmapViewBytesUpdateBytes is a View operation on a ReadOnlyBitmap and one or more BytesUpdateBytes.
type OpBitmapViewBytesUpdateBytes func( (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBytesUpdateBytes) Type(Bitmap) OpType { return OpTypeViewBytesUpdateBytes }
var zeroOpBitmapViewBytesUpdateBytes OpBitmapViewBytesUpdateBytes

// OpType to reflect.Type lookup table
var lookupBitmapFunctionTypes = [OpTypeMax]OpFunctionBitmap {
	OpTypeView: zeroOpBitmapView,
	OpTypeUpdate: zeroOpBitmapUpdate,
	OpTypeViewBit: zeroOpBitmapViewBit,
	OpTypeUpdateBit: zeroOpBitmapUpdateBit,
	OpTypeViewOther: zeroOpBitmapViewBitmap,
	OpTypeUpdateOther: zeroOpBitmapUpdateBitmap,
	OpTypeViewBits: zeroOpBitmapViewBits,
	OpTypeUpdateBits: zeroOpBitmapUpdateBits,
	OpTypeViewOthers: zeroOpBitmapViewBitmaps,
	OpTypeUpdateOthers: zeroOpBitmapUpdateBitmaps,
	OpTypeViewBytesUpdateBytes: zeroOpBitmapViewBytesUpdateBytes,
}
