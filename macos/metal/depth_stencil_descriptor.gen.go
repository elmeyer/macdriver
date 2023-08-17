// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DepthStencilDescriptor] class.
var DepthStencilDescriptorClass = _DepthStencilDescriptorClass{objc.GetClass("MTLDepthStencilDescriptor")}

type _DepthStencilDescriptorClass struct {
	objc.Class
}

// An interface definition for the [DepthStencilDescriptor] class.
type IDepthStencilDescriptor interface {
	objc.IObject
	BackFaceStencil() StencilDescriptor
	SetBackFaceStencil(value IStencilDescriptor)
	FrontFaceStencil() StencilDescriptor
	SetFrontFaceStencil(value IStencilDescriptor)
	DepthCompareFunction() CompareFunction
	SetDepthCompareFunction(value CompareFunction)
	Label() string
	SetLabel(value string)
	IsDepthWriteEnabled() bool
	SetDepthWriteEnabled(value bool)
}

// An object that configures new MTLDepthStencilState objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor?language=objc
type DepthStencilDescriptor struct {
	objc.Object
}

func DepthStencilDescriptorFrom(ptr unsafe.Pointer) DepthStencilDescriptor {
	return DepthStencilDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (dc _DepthStencilDescriptorClass) Alloc() DepthStencilDescriptor {
	rv := objc.Call[DepthStencilDescriptor](dc, objc.Sel("alloc"))
	return rv
}

func DepthStencilDescriptor_Alloc() DepthStencilDescriptor {
	return DepthStencilDescriptorClass.Alloc()
}

func (dc _DepthStencilDescriptorClass) New() DepthStencilDescriptor {
	rv := objc.Call[DepthStencilDescriptor](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDepthStencilDescriptor() DepthStencilDescriptor {
	return DepthStencilDescriptorClass.New()
}

func (d_ DepthStencilDescriptor) Init() DepthStencilDescriptor {
	rv := objc.Call[DepthStencilDescriptor](d_, objc.Sel("init"))
	return rv
}

// The stencil descriptor for back-facing primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462507-backfacestencil?language=objc
func (d_ DepthStencilDescriptor) BackFaceStencil() StencilDescriptor {
	rv := objc.Call[StencilDescriptor](d_, objc.Sel("backFaceStencil"))
	return rv
}

// The stencil descriptor for back-facing primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462507-backfacestencil?language=objc
func (d_ DepthStencilDescriptor) SetBackFaceStencil(value IStencilDescriptor) {
	objc.Call[objc.Void](d_, objc.Sel("setBackFaceStencil:"), objc.Ptr(value))
}

// The stencil descriptor for front-facing primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462476-frontfacestencil?language=objc
func (d_ DepthStencilDescriptor) FrontFaceStencil() StencilDescriptor {
	rv := objc.Call[StencilDescriptor](d_, objc.Sel("frontFaceStencil"))
	return rv
}

// The stencil descriptor for front-facing primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462476-frontfacestencil?language=objc
func (d_ DepthStencilDescriptor) SetFrontFaceStencil(value IStencilDescriptor) {
	objc.Call[objc.Void](d_, objc.Sel("setFrontFaceStencil:"), objc.Ptr(value))
}

// The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462463-depthcomparefunction?language=objc
func (d_ DepthStencilDescriptor) DepthCompareFunction() CompareFunction {
	rv := objc.Call[CompareFunction](d_, objc.Sel("depthCompareFunction"))
	return rv
}

// The comparison that is performed between a fragment’s depth value and the depth value in the attachment, which determines whether to discard the fragment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462463-depthcomparefunction?language=objc
func (d_ DepthStencilDescriptor) SetDepthCompareFunction(value CompareFunction) {
	objc.Call[objc.Void](d_, objc.Sel("setDepthCompareFunction:"), value)
}

// A string that identifies this object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462483-label?language=objc
func (d_ DepthStencilDescriptor) Label() string {
	rv := objc.Call[string](d_, objc.Sel("label"))
	return rv
}

// A string that identifies this object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462483-label?language=objc
func (d_ DepthStencilDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](d_, objc.Sel("setLabel:"), value)
}

// A Boolean value that indicates whether depth values can be written to the depth attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462501-depthwriteenabled?language=objc
func (d_ DepthStencilDescriptor) IsDepthWriteEnabled() bool {
	rv := objc.Call[bool](d_, objc.Sel("isDepthWriteEnabled"))
	return rv
}

// A Boolean value that indicates whether depth values can be written to the depth attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtldepthstencildescriptor/1462501-depthwriteenabled?language=objc
func (d_ DepthStencilDescriptor) SetDepthWriteEnabled(value bool) {
	objc.Call[objc.Void](d_, objc.Sel("setDepthWriteEnabled:"), value)
}