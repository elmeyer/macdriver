// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Container] class.
var ContainerClass = _ContainerClass{objc.GetClass("CNContainer")}

type _ContainerClass struct {
	objc.Class
}

// An interface definition for the [Container] class.
type IContainer interface {
	objc.IObject
	Name() string
	Type() ContainerType
	Identifier() string
}

// An immutable object that represents a collection of contacts. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer?language=objc
type Container struct {
	objc.Object
}

func ContainerFrom(ptr unsafe.Pointer) Container {
	return Container{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _ContainerClass) Alloc() Container {
	rv := objc.Call[Container](cc, objc.Sel("alloc"))
	return rv
}

func Container_Alloc() Container {
	return ContainerClass.Alloc()
}

func (cc _ContainerClass) New() Container {
	rv := objc.Call[Container](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewContainer() Container {
	return ContainerClass.New()
}

func (c_ Container) Init() Container {
	rv := objc.Call[Container](c_, objc.Sel("init"))
	return rv
}

// Returns a predicate to find the containers with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1402941-predicateforcontainerswithidenti?language=objc
func (cc _ContainerClass) PredicateForContainersWithIdentifiers(identifiers []string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContainersWithIdentifiers:"), identifiers)
	return rv
}

// Returns a predicate to find the containers with the specified identifiers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1402941-predicateforcontainerswithidenti?language=objc
func Container_PredicateForContainersWithIdentifiers(identifiers []string) foundation.Predicate {
	return ContainerClass.PredicateForContainersWithIdentifiers(identifiers)
}

// Returns a predicate to find the container of the specified contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1402785-predicateforcontainerofcontactwi?language=objc
func (cc _ContainerClass) PredicateForContainerOfContactWithIdentifier(contactIdentifier string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContainerOfContactWithIdentifier:"), contactIdentifier)
	return rv
}

// Returns a predicate to find the container of the specified contact. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1402785-predicateforcontainerofcontactwi?language=objc
func Container_PredicateForContainerOfContactWithIdentifier(contactIdentifier string) foundation.Predicate {
	return ContainerClass.PredicateForContainerOfContactWithIdentifier(contactIdentifier)
}

// Returns a predicate to find the container of the specified group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1403086-predicateforcontainerofgroupwith?language=objc
func (cc _ContainerClass) PredicateForContainerOfGroupWithIdentifier(groupIdentifier string) foundation.Predicate {
	rv := objc.Call[foundation.Predicate](cc, objc.Sel("predicateForContainerOfGroupWithIdentifier:"), groupIdentifier)
	return rv
}

// Returns a predicate to find the container of the specified group. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1403086-predicateforcontainerofgroupwith?language=objc
func Container_PredicateForContainerOfGroupWithIdentifier(groupIdentifier string) foundation.Predicate {
	return ContainerClass.PredicateForContainerOfGroupWithIdentifier(groupIdentifier)
}

// The name of the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1403082-name?language=objc
func (c_ Container) Name() string {
	rv := objc.Call[string](c_, objc.Sel("name"))
	return rv
}

// The type of the container. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1403412-type?language=objc
func (c_ Container) Type() ContainerType {
	rv := objc.Call[ContainerType](c_, objc.Sel("type"))
	return rv
}

// The unique identifier for a contacts container on the device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cncontainer/1403107-identifier?language=objc
func (c_ Container) Identifier() string {
	rv := objc.Call[string](c_, objc.Sel("identifier"))
	return rv
}