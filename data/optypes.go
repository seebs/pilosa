package data

//go:generate go run gen/main.go

// OpType represents the kind of operation we're looking at; whether it's
// a View (read-only) or Update (may modify) operation, and what its other
// parameter might be.
//
// OpType(View|Update)(Parameters)Gives(Results)
// If "Parameters" is omitted, it's empty.
//
// If "Gives(Results)" is omitted, it's (bool, int, object); for instance,
// (bool, int, Container) for Container, (bool, int64, Bitmap) for Bitmap.
//
// These types correspond to method names that would be used in specializations.
// For instance, if a given Bitmap implementation offers an efficient
// implementation of Union with other bitmaps, updating itself with the
// results, that method would be called UnionUpdateOthers. If it offers
// a specialized implementation of roaring bitmap export, that method
// would be called ExportRoaringViewGivesBytes.
type OpType int

// actual op type declarations live in optype_gen.go. They're generated
// from optype_list.txt.
