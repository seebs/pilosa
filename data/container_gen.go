package data

// GENERATED CODE, DO NOT EDIT
// Generic operations types for Container

// This interface exists to let us specify that something takes one of
// these functions, but not other function types, and avoid interface{}.
type OpFunctionContainer interface {
	Type(Container) OpType
}

// OpContainerView is a View operation on a ReadOnlyContainer with no other parameters.
type OpContainerView func(ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerView) Type(Container) OpType { return OpTypeView }

// OpContainerUpdate is an Update operation on a Container with no other parameters.
type OpContainerUpdate func(Container) (bool, int, Container)
func (OpContainerUpdate) Type(Container) OpType { return OpTypeUpdate }

// OpContainerViewBit is a View operation on a ReadOnlyContainer and one Bit.
type OpContainerViewBit func(ReadOnlyContainer, uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBit) Type(Container) OpType { return OpTypeViewBit }

// OpContainerUpdateBit is an Update operation on a Container and one Bit.
type OpContainerUpdateBit func(Container, uint16) (bool, int, Container)
func (OpContainerUpdateBit) Type(Container) OpType { return OpTypeUpdateBit }

// OpContainerViewContainer is a View operation on a ReadOnlyContainer and one other Container.
type OpContainerViewContainer func(ReadOnlyContainer, ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainer) Type(Container) OpType { return OpTypeViewOther }

// OpContainerUpdateContainer is an Update operation on a Container and one other Container.
type OpContainerUpdateContainer func(Container, ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainer) Type(Container) OpType { return OpTypeUpdateOther }

// OpContainerViewBits is a View operation on a ReadOnlyContainer and one or more Bits.
type OpContainerViewBits func(ReadOnlyContainer, ...uint16) (bool, int, ReadOnlyContainer)
func (OpContainerViewBits) Type(Container) OpType { return OpTypeViewBits }

// OpContainerUpdateBits is an Update operation on a Container and one or more Bits.
type OpContainerUpdateBits func(Container, ...uint16) (bool, int, Container)
func (OpContainerUpdateBits) Type(Container) OpType { return OpTypeUpdateBits }

// OpContainerViewContainers is a View operation on a ReadOnlyContainer and one or more other Containers.
type OpContainerViewContainers func(ReadOnlyContainer, ...ReadOnlyContainer) (bool, int, ReadOnlyContainer)
func (OpContainerViewContainers) Type(Container) OpType { return OpTypeViewOthers }

// OpContainerUpdateContainers is an Update operation on a Container and one or more other Containers.
type OpContainerUpdateContainers func(Container, ...ReadOnlyContainer) (bool, int, Container)
func (OpContainerUpdateContainers) Type(Container) OpType { return OpTypeUpdateOthers }
