package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Container (see gen/main.go)

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionContainer interface {
	Type(Container) OpType
}

// OpContainerView is a View operation on a ReadOnlyContainer with no other parameters.
type OpContainerView func() (bool, int, ReadOnlyContainer)
func (OpContainerView) Type(Container) OpType { return OpTypeView }
var zeroOpContainerView OpContainerView

// OpContainerUpdate is an Update operation on a Container with no other parameters.
type OpContainerUpdate func() (bool, int, Container)
func (OpContainerUpdate) Type(Container) OpType { return OpTypeUpdate }
var zeroOpContainerUpdate OpContainerUpdate

// OpContainerViewBit is a View operation on a ReadOnlyContainer and one Bit.
type OpContainerViewBit func(uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBit) Type(Container) OpType { return OpTypeViewBit }
var zeroOpContainerViewBit OpContainerViewBit

// OpContainerUpdateBit is an Update operation on a Container and one Bit.
type OpContainerUpdateBit func(uint16) (bool, int, Container)
func (OpContainerUpdateBit) Type(Container) OpType { return OpTypeUpdateBit }
var zeroOpContainerUpdateBit OpContainerUpdateBit

// OpContainerViewContainer is a View operation on a ReadOnlyContainer and one other Container.
type OpContainerViewContainer func(ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainer) Type(Container) OpType { return OpTypeViewOther }
var zeroOpContainerViewContainer OpContainerViewContainer

// OpContainerUpdateContainer is an Update operation on a Container and one other Container.
type OpContainerUpdateContainer func(ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainer) Type(Container) OpType { return OpTypeUpdateOther }
var zeroOpContainerUpdateContainer OpContainerUpdateContainer

// OpContainerViewBits is a View operation on a ReadOnlyContainer and one or more Bits.
type OpContainerViewBits func(...uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBits) Type(Container) OpType { return OpTypeViewBits }
var zeroOpContainerViewBits OpContainerViewBits

// OpContainerUpdateBits is an Update operation on a Container and one or more Bits.
type OpContainerUpdateBits func(...uint16) (bool, int, Container)
func (OpContainerUpdateBits) Type(Container) OpType { return OpTypeUpdateBits }
var zeroOpContainerUpdateBits OpContainerUpdateBits

// OpContainerViewContainers is a View operation on a ReadOnlyContainer and one or more other Containers.
type OpContainerViewContainers func(...ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainers) Type(Container) OpType { return OpTypeViewOthers }
var zeroOpContainerViewContainers OpContainerViewContainers

// OpContainerUpdateContainers is an Update operation on a Container and one or more other Containers.
type OpContainerUpdateContainers func(...ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainers) Type(Container) OpType { return OpTypeUpdateOthers }
var zeroOpContainerUpdateContainers OpContainerUpdateContainers

// OpContainerViewBytesUpdateBytes is a View operation on a ReadOnlyContainer and one or more BytesUpdateBytes.
type OpContainerViewBytesUpdateBytes func( (bool, int, ReadOnlyContainer)
func (OpContainerViewBytesUpdateBytes) Type(Container) OpType { return OpTypeViewBytesUpdateBytes }
var zeroOpContainerViewBytesUpdateBytes OpContainerViewBytesUpdateBytes

// OpType to reflect.Type lookup table
var lookupContainerFunctionTypes = [OpTypeMax]OpFunctionContainer {
	OpTypeView: zeroOpContainerView,
	OpTypeUpdate: zeroOpContainerUpdate,
	OpTypeViewBit: zeroOpContainerViewBit,
	OpTypeUpdateBit: zeroOpContainerUpdateBit,
	OpTypeViewOther: zeroOpContainerViewContainer,
	OpTypeUpdateOther: zeroOpContainerUpdateContainer,
	OpTypeViewBits: zeroOpContainerViewBits,
	OpTypeUpdateBits: zeroOpContainerUpdateBits,
	OpTypeViewOthers: zeroOpContainerViewContainers,
	OpTypeUpdateOthers: zeroOpContainerUpdateContainers,
	OpTypeViewBytesUpdateBytes: zeroOpContainerViewBytesUpdateBytes,
}
