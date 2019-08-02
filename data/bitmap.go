package data

import (
	"io"
)

type ReadOnlyBitmap interface {
	// GetContainer yields a Container holding the 1<<16 values starting at
	// key * (1<<16). For instance, if key is 1, it returns the bits from
	// 65,536 to 131,071 inclusive.
	//
	// If the bitmap is a ReadOnlyBitmap, do not attempt to modify the
	// container. The bitmap may not enforce this.
	GetContainer(key uint64) Container
	// ViewContainers calls the provided function on a series of containers,
	// stopping when done is true or err is non-nil, and returning err if it
	// was non-nil.
	ViewContainers(fn func(key uint64, c Container) (done bool, err error)) error
	// ViewContainersRange is like ViewContainers, but covers only the containers
	// matching the given range.
	ViewContainersRange(first, last uint64, fn func(key uint64, c Container) (done bool, err error)) error
	// WriteTo dumps the bitmap's contents to the given writer as roaring data using Pilosa's format.
	WriteRoaringTo(io.Writer) error
}

// WriteOnlyBitmap is an interface which lets us describe the composition of
// ReadOnly and regular Bitmaps, and then do the same composition on
// TransactionalBitmaps.
type WriteOnlyBitmap interface {
	// Add sets the given bit, reporting whether or not this was a change.
	Add(uint64) bool
	// Remove clears the given bit, reporting whether or not this was a change.
	Remove(uint64) bool
	// PutContainer overwrites the container at key with a new container.
	PutContainer(key uint64, c Container)
	// ImportRoaring, et al., handle importing bits from raw Roaring data.
	// The Import case takes an additional parameter for whether or not the
	// data is memory-mapped; if so, it should try to map container contents
	// to that data if possible/applicable. Otherwise, and in the other cases,
	// you should *not* use the provided data.
	//
	// ImportRoaring should indicate whether or not any containers are using
	// the provided storage, so the caller can unmap if it's unused.
	ImportRoaring(data []byte, mapped bool) (mappedAny bool, err error)
	// OpInPlaceRoaring does the same thing, given an arbitrary container op.
	OpInPlaceRoaring(data []byte, fn OpContainerUpdate) error
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

// Bitmap represents an indexable vector of 1<<64 bits, and supports write
// operations to it as well as read operations. Operations on a bitmap
// are not guaranteed to be concurrency-safe. If you want concurrency guarantees,
// you might need a TransactionalBitmap.
type Bitmap interface {
	ReadOnlyBitmap
	WriteOnlyBitmap
}

// Like WriteOnlyBitmap, OpsLogOnlyBitmap exists to be composable with the other interfaces.
type OpsLogOnlyBitmap interface {
	SetOpsLog(io.Writer)
	DisableOpsLog()
}

// OpsLogBitmap represents a bitmap which supports an operations log. An operations
// log is an io.Writer to which operations should be serialized. Operation logging
// can be disabled for performance reasons, but this is almost always a bad idea.
// By default, there's no concurrency guarantees on this behavior; don't run multiple
// simultaneous writes that would need ops logging.
type OpsLogBitmap interface {
	Bitmap
	OpsLogOnlyBitmap
}

// BitmapViewer is a function which operates on a read-only bitmap.
type BitmapViewer func(b ReadOnlyBitmap) (err error)

// BitmapUpdater is a function which operates on a bitmap. It should return
// write=true if it wishes to apply its update, and a non-nil error if an
// error occurred. If an error occurred, and write is true, the write should
// still occur (but the error will be passed back up to another caller.)
type BitmapUpdater func(b Bitmap) (write bool, err error)

type TransactionalReadOnlyBitmap interface {
	ReadOnlyBitmap
	View(first, last uint64, fn BitmapViewer) error
}

// TransactionalBitmap supports transactions which can run concurrently and
// safely.
//
// There exists an ordering on update operations for a given region of storage;
// if operation A happens before operation B, then anything that happens before
// operation A also happens before operation B.
//
// Any given view operation sees the state of the database after some set of
// updates, and before some set of updates. (Either set might be empty.) An update
// operation which starts after a view operation has begun executing will not
// affect the data seen by the view.
//
// How this is accomplished is an implementation detail. For instance, you could
// implement this with an RWMutex, with all views requesting a read lock, and
// updates requesting a write lock. You could implement it with MVCC.
//
// ImmediateUpdate is a special case; it acts like Update, but is blocked by
// and blocks View operations. For an implementation which would usually use
// MVCC, this might omit a lot of overhead (such as copying), and allow improved
// performance. However, it will likely degrade performance under a lot of
// other workloads.
//
// When running an ImmediateUpdate, if the write return value is false, that
// does not guarantee that the parent bitmap wasn't modified; modifications of
// containers could still have modified the parent's contents. Don't use
// ImmediateUpdate if you can't ensure that this isn't a problem.
type TransactionalBitmap interface {
	TransactionalReadOnlyBitmap
	WriteOnlyBitmap // lol
	Update(first, last uint64, fn BitmapUpdater) error
	ImmediateUpdate(first, last uint64, fn BitmapUpdater) error
}

// TransactionalOpsLogBitmap supports both transactions and operation logs.
// How they combine:
// If ops logging is enabled when an update occurs, the bitmap provided to
// the update callback function will also be an OpsLogBitmap. When the
// update process completes, if it returns with write=true, then data written to
// the ops log for that bitmap will also be appended to the ops log for the
// calling bitmap, otherwise, it won't. If ops logging is disabled, the
// bitmap provided to the update callabck will not be an OpsLogBitmap.
type TransactionalOpsLogBitmap interface {
	TransactionalBitmap
	OpsLogOnlyBitmap
}

func genericImportRoaring(target Bitmap, data []byte) (bool, uint64, Bitmap) {
	return false, 0, nil
}

func genericExportRoaring(target ReadOnlyBitmap) []byte {
	return nil
}

func genericAdd(target Bitmap, bit uint64) (bool, uint64, Bitmap) {
	return false, 0, nil
}

func genericRemove(target Bitmap, bit uint64) (bool, uint64, Bitmap) {
	return false, 0, nil
}

func genericAny(target ReadOnlyBitmap) bool {
	return false
}

func genericCount(target ReadOnlyBitmap) uint64 {
	return 0
}

func genericAnyRange(target ReadOnlyBitmap, first, last uint64) bool {
	return false
}

func genericCountRange(target ReadOnlyBitmap, first, last uint64) uint64 {
	return 0
}

func genericSliceRange(target ReadOnlyBitmap, first, last uint64) ([]uint64, bool) {
	return nil, false
}

func genericOffsetRange(target ReadOnlyBitmap, first, last uint64) Bitmap {
	return nil
}
