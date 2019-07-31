package data

//go:generate go run gen/main.go -types "Container:16:,Bitmap:64:64" View Update ViewBit UpdateBit ViewOther UpdateOther ViewBits UpdateBits ViewOthers UpdateOthers ViewBytesUpdateBytes

// OpType represents the kind of operation we're looking at; whether it's
// a View (read-only) or Update (may modify) operation, and what its other
// parameter might be.
//
// Every OpType should exist in View and Update forms, and they should alternate
// in the iota-based list below so that OpTypeViewX|1 == OpTypeUpdateX.
type OpType int

const (
	// OpTypeView and OpTypeUpdate are unary ops -- no other parameters.
	OpTypeView = OpType(iota)
	OpTypeUpdate
	// OpTypeViewBit and OpTypeUpdateBit are binary ops, where the second
	// parameter is a bit (uint16 for a Container, uint64 for a Bitmap),
	// such as add/remove.
	OpTypeViewBit
	OpTypeUpdateBit
	// OpTypeViewOther and OpTypeUpdateOther are binary ops, where the
	// second parameter is another (read-only) object of the same sort --
	// a Container or Bitmap.
	OpTypeViewOther
	OpTypeUpdateOther
	// OpTypeViewBits and OpTypeUpdateBits are N-ary ops, where the other
	// parameters are bits (uint16/uint64). Examples would be AddN for a
	// bitmap or container.
	OpTypeViewBits
	OpTypeUpdateBits
	// OpTypeViewOthers and OpTypeUpdateOthers are N-ary, where the other
	// parameters are containers or bitmaps respectively. This is only
	// meaningful for some operation types, but for instance, you can
	// Union or Intersect large sets.
	OpTypeViewOthers
	OpTypeUpdateOthers
	// OpTypeViewBytes and OpTypeUpdateBytes take a slice of bytes, which
	// is written as "...byte" because that matches the spelling of the
	// other cases, but the assumption is you'd call with `slice...`
	// parameters. These are provided to implement things like ImportRoaring.
	OpTypeViewBytes
	OpTypeUpdateBytes
	// OpTypeMax is not a valid OpType; it's the exclusive upper bound
	// of the set. Thus, `[OpTypeMax]foo` is an array-of-foo in which
	// every valid OpType is a valid index.
	OpTypeMax
	// If you are adding an OpType below this line, you need to re-read
	// that last comment.
)
