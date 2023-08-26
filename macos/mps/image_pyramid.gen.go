// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImagePyramid] class.
var ImagePyramidClass = _ImagePyramidClass{objc.GetClass("MPSImagePyramid")}

type _ImagePyramidClass struct {
	objc.Class
}

// An interface definition for the [ImagePyramid] class.
type IImagePyramid interface {
	IUnaryImageKernel
	KernelHeight() uint
	KernelWidth() uint
}

// A base class for creating different kinds of pyramid images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid?language=objc
type ImagePyramid struct {
	UnaryImageKernel
}

func ImagePyramidFrom(ptr unsafe.Pointer) ImagePyramid {
	return ImagePyramid{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImagePyramid) InitWithDevice(device metal.PDevice) ImagePyramid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImagePyramid](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a downwards 5-tap image pyramid with the default filter kernel and device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648935-initwithdevice?language=objc
func NewImagePyramidWithDevice(device metal.PDevice) ImagePyramid {
	instance := ImagePyramidClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImagePyramidClass) Alloc() ImagePyramid {
	rv := objc.Call[ImagePyramid](ic, objc.Sel("alloc"))
	return rv
}

func ImagePyramid_Alloc() ImagePyramid {
	return ImagePyramidClass.Alloc()
}

func (ic _ImagePyramidClass) New() ImagePyramid {
	rv := objc.Call[ImagePyramid](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImagePyramid() ImagePyramid {
	return ImagePyramidClass.New()
}

func (i_ ImagePyramid) Init() ImagePyramid {
	rv := objc.Call[ImagePyramid](i_, objc.Sel("init"))
	return rv
}

func (i_ ImagePyramid) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImagePyramid {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImagePyramid](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImagePyramid_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImagePyramid {
	instance := ImagePyramidClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The height of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648863-kernelheight?language=objc
func (i_ ImagePyramid) KernelHeight() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagepyramid/1648842-kernelwidth?language=objc
func (i_ ImagePyramid) KernelWidth() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelWidth"))
	return rv
}