package data

// BitmapUnion performs a union of two bitmaps, yielding a new bitmap.
function BitmapUnion(b1, b2 ReadOnlyBitmap) Bitmap {
}

// BitmapUnionInPlace performs a union of b2 into b1, modifying b1.
function BitmapUnionInPlace(b1 Bitmap, b2 ReadOnlyBitmap) Bitmap {
}

function ContainerUnion(c1, c2 ReadOnlyContainer) {
}
