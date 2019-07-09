package data

// Bitmap represents an indexable vector of 1<<64 bits.
type Bitmap interface {
	// Any reports whether at least one bit is set.
	Any() bool
	// AnyRange reports whether at least one bit is set between first and last, inclusive.
	AnyRange(first, last uint64) bool
	// Count reports the number of bits set.
	Count() uint64
	// CountRange reports the number of bits set between first and last, inclusive.
	CountRange(first, last uint64) uint64
	// Slice() yields the bits as a slice of uint64, up to some limit;
	// past that, it will report overflow. The limit is up to the
	// implementation, but should be at least 2<<16. (This arbitrary
	// value exists to make test case writing easier.)
	Slice() (values []uint64, overflow bool)
	// SliceRange is like Slice(), but within a limited range.
	SliceRange(first, last uint64) (values []uint64, overflow bool)
	// Add sets the given bit, reporting whether or not this was a change.
	Add(uint64) bool
	// Remove clears the given bit, reporting whether or not this was a change.
	Remove(uint64) bool
	// GetContainer yields a Container holding the 1<<16 values starting at
	// key * (1<<16). For instance, if key is 1, it returns the bits from
	// 65,536 to 131,071 inclusive.
	GetContainer(key uint64) Container
	// PutContainer overwrites the container at key with a new container.
	PutContainer(key uint64, c Container)
	// Union, etc, do what they say they would do. InPlace forms modify the given
	// bitmap rather than creating a new one.
	Union(...Bitmap) Bitmap
	UnionInPlace(...Bitmap)
	Intersection(...Bitmap) Bitmap
	IntersectionInPlace(...Bitmap)
	Difference(...Bitmap) Bitmap
	DifferenceInPlace(...Bitmap)
	Xor(...Bitmap) Bitmap
	XorInPlace(...Bitmap)
	// OffsetRange yields a new bitmap containing values from first to last,
	// with all positions increased by offset. All three values must be multiples
	// of 1<<16.
	OffsetRange(offset, first, last uint64) Bitmap
	// utility functions for updates
	InvertRangeInPlace(first, last uint64)
	// ImportRoaring, et al., handle importing bits from raw Roaring data.
	// The Import case takes an additional parameter for whether or not the
	// data is memory-mapped; if so, it should try to map container contents
	// to that data if possible/applicable.
	ImportRoaring(data []byte, mapped bool) error
	UnionInPlaceRoaring(data []byte) error
	IntersectionInPlaceRoaring(data []byte) error
	DifferenceInPlaceRoaring(data []byte) error
	XorInPlaceRoaring(data []byte) error
	// ProcessContainers iterates through the containers present in the bitmap calling
	// func for each one. If func returns write == true, newC replaces the previous container
	// for that key. Process containers returns when it runs out of containers, done is true,
	// or err is non-nil, returning a non-nil err if one was given.
	ProcessContainers(fn func(key uint64, c Container) (newC Container, write bool, done bool, err error)) error
	// ProcessContainersRange does the same thing, but only for containers which contain bits in the provided
	// range. For instance, if called with first 0, last 1<<16, it will process containers at keys 0 and 1.
	// It doesn't care whether the bits in the range are set, just whether the containers would contain them
	// if they were.
	ProcessContainersRange(first, last uint64, fn func(key uint64, c Container) (newC Container, write bool, done bool, err error)) error
}
