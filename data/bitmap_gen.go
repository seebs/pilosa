package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Bitmap

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionBitmap interface {
	Type(Bitmap) OpType
}

// OpBitmapView is a View operation on a ReadOnlyBitmap with no other parameters.
type OpBitmapView func(ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapView) Type(Bitmap) OpType { return OpTypeView }

// OpBitmapUpdate is an Update operation on a Bitmap with no other parameters.
type OpBitmapUpdate func(Bitmap) (bool, int64, Bitmap)
func (OpBitmapUpdate) Type(Bitmap) OpType { return OpTypeUpdate }

// OpBitmapViewBit is a View operation on a ReadOnlyBitmap and one Bit.
type OpBitmapViewBit func(ReadOnlyBitmap, uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBit) Type(Bitmap) OpType { return OpTypeViewBit }

// OpBitmapUpdateBit is an Update operation on a Bitmap and one Bit.
type OpBitmapUpdateBit func(Bitmap, uint64) (bool, int64, Bitmap)
func (OpBitmapUpdateBit) Type(Bitmap) OpType { return OpTypeUpdateBit }

// OpBitmapViewBitmap is a View operation on a ReadOnlyBitmap and one other Bitmap.
type OpBitmapViewBitmap func(ReadOnlyBitmap, ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBitmap) Type(Bitmap) OpType { return OpTypeViewOther }

// OpBitmapUpdateBitmap is an Update operation on a Bitmap and one other Bitmap.
type OpBitmapUpdateBitmap func(Bitmap, ReadOnlyBitmap) (bool, int64, Bitmap)
func (OpBitmapUpdateBitmap) Type(Bitmap) OpType { return OpTypeUpdateOther }

// OpBitmapViewBits is a View operation on a ReadOnlyBitmap and one or more Bits.
type OpBitmapViewBits func(ReadOnlyBitmap, ...uint64) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBits) Type(Bitmap) OpType { return OpTypeViewBits }

// OpBitmapUpdateBits is an Update operation on a Bitmap and one or more Bits.
type OpBitmapUpdateBits func(Bitmap, ...uint64) (bool, int64, Bitmap)
func (OpBitmapUpdateBits) Type(Bitmap) OpType { return OpTypeUpdateBits }

// OpBitmapViewBitmaps is a View operation on a ReadOnlyBitmap and one or more other Bitmaps.
type OpBitmapViewBitmaps func(ReadOnlyBitmap, ...ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap)
func (OpBitmapViewBitmaps) Type(Bitmap) OpType { return OpTypeViewOthers }

// OpBitmapUpdateBitmaps is an Update operation on a Bitmap and one or more other Bitmaps.
type OpBitmapUpdateBitmaps func(Bitmap, ...ReadOnlyBitmap) (bool, int64, Bitmap)
func (OpBitmapUpdateBitmaps) Type(Bitmap) OpType { return OpTypeUpdateOthers }
