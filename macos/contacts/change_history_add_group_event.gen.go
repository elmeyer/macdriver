// AUTO-GENERATED CODE, DO NOT MODIFY

package contacts

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ChangeHistoryAddGroupEvent] class.
var ChangeHistoryAddGroupEventClass = _ChangeHistoryAddGroupEventClass{objc.GetClass("CNChangeHistoryAddGroupEvent")}

type _ChangeHistoryAddGroupEventClass struct {
	objc.Class
}

// An interface definition for the [ChangeHistoryAddGroupEvent] class.
type IChangeHistoryAddGroupEvent interface {
	IChangeHistoryEvent
	Group() Group
	ContainerIdentifier() string
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddgroupevent?language=objc
type ChangeHistoryAddGroupEvent struct {
	ChangeHistoryEvent
}

func ChangeHistoryAddGroupEventFrom(ptr unsafe.Pointer) ChangeHistoryAddGroupEvent {
	return ChangeHistoryAddGroupEvent{
		ChangeHistoryEvent: ChangeHistoryEventFrom(ptr),
	}
}

func (cc _ChangeHistoryAddGroupEventClass) Alloc() ChangeHistoryAddGroupEvent {
	rv := objc.Call[ChangeHistoryAddGroupEvent](cc, objc.Sel("alloc"))
	return rv
}

func ChangeHistoryAddGroupEvent_Alloc() ChangeHistoryAddGroupEvent {
	return ChangeHistoryAddGroupEventClass.Alloc()
}

func (cc _ChangeHistoryAddGroupEventClass) New() ChangeHistoryAddGroupEvent {
	rv := objc.Call[ChangeHistoryAddGroupEvent](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewChangeHistoryAddGroupEvent() ChangeHistoryAddGroupEvent {
	return ChangeHistoryAddGroupEventClass.New()
}

func (c_ ChangeHistoryAddGroupEvent) Init() ChangeHistoryAddGroupEvent {
	rv := objc.Call[ChangeHistoryAddGroupEvent](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddgroupevent/3113254-group?language=objc
func (c_ ChangeHistoryAddGroupEvent) Group() Group {
	rv := objc.Call[Group](c_, objc.Sel("group"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/contacts/cnchangehistoryaddgroupevent/3113253-containeridentifier?language=objc
func (c_ ChangeHistoryAddGroupEvent) ContainerIdentifier() string {
	rv := objc.Call[string](c_, objc.Sel("containerIdentifier"))
	return rv
}