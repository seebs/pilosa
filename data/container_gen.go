package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Container (see gen/main.go). These are expressed
// as method signatures -- the Container they operate on is an implicit
// receiver not shown in the signature.

import (
	"io"
	"reflect"
)

// OpFunctionContainer exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionContainer interface {
	ContainerOpType() OpType
}

// OpTableGenericContainer, similarly, lets us specify a range of types -- in
// this case, map[string]OpFunctionType, where the type is one of the
// specific op function types defined.
type OpTableContainerGeneric interface {
	ContainerOpTypeTable() OpType
}

// ContainerHasOpLookup indicates that you need per-item operation
// lookups. Implement this if, for instance, your implementation wraps
// another implementation and you do forwarding for arbitrary methods in
// some fancy way.
type ContainerHasOpLookup interface {
	OpLookup(OpType, string) OpFunctionContainer
}

// OpTableContainer is a slice mapping optypes to map[string]opFunc,
// where any specific map will actually be a map with a concrete type of
// op function. We defined the
type OpTableContainer []OpTableContainerGeneric

// Implementation stuff for OpContainerView, the Container-specific
// form of OpTypeView.
type OpContainerView func() (bool, int, ReadOnlyContainer)

func (OpContainerView) ContainerOpType() OpType {
	return OpTypeView
}

type OpTableContainerView map[string]OpContainerView

func (OpTableContainerView) ContainerOpTypeTable() OpType {
	return OpTypeView
}

func OpLookupContainerView(target ReadOnlyContainer, name string) OpContainerView {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeView, name)
		if method != nil {
			return method.(OpContainerView)
		}
	}
	return OpLookupDirectContainerView(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerView(target ReadOnlyContainer, name string) OpContainerView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, ReadOnlyContainer))
		return OpContainerView(fn)
	}
	return nil
}

// OpWrapContainerView takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerView(wrapped OpContainerView, fn func(inner func())) OpContainerView {
	return func() (out1 bool, out2 int, out3 ReadOnlyContainer) {
		inner := func() {
			out1, out2, out3 = wrapped()
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerView(table OpTableContainer, name string) OpContainerView {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeView]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerView)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewGivesBool, the Container-specific
// form of OpTypeViewGivesBool.
type OpContainerViewGivesBool func() bool

func (OpContainerViewGivesBool) ContainerOpType() OpType {
	return OpTypeViewGivesBool
}

type OpTableContainerViewGivesBool map[string]OpContainerViewGivesBool

func (OpTableContainerViewGivesBool) ContainerOpTypeTable() OpType {
	return OpTypeViewGivesBool
}

func OpLookupContainerViewGivesBool(target ReadOnlyContainer, name string) OpContainerViewGivesBool {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewGivesBool, name)
		if method != nil {
			return method.(OpContainerViewGivesBool)
		}
	}
	return OpLookupDirectContainerViewGivesBool(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewGivesBool(target ReadOnlyContainer, name string) OpContainerViewGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func() bool)
		return OpContainerViewGivesBool(fn)
	}
	return nil
}

// OpWrapContainerViewGivesBool takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewGivesBool(wrapped OpContainerViewGivesBool, fn func(inner func())) OpContainerViewGivesBool {
	return func() (out1 bool) {
		inner := func() {
			out1 = wrapped()
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewGivesBool(table OpTableContainer, name string) OpContainerViewGivesBool {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewGivesBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewGivesBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Any performs a default ContainerViewGivesBool on a Container
type interfaceContainerHasAny interface {
	AnyViewGivesBool() bool
}

func ContainerAny(target ReadOnlyContainer) bool {
	if target, ok := target.(interfaceContainerHasAny); ok {
		return target.AnyViewGivesBool()
	}
	return genericContainerAny(target)
}

// Implementation stuff for OpContainerViewGivesBit, the Container-specific
// form of OpTypeViewGivesBit.
type OpContainerViewGivesBit func() int

func (OpContainerViewGivesBit) ContainerOpType() OpType {
	return OpTypeViewGivesBit
}

type OpTableContainerViewGivesBit map[string]OpContainerViewGivesBit

func (OpTableContainerViewGivesBit) ContainerOpTypeTable() OpType {
	return OpTypeViewGivesBit
}

func OpLookupContainerViewGivesBit(target ReadOnlyContainer, name string) OpContainerViewGivesBit {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewGivesBit, name)
		if method != nil {
			return method.(OpContainerViewGivesBit)
		}
	}
	return OpLookupDirectContainerViewGivesBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewGivesBit(target ReadOnlyContainer, name string) OpContainerViewGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func() int)
		return OpContainerViewGivesBit(fn)
	}
	return nil
}

// OpWrapContainerViewGivesBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewGivesBit(wrapped OpContainerViewGivesBit, fn func(inner func())) OpContainerViewGivesBit {
	return func() (out1 int) {
		inner := func() {
			out1 = wrapped()
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewGivesBit(table OpTableContainer, name string) OpContainerViewGivesBit {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewGivesBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewGivesBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Count performs a default ContainerViewGivesBit on a Container
type interfaceContainerHasCount interface {
	CountViewGivesBit() int
}

func ContainerCount(target ReadOnlyContainer) int {
	if target, ok := target.(interfaceContainerHasCount); ok {
		return target.CountViewGivesBit()
	}
	return genericContainerCount(target)
}

// Implementation stuff for OpContainerViewRangeGivesBool, the Container-specific
// form of OpTypeViewRangeGivesBool.
type OpContainerViewRangeGivesBool func(uint16, uint16) bool

func (OpContainerViewRangeGivesBool) ContainerOpType() OpType {
	return OpTypeViewRangeGivesBool
}

type OpTableContainerViewRangeGivesBool map[string]OpContainerViewRangeGivesBool

func (OpTableContainerViewRangeGivesBool) ContainerOpTypeTable() OpType {
	return OpTypeViewRangeGivesBool
}

func OpLookupContainerViewRangeGivesBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBool {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesBool, name)
		if method != nil {
			return method.(OpContainerViewRangeGivesBool)
		}
	}
	return OpLookupDirectContainerViewRangeGivesBool(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewRangeGivesBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) bool)
		return OpContainerViewRangeGivesBool(fn)
	}
	return nil
}

// OpWrapContainerViewRangeGivesBool takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewRangeGivesBool(wrapped OpContainerViewRangeGivesBool, fn func(inner func())) OpContainerViewRangeGivesBool {
	return func(in1 uint16, in2 uint16) (out1 bool) {
		inner := func() {
			out1 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewRangeGivesBool(table OpTableContainer, name string) OpContainerViewRangeGivesBool {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewRangeGivesBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewRangeGivesBit, the Container-specific
// form of OpTypeViewRangeGivesBit.
type OpContainerViewRangeGivesBit func(uint16, uint16) int

func (OpContainerViewRangeGivesBit) ContainerOpType() OpType {
	return OpTypeViewRangeGivesBit
}

type OpTableContainerViewRangeGivesBit map[string]OpContainerViewRangeGivesBit

func (OpTableContainerViewRangeGivesBit) ContainerOpTypeTable() OpType {
	return OpTypeViewRangeGivesBit
}

func OpLookupContainerViewRangeGivesBit(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBit {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesBit, name)
		if method != nil {
			return method.(OpContainerViewRangeGivesBit)
		}
	}
	return OpLookupDirectContainerViewRangeGivesBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewRangeGivesBit(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) int)
		return OpContainerViewRangeGivesBit(fn)
	}
	return nil
}

// OpWrapContainerViewRangeGivesBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewRangeGivesBit(wrapped OpContainerViewRangeGivesBit, fn func(inner func())) OpContainerViewRangeGivesBit {
	return func(in1 uint16, in2 uint16) (out1 int) {
		inner := func() {
			out1 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewRangeGivesBit(table OpTableContainer, name string) OpContainerViewRangeGivesBit {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewRangeGivesBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewRangeGivesContainer, the Container-specific
// form of OpTypeViewRangeGivesOther.
type OpContainerViewRangeGivesContainer func(uint16, uint16) ReadOnlyContainer

func (OpContainerViewRangeGivesContainer) ContainerOpType() OpType {
	return OpTypeViewRangeGivesOther
}

type OpTableContainerViewRangeGivesContainer map[string]OpContainerViewRangeGivesContainer

func (OpTableContainerViewRangeGivesContainer) ContainerOpTypeTable() OpType {
	return OpTypeViewRangeGivesOther
}

func OpLookupContainerViewRangeGivesContainer(target ReadOnlyContainer, name string) OpContainerViewRangeGivesContainer {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesOther, name)
		if method != nil {
			return method.(OpContainerViewRangeGivesContainer)
		}
	}
	return OpLookupDirectContainerViewRangeGivesContainer(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewRangeGivesContainer(target ReadOnlyContainer, name string) OpContainerViewRangeGivesContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) ReadOnlyContainer)
		return OpContainerViewRangeGivesContainer(fn)
	}
	return nil
}

// OpWrapContainerViewRangeGivesContainer takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewRangeGivesContainer(wrapped OpContainerViewRangeGivesContainer, fn func(inner func())) OpContainerViewRangeGivesContainer {
	return func(in1 uint16, in2 uint16) (out1 ReadOnlyContainer) {
		inner := func() {
			out1 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewRangeGivesContainer(table OpTableContainer, name string) OpContainerViewRangeGivesContainer {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewRangeGivesOther]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesContainer)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewRangeGivesBitsBool, the Container-specific
// form of OpTypeViewRangeGivesBitsBool.
type OpContainerViewRangeGivesBitsBool func(uint16, uint16) ([]int, bool)

func (OpContainerViewRangeGivesBitsBool) ContainerOpType() OpType {
	return OpTypeViewRangeGivesBitsBool
}

type OpTableContainerViewRangeGivesBitsBool map[string]OpContainerViewRangeGivesBitsBool

func (OpTableContainerViewRangeGivesBitsBool) ContainerOpTypeTable() OpType {
	return OpTypeViewRangeGivesBitsBool
}

func OpLookupContainerViewRangeGivesBitsBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBitsBool {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRangeGivesBitsBool, name)
		if method != nil {
			return method.(OpContainerViewRangeGivesBitsBool)
		}
	}
	return OpLookupDirectContainerViewRangeGivesBitsBool(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewRangeGivesBitsBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBitsBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitsBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) ([]int, bool))
		return OpContainerViewRangeGivesBitsBool(fn)
	}
	return nil
}

// OpWrapContainerViewRangeGivesBitsBool takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewRangeGivesBitsBool(wrapped OpContainerViewRangeGivesBitsBool, fn func(inner func())) OpContainerViewRangeGivesBitsBool {
	return func(in1 uint16, in2 uint16) (out1 []int, out2 bool) {
		inner := func() {
			out1, out2 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewRangeGivesBitsBool(table OpTableContainer, name string) OpContainerViewRangeGivesBitsBool {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewRangeGivesBitsBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesBitsBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerUpdate, the Container-specific
// form of OpTypeUpdate.
type OpContainerUpdate func() (bool, int, Container)

func (OpContainerUpdate) ContainerOpType() OpType {
	return OpTypeUpdate
}

type OpTableContainerUpdate map[string]OpContainerUpdate

func (OpTableContainerUpdate) ContainerOpTypeTable() OpType {
	return OpTypeUpdate
}

func OpLookupContainerUpdate(target ReadOnlyContainer, name string) OpContainerUpdate {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdate, name)
		if method != nil {
			return method.(OpContainerUpdate)
		}
	}
	return OpLookupDirectContainerUpdate(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerUpdate(target ReadOnlyContainer, name string) OpContainerUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, Container))
		return OpContainerUpdate(fn)
	}
	return nil
}

// OpWrapContainerUpdate takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerUpdate(wrapped OpContainerUpdate, fn func(inner func())) OpContainerUpdate {
	return func() (out1 bool, out2 int, out3 Container) {
		inner := func() {
			out1, out2, out3 = wrapped()
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerUpdate(table OpTableContainer, name string) OpContainerUpdate {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeUpdate]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdate)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewRange, the Container-specific
// form of OpTypeViewRange.
type OpContainerViewRange func(uint16, uint16) (bool, int, ReadOnlyContainer)

func (OpContainerViewRange) ContainerOpType() OpType {
	return OpTypeViewRange
}

type OpTableContainerViewRange map[string]OpContainerViewRange

func (OpTableContainerViewRange) ContainerOpTypeTable() OpType {
	return OpTypeViewRange
}

func OpLookupContainerViewRange(target ReadOnlyContainer, name string) OpContainerViewRange {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewRange, name)
		if method != nil {
			return method.(OpContainerViewRange)
		}
	}
	return OpLookupDirectContainerViewRange(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewRange(target ReadOnlyContainer, name string) OpContainerViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewRange(fn)
	}
	return nil
}

// OpWrapContainerViewRange takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewRange(wrapped OpContainerViewRange, fn func(inner func())) OpContainerViewRange {
	return func(in1 uint16, in2 uint16) (out1 bool, out2 int, out3 ReadOnlyContainer) {
		inner := func() {
			out1, out2, out3 = wrapped(in1, in2)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewRange(table OpTableContainer, name string) OpContainerViewRange {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewRange]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRange)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewBit, the Container-specific
// form of OpTypeViewBit.
type OpContainerViewBit func(uint16) (bool, int, ReadOnlyContainer)

func (OpContainerViewBit) ContainerOpType() OpType {
	return OpTypeViewBit
}

type OpTableContainerViewBit map[string]OpContainerViewBit

func (OpTableContainerViewBit) ContainerOpTypeTable() OpType {
	return OpTypeViewBit
}

func OpLookupContainerViewBit(target ReadOnlyContainer, name string) OpContainerViewBit {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewBit, name)
		if method != nil {
			return method.(OpContainerViewBit)
		}
	}
	return OpLookupDirectContainerViewBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewBit(target ReadOnlyContainer, name string) OpContainerViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewBit(fn)
	}
	return nil
}

// OpWrapContainerViewBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewBit(wrapped OpContainerViewBit, fn func(inner func())) OpContainerViewBit {
	return func(in1 uint16) (out1 bool, out2 int, out3 ReadOnlyContainer) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewBit(table OpTableContainer, name string) OpContainerViewBit {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerUpdateBit, the Container-specific
// form of OpTypeUpdateBit.
type OpContainerUpdateBit func(uint16) (bool, int, Container)

func (OpContainerUpdateBit) ContainerOpType() OpType {
	return OpTypeUpdateBit
}

type OpTableContainerUpdateBit map[string]OpContainerUpdateBit

func (OpTableContainerUpdateBit) ContainerOpTypeTable() OpType {
	return OpTypeUpdateBit
}

func OpLookupContainerUpdateBit(target ReadOnlyContainer, name string) OpContainerUpdateBit {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateBit, name)
		if method != nil {
			return method.(OpContainerUpdateBit)
		}
	}
	return OpLookupDirectContainerUpdateBit(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerUpdateBit(target ReadOnlyContainer, name string) OpContainerUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, Container))
		return OpContainerUpdateBit(fn)
	}
	return nil
}

// OpWrapContainerUpdateBit takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerUpdateBit(wrapped OpContainerUpdateBit, fn func(inner func())) OpContainerUpdateBit {
	return func(in1 uint16) (out1 bool, out2 int, out3 Container) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerUpdateBit(table OpTableContainer, name string) OpContainerUpdateBit {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeUpdateBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewContainer, the Container-specific
// form of OpTypeViewOther.
type OpContainerViewContainer func(ReadOnlyContainer) (bool, int, ReadOnlyContainer)

func (OpContainerViewContainer) ContainerOpType() OpType {
	return OpTypeViewOther
}

type OpTableContainerViewContainer map[string]OpContainerViewContainer

func (OpTableContainerViewContainer) ContainerOpTypeTable() OpType {
	return OpTypeViewOther
}

func OpLookupContainerViewContainer(target ReadOnlyContainer, name string) OpContainerViewContainer {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewOther, name)
		if method != nil {
			return method.(OpContainerViewContainer)
		}
	}
	return OpLookupDirectContainerViewContainer(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewContainer(target ReadOnlyContainer, name string) OpContainerViewContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyContainer) (bool, int, ReadOnlyContainer))
		return OpContainerViewContainer(fn)
	}
	return nil
}

// OpWrapContainerViewContainer takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewContainer(wrapped OpContainerViewContainer, fn func(inner func())) OpContainerViewContainer {
	return func(in1 ReadOnlyContainer) (out1 bool, out2 int, out3 ReadOnlyContainer) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewContainer(table OpTableContainer, name string) OpContainerViewContainer {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewOther]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewContainer)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerUpdateContainer, the Container-specific
// form of OpTypeUpdateOther.
type OpContainerUpdateContainer func(ReadOnlyContainer) (bool, int, Container)

func (OpContainerUpdateContainer) ContainerOpType() OpType {
	return OpTypeUpdateOther
}

type OpTableContainerUpdateContainer map[string]OpContainerUpdateContainer

func (OpTableContainerUpdateContainer) ContainerOpTypeTable() OpType {
	return OpTypeUpdateOther
}

func OpLookupContainerUpdateContainer(target ReadOnlyContainer, name string) OpContainerUpdateContainer {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateOther, name)
		if method != nil {
			return method.(OpContainerUpdateContainer)
		}
	}
	return OpLookupDirectContainerUpdateContainer(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerUpdateContainer(target ReadOnlyContainer, name string) OpContainerUpdateContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyContainer) (bool, int, Container))
		return OpContainerUpdateContainer(fn)
	}
	return nil
}

// OpWrapContainerUpdateContainer takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerUpdateContainer(wrapped OpContainerUpdateContainer, fn func(inner func())) OpContainerUpdateContainer {
	return func(in1 ReadOnlyContainer) (out1 bool, out2 int, out3 Container) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerUpdateContainer(table OpTableContainer, name string) OpContainerUpdateContainer {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeUpdateOther]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateContainer)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewBits, the Container-specific
// form of OpTypeViewBits.
type OpContainerViewBits func([]uint16) (bool, int, ReadOnlyContainer)

func (OpContainerViewBits) ContainerOpType() OpType {
	return OpTypeViewBits
}

type OpTableContainerViewBits map[string]OpContainerViewBits

func (OpTableContainerViewBits) ContainerOpTypeTable() OpType {
	return OpTypeViewBits
}

func OpLookupContainerViewBits(target ReadOnlyContainer, name string) OpContainerViewBits {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewBits, name)
		if method != nil {
			return method.(OpContainerViewBits)
		}
	}
	return OpLookupDirectContainerViewBits(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewBits(target ReadOnlyContainer, name string) OpContainerViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewBits(fn)
	}
	return nil
}

// OpWrapContainerViewBits takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewBits(wrapped OpContainerViewBits, fn func(inner func())) OpContainerViewBits {
	return func(in1 []uint16) (out1 bool, out2 int, out3 ReadOnlyContainer) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewBits(table OpTableContainer, name string) OpContainerViewBits {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerUpdateBits, the Container-specific
// form of OpTypeUpdateBits.
type OpContainerUpdateBits func([]uint16) (bool, int, Container)

func (OpContainerUpdateBits) ContainerOpType() OpType {
	return OpTypeUpdateBits
}

type OpTableContainerUpdateBits map[string]OpContainerUpdateBits

func (OpTableContainerUpdateBits) ContainerOpTypeTable() OpType {
	return OpTypeUpdateBits
}

func OpLookupContainerUpdateBits(target ReadOnlyContainer, name string) OpContainerUpdateBits {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateBits, name)
		if method != nil {
			return method.(OpContainerUpdateBits)
		}
	}
	return OpLookupDirectContainerUpdateBits(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerUpdateBits(target ReadOnlyContainer, name string) OpContainerUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, Container))
		return OpContainerUpdateBits(fn)
	}
	return nil
}

// OpWrapContainerUpdateBits takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerUpdateBits(wrapped OpContainerUpdateBits, fn func(inner func())) OpContainerUpdateBits {
	return func(in1 []uint16) (out1 bool, out2 int, out3 Container) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerUpdateBits(table OpTableContainer, name string) OpContainerUpdateBits {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeUpdateBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewContainers, the Container-specific
// form of OpTypeViewOthers.
type OpContainerViewContainers func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer)

func (OpContainerViewContainers) ContainerOpType() OpType {
	return OpTypeViewOthers
}

type OpTableContainerViewContainers map[string]OpContainerViewContainers

func (OpTableContainerViewContainers) ContainerOpTypeTable() OpType {
	return OpTypeViewOthers
}

func OpLookupContainerViewContainers(target ReadOnlyContainer, name string) OpContainerViewContainers {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewOthers, name)
		if method != nil {
			return method.(OpContainerViewContainers)
		}
	}
	return OpLookupDirectContainerViewContainers(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewContainers(target ReadOnlyContainer, name string) OpContainerViewContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer))
		return OpContainerViewContainers(fn)
	}
	return nil
}

// OpWrapContainerViewContainers takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewContainers(wrapped OpContainerViewContainers, fn func(inner func())) OpContainerViewContainers {
	return func(in1 []ReadOnlyContainer) (out1 bool, out2 int, out3 ReadOnlyContainer) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewContainers(table OpTableContainer, name string) OpContainerViewContainers {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewOthers]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewContainers)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerUpdateContainers, the Container-specific
// form of OpTypeUpdateOthers.
type OpContainerUpdateContainers func([]ReadOnlyContainer) (bool, int, Container)

func (OpContainerUpdateContainers) ContainerOpType() OpType {
	return OpTypeUpdateOthers
}

type OpTableContainerUpdateContainers map[string]OpContainerUpdateContainers

func (OpTableContainerUpdateContainers) ContainerOpTypeTable() OpType {
	return OpTypeUpdateOthers
}

func OpLookupContainerUpdateContainers(target ReadOnlyContainer, name string) OpContainerUpdateContainers {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateOthers, name)
		if method != nil {
			return method.(OpContainerUpdateContainers)
		}
	}
	return OpLookupDirectContainerUpdateContainers(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerUpdateContainers(target ReadOnlyContainer, name string) OpContainerUpdateContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyContainer) (bool, int, Container))
		return OpContainerUpdateContainers(fn)
	}
	return nil
}

// OpWrapContainerUpdateContainers takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerUpdateContainers(wrapped OpContainerUpdateContainers, fn func(inner func())) OpContainerUpdateContainers {
	return func(in1 []ReadOnlyContainer) (out1 bool, out2 int, out3 Container) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerUpdateContainers(table OpTableContainer, name string) OpContainerUpdateContainers {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeUpdateOthers]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateContainers)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerUpdateBytes, the Container-specific
// form of OpTypeUpdateBytes.
type OpContainerUpdateBytes func([]byte) (bool, int, Container)

func (OpContainerUpdateBytes) ContainerOpType() OpType {
	return OpTypeUpdateBytes
}

type OpTableContainerUpdateBytes map[string]OpContainerUpdateBytes

func (OpTableContainerUpdateBytes) ContainerOpTypeTable() OpType {
	return OpTypeUpdateBytes
}

func OpLookupContainerUpdateBytes(target ReadOnlyContainer, name string) OpContainerUpdateBytes {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeUpdateBytes, name)
		if method != nil {
			return method.(OpContainerUpdateBytes)
		}
	}
	return OpLookupDirectContainerUpdateBytes(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerUpdateBytes(target ReadOnlyContainer, name string) OpContainerUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, int, Container))
		return OpContainerUpdateBytes(fn)
	}
	return nil
}

// OpWrapContainerUpdateBytes takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerUpdateBytes(wrapped OpContainerUpdateBytes, fn func(inner func())) OpContainerUpdateBytes {
	return func(in1 []byte) (out1 bool, out2 int, out3 Container) {
		inner := func() {
			out1, out2, out3 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerUpdateBytes(table OpTableContainer, name string) OpContainerUpdateBytes {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeUpdateBytes]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateBytes)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// Implementation stuff for OpContainerViewWriterGivesError, the Container-specific
// form of OpTypeViewWriterGivesError.
type OpContainerViewWriterGivesError func(io.Writer) error

func (OpContainerViewWriterGivesError) ContainerOpType() OpType {
	return OpTypeViewWriterGivesError
}

type OpTableContainerViewWriterGivesError map[string]OpContainerViewWriterGivesError

func (OpTableContainerViewWriterGivesError) ContainerOpTypeTable() OpType {
	return OpTypeViewWriterGivesError
}

func OpLookupContainerViewWriterGivesError(target ReadOnlyContainer, name string) OpContainerViewWriterGivesError {
	if target, ok := target.((ContainerHasOpLookup)); ok {
		method := target.OpLookup(OpTypeViewWriterGivesError, name)
		if method != nil {
			return method.(OpContainerViewWriterGivesError)
		}
	}
	return OpLookupDirectContainerViewWriterGivesError(target, name)
}

// OpLookupDirect disregards any OpLookup method. It's there to be used in
// cases where you don't want to risk recursive lookups because you're
// already in a lookup of some kind.
func OpLookupDirectContainerViewWriterGivesError(target ReadOnlyContainer, name string) OpContainerViewWriterGivesError {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewWriterGivesError")
	if method.IsValid() {
		fn, _ := method.Interface().(func(io.Writer) error)
		return OpContainerViewWriterGivesError(fn)
	}
	return nil
}

// OpWrapContainerViewWriterGivesError takes a function which takes a function, and gives
// you a function which wraps a provided operation in that function.
func OpWrapContainerViewWriterGivesError(wrapped OpContainerViewWriterGivesError, fn func(inner func())) OpContainerViewWriterGivesError {
	return func(in1 io.Writer) (out1 error) {
		inner := func() {
			out1 = wrapped(in1)
		}
		fn(inner)
		return
	}
}

func LookupTableOpContainerViewWriterGivesError(table OpTableContainer, name string) OpContainerViewWriterGivesError {
	if table == nil {
		return nil
	}
	subTable := table[OpTypeViewWriterGivesError]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewWriterGivesError)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}

// OpContainerLookupGeneric is a generic lookup function which will
// use any provided OpLookup functionality of its target, falling back on
// the default name-based lookup.
func OpContainerLookupGeneric(target ReadOnlyContainer, typ OpType, name string) OpFunctionContainer {
	switch typ {
	case OpTypeView:
		return OpLookupContainerView(target, name)
	case OpTypeViewGivesBool:
		return OpLookupContainerViewGivesBool(target, name)
	case OpTypeViewGivesBit:
		return OpLookupContainerViewGivesBit(target, name)
	case OpTypeViewRangeGivesBool:
		return OpLookupContainerViewRangeGivesBool(target, name)
	case OpTypeViewRangeGivesBit:
		return OpLookupContainerViewRangeGivesBit(target, name)
	case OpTypeViewRangeGivesOther:
		return OpLookupContainerViewRangeGivesContainer(target, name)
	case OpTypeViewRangeGivesBitsBool:
		return OpLookupContainerViewRangeGivesBitsBool(target, name)
	case OpTypeUpdate:
		return OpLookupContainerUpdate(target, name)
	case OpTypeViewRange:
		return OpLookupContainerViewRange(target, name)
	case OpTypeViewBit:
		return OpLookupContainerViewBit(target, name)
	case OpTypeUpdateBit:
		return OpLookupContainerUpdateBit(target, name)
	case OpTypeViewOther:
		return OpLookupContainerViewContainer(target, name)
	case OpTypeUpdateOther:
		return OpLookupContainerUpdateContainer(target, name)
	case OpTypeViewBits:
		return OpLookupContainerViewBits(target, name)
	case OpTypeUpdateBits:
		return OpLookupContainerUpdateBits(target, name)
	case OpTypeViewOthers:
		return OpLookupContainerViewContainers(target, name)
	case OpTypeUpdateOthers:
		return OpLookupContainerUpdateContainers(target, name)
	case OpTypeUpdateBytes:
		return OpLookupContainerUpdateBytes(target, name)
	case OpTypeViewWriterGivesError:
		return OpLookupContainerViewWriterGivesError(target, name)
	}
	return nil
}

// OpContainerLookupGenericDirect is a generic lookup function which ignores
// any OpLookup functionality of the target and just does the reflect stuff.
func OpContainerLookupGenericDirect(target ReadOnlyContainer, typ OpType, name string) OpFunctionContainer {
	switch typ {
	case OpTypeView:
		return OpLookupDirectContainerView(target, name)
	case OpTypeViewGivesBool:
		return OpLookupDirectContainerViewGivesBool(target, name)
	case OpTypeViewGivesBit:
		return OpLookupDirectContainerViewGivesBit(target, name)
	case OpTypeViewRangeGivesBool:
		return OpLookupDirectContainerViewRangeGivesBool(target, name)
	case OpTypeViewRangeGivesBit:
		return OpLookupDirectContainerViewRangeGivesBit(target, name)
	case OpTypeViewRangeGivesOther:
		return OpLookupDirectContainerViewRangeGivesContainer(target, name)
	case OpTypeViewRangeGivesBitsBool:
		return OpLookupDirectContainerViewRangeGivesBitsBool(target, name)
	case OpTypeUpdate:
		return OpLookupDirectContainerUpdate(target, name)
	case OpTypeViewRange:
		return OpLookupDirectContainerViewRange(target, name)
	case OpTypeViewBit:
		return OpLookupDirectContainerViewBit(target, name)
	case OpTypeUpdateBit:
		return OpLookupDirectContainerUpdateBit(target, name)
	case OpTypeViewOther:
		return OpLookupDirectContainerViewContainer(target, name)
	case OpTypeUpdateOther:
		return OpLookupDirectContainerUpdateContainer(target, name)
	case OpTypeViewBits:
		return OpLookupDirectContainerViewBits(target, name)
	case OpTypeUpdateBits:
		return OpLookupDirectContainerUpdateBits(target, name)
	case OpTypeViewOthers:
		return OpLookupDirectContainerViewContainers(target, name)
	case OpTypeUpdateOthers:
		return OpLookupDirectContainerUpdateContainers(target, name)
	case OpTypeUpdateBytes:
		return OpLookupDirectContainerUpdateBytes(target, name)
	case OpTypeViewWriterGivesError:
		return OpLookupDirectContainerViewWriterGivesError(target, name)
	}
	return nil
}
