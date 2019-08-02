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
	// ViewContainersRange is like ViewContainers, but covers only the containers
	// matching the given range. It starts with the first container with a
	// key not less than first, and stops with the last container with
	// a key not greater than last.
	ViewContainersRange(first, last uint64, fn func(key uint64, c Container) (done bool, err error)) error
}

// WriteOnlyBitmap is an interface which lets us describe the composition of
// ReadOnly and regular Bitmaps, and then do the same composition on
// TransactionalBitmaps.
type WriteOnlyBitmap interface {
	// PutContainer overwrites the container at key with a new container.
	PutContainer(key uint64, c Container)
	// ProcessContainersRange iterates through the containers present in the bitmap calling
	// func for each one. If func returns write == true, newC replaces the previous container
	// for that key. Process containers returns when it runs out of containers, done is true,
	// or err is non-nil, returning a non-nil err if one was given.
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

func genericExportRoaring(target ReadOnlyBitmap, w io.Writer) error {
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
