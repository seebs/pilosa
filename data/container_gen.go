package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Container (see gen/main.go). These are expressed
// as method signatures -- the Container they operate on is an implicit
// receiver not shown in the signature.

import (
	"io"
	"reflect"
)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionContainer interface {
	ContainerOpType() OpType
}

type OpContainerView func() (bool, int, ReadOnlyContainer)

func (OpContainerView) ContainerOpType() OpType { return OpTypeView }

func LookupOpContainerView(target ReadOnlyContainer, name string) OpContainerView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, ReadOnlyContainer))
		return OpContainerView(fn)
	}
	return nil
}

type OpContainerViewGivesBool func() bool

func (OpContainerViewGivesBool) ContainerOpType() OpType { return OpTypeViewGivesBool }

func LookupOpContainerViewGivesBool(target ReadOnlyContainer, name string) OpContainerViewGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func() bool)
		return OpContainerViewGivesBool(fn)
	}
	return nil
}

// Any performs a default ContainerViewGivesBool on a Container.
type interfaceContainerHasAny interface {
	AnyViewGivesBool() bool
}

func ContainerAny(target ReadOnlyContainer) bool {
	if target, ok := target.(interfaceContainerHasAny); ok {
		return target.AnyViewGivesBool()
	}
	return genericContainerAny(target)
}

type OpContainerViewGivesBit func() int

func (OpContainerViewGivesBit) ContainerOpType() OpType { return OpTypeViewGivesBit }

func LookupOpContainerViewGivesBit(target ReadOnlyContainer, name string) OpContainerViewGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func() int)
		return OpContainerViewGivesBit(fn)
	}
	return nil
}

// Count performs a default ContainerViewGivesBit on a Container.
type interfaceContainerHasCount interface {
	CountViewGivesBit() int
}

func ContainerCount(target ReadOnlyContainer) int {
	if target, ok := target.(interfaceContainerHasCount); ok {
		return target.CountViewGivesBit()
	}
	return genericContainerCount(target)
}

type OpContainerViewRangeGivesBool func(uint16, uint16) bool

func (OpContainerViewRangeGivesBool) ContainerOpType() OpType { return OpTypeViewRangeGivesBool }

func LookupOpContainerViewRangeGivesBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) bool)
		return OpContainerViewRangeGivesBool(fn)
	}
	return nil
}

type OpContainerViewRangeGivesBit func(uint16, uint16) int

func (OpContainerViewRangeGivesBit) ContainerOpType() OpType { return OpTypeViewRangeGivesBit }

func LookupOpContainerViewRangeGivesBit(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) int)
		return OpContainerViewRangeGivesBit(fn)
	}
	return nil
}

type OpContainerViewRangeGivesContainer func(uint16, uint16) ReadOnlyContainer

func (OpContainerViewRangeGivesContainer) ContainerOpType() OpType { return OpTypeViewRangeGivesOther }

func LookupOpContainerViewRangeGivesContainer(target ReadOnlyContainer, name string) OpContainerViewRangeGivesContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) ReadOnlyContainer)
		return OpContainerViewRangeGivesContainer(fn)
	}
	return nil
}

type OpContainerViewRangeGivesBitsBool func(uint16, uint16) ([]int, bool)

func (OpContainerViewRangeGivesBitsBool) ContainerOpType() OpType { return OpTypeViewRangeGivesBitsBool }

func LookupOpContainerViewRangeGivesBitsBool(target ReadOnlyContainer, name string) OpContainerViewRangeGivesBitsBool {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesBitsBool")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) ([]int, bool))
		return OpContainerViewRangeGivesBitsBool(fn)
	}
	return nil
}

type OpContainerUpdate func() (bool, int, Container)

func (OpContainerUpdate) ContainerOpType() OpType { return OpTypeUpdate }

func LookupOpContainerUpdate(target ReadOnlyContainer, name string) OpContainerUpdate {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "Update")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, Container))
		return OpContainerUpdate(fn)
	}
	return nil
}

type OpContainerViewRange func(uint16, uint16) (bool, int, ReadOnlyContainer)

func (OpContainerViewRange) ContainerOpType() OpType { return OpTypeViewRange }

func LookupOpContainerViewRange(target ReadOnlyContainer, name string) OpContainerViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewRange(fn)
	}
	return nil
}

type OpContainerViewBit func(uint16) (bool, int, ReadOnlyContainer)

func (OpContainerViewBit) ContainerOpType() OpType { return OpTypeViewBit }

func LookupOpContainerViewBit(target ReadOnlyContainer, name string) OpContainerViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewBit(fn)
	}
	return nil
}

type OpContainerUpdateBit func(uint16) (bool, int, Container)

func (OpContainerUpdateBit) ContainerOpType() OpType { return OpTypeUpdateBit }

func LookupOpContainerUpdateBit(target ReadOnlyContainer, name string) OpContainerUpdateBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, Container))
		return OpContainerUpdateBit(fn)
	}
	return nil
}

type OpContainerViewContainer func(ReadOnlyContainer) (bool, int, ReadOnlyContainer)

func (OpContainerViewContainer) ContainerOpType() OpType { return OpTypeViewOther }

func LookupOpContainerViewContainer(target ReadOnlyContainer, name string) OpContainerViewContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyContainer) (bool, int, ReadOnlyContainer))
		return OpContainerViewContainer(fn)
	}
	return nil
}

type OpContainerUpdateContainer func(ReadOnlyContainer) (bool, int, Container)

func (OpContainerUpdateContainer) ContainerOpType() OpType { return OpTypeUpdateOther }

func LookupOpContainerUpdateContainer(target ReadOnlyContainer, name string) OpContainerUpdateContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(ReadOnlyContainer) (bool, int, Container))
		return OpContainerUpdateContainer(fn)
	}
	return nil
}

type OpContainerViewBits func([]uint16) (bool, int, ReadOnlyContainer)

func (OpContainerViewBits) ContainerOpType() OpType { return OpTypeViewBits }

func LookupOpContainerViewBits(target ReadOnlyContainer, name string) OpContainerViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, ReadOnlyContainer))
		return OpContainerViewBits(fn)
	}
	return nil
}

type OpContainerUpdateBits func([]uint16) (bool, int, Container)

func (OpContainerUpdateBits) ContainerOpType() OpType { return OpTypeUpdateBits }

func LookupOpContainerUpdateBits(target ReadOnlyContainer, name string) OpContainerUpdateBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, Container))
		return OpContainerUpdateBits(fn)
	}
	return nil
}

type OpContainerViewContainers func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer)

func (OpContainerViewContainers) ContainerOpType() OpType { return OpTypeViewOthers }

func LookupOpContainerViewContainers(target ReadOnlyContainer, name string) OpContainerViewContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer))
		return OpContainerViewContainers(fn)
	}
	return nil
}

type OpContainerUpdateContainers func([]ReadOnlyContainer) (bool, int, Container)

func (OpContainerUpdateContainers) ContainerOpType() OpType { return OpTypeUpdateOthers }

func LookupOpContainerUpdateContainers(target ReadOnlyContainer, name string) OpContainerUpdateContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]ReadOnlyContainer) (bool, int, Container))
		return OpContainerUpdateContainers(fn)
	}
	return nil
}

type OpContainerUpdateBytes func([]byte) (bool, int, Container)

func (OpContainerUpdateBytes) ContainerOpType() OpType { return OpTypeUpdateBytes }

func LookupOpContainerUpdateBytes(target ReadOnlyContainer, name string) OpContainerUpdateBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]byte) (bool, int, Container))
		return OpContainerUpdateBytes(fn)
	}
	return nil
}

type OpContainerViewWriterGivesError func(io.Writer) error

func (OpContainerViewWriterGivesError) ContainerOpType() OpType { return OpTypeViewWriterGivesError }

func LookupOpContainerViewWriterGivesError(target ReadOnlyContainer, name string) OpContainerViewWriterGivesError {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewWriterGivesError")
	if method.IsValid() {
		fn, _ := method.Interface().(func(io.Writer) error)
		return OpContainerViewWriterGivesError(fn)
	}
	return nil
}
