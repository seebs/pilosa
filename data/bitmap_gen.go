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

// BitmapHasOpLookup indicates that you need per-item operation
// lookups. Implement this if, for instance, your implementation wraps
// another implementation and you do forwarding for arbitrary methods in
// some fancy way.
type BitmapHasOpLookup interface {
	OpLookup(OpType, string) OpFunctionBitmap
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeView, name)
		if method != nil {
			return method.(OpBitmapView)
		}
	}
	return OpLookupDirectBitmapView(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapView(target ReadOnlyBitmap, name string) OpBitmapView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, uint64, ReadOnlyBitmap))
		return OpBitmapView(fn)
	}
	return nil
}

// OpWrapBitmapView takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapView(wrapped OpBitmapView, fn func(inner func())) OpBitmapView {
	return func() (out1 bool, out2 uint64, out3 ReadOnlyBitmap) {
		inner := func() {
			out1, out2, out3 = wrapped()
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewGivesBool, name)
		if method != nil {
			return method.(OpBitmapViewGivesBool)
		}
	}
	return OpLookupDirectBitmapViewGivesBool(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func() bool)
		return OpBitmapViewGivesBool(fn)
	}
	return nil
}

// OpWrapBitmapViewGivesBool takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewGivesBool(wrapped OpBitmapViewGivesBool, fn func(inner func())) OpBitmapViewGivesBool {
	return func() (out1 bool) {
		inner := func() {
			out1 = wrapped()
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewGivesBit, name)
		if method != nil {
			return method.(OpBitmapViewGivesBit)
		}
	}
	return OpLookupDirectBitmapViewGivesBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func() uint64)
		return OpBitmapViewGivesBit(fn)
	}
	return nil
}

// OpWrapBitmapViewGivesBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewGivesBit(wrapped OpBitmapViewGivesBit, fn func(inner func())) OpBitmapViewGivesBit {
	return func() (out1 uint64) {
		inner := func() {
			out1 = wrapped()
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesBool, name)
		if method != nil {
			return method.(OpBitmapViewRangeGivesBool)
		}
	}
	return OpLookupDirectBitmapViewRangeGivesBool(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewRangeGivesBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) bool)
		return OpBitmapViewRangeGivesBool(fn)
	}
	return nil
}

// OpWrapBitmapViewRangeGivesBool takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewRangeGivesBool(wrapped OpBitmapViewRangeGivesBool, fn func(inner func())) OpBitmapViewRangeGivesBool {
	return func(in1 uint64, in2 uint64) (out1 bool) {
		inner := func() {
			out1 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesBit, name)
		if method != nil {
			return method.(OpBitmapViewRangeGivesBit)
		}
	}
	return OpLookupDirectBitmapViewRangeGivesBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewRangeGivesBit(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) uint64)
		return OpBitmapViewRangeGivesBit(fn)
	}
	return nil
}

// OpWrapBitmapViewRangeGivesBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewRangeGivesBit(wrapped OpBitmapViewRangeGivesBit, fn func(inner func())) OpBitmapViewRangeGivesBit {
	return func(in1 uint64, in2 uint64) (out1 uint64) {
		inner := func() {
			out1 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesOther, name)
		if method != nil {
			return method.(OpBitmapViewRangeGivesBitmap)
		}
	}
	return OpLookupDirectBitmapViewRangeGivesBitmap(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewRangeGivesBitmap(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) ReadOnlyBitmap)
		return OpBitmapViewRangeGivesBitmap(fn)
	}
	return nil
}

// OpWrapBitmapViewRangeGivesBitmap takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewRangeGivesBitmap(wrapped OpBitmapViewRangeGivesBitmap, fn func(inner func())) OpBitmapViewRangeGivesBitmap {
	return func(in1 uint64, in2 uint64) (out1 ReadOnlyBitmap) {
		inner := func() {
			out1 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesBitsBool, name)
		if method != nil {
			return method.(OpBitmapViewRangeGivesBitsBool)
		}
	}
	return OpLookupDirectBitmapViewRangeGivesBitsBool(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewRangeGivesBitsBool(target ReadOnlyBitmap, name string) OpBitmapViewRangeGivesBitsBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitsBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) ([]uint64, bool))
		return OpBitmapViewRangeGivesBitsBool(fn)
	}
	return nil
}

// OpWrapBitmapViewRangeGivesBitsBool takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewRangeGivesBitsBool(wrapped OpBitmapViewRangeGivesBitsBool, fn func(inner func())) OpBitmapViewRangeGivesBitsBool {
	return func(in1 uint64, in2 uint64) (out1 []uint64, out2 bool) {
		inner := func() {
			out1, out2 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdate, name)
		if method != nil {
			return method.(OpBitmapUpdate)
		}
	}
	return OpLookupDirectBitmapUpdate(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapUpdate(target ReadOnlyBitmap, name string) OpBitmapUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, uint64, Bitmap))
		return OpBitmapUpdate(fn)
	}
	return nil
}

// OpWrapBitmapUpdate takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapUpdate(wrapped OpBitmapUpdate, fn func(inner func())) OpBitmapUpdate {
	return func() (out1 bool, out2 uint64, out3 Bitmap) {
		inner := func() {
			out1, out2, out3 = wrapped()
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRange, name)
		if method != nil {
			return method.(OpBitmapViewRange)
		}
	}
	return OpLookupDirectBitmapViewRange(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewRange(target ReadOnlyBitmap, name string) OpBitmapViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64, uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewRange(fn)
	}
	return nil
}

// OpWrapBitmapViewRange takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewRange(wrapped OpBitmapViewRange, fn func(inner func())) OpBitmapViewRange {
	return func(in1 uint64, in2 uint64) (out1 bool, out2 uint64, out3 ReadOnlyBitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewBit, name)
		if method != nil {
			return method.(OpBitmapViewBit)
		}
	}
	return OpLookupDirectBitmapViewBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewBit(target ReadOnlyBitmap, name string) OpBitmapViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBit(fn)
	}
	return nil
}

// OpWrapBitmapViewBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewBit(wrapped OpBitmapViewBit, fn func(inner func())) OpBitmapViewBit {
	return func(in1 uint64) (out1 bool, out2 uint64, out3 ReadOnlyBitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateBit, name)
		if method != nil {
			return method.(OpBitmapUpdateBit)
		}
	}
	return OpLookupDirectBitmapUpdateBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapUpdateBit(target ReadOnlyBitmap, name string) OpBitmapUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint64) (bool, uint64, Bitmap))
		return OpBitmapUpdateBit(fn)
	}
	return nil
}

// OpWrapBitmapUpdateBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapUpdateBit(wrapped OpBitmapUpdateBit, fn func(inner func())) OpBitmapUpdateBit {
	return func(in1 uint64) (out1 bool, out2 uint64, out3 Bitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewOther, name)
		if method != nil {
			return method.(OpBitmapViewBitmap)
		}
	}
	return OpLookupDirectBitmapViewBitmap(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewBitmap(target ReadOnlyBitmap, name string) OpBitmapViewBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBitmap(fn)
	}
	return nil
}

// OpWrapBitmapViewBitmap takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewBitmap(wrapped OpBitmapViewBitmap, fn func(inner func())) OpBitmapViewBitmap {
	return func(in1 ReadOnlyBitmap) (out1 bool, out2 uint64, out3 ReadOnlyBitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateOther, name)
		if method != nil {
			return method.(OpBitmapUpdateBitmap)
		}
	}
	return OpLookupDirectBitmapUpdateBitmap(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapUpdateBitmap(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmap {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmap")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyBitmap) (bool, uint64, Bitmap))
		return OpBitmapUpdateBitmap(fn)
	}
	return nil
}

// OpWrapBitmapUpdateBitmap takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapUpdateBitmap(wrapped OpBitmapUpdateBitmap, fn func(inner func())) OpBitmapUpdateBitmap {
	return func(in1 ReadOnlyBitmap) (out1 bool, out2 uint64, out3 Bitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewBits, name)
		if method != nil {
			return method.(OpBitmapViewBits)
		}
	}
	return OpLookupDirectBitmapViewBits(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewBits(target ReadOnlyBitmap, name string) OpBitmapViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBits(fn)
	}
	return nil
}

// OpWrapBitmapViewBits takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewBits(wrapped OpBitmapViewBits, fn func(inner func())) OpBitmapViewBits {
	return func(in1 []uint64) (out1 bool, out2 uint64, out3 ReadOnlyBitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateBits, name)
		if method != nil {
			return method.(OpBitmapUpdateBits)
		}
	}
	return OpLookupDirectBitmapUpdateBits(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapUpdateBits(target ReadOnlyBitmap, name string) OpBitmapUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint64) (bool, uint64, Bitmap))
		return OpBitmapUpdateBits(fn)
	}
	return nil
}

// OpWrapBitmapUpdateBits takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapUpdateBits(wrapped OpBitmapUpdateBits, fn func(inner func())) OpBitmapUpdateBits {
	return func(in1 []uint64) (out1 bool, out2 uint64, out3 Bitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewOthers, name)
		if method != nil {
			return method.(OpBitmapViewBitmaps)
		}
	}
	return OpLookupDirectBitmapViewBitmaps(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewBitmaps(target ReadOnlyBitmap, name string) OpBitmapViewBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, uint64, ReadOnlyBitmap))
		return OpBitmapViewBitmaps(fn)
	}
	return nil
}

// OpWrapBitmapViewBitmaps takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewBitmaps(wrapped OpBitmapViewBitmaps, fn func(inner func())) OpBitmapViewBitmaps {
	return func(in1 []ReadOnlyBitmap) (out1 bool, out2 uint64, out3 ReadOnlyBitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateOthers, name)
		if method != nil {
			return method.(OpBitmapUpdateBitmaps)
		}
	}
	return OpLookupDirectBitmapUpdateBitmaps(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapUpdateBitmaps(target ReadOnlyBitmap, name string) OpBitmapUpdateBitmaps {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBitmaps")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyBitmap) (bool, uint64, Bitmap))
		return OpBitmapUpdateBitmaps(fn)
	}
	return nil
}

// OpWrapBitmapUpdateBitmaps takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapUpdateBitmaps(wrapped OpBitmapUpdateBitmaps, fn func(inner func())) OpBitmapUpdateBitmaps {
	return func(in1 []ReadOnlyBitmap) (out1 bool, out2 uint64, out3 Bitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateBytes, name)
		if method != nil {
			return method.(OpBitmapUpdateBytes)
		}
	}
	return OpLookupDirectBitmapUpdateBytes(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapUpdateBytes(target ReadOnlyBitmap, name string) OpBitmapUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, uint64, Bitmap))
		return OpBitmapUpdateBytes(fn)
	}
	return nil
}

// OpWrapBitmapUpdateBytes takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapUpdateBytes(wrapped OpBitmapUpdateBytes, fn func(inner func())) OpBitmapUpdateBytes {
	return func(in1 []byte) (out1 bool, out2 uint64, out3 Bitmap) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
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
	if target, ok := target.((BitmapHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewWriterGivesError, name)
		if method != nil {
			return method.(OpBitmapViewWriterGivesError)
		}
	}
	return OpLookupDirectBitmapViewWriterGivesError(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectBitmapViewWriterGivesError(target ReadOnlyBitmap, name string) OpBitmapViewWriterGivesError {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewWriterGivesError")
	if method.IsValid() {
		fn, _ := method.Interface().(func(io.Writer) error)
		return OpBitmapViewWriterGivesError(fn)
	}
	return nil
}

// OpWrapBitmapViewWriterGivesError takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapBitmapViewWriterGivesError(wrapped OpBitmapViewWriterGivesError, fn func(inner func())) OpBitmapViewWriterGivesError {
	return func(in1 io.Writer) (out1 error) {
		inner := func() {
			out1 = wrapped(in1)
		}
		fn(inner)
		return
	}
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

// OpBitmapLookupGeneric is a generic lookup function which will
// use any provided OpLookup functionality of its target, falling back on
// the default name-based lookup.
func OpBitmapLookupGeneric(target ReadOnlyBitmap, typ OpType, name string) OpFunctionBitmap {
	switch typ {
	case OpTypeView:
		return OpLookupBitmapView(target, name)
	case OpTypeViewGivesBool:
		return OpLookupBitmapViewGivesBool(target, name)
	case OpTypeViewGivesBit:
		return OpLookupBitmapViewGivesBit(target, name)
	case OpTypeViewRangeGivesBool:
		return OpLookupBitmapViewRangeGivesBool(target, name)
	case OpTypeViewRangeGivesBit:
		return OpLookupBitmapViewRangeGivesBit(target, name)
	case OpTypeViewRangeGivesOther:
		return OpLookupBitmapViewRangeGivesBitmap(target, name)
	case OpTypeViewRangeGivesBitsBool:
		return OpLookupBitmapViewRangeGivesBitsBool(target, name)
	case OpTypeUpdate:
		return OpLookupBitmapUpdate(target, name)
	case OpTypeViewRange:
		return OpLookupBitmapViewRange(target, name)
	case OpTypeViewBit:
		return OpLookupBitmapViewBit(target, name)
	case OpTypeUpdateBit:
		return OpLookupBitmapUpdateBit(target, name)
	case OpTypeViewOther:
		return OpLookupBitmapViewBitmap(target, name)
	case OpTypeUpdateOther:
		return OpLookupBitmapUpdateBitmap(target, name)
	case OpTypeViewBits:
		return OpLookupBitmapViewBits(target, name)
	case OpTypeUpdateBits:
		return OpLookupBitmapUpdateBits(target, name)
	case OpTypeViewOthers:
		return OpLookupBitmapViewBitmaps(target, name)
	case OpTypeUpdateOthers:
		return OpLookupBitmapUpdateBitmaps(target, name)
	case OpTypeUpdateBytes:
		return OpLookupBitmapUpdateBytes(target, name)
	case OpTypeViewWriterGivesError:
		return OpLookupBitmapViewWriterGivesError(target, name)
	}
	return nil
}

// OpBitmapLookupGenericDirect is a generic lookup function which ignores
// any OpLookup functionality of the target and just does the reflect stuff.
func OpBitmapLookupGenericDirect(target ReadOnlyBitmap, typ OpType, name string) OpFunctionBitmap {
	switch typ {
	case OpTypeView:
		return OpLookupDirectBitmapView(target, name)
	case OpTypeViewGivesBool:
		return OpLookupDirectBitmapViewGivesBool(target, name)
	case OpTypeViewGivesBit:
		return OpLookupDirectBitmapViewGivesBit(target, name)
	case OpTypeViewRangeGivesBool:
		return OpLookupDirectBitmapViewRangeGivesBool(target, name)
	case OpTypeViewRangeGivesBit:
		return OpLookupDirectBitmapViewRangeGivesBit(target, name)
	case OpTypeViewRangeGivesOther:
		return OpLookupDirectBitmapViewRangeGivesBitmap(target, name)
	case OpTypeViewRangeGivesBitsBool:
		return OpLookupDirectBitmapViewRangeGivesBitsBool(target, name)
	case OpTypeUpdate:
		return OpLookupDirectBitmapUpdate(target, name)
	case OpTypeViewRange:
		return OpLookupDirectBitmapViewRange(target, name)
	case OpTypeViewBit:
		return OpLookupDirectBitmapViewBit(target, name)
	case OpTypeUpdateBit:
		return OpLookupDirectBitmapUpdateBit(target, name)
	case OpTypeViewOther:
		return OpLookupDirectBitmapViewBitmap(target, name)
	case OpTypeUpdateOther:
		return OpLookupDirectBitmapUpdateBitmap(target, name)
	case OpTypeViewBits:
		return OpLookupDirectBitmapViewBits(target, name)
	case OpTypeUpdateBits:
		return OpLookupDirectBitmapUpdateBits(target, name)
	case OpTypeViewOthers:
		return OpLookupDirectBitmapViewBitmaps(target, name)
	case OpTypeUpdateOthers:
		return OpLookupDirectBitmapUpdateBitmaps(target, name)
	case OpTypeUpdateBytes:
		return OpLookupDirectBitmapUpdateBytes(target, name)
	case OpTypeViewWriterGivesError:
		return OpLookupDirectBitmapViewWriterGivesError(target, name)
	}
	return nil
}
