package data

//go:generate go run gen/main.go -types "Container:16:,Bitmap:64:64" View Update ViewBit UpdateBit ViewOther UpdateOther ViewBits UpdateBits ViewOthers UpdateOthers

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
)
