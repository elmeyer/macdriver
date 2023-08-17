// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TextAttachmentViewProvider] class.
var TextAttachmentViewProviderClass = _TextAttachmentViewProviderClass{objc.GetClass("NSTextAttachmentViewProvider")}

type _TextAttachmentViewProviderClass struct {
	objc.Class
}

// An interface definition for the [TextAttachmentViewProvider] class.
type ITextAttachmentViewProvider interface {
	objc.IObject
	AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes map[foundation.AttributedStringKey]objc.IObject, location PTextLocation, textContainer ITextContainer, proposedLineFragment coregraphics.Rect, position coregraphics.Point) coregraphics.Rect
	AttachmentBoundsForAttributesLocationObjectTextContainerProposedLineFragmentPosition(attributes map[foundation.AttributedStringKey]objc.IObject, locationObject objc.IObject, textContainer ITextContainer, proposedLineFragment coregraphics.Rect, position coregraphics.Point) coregraphics.Rect
	LoadView()
	TextLayoutManager() TextLayoutManager
	Location() TextLocationWrapper
	View() View
	SetView(value IView)
	TracksTextAttachmentViewBounds() bool
	SetTracksTextAttachmentViewBounds(value bool)
	TextAttachment() TextAttachment
}

// A container object that associates a text attachment at a particular document location with a view object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider?language=objc
type TextAttachmentViewProvider struct {
	objc.Object
}

func TextAttachmentViewProviderFrom(ptr unsafe.Pointer) TextAttachmentViewProvider {
	return TextAttachmentViewProvider{
		Object: objc.ObjectFrom(ptr),
	}
}

func (t_ TextAttachmentViewProvider) InitWithTextAttachmentParentViewTextLayoutManagerLocation(textAttachment ITextAttachment, parentView IView, textLayoutManager ITextLayoutManager, location PTextLocation) TextAttachmentViewProvider {
	po3 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[TextAttachmentViewProvider](t_, objc.Sel("initWithTextAttachment:parentView:textLayoutManager:location:"), objc.Ptr(textAttachment), objc.Ptr(parentView), objc.Ptr(textLayoutManager), po3)
	return rv
}

// Creates a new text attachment view whose content starts at the location you provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857595-initwithtextattachment?language=objc
func TextAttachmentViewProvider_InitWithTextAttachmentParentViewTextLayoutManagerLocation(textAttachment ITextAttachment, parentView IView, textLayoutManager ITextLayoutManager, location PTextLocation) TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.Alloc().InitWithTextAttachmentParentViewTextLayoutManagerLocation(textAttachment, parentView, textLayoutManager, location)
}

func (tc _TextAttachmentViewProviderClass) Alloc() TextAttachmentViewProvider {
	rv := objc.Call[TextAttachmentViewProvider](tc, objc.Sel("alloc"))
	return rv
}

func TextAttachmentViewProvider_Alloc() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.Alloc()
}

func (tc _TextAttachmentViewProviderClass) New() TextAttachmentViewProvider {
	rv := objc.Call[TextAttachmentViewProvider](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTextAttachmentViewProvider() TextAttachmentViewProvider {
	return TextAttachmentViewProviderClass.New()
}

func (t_ TextAttachmentViewProvider) Init() TextAttachmentViewProvider {
	rv := objc.Call[TextAttachmentViewProvider](t_, objc.Sel("init"))
	return rv
}

// Returns the layout bounds for an attachment at a specific text location that contains the text attributes you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857594-attachmentboundsforattributes?language=objc
func (t_ TextAttachmentViewProvider) AttachmentBoundsForAttributesLocationTextContainerProposedLineFragmentPosition(attributes map[foundation.AttributedStringKey]objc.IObject, location PTextLocation, textContainer ITextContainer, proposedLineFragment coregraphics.Rect, position coregraphics.Point) coregraphics.Rect {
	po1 := objc.WrapAsProtocol("NSTextLocation", location)
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"), attributes, po1, objc.Ptr(textContainer), proposedLineFragment, position)
	return rv
}

// Returns the layout bounds for an attachment at a specific text location that contains the text attributes you specify. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857594-attachmentboundsforattributes?language=objc
func (t_ TextAttachmentViewProvider) AttachmentBoundsForAttributesLocationObjectTextContainerProposedLineFragmentPosition(attributes map[foundation.AttributedStringKey]objc.IObject, locationObject objc.IObject, textContainer ITextContainer, proposedLineFragment coregraphics.Rect, position coregraphics.Point) coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](t_, objc.Sel("attachmentBoundsForAttributes:location:textContainer:proposedLineFragment:position:"), attributes, objc.Ptr(locationObject), objc.Ptr(textContainer), proposedLineFragment, position)
	return rv
}

// Draws the custom view hierarchy that text attachment view subclasses implement. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857596-loadview?language=objc
func (t_ TextAttachmentViewProvider) LoadView() {
	objc.Call[objc.Void](t_, objc.Sel("loadView"))
}

// The text layout manager for this view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857599-textlayoutmanager?language=objc
func (t_ TextAttachmentViewProvider) TextLayoutManager() TextLayoutManager {
	rv := objc.Call[TextLayoutManager](t_, objc.Sel("textLayoutManager"))
	return rv
}

// The location that indicates the start of the text attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857597-location?language=objc
func (t_ TextAttachmentViewProvider) Location() TextLocationWrapper {
	rv := objc.Call[TextLocationWrapper](t_, objc.Sel("location"))
	return rv
}

// The text attachment’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857601-view?language=objc
func (t_ TextAttachmentViewProvider) View() View {
	rv := objc.Call[View](t_, objc.Sel("view"))
	return rv
}

// The text attachment’s view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857601-view?language=objc
func (t_ TextAttachmentViewProvider) SetView(value IView) {
	objc.Call[objc.Void](t_, objc.Sel("setView:"), objc.Ptr(value))
}

// A Boolean value that determines the text attachment’s bounds policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857600-trackstextattachmentviewbounds?language=objc
func (t_ TextAttachmentViewProvider) TracksTextAttachmentViewBounds() bool {
	rv := objc.Call[bool](t_, objc.Sel("tracksTextAttachmentViewBounds"))
	return rv
}

// A Boolean value that determines the text attachment’s bounds policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857600-trackstextattachmentviewbounds?language=objc
func (t_ TextAttachmentViewProvider) SetTracksTextAttachmentViewBounds(value bool) {
	objc.Call[objc.Void](t_, objc.Sel("setTracksTextAttachmentViewBounds:"), value)
}

// The text attachment for this view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/uikit/nstextattachmentviewprovider/3857598-textattachment?language=objc
func (t_ TextAttachmentViewProvider) TextAttachment() TextAttachment {
	rv := objc.Call[TextAttachment](t_, objc.Sel("textAttachment"))
	return rv
}