package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Container (see gen/main.go). These are expressed
// as method signatures -- the Container they operate on is an implicit
// receiver not shown in the signature.

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionContainer interface {
	ContainerOpType() OpType
}

type OpContainerView func() (bool, int, ReadOnlyContainer)
func (OpContainerView) ContainerOpType() OpType { return OpTypeView }
var zeroOpContainerView OpContainerView

type OpContainerUpdate func() (bool, int, Container)
func (OpContainerUpdate) ContainerOpType() OpType { return OpTypeUpdate }
var zeroOpContainerUpdate OpContainerUpdate

type OpContainerViewRange func(uint16, uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewRange) ContainerOpType() OpType { return OpTypeViewRange }
var zeroOpContainerViewRange OpContainerViewRange

type OpContainerViewBit func(uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBit) ContainerOpType() OpType { return OpTypeViewBit }
var zeroOpContainerViewBit OpContainerViewBit

type OpContainerUpdateBit func(uint16) (bool, int, Container)
func (OpContainerUpdateBit) ContainerOpType() OpType { return OpTypeUpdateBit }
var zeroOpContainerUpdateBit OpContainerUpdateBit

type OpContainerViewContainer func(ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainer) ContainerOpType() OpType { return OpTypeViewOther }
var zeroOpContainerViewContainer OpContainerViewContainer

type OpContainerUpdateContainer func(ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainer) ContainerOpType() OpType { return OpTypeUpdateOther }
var zeroOpContainerUpdateContainer OpContainerUpdateContainer

type OpContainerViewBits func([]uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBits) ContainerOpType() OpType { return OpTypeViewBits }
var zeroOpContainerViewBits OpContainerViewBits

type OpContainerUpdateBits func([]uint16) (bool, int, Container)
func (OpContainerUpdateBits) ContainerOpType() OpType { return OpTypeUpdateBits }
var zeroOpContainerUpdateBits OpContainerUpdateBits

type OpContainerViewContainers func([]ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainers) ContainerOpType() OpType { return OpTypeViewOthers }
var zeroOpContainerViewContainers OpContainerViewContainers

type OpContainerUpdateContainers func([]ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainers) ContainerOpType() OpType { return OpTypeUpdateOthers }
var zeroOpContainerUpdateContainers OpContainerUpdateContainers

type OpContainerUpdateBytes func([]byte) (bool, int, Container)
func (OpContainerUpdateBytes) ContainerOpType() OpType { return OpTypeUpdateBytes }
var zeroOpContainerUpdateBytes OpContainerUpdateBytes

type OpContainerViewGivesBytes func() ([]byte)
func (OpContainerViewGivesBytes) ContainerOpType() OpType { return OpTypeViewGivesBytes }
var zeroOpContainerViewGivesBytes OpContainerViewGivesBytes

// OpType to reflect.Type lookup table
var lookupContainerFunctionTypes = [OpTypeMax]OpFunctionContainer {
	OpTypeView: zeroOpContainerView,
	OpTypeUpdate: zeroOpContainerUpdate,
	OpTypeViewRange: zeroOpContainerViewRange,
	OpTypeViewBit: zeroOpContainerViewBit,
	OpTypeUpdateBit: zeroOpContainerUpdateBit,
	OpTypeViewOther: zeroOpContainerViewContainer,
	OpTypeUpdateOther: zeroOpContainerUpdateContainer,
	OpTypeViewBits: zeroOpContainerViewBits,
	OpTypeUpdateBits: zeroOpContainerUpdateBits,
	OpTypeViewOthers: zeroOpContainerViewContainers,
	OpTypeUpdateOthers: zeroOpContainerUpdateContainers,
	OpTypeUpdateBytes: zeroOpContainerUpdateBytes,
	OpTypeViewGivesBytes: zeroOpContainerViewGivesBytes,
}
