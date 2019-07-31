package data

// BitmapUnion performs a union of two bitmaps, yielding a new bitmap.
func BitmapUnion(b1, b2 ReadOnlyBitmap) (bool, int64, ReadOnlyBitmap) {
	return false, 0, nil
}

// BitmapUnionInPlace performs a union of b2 into b1, modifying b1.
func BitmapUnionInPlace(b1 Bitmap, b2 ReadOnlyBitmap) (bool, int64, Bitmap) {
	return false, 0, nil
}

func ContainerUnion(c1, c2 ReadOnlyContainer) {
}
