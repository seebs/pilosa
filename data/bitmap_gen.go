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
type OpTableBitmapGeneric interface {
	BitmapOpTypeTable() OpType
}

// OpTableBitmap is a slice mapping optypes to map[string]opFunc,
// where any specific map will actually be a map with a concrete type of
// op function. We defined the
type OpTableBitmap []OpTableBitmapGeneric

// Implementation stuff for OpBitmapView, the Bitmap-specific
// form of OpTypeView.
type OpBitmapView func() (bool, uint64, ReadOnlyBitmap)

func (OpBitmapView) BitmapOpType() OpType {
	return OpTypeView
}

type OpTableBitmapView map[string]OpBitmapView

func (OpTableBitmapView) BitmapOpTypeTable() OpType {
	return OpTypeView
}

func OpLookupBitmapView(target ReadOnlyBitmap, name string) OpBitmapView {
	method := target.OpLookup(OpTypeView, name)
	if method != nil {
		return method.(OpBitmapView)
	}
	return nil
}

func OpLookupGenericBitmapView(target ReadOnlyBitmap, name string) OpBitmapView {
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
	subTable := table[OpTypeView]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapView)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapViewGivesBool, the Bitmap-specific
// form of OpTypeViewGivesBool.
type OpBitmapViewGivesBool func() bool

func (OpBitmapViewGivesBool) BitmapOpType() OpType {
	return OpTypeViewGivesBool
}

type OpTableBitmapViewGivesBool map[string]OpBitmapViewGivesBool

func (OpTableBitmapViewGivesBool) BitmapOpTypeTable() OpType {
	return OpTypeViewGivesBool
}

func OpLookupBitmapViewGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewGivesBool {
	method := target.OpLookup(OpTypeViewGivesBool, name)
	if method != nil {
		return method.(OpBitmapViewGivesBool)
	}
	return nil
}

func OpLookupGenericBitmapViewGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewGivesBool {
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
	subTable := table[OpTypeViewGivesBool]
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

// Implementation stuff for OpBitmapViewGivesBit, the Bitmap-specific
// form of OpTypeViewGivesBit.
type OpBitmapViewGivesBit func() uint64

func (OpBitmapViewGivesBit) BitmapOpType() OpType {
	return OpTypeViewGivesBit
}

type OpTableBitmapViewGivesBit map[string]OpBitmapViewGivesBit

func (OpTableBitmapViewGivesBit) BitmapOpTypeTable() OpType {
	return OpTypeViewGivesBit
}

func OpLookupBitmapViewGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewGivesBit {
	method := target.OpLookup(OpTypeViewGivesBit, name)
	if method != nil {
		return method.(OpBitmapViewGivesBit)
	}
	return nil
}

func OpLookupGenericBitmapViewGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewGivesBit {
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
	subTable := table[OpTypeViewGivesBit]
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

// Implementation stuff for OpBitmapViewRangeGivesBool, the Bitmap-specific
// form of OpTypeViewRangeGivesBool.
type OpBitmapViewRangeGivesBool func(uint64, uint64) bool

func (OpBitmapViewRangeGivesBool) BitmapOpType() OpType {
	return OpTypeViewRangeGivesBool
}

type OpTableBitmapViewRangeGivesBool map[string]OpBitmapViewRangeGivesBool

func (OpTableBitmapViewRangeGivesBool) BitmapOpTypeTable() OpType {
	return OpTypeViewRangeGivesBool
}

func OpLookupBitmapViewRangeGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBool {
	method := target.OpLookup(OpTypeViewRangeGivesBool, name)
	if method != nil {
		return method.(OpBitmapViewRangeGivesBool)
	}
	return nil
}

func OpLookupGenericBitmapViewRangeGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBool {
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
	subTable := table[OpTypeViewRangeGivesBool]
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

// Implementation stuff for OpBitmapViewRangeGivesBit, the Bitmap-specific
// form of OpTypeViewRangeGivesBit.
type OpBitmapViewRangeGivesBit func(uint64, uint64) uint64

func (OpBitmapViewRangeGivesBit) BitmapOpType() OpType {
	return OpTypeViewRangeGivesBit
}

type OpTableBitmapViewRangeGivesBit map[string]OpBitmapViewRangeGivesBit

func (OpTableBitmapViewRangeGivesBit) BitmapOpTypeTable() OpType {
	return OpTypeViewRangeGivesBit
}

func OpLookupBitmapViewRangeGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBit {
	method := target.OpLookup(OpTypeViewRangeGivesBit, name)
	if method != nil {
		return method.(OpBitmapViewRangeGivesBit)
	}
	return nil
}

func OpLookupGenericBitmapViewRangeGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBit {
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
	subTable := table[OpTypeViewRangeGivesBit]
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

// Implementation stuff for OpBitmapViewRangeGivesBitmap, the Bitmap-specific
// form of OpTypeViewRangeGivesOther.
type OpBitmapViewRangeGivesBitmap func(uint64, uint64) ReadOnlyBitmap

func (OpBitmapViewRangeGivesBitmap) BitmapOpType() OpType {
	return OpTypeViewRangeGivesOther
}

type OpTableBitmapViewRangeGivesBitmap map[string]OpBitmapViewRangeGivesBitmap

func (OpTableBitmapViewRangeGivesBitmap) BitmapOpTypeTable() OpType {
	return OpTypeViewRangeGivesOther
}

func OpLookupBitmapViewRangeGivesBitmap(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitmap {
	method := target.OpLookup(OpTypeViewRangeGivesOther, name)
	if method != nil {
		return method.(OpBitmapViewRangeGivesBitmap)
	}
	return nil
}

func OpLookupGenericBitmapViewRangeGivesBitmap(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitmap {
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
	subTable := table[OpTypeViewRangeGivesOther]
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

// Implementation stuff for OpBitmapViewRangeGivesBitsBool, the Bitmap-specific
// form of OpTypeViewRangeGivesBitsBool.
type OpBitmapViewRangeGivesBitsBool func(uint64, uint64) ([]uint64, bool)

func (OpBitmapViewRangeGivesBitsBool) BitmapOpType() OpType {
	return OpTypeViewRangeGivesBitsBool
}

type OpTableBitmapViewRangeGivesBitsBool map[string]OpBitmapViewRangeGivesBitsBool

func (OpTableBitmapViewRangeGivesBitsBool) BitmapOpTypeTable() OpType {
	return OpTypeViewRangeGivesBitsBool
}

func OpLookupBitmapViewRangeGivesBitsBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitsBool {
	method := target.OpLookup(OpTypeViewRangeGivesBitsBool, name)
	if method != nil {
		return method.(OpBitmapViewRangeGivesBitsBool)
	}
	return nil
}

func OpLookupGenericBitmapViewRangeGivesBitsBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitsBool {
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
	subTable := table[OpTypeViewRangeGivesBitsBool]
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

// Implementation stuff for OpBitmapUpdate, the Bitmap-specific
// form of OpTypeUpdate.
type OpBitmapUpdate func() (bool, uint64, Bitmap)

func (OpBitmapUpdate) BitmapOpType() OpType {
	return OpTypeUpdate
}

type OpTableBitmapUpdate map[string]OpBitmapUpdate

func (OpTableBitmapUpdate) BitmapOpTypeTable() OpType {
	return OpTypeUpdate
}

func OpLookupBitmapUpdate(target ReadOnlyBitmap, name string) OpBitmapUpdate {
	method := target.OpLookup(OpTypeUpdate, name)
	if method != nil {
		return method.(OpBitmapUpdate)
	}
	return nil
}

func OpLookupGenericBitmapUpdate(target ReadOnlyBitmap, name string) OpBitmapUpdate {
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
	subTable := table[OpTypeUpdate]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdate)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapViewRange, the Bitmap-specific
// form of OpTypeViewRange.
type OpBitmapViewRange func(uint64, uint64) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewRange) BitmapOpType() OpType {
	return OpTypeViewRange
}

type OpTableBitmapViewRange map[string]OpBitmapViewRange

func (OpTableBitmapViewRange) BitmapOpTypeTable() OpType {
	return OpTypeViewRange
}

func OpLookupBitmapViewRange(target ReadOnlyBitmap, name string) OpBitmapViewRange {
	method := target.OpLookup(OpTypeViewRange, name)
	if method != nil {
		return method.(OpBitmapViewRange)
	}
	return nil
}

func OpLookupGenericBitmapViewRange(target ReadOnlyBitmap, name string) OpBitmapViewRange {
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
	subTable := table[OpTypeViewRange]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewRange)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapViewBit, the Bitmap-specific
// form of OpTypeViewBit.
type OpBitmapViewBit func(uint64) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBit) BitmapOpType() OpType {
	return OpTypeViewBit
}

type OpTableBitmapViewBit map[string]OpBitmapViewBit

func (OpTableBitmapViewBit) BitmapOpTypeTable() OpType {
	return OpTypeViewBit
}

func OpLookupBitmapViewBit(target ReadOnlyBitmap, name string) OpBitmapViewBit {
	method := target.OpLookup(OpTypeViewBit, name)
	if method != nil {
		return method.(OpBitmapViewBit)
	}
	return nil
}

func OpLookupGenericBitmapViewBit(target ReadOnlyBitmap, name string) OpBitmapViewBit {
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
	subTable := table[OpTypeViewBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapUpdateBit, the Bitmap-specific
// form of OpTypeUpdateBit.
type OpBitmapUpdateBit func(uint64) (bool, uint64, Bitmap)

func (OpBitmapUpdateBit) BitmapOpType() OpType {
	return OpTypeUpdateBit
}

type OpTableBitmapUpdateBit map[string]OpBitmapUpdateBit

func (OpTableBitmapUpdateBit) BitmapOpTypeTable() OpType {
	return OpTypeUpdateBit
}

func OpLookupBitmapUpdateBit(target ReadOnlyBitmap, name string) OpBitmapUpdateBit {
	method := target.OpLookup(OpTypeUpdateBit, name)
	if method != nil {
		return method.(OpBitmapUpdateBit)
	}
	return nil
}

func OpLookupGenericBitmapUpdateBit(target ReadOnlyBitmap, name string) OpBitmapUpdateBit {
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
	subTable := table[OpTypeUpdateBit]
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

// Implementation stuff for OpBitmapViewBitmap, the Bitmap-specific
// form of OpTypeViewOther.
type OpBitmapViewBitmap func(ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBitmap) BitmapOpType() OpType {
	return OpTypeViewOther
}

type OpTableBitmapViewBitmap map[string]OpBitmapViewBitmap

func (OpTableBitmapViewBitmap) BitmapOpTypeTable() OpType {
	return OpTypeViewOther
}

func OpLookupBitmapViewBitmap(target ReadOnlyBitmap, name string) OpBitmapViewBitmap {
	method := target.OpLookup(OpTypeViewOther, name)
	if method != nil {
		return method.(OpBitmapViewBitmap)
	}
	return nil
}

func OpLookupGenericBitmapViewBitmap(target ReadOnlyBitmap, name string) OpBitmapViewBitmap {
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
	subTable := table[OpTypeViewOther]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBitmap)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapUpdateBitmap, the Bitmap-specific
// form of OpTypeUpdateOther.
type OpBitmapUpdateBitmap func(ReadOnlyBitmap) (bool, uint64, Bitmap)

func (OpBitmapUpdateBitmap) BitmapOpType() OpType {
	return OpTypeUpdateOther
}

type OpTableBitmapUpdateBitmap map[string]OpBitmapUpdateBitmap

func (OpTableBitmapUpdateBitmap) BitmapOpTypeTable() OpType {
	return OpTypeUpdateOther
}

func OpLookupBitmapUpdateBitmap(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmap {
	method := target.OpLookup(OpTypeUpdateOther, name)
	if method != nil {
		return method.(OpBitmapUpdateBitmap)
	}
	return nil
}

func OpLookupGenericBitmapUpdateBitmap(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmap {
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
	subTable := table[OpTypeUpdateOther]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBitmap)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapViewBits, the Bitmap-specific
// form of OpTypeViewBits.
type OpBitmapViewBits func([]uint64) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBits) BitmapOpType() OpType {
	return OpTypeViewBits
}

type OpTableBitmapViewBits map[string]OpBitmapViewBits

func (OpTableBitmapViewBits) BitmapOpTypeTable() OpType {
	return OpTypeViewBits
}

func OpLookupBitmapViewBits(target ReadOnlyBitmap, name string) OpBitmapViewBits {
	method := target.OpLookup(OpTypeViewBits, name)
	if method != nil {
		return method.(OpBitmapViewBits)
	}
	return nil
}

func OpLookupGenericBitmapViewBits(target ReadOnlyBitmap, name string) OpBitmapViewBits {
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
	subTable := table[OpTypeViewBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapUpdateBits, the Bitmap-specific
// form of OpTypeUpdateBits.
type OpBitmapUpdateBits func([]uint64) (bool, uint64, Bitmap)

func (OpBitmapUpdateBits) BitmapOpType() OpType {
	return OpTypeUpdateBits
}

type OpTableBitmapUpdateBits map[string]OpBitmapUpdateBits

func (OpTableBitmapUpdateBits) BitmapOpTypeTable() OpType {
	return OpTypeUpdateBits
}

func OpLookupBitmapUpdateBits(target ReadOnlyBitmap, name string) OpBitmapUpdateBits {
	method := target.OpLookup(OpTypeUpdateBits, name)
	if method != nil {
		return method.(OpBitmapUpdateBits)
	}
	return nil
}

func OpLookupGenericBitmapUpdateBits(target ReadOnlyBitmap, name string) OpBitmapUpdateBits {
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
	subTable := table[OpTypeUpdateBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapViewBitmaps, the Bitmap-specific
// form of OpTypeViewOthers.
type OpBitmapViewBitmaps func([]ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap)

func (OpBitmapViewBitmaps) BitmapOpType() OpType {
	return OpTypeViewOthers
}

type OpTableBitmapViewBitmaps map[string]OpBitmapViewBitmaps

func (OpTableBitmapViewBitmaps) BitmapOpTypeTable() OpType {
	return OpTypeViewOthers
}

func OpLookupBitmapViewBitmaps(target ReadOnlyBitmap, name string) OpBitmapViewBitmaps {
	method := target.OpLookup(OpTypeViewOthers, name)
	if method != nil {
		return method.(OpBitmapViewBitmaps)
	}
	return nil
}

func OpLookupGenericBitmapViewBitmaps(target ReadOnlyBitmap, name string) OpBitmapViewBitmaps {
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
	subTable := table[OpTypeViewOthers]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapViewBitmaps)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapUpdateBitmaps, the Bitmap-specific
// form of OpTypeUpdateOthers.
type OpBitmapUpdateBitmaps func([]ReadOnlyBitmap) (bool, uint64, Bitmap)

func (OpBitmapUpdateBitmaps) BitmapOpType() OpType {
	return OpTypeUpdateOthers
}

type OpTableBitmapUpdateBitmaps map[string]OpBitmapUpdateBitmaps

func (OpTableBitmapUpdateBitmaps) BitmapOpTypeTable() OpType {
	return OpTypeUpdateOthers
}

func OpLookupBitmapUpdateBitmaps(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmaps {
	method := target.OpLookup(OpTypeUpdateOthers, name)
	if method != nil {
		return method.(OpBitmapUpdateBitmaps)
	}
	return nil
}

func OpLookupGenericBitmapUpdateBitmaps(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmaps {
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
	subTable := table[OpTypeUpdateOthers]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableBitmapUpdateBitmaps)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpBitmapUpdateBytes, the Bitmap-specific
// form of OpTypeUpdateBytes.
type OpBitmapUpdateBytes func([]byte) (bool, uint64, Bitmap)

func (OpBitmapUpdateBytes) BitmapOpType() OpType {
	return OpTypeUpdateBytes
}

type OpTableBitmapUpdateBytes map[string]OpBitmapUpdateBytes

func (OpTableBitmapUpdateBytes) BitmapOpTypeTable() OpType {
	return OpTypeUpdateBytes
}

func OpLookupBitmapUpdateBytes(target ReadOnlyBitmap, name string) OpBitmapUpdateBytes {
	method := target.OpLookup(OpTypeUpdateBytes, name)
	if method != nil {
		return method.(OpBitmapUpdateBytes)
	}
	return nil
}

func OpLookupGenericBitmapUpdateBytes(target ReadOnlyBitmap, name string) OpBitmapUpdateBytes {
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
	subTable := table[OpTypeUpdateBytes]
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

// Implementation stuff for OpBitmapViewWriterGivesError, the Bitmap-specific
// form of OpTypeViewWriterGivesError.
type OpBitmapViewWriterGivesError func(io.Writer) error

func (OpBitmapViewWriterGivesError) BitmapOpType() OpType {
	return OpTypeViewWriterGivesError
}

type OpTableBitmapViewWriterGivesError map[string]OpBitmapViewWriterGivesError

func (OpTableBitmapViewWriterGivesError) BitmapOpTypeTable() OpType {
	return OpTypeViewWriterGivesError
}

func OpLookupBitmapViewWriterGivesError(target ReadOnlyBitmap, name string) OpBitmapViewWriterGivesError {
	method := target.OpLookup(OpTypeViewWriterGivesError, name)
	if method != nil {
		return method.(OpBitmapViewWriterGivesError)
	}
	return nil
}

func OpLookupGenericBitmapViewWriterGivesError(target ReadOnlyBitmap, name string) OpBitmapViewWriterGivesError {
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
	subTable := table[OpTypeViewWriterGivesError]
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

// OpBitmapLookupGeneric is a generic lookup function which just looks things up by
// name, using code-generation and a naming convention.
func OpBitmapLookupGeneric(target ReadOnlyBitmap, typ OpType, name string) OpFunctionBitmap {
	switch typ {
	case OpTypeView:
		return OpLookupGenericBitmapView(target, name)
	case OpTypeViewGivesBool:
		return OpLookupGenericBitmapViewGivesBool(target, name)
	case OpTypeViewGivesBit:
		return OpLookupGenericBitmapViewGivesBit(target, name)
	case OpTypeViewRangeGivesBool:
		return OpLookupGenericBitmapViewRangeGivesBool(target, name)
	case OpTypeViewRangeGivesBit:
		return OpLookupGenericBitmapViewRangeGivesBit(target, name)
	case OpTypeViewRangeGivesOther:
		return OpLookupGenericBitmapViewRangeGivesBitmap(target, name)
	case OpTypeViewRangeGivesBitsBool:
		return OpLookupGenericBitmapViewRangeGivesBitsBool(target, name)
	case OpTypeUpdate:
		return OpLookupGenericBitmapUpdate(target, name)
	case OpTypeViewRange:
		return OpLookupGenericBitmapViewRange(target, name)
	case OpTypeViewBit:
		return OpLookupGenericBitmapViewBit(target, name)
	case OpTypeUpdateBit:
		return OpLookupGenericBitmapUpdateBit(target, name)
	case OpTypeViewOther:
		return OpLookupGenericBitmapViewBitmap(target, name)
	case OpTypeUpdateOther:
		return OpLookupGenericBitmapUpdateBitmap(target, name)
	case OpTypeViewBits:
		return OpLookupGenericBitmapViewBits(target, name)
	case OpTypeUpdateBits:
		return OpLookupGenericBitmapUpdateBits(target, name)
	case OpTypeViewOthers:
		return OpLookupGenericBitmapViewBitmaps(target, name)
	case OpTypeUpdateOthers:
		return OpLookupGenericBitmapUpdateBitmaps(target, name)
	case OpTypeUpdateBytes:
		return OpLookupGenericBitmapUpdateBytes(target, name)
	case OpTypeViewWriterGivesError:
		return OpLookupGenericBitmapViewWriterGivesError(target, name)
	}
	return nil
}
