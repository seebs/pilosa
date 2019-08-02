package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Container (see gen/main.go). These are expressed
// as method signatures -- the Container they operate on is an implicit
// receiver not shown in the signature.

import (
	"reflect"
)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionContainer interface {
	ContainerOpType() OpType
}

type OpContainerView func() (bool, int, Container)

func (OpContainerView) ContainerOpType() OpType { return OpTypeView }

func LookupOpContainerView(target ReadOnlyContainer, name string) OpContainerView {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "View")
	if method.IsValid() {
		fn, _ := method.Interface().(func() (bool, int, Container))
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

type OpContainerViewRangeGivesContainer func(uint16, uint16) Container

func (OpContainerViewRangeGivesContainer) ContainerOpType() OpType { return OpTypeViewRangeGivesOther }

func LookupOpContainerViewRangeGivesContainer(target ReadOnlyContainer, name string) OpContainerViewRangeGivesContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRangeGivesContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) Container)
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

type OpContainerViewRange func(uint16, uint16) (bool, int, Container)

func (OpContainerViewRange) ContainerOpType() OpType { return OpTypeViewRange }

func LookupOpContainerViewRange(target ReadOnlyContainer, name string) OpContainerViewRange {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewRange")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16, uint16) (bool, int, Container))
		return OpContainerViewRange(fn)
	}
	return nil
}

type OpContainerViewBit func(uint16) (bool, int, Container)

func (OpContainerViewBit) ContainerOpType() OpType { return OpTypeViewBit }

func LookupOpContainerViewBit(target ReadOnlyContainer, name string) OpContainerViewBit {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBit")
	if method.IsValid() {
		fn, _ := method.Interface().(func(uint16) (bool, int, Container))
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

type OpContainerViewContainer func(Container) (bool, int, Container)

func (OpContainerViewContainer) ContainerOpType() OpType { return OpTypeViewOther }

func LookupOpContainerViewContainer(target ReadOnlyContainer, name string) OpContainerViewContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(Container) (bool, int, Container))
		return OpContainerViewContainer(fn)
	}
	return nil
}

type OpContainerUpdateContainer func(Container) (bool, int, Container)

func (OpContainerUpdateContainer) ContainerOpType() OpType { return OpTypeUpdateOther }

func LookupOpContainerUpdateContainer(target ReadOnlyContainer, name string) OpContainerUpdateContainer {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainer")
	if method.IsValid() {
		fn, _ := method.Interface().(func(Container) (bool, int, Container))
		return OpContainerUpdateContainer(fn)
	}
	return nil
}

type OpContainerViewBits func([]uint16) (bool, int, Container)

func (OpContainerViewBits) ContainerOpType() OpType { return OpTypeViewBits }

func LookupOpContainerViewBits(target ReadOnlyContainer, name string) OpContainerViewBits {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewBits")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]uint16) (bool, int, Container))
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

type OpContainerViewContainers func([]Container) (bool, int, Container)

func (OpContainerViewContainers) ContainerOpType() OpType { return OpTypeViewOthers }

func LookupOpContainerViewContainers(target ReadOnlyContainer, name string) OpContainerViewContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]Container) (bool, int, Container))
		return OpContainerViewContainers(fn)
	}
	return nil
}

type OpContainerUpdateContainers func([]Container) (bool, int, Container)

func (OpContainerUpdateContainers) ContainerOpType() OpType { return OpTypeUpdateOthers }

func LookupOpContainerUpdateContainers(target ReadOnlyContainer, name string) OpContainerUpdateContainers {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "UpdateContainers")
	if method.IsValid() {
		fn, _ := method.Interface().(func([]Container) (bool, int, Container))
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

type OpContainerViewGivesBytes func() []byte

func (OpContainerViewGivesBytes) ContainerOpType() OpType { return OpTypeViewGivesBytes }

func LookupOpContainerViewGivesBytes(target ReadOnlyContainer, name string) OpContainerViewGivesBytes {
	val := reflect.ValueOf(target)
	method := val.MethodByName(name + "ViewGivesBytes")
	if method.IsValid() {
		fn, _ := method.Interface().(func() []byte)
		return OpContainerViewGivesBytes(fn)
	}
	return nil
}
