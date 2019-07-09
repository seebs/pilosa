package data

// Container represents a 2^16 block of bit values, numbered 0..65535.
type Container interface {
	// PreferredType is a hint as to which of the container representations will be cheapest to request.
	PreferredType() ContainerType
	// BitVec yields the container's contents as a bit vector in a slice of 1024 uint64.
	BitVec() BitVec
	// Slice yields the container's contents as a slice of individual uint16.
	Slice() Slice
	// Runs yields the container's contents as a slice of runs.
	Runs() Runs
	// Any indicates whether at least one bit is set.
	Any() bool
	// Count reports the number of bits set.
	Count() int64
	// Add sets the given bit, yielding a possibly-new container and a boolean indicating whether this was a change.
	Add(uint16) (Container, bool)
	// Remove clears the given bit, yielding a possibly-new container and a boolean indicating whether this was a change.
	Remove(uint16) (Container, bool)
	// Union returns a container holding the union of this container and the other.
	Union(Container) Container
	// Intersection returns a container holding the intersection of this container and the other.
	Intersection(Container) Container
	// Difference returns a container holding the difference of this container and the other. (Thus, clearing any
	// bits set in the other.)
	Difference(Container) Container
	// Xor returns a container holding the exclusive or of this container and the other.
	Xor(Container) Container
	// UnionInPlace returns a container holding the union of this container and the other,
	// but may overwrite this container.
	UnionInPlace(Container) Container
	// Intersection returns a container holding the intersection of this container and the other,
	// but may overwrite this container.
	IntersectionInPlace(Container) Container
	// Difference returns a container holding the difference of this container and the other, but
	// may overwrite this container.
	DifferenceInPlace(Container) Container
	// Xor returns a container holding the exclusive or of this container and the other, but
	// may overwrite this container.
	XorInPlace(Container) Container
	// Inverse returns a container with all bits inverted.
	Inverse(Container) Container
	// InverseInPlace returns a container with all bits inverted, but may overwrite the
	// original container.
	InverseInPlace(Container) Container
}

// ContainerType is a representation format used by a container. The set of
// ContainerType values is potentially open-ended, but usually you should be
// using one of the named "standard" types.
type ContainerType byte

const (
	ContainerBitVec = ContainerType(iota)
	ContainerSlice
	ContainerRuns
)

// BitVec represents a slice of 1024 uint64. It's a slice, not an array,
// because slices are smaller to copy.
type BitVec []uint64

// Slice represents a sorted slice of uint16 values. Usually it would be used
// only for N < 4096 (thus, under 8KB of storage).
type Slice []uint16

// Run represents a series of bits, from first to last.
type Run struct{ First, Last uint16 }

// Runs represents a container's contents as a sorted slice of Runs.
type Runs []Run