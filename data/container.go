package data

type ContainerViewOp func(ReadOnlyContainer, ReadOnlyContainer) ReadOnlyContainer
type ContainerUpdateOp func(Container, ReadOnlyContainer) Container

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
	// Any indicates whether at least one bit is set.
	Any() bool
	// Count reports the number of bits set.
	Count() int64
	// Clone yields a definitely-new containe which is writeable.
	Clone() Container
}

// Container represents a writeable container -- or at least, a container which
// provides write operations. The results of a write operation will sometimes be
// a new container for performance/safety reasons.
type Container interface {
	ReadOnlyContainer
	// Add sets the given bit, yielding a possibly-new container and a boolean indicating whether this was a change.
	Add(uint16) (Container, bool)
	// Remove clears the given bit, yielding a possibly-new container and a boolean indicating whether this was a change.
	Remove(uint16) (Container, bool)
	// Freeze yields a read-only container identical to this one. It may
	// actually be this one, now marked read-only -- if so, it will have
	// been changed in such a way that future write operations will actually
	// yield a new container.
	Freeze() ReadOnlyContainer
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
