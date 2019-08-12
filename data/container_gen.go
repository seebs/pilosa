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
type OpTableGenericContainer interface {
	ContainerOpTypeTable() OpType
}

// OpTableContainer is a slice mapping optypes to map[string]opFunc,
// where any specific map will actually be a map with a concrete type of
// op function. We defined the
type OpTableContainer []OpTableGeneric

// Implementation stuff for ContainerView, the Container-specific
// form of OpTypeView.
type OpContainerView func() (bool, int, ReadOnlyContainer)
func (OpContainerView) ContainerOpType() { return OpTypeView }
type OpTableContainerView map[string]opContainerView
func (OpTableContainerView) ContainerOpTypeTable() OpType { return OpTypeView }

func LookupOpContainerView(target ReadOnlyContainer, name string) OpContainerView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, ReadOnlyContainer))
		return OpContainerView(fn)
	}
	return nil
}

func LookupTableOpContainerView(table OpTableContainer, name string) OpContainerView {
	if table == nil {
		return nil
	}
	subTable := table[ContainerView]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerView)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewGivesBool, the Container-specific
// form of OpTypeViewGivesBool.
type OpContainerViewGivesBool func() bool
func (OpContainerViewGivesBool) ContainerOpType() { return OpTypeViewGivesBool }
type OpTableContainerViewGivesBool map[string]opContainerViewGivesBool
func (OpTableContainerViewGivesBool) ContainerOpTypeTable() OpType { return OpTypeViewGivesBool }

func LookupOpContainerViewGivesBool(target ReadOnlyContainer, name string) OpContainerViewGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func() bool)
		return OpContainerViewGivesBool(fn)
	}
	return nil
}

func LookupTableOpContainerViewGivesBool(table OpTableContainer, name string) OpContainerViewGivesBool {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewGivesBool]
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

// Implementation stuff for ContainerViewGivesBit, the Container-specific
// form of OpTypeViewGivesBit.
type OpContainerViewGivesBit func() int
func (OpContainerViewGivesBit) ContainerOpType() { return OpTypeViewGivesBit }
type OpTableContainerViewGivesBit map[string]opContainerViewGivesBit
func (OpTableContainerViewGivesBit) ContainerOpTypeTable() OpType { return OpTypeViewGivesBit }

func LookupOpContainerViewGivesBit(target ReadOnlyContainer, name string) OpContainerViewGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func() int)
		return OpContainerViewGivesBit(fn)
	}
	return nil
}

func LookupTableOpContainerViewGivesBit(table OpTableContainer, name string) OpContainerViewGivesBit {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewGivesBit]
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

// Implementation stuff for ContainerViewRangeGivesBool, the Container-specific
// form of OpTypeViewRangeGivesBool.
type OpContainerViewRangeGivesBool func(uint16, uint16) bool
func (OpContainerViewRangeGivesBool) ContainerOpType() { return OpTypeViewRangeGivesBool }
type OpTableContainerViewRangeGivesBool map[string]opContainerViewRangeGivesBool
func (OpTableContainerViewRangeGivesBool) ContainerOpTypeTable() OpType { return OpTypeViewRangeGivesBool }

func LookupOpContainerViewRangeGivesBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) bool)
		return OpContainerViewRangeGivesBool(fn)
	}
	return nil
}

func LookupTableOpContainerViewRangeGivesBool(table OpTableContainer, name string) OpContainerViewRangeGivesBool {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewRangeGivesBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewRangeGivesBit, the Container-specific
// form of OpTypeViewRangeGivesBit.
type OpContainerViewRangeGivesBit func(uint16, uint16) int
func (OpContainerViewRangeGivesBit) ContainerOpType() { return OpTypeViewRangeGivesBit }
type OpTableContainerViewRangeGivesBit map[string]opContainerViewRangeGivesBit
func (OpTableContainerViewRangeGivesBit) ContainerOpTypeTable() OpType { return OpTypeViewRangeGivesBit }

func LookupOpContainerViewRangeGivesBit(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) int)
		return OpContainerViewRangeGivesBit(fn)
	}
	return nil
}

func LookupTableOpContainerViewRangeGivesBit(table OpTableContainer, name string) OpContainerViewRangeGivesBit {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewRangeGivesBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewRangeGivesContainer, the Container-specific
// form of OpTypeViewRangeGivesOther.
type OpContainerViewRangeGivesContainer func(uint16, uint16) ReadOnlyContainer
func (OpContainerViewRangeGivesContainer) ContainerOpType() { return OpTypeViewRangeGivesOther }
type OpTableContainerViewRangeGivesContainer map[string]opContainerViewRangeGivesContainer
func (OpTableContainerViewRangeGivesContainer) ContainerOpTypeTable() OpType { return OpTypeViewRangeGivesOther }

func LookupOpContainerViewRangeGivesContainer(target ReadOnlyContainer, name string) OpContainerViewRangeGivesContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) ReadOnlyContainer)
		return OpContainerViewRangeGivesContainer(fn)
	}
	return nil
}

func LookupTableOpContainerViewRangeGivesContainer(table OpTableContainer, name string) OpContainerViewRangeGivesContainer {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewRangeGivesContainer]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesContainer)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewRangeGivesBitsBool, the Container-specific
// form of OpTypeViewRangeGivesBitsBool.
type OpContainerViewRangeGivesBitsBool func(uint16, uint16) ([]int, bool)
func (OpContainerViewRangeGivesBitsBool) ContainerOpType() { return OpTypeViewRangeGivesBitsBool }
type OpTableContainerViewRangeGivesBitsBool map[string]opContainerViewRangeGivesBitsBool
func (OpTableContainerViewRangeGivesBitsBool) ContainerOpTypeTable() OpType { return OpTypeViewRangeGivesBitsBool }

func LookupOpContainerViewRangeGivesBitsBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBitsBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitsBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) ([]int, bool))
		return OpContainerViewRangeGivesBitsBool(fn)
	}
	return nil
}

func LookupTableOpContainerViewRangeGivesBitsBool(table OpTableContainer, name string) OpContainerViewRangeGivesBitsBool {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewRangeGivesBitsBool]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRangeGivesBitsBool)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerUpdate, the Container-specific
// form of OpTypeUpdate.
type OpContainerUpdate func() (bool, int, Container)
func (OpContainerUpdate) ContainerOpType() { return OpTypeUpdate }
type OpTableContainerUpdate map[string]opContainerUpdate
func (OpTableContainerUpdate) ContainerOpTypeTable() OpType { return OpTypeUpdate }

func LookupOpContainerUpdate(target ReadOnlyContainer, name string) OpContainerUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, Container))
		return OpContainerUpdate(fn)
	}
	return nil
}

func LookupTableOpContainerUpdate(table OpTableContainer, name string) OpContainerUpdate {
	if table == nil {
		return nil
	}
	subTable := table[ContainerUpdate]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdate)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewRange, the Container-specific
// form of OpTypeViewRange.
type OpContainerViewRange func(uint16, uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewRange) ContainerOpType() { return OpTypeViewRange }
type OpTableContainerViewRange map[string]opContainerViewRange
func (OpTableContainerViewRange) ContainerOpTypeTable() OpType { return OpTypeViewRange }

func LookupOpContainerViewRange(target ReadOnlyContainer, name string) OpContainerViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewRange(fn)
	}
	return nil
}

func LookupTableOpContainerViewRange(table OpTableContainer, name string) OpContainerViewRange {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewRange]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewRange)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewBit, the Container-specific
// form of OpTypeViewBit.
type OpContainerViewBit func(uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBit) ContainerOpType() { return OpTypeViewBit }
type OpTableContainerViewBit map[string]opContainerViewBit
func (OpTableContainerViewBit) ContainerOpTypeTable() OpType { return OpTypeViewBit }

func LookupOpContainerViewBit(target ReadOnlyContainer, name string) OpContainerViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewBit(fn)
	}
	return nil
}

func LookupTableOpContainerViewBit(table OpTableContainer, name string) OpContainerViewBit {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerUpdateBit, the Container-specific
// form of OpTypeUpdateBit.
type OpContainerUpdateBit func(uint16) (bool, int, Container)
func (OpContainerUpdateBit) ContainerOpType() { return OpTypeUpdateBit }
type OpTableContainerUpdateBit map[string]opContainerUpdateBit
func (OpTableContainerUpdateBit) ContainerOpTypeTable() OpType { return OpTypeUpdateBit }

func LookupOpContainerUpdateBit(target ReadOnlyContainer, name string) OpContainerUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, Container))
		return OpContainerUpdateBit(fn)
	}
	return nil
}

func LookupTableOpContainerUpdateBit(table OpTableContainer, name string) OpContainerUpdateBit {
	if table == nil {
		return nil
	}
	subTable := table[ContainerUpdateBit]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateBit)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewContainer, the Container-specific
// form of OpTypeViewOther.
type OpContainerViewContainer func(ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainer) ContainerOpType() { return OpTypeViewOther }
type OpTableContainerViewContainer map[string]opContainerViewContainer
func (OpTableContainerViewContainer) ContainerOpTypeTable() OpType { return OpTypeViewOther }

func LookupOpContainerViewContainer(target ReadOnlyContainer, name string) OpContainerViewContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyContainer) (bool, int, ReadOnlyContainer))
		return OpContainerViewContainer(fn)
	}
	return nil
}

func LookupTableOpContainerViewContainer(table OpTableContainer, name string) OpContainerViewContainer {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewContainer]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewContainer)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerUpdateContainer, the Container-specific
// form of OpTypeUpdateOther.
type OpContainerUpdateContainer func(ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainer) ContainerOpType() { return OpTypeUpdateOther }
type OpTableContainerUpdateContainer map[string]opContainerUpdateContainer
func (OpTableContainerUpdateContainer) ContainerOpTypeTable() OpType { return OpTypeUpdateOther }

func LookupOpContainerUpdateContainer(target ReadOnlyContainer, name string) OpContainerUpdateContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyContainer) (bool, int, Container))
		return OpContainerUpdateContainer(fn)
	}
	return nil
}

func LookupTableOpContainerUpdateContainer(table OpTableContainer, name string) OpContainerUpdateContainer {
	if table == nil {
		return nil
	}
	subTable := table[ContainerUpdateContainer]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateContainer)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewBits, the Container-specific
// form of OpTypeViewBits.
type OpContainerViewBits func([]uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBits) ContainerOpType() { return OpTypeViewBits }
type OpTableContainerViewBits map[string]opContainerViewBits
func (OpTableContainerViewBits) ContainerOpTypeTable() OpType { return OpTypeViewBits }

func LookupOpContainerViewBits(target ReadOnlyContainer, name string) OpContainerViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewBits(fn)
	}
	return nil
}

func LookupTableOpContainerViewBits(table OpTableContainer, name string) OpContainerViewBits {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerUpdateBits, the Container-specific
// form of OpTypeUpdateBits.
type OpContainerUpdateBits func([]uint16) (bool, int, Container)
func (OpContainerUpdateBits) ContainerOpType() { return OpTypeUpdateBits }
type OpTableContainerUpdateBits map[string]opContainerUpdateBits
func (OpTableContainerUpdateBits) ContainerOpTypeTable() OpType { return OpTypeUpdateBits }

func LookupOpContainerUpdateBits(target ReadOnlyContainer, name string) OpContainerUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, Container))
		return OpContainerUpdateBits(fn)
	}
	return nil
}

func LookupTableOpContainerUpdateBits(table OpTableContainer, name string) OpContainerUpdateBits {
	if table == nil {
		return nil
	}
	subTable := table[ContainerUpdateBits]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateBits)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewContainers, the Container-specific
// form of OpTypeViewOthers.
type OpContainerViewContainers func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainers) ContainerOpType() { return OpTypeViewOthers }
type OpTableContainerViewContainers map[string]opContainerViewContainers
func (OpTableContainerViewContainers) ContainerOpTypeTable() OpType { return OpTypeViewOthers }

func LookupOpContainerViewContainers(target ReadOnlyContainer, name string) OpContainerViewContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer))
		return OpContainerViewContainers(fn)
	}
	return nil
}

func LookupTableOpContainerViewContainers(table OpTableContainer, name string) OpContainerViewContainers {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewContainers]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewContainers)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerUpdateContainers, the Container-specific
// form of OpTypeUpdateOthers.
type OpContainerUpdateContainers func([]ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainers) ContainerOpType() { return OpTypeUpdateOthers }
type OpTableContainerUpdateContainers map[string]opContainerUpdateContainers
func (OpTableContainerUpdateContainers) ContainerOpTypeTable() OpType { return OpTypeUpdateOthers }

func LookupOpContainerUpdateContainers(target ReadOnlyContainer, name string) OpContainerUpdateContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyContainer) (bool, int, Container))
		return OpContainerUpdateContainers(fn)
	}
	return nil
}

func LookupTableOpContainerUpdateContainers(table OpTableContainer, name string) OpContainerUpdateContainers {
	if table == nil {
		return nil
	}
	subTable := table[ContainerUpdateContainers]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateContainers)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerUpdateBytes, the Container-specific
// form of OpTypeUpdateBytes.
type OpContainerUpdateBytes func([]byte) (bool, int, Container)
func (OpContainerUpdateBytes) ContainerOpType() { return OpTypeUpdateBytes }
type OpTableContainerUpdateBytes map[string]opContainerUpdateBytes
func (OpTableContainerUpdateBytes) ContainerOpTypeTable() OpType { return OpTypeUpdateBytes }

func LookupOpContainerUpdateBytes(target ReadOnlyContainer, name string) OpContainerUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, int, Container))
		return OpContainerUpdateBytes(fn)
	}
	return nil
}

func LookupTableOpContainerUpdateBytes(table OpTableContainer, name string) OpContainerUpdateBytes {
	if table == nil {
		return nil
	}
	subTable := table[ContainerUpdateBytes]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerUpdateBytes)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults

// Implementation stuff for ContainerViewWriterGivesError, the Container-specific
// form of OpTypeViewWriterGivesError.
type OpContainerViewWriterGivesError func(io.Writer) error
func (OpContainerViewWriterGivesError) ContainerOpType() { return OpTypeViewWriterGivesError }
type OpTableContainerViewWriterGivesError map[string]opContainerViewWriterGivesError
func (OpTableContainerViewWriterGivesError) ContainerOpTypeTable() OpType { return OpTypeViewWriterGivesError }

func LookupOpContainerViewWriterGivesError(target ReadOnlyContainer, name string) OpContainerViewWriterGivesError {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewWriterGivesError")
	if method.IsValid() {
		fn, _ := method.Interface().(func(io.Writer) error)
		return OpContainerViewWriterGivesError(fn)
	}
	return nil
}

func LookupTableOpContainerViewWriterGivesError(table OpTableContainer, name string) OpContainerViewWriterGivesError {
	if table == nil {
		return nil
	}
	subTable := table[ContainerViewWriterGivesError]
	if subTable == nil {
		return nil
	}
	tab, ok := subTable.(OpTableContainerViewWriterGivesError)
	if tab == nil || !ok {
		return nil
	}
	return tab[name]
}
// No Defaults
}
