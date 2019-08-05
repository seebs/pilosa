package data

// likely generic functions, not yet provided:
// Clone(ReadOnlyContainer) Container
// Freeze(Container) ReadOnlyContainer
// Any(Container) bool
// Count(Container) int

// ReadOnlyContainer represents a 2^16 block of bit values, numbered 0..65535,
// which may not be safe to write to.
type ReadOnlyContainer interface {
	// PreferredType is a hint as to which of the container representations will be cheapest to request.
	PreferredType() ContainerType
	// BitVec yields the container's contents as a bit vector in a slice of 1024 uint64.
	BitVec() BitVec
	// Slice yields the container's contents as a slice of individual uint16.
	Slice() Slice
	// Runs yields the container's contents as a slice of runs.
	Runs() Runs
}

// Container represents a writeable container -- or at least, a container which
// provides write operations. The results of a write operation will sometimes be
// a new container for performance/safety reasons.
type Container interface {
	ReadOnlyContainer
	// Add sets the given bit, yielding a possibly-new container and a boolean indicating whether this was a change.
	Add(uint16) (bool, int, Container)
	// Remove clears the given bit, yielding a possibly-new container and a boolean indicating whether this was a change.
	Remove(uint16) (bool, int, Container)
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

func genericContainerAny(target ReadOnlyContainer) bool {
	return false
}

func genericContainerCount(target ReadOnlyContainer) int {
	return 0
}
