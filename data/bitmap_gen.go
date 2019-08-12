package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Bitmap (see gen/main.go). These are expressed
// as method signatures -- the Bitmap they operate on is an implicit
// receiver not shown in the signature.

import (
	"io"
	"reflect"
)

// OpFunctionBitmap exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionBitmap interface {
	BitmapOpType() OpType
}

// OpTableGenericBitmap, similarly, lets us specify a range of types -- in
// this case, map[string]OpFunctionType, where the type is one of the
// specific op function types defined.
type OpTableGenericBitmap interface {
	BitmapOpTypeTable() OpType
}

// OpTableBitmap is a slice mapping optypes to map[string]opFunc,
// where any specific map will actually be a map with a concrete type of
// op function. We defined the
type OpTableBitmap []OpTableGeneric

// Implementation stuff for BitmapView, the Bitmap-specific
// form of OpTypeView.
type OpBitmapView func() (bool, uint64, ReadOnlyBitmap)
func (OpBitmapView) BitmapOpType() { return OpTypeView }
type OpTableBitmapView map[string]opBitmapView
func (OpTableBitmapView) BitmapOpTypeTable() OpType { return OpTypeView }

func LookupOpBitmapView(target ReadOnlyBitmap, name string) OpBitmapView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, uint64, ReadOnlyBitmap))
		return OpBitmapView(fn)
	}
	return nil
}

func LookupTableOpBitmapView(table OpTableBitmap, name string) OpBitmapView {
	if table == nil {
		return nil
	}
	subTable := table[BitmapView]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapView)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapViewGivesBool, the Bitmap-specific
// form of OpTypeViewGivesBool.
type OpBitmapViewGivesBool func() bool
func (OpBitmapViewGivesBool) BitmapOpType() { return OpTypeViewGivesBool }
type OpTableBitmapViewGivesBool map[string]opBitmapViewGivesBool
func (OpTableBitmapViewGivesBool) BitmapOpTypeTable() OpType { return OpTypeViewGivesBool }

func LookupOpBitmapViewGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func() bool)
		return OpBitmapViewGivesBool(fn)
	}
	return nil
}

func LookupTableOpBitmapViewGivesBool(table OpTableBitmap, name string) OpBitmapViewGivesBool {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewGivesBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewGivesBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Any performs a default BitmapViewGivesBool on a Bitmap
type interfaceBitmapHasAny interface {
	AnyViewGivesBool() bool
}
func Any(target ReadOnlyBitmap) bool {
	if target, ok := target.(interfaceBitmapHasAny); ok {
		return target.AnyViewGivesBool()
	}
	return genericAny(target)
}

// Implementation stuff for BitmapViewGivesBit, the Bitmap-specific
// form of OpTypeViewGivesBit.
type OpBitmapViewGivesBit func() uint64
func (OpBitmapViewGivesBit) BitmapOpType() { return OpTypeViewGivesBit }
type OpTableBitmapViewGivesBit map[string]opBitmapViewGivesBit
func (OpTableBitmapViewGivesBit) BitmapOpTypeTable() OpType { return OpTypeViewGivesBit }

func LookupOpBitmapViewGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func() uint64)
		return OpBitmapViewGivesBit(fn)
	}
	return nil
}

func LookupTableOpBitmapViewGivesBit(table OpTableBitmap, name string) OpBitmapViewGivesBit {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewGivesBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewGivesBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Count performs a default BitmapViewGivesBit on a Bitmap
type interfaceBitmapHasCount interface {
	CountViewGivesBit() uint64
}
func Count(target ReadOnlyBitmap) uint64 {
	if target, ok := target.(interfaceBitmapHasCount); ok {
		return target.CountViewGivesBit()
	}
	return genericCount(target)
}

// Implementation stuff for BitmapViewRangeGivesBool, the Bitmap-specific
// form of OpTypeViewRangeGivesBool.
type OpBitmapViewRangeGivesBool func(uint64, uint64) bool
func (OpBitmapViewRangeGivesBool) BitmapOpType() { return OpTypeViewRangeGivesBool }
type OpTableBitmapViewRangeGivesBool map[string]opBitmapViewRangeGivesBool
func (OpTableBitmapViewRangeGivesBool) BitmapOpTypeTable() OpType { return OpTypeViewRangeGivesBool }

func LookupOpBitmapViewRangeGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) bool)
		return OpBitmapViewRangeGivesBool(fn)
	}
	return nil
}

func LookupTableOpBitmapViewRangeGivesBool(table OpTableBitmap, name string) OpBitmapViewRangeGivesBool {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewRangeGivesBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewRangeGivesBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// AnyRange performs a default BitmapViewRangeGivesBool on a Bitmap
type interfaceBitmapHasAnyRange interface {
	AnyRangeViewRangeGivesBool(uint64, uint64) bool
}
func AnyRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) bool {
	if target, ok := target.(interfaceBitmapHasAnyRange); ok {
		return target.AnyRangeViewRangeGivesBool(in1, in2)
	}
	return genericAnyRange(target, in1, in2)
}

// Implementation stuff for BitmapViewRangeGivesBit, the Bitmap-specific
// form of OpTypeViewRangeGivesBit.
type OpBitmapViewRangeGivesBit func(uint64, uint64) uint64
func (OpBitmapViewRangeGivesBit) BitmapOpType() { return OpTypeViewRangeGivesBit }
type OpTableBitmapViewRangeGivesBit map[string]opBitmapViewRangeGivesBit
func (OpTableBitmapViewRangeGivesBit) BitmapOpTypeTable() OpType { return OpTypeViewRangeGivesBit }

func LookupOpBitmapViewRangeGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) uint64)
		return OpBitmapViewRangeGivesBit(fn)
	}
	return nil
}

func LookupTableOpBitmapViewRangeGivesBit(table OpTableBitmap, name string) OpBitmapViewRangeGivesBit {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewRangeGivesBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewRangeGivesBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// CountRange performs a default BitmapViewRangeGivesBit on a Bitmap
type interfaceBitmapHasCountRange interface {
	CountRangeViewRangeGivesBit(uint64, uint64) uint64
}
func CountRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) uint64 {
	if target, ok := target.(interfaceBitmapHasCountRange); ok {
		return target.CountRangeViewRangeGivesBit(in1, in2)
	}
	return genericCountRange(target, in1, in2)
}

// Implementation stuff for BitmapViewRangeGivesBitmap, the Bitmap-specific
// form of OpTypeViewRangeGivesOther.
type OpBitmapViewRangeGivesBitmap func(uint64, uint64) ReadOnlyBitmap
func (OpBitmapViewRangeGivesBitmap) BitmapOpType() { return OpTypeViewRangeGivesOther }
type OpTableBitmapViewRangeGivesBitmap map[string]opBitmapViewRangeGivesBitmap
func (OpTableBitmapViewRangeGivesBitmap) BitmapOpTypeTable() OpType { return OpTypeViewRangeGivesOther }

func LookupOpBitmapViewRangeGivesBitmap(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) ReadOnlyBitmap)
		return OpBitmapViewRangeGivesBitmap(fn)
	}
	return nil
}

func LookupTableOpBitmapViewRangeGivesBitmap(table OpTableBitmap, name string) OpBitmapViewRangeGivesBitmap {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewRangeGivesBitmap]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewRangeGivesBitmap)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// OffsetRange performs a default BitmapViewRangeGivesBitmap on a Bitmap
type interfaceBitmapHasOffsetRange interface {
	OffsetRangeViewRangeGivesBitmap(uint64, uint64) ReadOnlyBitmap
}
func OffsetRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) ReadOnlyBitmap {
	if target, ok := target.(interfaceBitmapHasOffsetRange); ok {
		return target.OffsetRangeViewRangeGivesBitmap(in1, in2)
	}
	return genericOffsetRange(target, in1, in2)
}

// Implementation stuff for BitmapViewRangeGivesBitsBool, the Bitmap-specific
// form of OpTypeViewRangeGivesBitsBool.
type OpBitmapViewRangeGivesBitsBool func(uint64, uint64) ([]uint64, bool)
func (OpBitmapViewRangeGivesBitsBool) BitmapOpType() { return OpTypeViewRangeGivesBitsBool }
type OpTableBitmapViewRangeGivesBitsBool map[string]opBitmapViewRangeGivesBitsBool
func (OpTableBitmapViewRangeGivesBitsBool) BitmapOpTypeTable() OpType { return OpTypeViewRangeGivesBitsBool }

func LookupOpBitmapViewRangeGivesBitsBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitsBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitsBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) ([]uint64, bool))
		return OpBitmapViewRangeGivesBitsBool(fn)
	}
	return nil
}

func LookupTableOpBitmapViewRangeGivesBitsBool(table OpTableBitmap, name string) OpBitmapViewRangeGivesBitsBool {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewRangeGivesBitsBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewRangeGivesBitsBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// SliceRange performs a default BitmapViewRangeGivesBitsBool on a Bitmap
type interfaceBitmapHasSliceRange interface {
	SliceRangeViewRangeGivesBitsBool(uint64, uint64) ([]uint64, bool)
}
func SliceRange(target ReadOnlyBitmap, in1 uint64, in2 uint64) ([]uint64, bool) {
	if target, ok := target.(interfaceBitmapHasSliceRange); ok {
		return target.SliceRangeViewRangeGivesBitsBool(in1, in2)
	}
	return genericSliceRange(target, in1, in2)
}

// Implementation stuff for BitmapUpdate, the Bitmap-specific
// form of OpTypeUpdate.
type OpBitmapUpdate func() (bool, uint64, Bitmap)
func (OpBitmapUpdate) BitmapOpType() { return OpTypeUpdate }
type OpTableBitmapUpdate map[string]opBitmapUpdate
func (OpTableBitmapUpdate) BitmapOpTypeTable() OpType { return OpTypeUpdate }

func LookupOpBitmapUpdate(target ReadOnlyBitmap, name string) OpBitmapUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, uint64, Bitmap))
		return OpBitmapUpdate(fn)
	}
	return nil
}

func LookupTableOpBitmapUpdate(table OpTableBitmap, name string) OpBitmapUpdate {
	if table == nil {
		return nil
	}
	subTable := table[BitmapUpdate]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdate)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapViewRange, the Bitmap-specific
// form of OpTypeViewRange.
type OpBitmapViewRange func(uint64, uint64) (bool, uint64, ReadOnlyBitmap)
func (OpBitmapViewRange) BitmapOpType() { return OpTypeViewRange }
type OpTableBitmapViewRange map[string]opBitmapViewRange
func (OpTableBitmapViewRange) BitmapOpTypeTable() OpType { return OpTypeViewRange }

func LookupOpBitmapViewRange(target ReadOnlyBitmap, name string) OpBitmapViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewRange(fn)
	}
	return nil
}

func LookupTableOpBitmapViewRange(table OpTableBitmap, name string) OpBitmapViewRange {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewRange]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewRange)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapViewBit, the Bitmap-specific
// form of OpTypeViewBit.
type OpBitmapViewBit func(uint64) (bool, uint64, ReadOnlyBitmap)
func (OpBitmapViewBit) BitmapOpType() { return OpTypeViewBit }
type OpTableBitmapViewBit map[string]opBitmapViewBit
func (OpTableBitmapViewBit) BitmapOpTypeTable() OpType { return OpTypeViewBit }

func LookupOpBitmapViewBit(target ReadOnlyBitmap, name string) OpBitmapViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBit(fn)
	}
	return nil
}

func LookupTableOpBitmapViewBit(table OpTableBitmap, name string) OpBitmapViewBit {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapUpdateBit, the Bitmap-specific
// form of OpTypeUpdateBit.
type OpBitmapUpdateBit func(uint64) (bool, uint64, Bitmap)
func (OpBitmapUpdateBit) BitmapOpType() { return OpTypeUpdateBit }
type OpTableBitmapUpdateBit map[string]opBitmapUpdateBit
func (OpTableBitmapUpdateBit) BitmapOpTypeTable() OpType { return OpTypeUpdateBit }

func LookupOpBitmapUpdateBit(target ReadOnlyBitmap, name string) OpBitmapUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, uint64, Bitmap))
		return OpBitmapUpdateBit(fn)
	}
	return nil
}

func LookupTableOpBitmapUpdateBit(table OpTableBitmap, name string) OpBitmapUpdateBit {
	if table == nil {
		return nil
	}
	subTable := table[BitmapUpdateBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Add performs a default BitmapUpdateBit on a Bitmap
type interfaceBitmapHasAdd interface {
	AddUpdateBit(uint64) (bool, uint64, Bitmap)
}
func Add(target Bitmap, in1 uint64) (bool, uint64, Bitmap) {
	if target, ok := target.(interfaceBitmapHasAdd); ok {
		return target.AddUpdateBit(in1)
	}
	return genericAdd(target, in1)
}
// Remove performs a default BitmapUpdateBit on a Bitmap
type interfaceBitmapHasRemove interface {
	RemoveUpdateBit(uint64) (bool, uint64, Bitmap)
}
func Remove(target Bitmap, in1 uint64) (bool, uint64, Bitmap) {
	if target, ok := target.(interfaceBitmapHasRemove); ok {
		return target.RemoveUpdateBit(in1)
	}
	return genericRemove(target, in1)
}

// Implementation stuff for BitmapViewBitmap, the Bitmap-specific
// form of OpTypeViewOther.
type OpBitmapViewBitmap func(ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap)
func (OpBitmapViewBitmap) BitmapOpType() { return OpTypeViewOther }
type OpTableBitmapViewBitmap map[string]opBitmapViewBitmap
func (OpTableBitmapViewBitmap) BitmapOpTypeTable() OpType { return OpTypeViewOther }

func LookupOpBitmapViewBitmap(target ReadOnlyBitmap, name string) OpBitmapViewBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBitmap(fn)
	}
	return nil
}

func LookupTableOpBitmapViewBitmap(table OpTableBitmap, name string) OpBitmapViewBitmap {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewBitmap]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBitmap)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapUpdateBitmap, the Bitmap-specific
// form of OpTypeUpdateOther.
type OpBitmapUpdateBitmap func(ReadOnlyBitmap) (bool, uint64, Bitmap)
func (OpBitmapUpdateBitmap) BitmapOpType() { return OpTypeUpdateOther }
type OpTableBitmapUpdateBitmap map[string]opBitmapUpdateBitmap
func (OpTableBitmapUpdateBitmap) BitmapOpTypeTable() OpType { return OpTypeUpdateOther }

func LookupOpBitmapUpdateBitmap(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, uint64, Bitmap))
		return OpBitmapUpdateBitmap(fn)
	}
	return nil
}

func LookupTableOpBitmapUpdateBitmap(table OpTableBitmap, name string) OpBitmapUpdateBitmap {
	if table == nil {
		return nil
	}
	subTable := table[BitmapUpdateBitmap]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBitmap)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapViewBits, the Bitmap-specific
// form of OpTypeViewBits.
type OpBitmapViewBits func([]uint64) (bool, uint64, ReadOnlyBitmap)
func (OpBitmapViewBits) BitmapOpType() { return OpTypeViewBits }
type OpTableBitmapViewBits map[string]opBitmapViewBits
func (OpTableBitmapViewBits) BitmapOpTypeTable() OpType { return OpTypeViewBits }

func LookupOpBitmapViewBits(target ReadOnlyBitmap, name string) OpBitmapViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBits(fn)
	}
	return nil
}

func LookupTableOpBitmapViewBits(table OpTableBitmap, name string) OpBitmapViewBits {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapUpdateBits, the Bitmap-specific
// form of OpTypeUpdateBits.
type OpBitmapUpdateBits func([]uint64) (bool, uint64, Bitmap)
func (OpBitmapUpdateBits) BitmapOpType() { return OpTypeUpdateBits }
type OpTableBitmapUpdateBits map[string]opBitmapUpdateBits
func (OpTableBitmapUpdateBits) BitmapOpTypeTable() OpType { return OpTypeUpdateBits }

func LookupOpBitmapUpdateBits(target ReadOnlyBitmap, name string) OpBitmapUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, uint64, Bitmap))
		return OpBitmapUpdateBits(fn)
	}
	return nil
}

func LookupTableOpBitmapUpdateBits(table OpTableBitmap, name string) OpBitmapUpdateBits {
	if table == nil {
		return nil
	}
	subTable := table[BitmapUpdateBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapViewBitmaps, the Bitmap-specific
// form of OpTypeViewOthers.
type OpBitmapViewBitmaps func([]ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap)
func (OpBitmapViewBitmaps) BitmapOpType() { return OpTypeViewOthers }
type OpTableBitmapViewBitmaps map[string]opBitmapViewBitmaps
func (OpTableBitmapViewBitmaps) BitmapOpTypeTable() OpType { return OpTypeViewOthers }

func LookupOpBitmapViewBitmaps(target ReadOnlyBitmap, name string) OpBitmapViewBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBitmaps(fn)
	}
	return nil
}

func LookupTableOpBitmapViewBitmaps(table OpTableBitmap, name string) OpBitmapViewBitmaps {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewBitmaps]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBitmaps)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapUpdateBitmaps, the Bitmap-specific
// form of OpTypeUpdateOthers.
type OpBitmapUpdateBitmaps func([]ReadOnlyBitmap) (bool, uint64, Bitmap)
func (OpBitmapUpdateBitmaps) BitmapOpType() { return OpTypeUpdateOthers }
type OpTableBitmapUpdateBitmaps map[string]opBitmapUpdateBitmaps
func (OpTableBitmapUpdateBitmaps) BitmapOpTypeTable() OpType { return OpTypeUpdateOthers }

func LookupOpBitmapUpdateBitmaps(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, uint64, Bitmap))
		return OpBitmapUpdateBitmaps(fn)
	}
	return nil
}

func LookupTableOpBitmapUpdateBitmaps(table OpTableBitmap, name string) OpBitmapUpdateBitmaps {
	if table == nil {
		return nil
	}
	subTable := table[BitmapUpdateBitmaps]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBitmaps)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for BitmapUpdateBytes, the Bitmap-specific
// form of OpTypeUpdateBytes.
type OpBitmapUpdateBytes func([]byte) (bool, uint64, Bitmap)
func (OpBitmapUpdateBytes) BitmapOpType() { return OpTypeUpdateBytes }
type OpTableBitmapUpdateBytes map[string]opBitmapUpdateBytes
func (OpTableBitmapUpdateBytes) BitmapOpTypeTable() OpType { return OpTypeUpdateBytes }

func LookupOpBitmapUpdateBytes(target ReadOnlyBitmap, name string) OpBitmapUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, uint64, Bitmap))
		return OpBitmapUpdateBytes(fn)
	}
	return nil
}

func LookupTableOpBitmapUpdateBytes(table OpTableBitmap, name string) OpBitmapUpdateBytes {
	if table == nil {
		return nil
	}
	subTable := table[BitmapUpdateBytes]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBytes)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// ImportRoaring performs a default BitmapUpdateBytes on a Bitmap
type interfaceBitmapHasImportRoaring interface {
	ImportRoaringUpdateBytes([]byte) (bool, uint64, Bitmap)
}
func ImportRoaring(target Bitmap, in1 []byte) (bool, uint64, Bitmap) {
	if target, ok := target.(interfaceBitmapHasImportRoaring); ok {
		return target.ImportRoaringUpdateBytes(in1)
	}
	return genericImportRoaring(target, in1)
}

// Implementation stuff for BitmapViewWriterGivesError, the Bitmap-specific
// form of OpTypeViewWriterGivesError.
type OpBitmapViewWriterGivesError func(io.Writer) error
func (OpBitmapViewWriterGivesError) BitmapOpType() { return OpTypeViewWriterGivesError }
type OpTableBitmapViewWriterGivesError map[string]opBitmapViewWriterGivesError
func (OpTableBitmapViewWriterGivesError) BitmapOpTypeTable() OpType { return OpTypeViewWriterGivesError }

func LookupOpBitmapViewWriterGivesError(target ReadOnlyBitmap, name string) OpBitmapViewWriterGivesError {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewWriterGivesError")
	if method.IsValid() {
		fn, _ := method.Interface().(func(io.Writer) error)
		return OpBitmapViewWriterGivesError(fn)
	}
	return nil
}

func LookupTableOpBitmapViewWriterGivesError(table OpTableBitmap, name string) OpBitmapViewWriterGivesError {
	if table == nil {
		return nil
	}
	subTable := table[BitmapViewWriterGivesError]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewWriterGivesError)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// ExportRoaring performs a default BitmapViewWriterGivesError on a Bitmap
type interfaceBitmapHasExportRoaring interface {
	ExportRoaringViewWriterGivesError(io.Writer) error
}
func ExportRoaring(target ReadOnlyBitmap, in1 io.Writer) error {
	if target, ok := target.(interfaceBitmapHasExportRoaring); ok {
		return target.ExportRoaringViewWriterGivesError(in1)
	}
	return genericExportRoaring(target, in1)
}
}
