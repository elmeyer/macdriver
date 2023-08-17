// AUTO-GENERATED CODE, DO NOT MODIFY

package appkit

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CollectionViewLayoutAttributes] class.
var CollectionViewLayoutAttributesClass = _CollectionViewLayoutAttributesClass{objc.GetClass("NSCollectionViewLayoutAttributes")}

type _CollectionViewLayoutAttributesClass struct {
	objc.Class
}

// An interface definition for the [CollectionViewLayoutAttributes] class.
type ICollectionViewLayoutAttributes interface {
	objc.IObject
	IsHidden() bool
	SetHidden(value bool)
	RepresentedElementKind() string
	RepresentedElementCategory() CollectionElementCategory
	Alpha() float64
	SetAlpha(value float64)
	IndexPath() foundation.IndexPath
	SetIndexPath(value foundation.IIndexPath)
	Frame() foundation.Rect
	SetFrame(value foundation.Rect)
	ZIndex() int
	SetZIndex(value int)
	Size() foundation.Size
	SetSize(value foundation.Size)
}

// An object that contains layout-related attributes for an element in a collection view. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes?language=objc
type CollectionViewLayoutAttributes struct {
	objc.Object
}

func CollectionViewLayoutAttributesFrom(ptr unsafe.Pointer) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributes{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](cc, objc.Sel("layoutAttributesForInterItemGapBeforeIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Creates and returns a layout attributes object for an inter-item gap view at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1534062-layoutattributesforinteritemgapb?language=objc
func CollectionViewLayoutAttributes_LayoutAttributesForInterItemGapBeforeIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.LayoutAttributesForInterItemGapBeforeIndexPath(indexPath)
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](cc, objc.Sel("layoutAttributesForDecorationViewOfKind:withIndexPath:"), decorationViewKind, objc.Ptr(indexPath))
	return rv
}

// Creates and returns a layout attributes object for a decoration view based on the specified information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1535736-layoutattributesfordecorationvie?language=objc
func CollectionViewLayoutAttributes_LayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind CollectionViewDecorationElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.LayoutAttributesForDecorationViewOfKindWithIndexPath(decorationViewKind, indexPath)
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](cc, objc.Sel("layoutAttributesForSupplementaryViewOfKind:withIndexPath:"), elementKind, objc.Ptr(indexPath))
	return rv
}

// Creates and returns a layout attributes object for a supplementary view based on the specified information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1529406-layoutattributesforsupplementary?language=objc
func CollectionViewLayoutAttributes_LayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind CollectionViewSupplementaryElementKind, indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.LayoutAttributesForSupplementaryViewOfKindWithIndexPath(elementKind, indexPath)
}

func (cc _CollectionViewLayoutAttributesClass) LayoutAttributesForItemWithIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](cc, objc.Sel("layoutAttributesForItemWithIndexPath:"), objc.Ptr(indexPath))
	return rv
}

// Creates and returns a layout attributes object for the item at the specified index path. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1529886-layoutattributesforitemwithindex?language=objc
func CollectionViewLayoutAttributes_LayoutAttributesForItemWithIndexPath(indexPath foundation.IIndexPath) CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.LayoutAttributesForItemWithIndexPath(indexPath)
}

func (cc _CollectionViewLayoutAttributesClass) Alloc() CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](cc, objc.Sel("alloc"))
	return rv
}

func CollectionViewLayoutAttributes_Alloc() CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.Alloc()
}

func (cc _CollectionViewLayoutAttributesClass) New() CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCollectionViewLayoutAttributes() CollectionViewLayoutAttributes {
	return CollectionViewLayoutAttributesClass.New()
}

func (c_ CollectionViewLayoutAttributes) Init() CollectionViewLayoutAttributes {
	rv := objc.Call[CollectionViewLayoutAttributes](c_, objc.Sel("init"))
	return rv
}

// A Boolean value indicating whether the element is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1535336-hidden?language=objc
func (c_ CollectionViewLayoutAttributes) IsHidden() bool {
	rv := objc.Call[bool](c_, objc.Sel("isHidden"))
	return rv
}

// A Boolean value indicating whether the element is hidden. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1535336-hidden?language=objc
func (c_ CollectionViewLayoutAttributes) SetHidden(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setHidden:"), value)
}

// The identifier for specific elements of your collection view interface. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1533826-representedelementkind?language=objc
func (c_ CollectionViewLayoutAttributes) RepresentedElementKind() string {
	rv := objc.Call[string](c_, objc.Sel("representedElementKind"))
	return rv
}

// The type of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1529026-representedelementcategory?language=objc
func (c_ CollectionViewLayoutAttributes) RepresentedElementCategory() CollectionElementCategory {
	rv := objc.Call[CollectionElementCategory](c_, objc.Sel("representedElementCategory"))
	return rv
}

// The transparency of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1525453-alpha?language=objc
func (c_ CollectionViewLayoutAttributes) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

// The transparency of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1525453-alpha?language=objc
func (c_ CollectionViewLayoutAttributes) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}

// The index path of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1531306-indexpath?language=objc
func (c_ CollectionViewLayoutAttributes) IndexPath() foundation.IndexPath {
	rv := objc.Call[foundation.IndexPath](c_, objc.Sel("indexPath"))
	return rv
}

// The index path of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1531306-indexpath?language=objc
func (c_ CollectionViewLayoutAttributes) SetIndexPath(value foundation.IIndexPath) {
	objc.Call[objc.Void](c_, objc.Sel("setIndexPath:"), objc.Ptr(value))
}

// The frame rectangle of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1532636-frame?language=objc
func (c_ CollectionViewLayoutAttributes) Frame() foundation.Rect {
	rv := objc.Call[foundation.Rect](c_, objc.Sel("frame"))
	return rv
}

// The frame rectangle of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1532636-frame?language=objc
func (c_ CollectionViewLayoutAttributes) SetFrame(value foundation.Rect) {
	objc.Call[objc.Void](c_, objc.Sel("setFrame:"), value)
}

// The element’s position on the z axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1531553-zindex?language=objc
func (c_ CollectionViewLayoutAttributes) ZIndex() int {
	rv := objc.Call[int](c_, objc.Sel("zIndex"))
	return rv
}

// The element’s position on the z axis. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1531553-zindex?language=objc
func (c_ CollectionViewLayoutAttributes) SetZIndex(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setZIndex:"), value)
}

// The size of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1528769-size?language=objc
func (c_ CollectionViewLayoutAttributes) Size() foundation.Size {
	rv := objc.Call[foundation.Size](c_, objc.Sel("size"))
	return rv
}

// The size of the element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/appkit/nscollectionviewlayoutattributes/1528769-size?language=objc
func (c_ CollectionViewLayoutAttributes) SetSize(value foundation.Size) {
	objc.Call[objc.Void](c_, objc.Sel("setSize:"), value)
}