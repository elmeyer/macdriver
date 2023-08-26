// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coremedia"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CaptureDeviceFormat] class.
var CaptureDeviceFormatClass = _CaptureDeviceFormatClass{objc.GetClass("AVCaptureDeviceFormat")}

type _CaptureDeviceFormatClass struct {
	objc.Class
}

// An interface definition for the [CaptureDeviceFormat] class.
type ICaptureDeviceFormat interface {
	objc.IObject
	VideoSupportedFrameRateRanges() []FrameRateRange
	AutoFocusSystem() CaptureAutoFocusSystem
	IsHighPhotoQualitySupported() bool
	VideoMaxZoomFactorForCenterStage() float64
	VideoFrameRateRangeForPortraitEffect() FrameRateRange
	MediaType() MediaType
	IsPortraitEffectSupported() bool
	SupportedColorSpaces() []foundation.Number
	IsCenterStageSupported() bool
	VideoFrameRateRangeForCenterStage() FrameRateRange
	VideoMinZoomFactorForCenterStage() float64
	FormatDescription() coremedia.FormatDescriptionRef
}

// A class that defines media formats and capture settings that capture devices support. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat?language=objc
type CaptureDeviceFormat struct {
	objc.Object
}

func CaptureDeviceFormatFrom(ptr unsafe.Pointer) CaptureDeviceFormat {
	return CaptureDeviceFormat{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CaptureDeviceFormatClass) Alloc() CaptureDeviceFormat {
	rv := objc.Call[CaptureDeviceFormat](cc, objc.Sel("alloc"))
	return rv
}

func CaptureDeviceFormat_Alloc() CaptureDeviceFormat {
	return CaptureDeviceFormatClass.Alloc()
}

func (cc _CaptureDeviceFormatClass) New() CaptureDeviceFormat {
	rv := objc.Call[CaptureDeviceFormat](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCaptureDeviceFormat() CaptureDeviceFormat {
	return CaptureDeviceFormatClass.New()
}

func (c_ CaptureDeviceFormat) Init() CaptureDeviceFormat {
	rv := objc.Call[CaptureDeviceFormat](c_, objc.Sel("init"))
	return rv
}

// A list of frame rate ranges that a format supports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/1387592-videosupportedframerateranges?language=objc
func (c_ CaptureDeviceFormat) VideoSupportedFrameRateRanges() []FrameRateRange {
	rv := objc.Call[[]FrameRateRange](c_, objc.Sel("videoSupportedFrameRateRanges"))
	return rv
}

// The auto focus system the format uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/1624600-autofocussystem?language=objc
func (c_ CaptureDeviceFormat) AutoFocusSystem() CaptureAutoFocusSystem {
	rv := objc.Call[CaptureAutoFocusSystem](c_, objc.Sel("autoFocusSystem"))
	return rv
}

// A Boolean value that indicates whether this format supports high-quality capture with the current quality prioritization setting. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3751763-highphotoqualitysupported?language=objc
func (c_ CaptureDeviceFormat) IsHighPhotoQualitySupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isHighPhotoQualitySupported"))
	return rv
}

// The maximum zoom factor available when Center Stage is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3738422-videomaxzoomfactorforcenterstage?language=objc
func (c_ CaptureDeviceFormat) VideoMaxZoomFactorForCenterStage() float64 {
	rv := objc.Call[float64](c_, objc.Sel("videoMaxZoomFactorForCenterStage"))
	return rv
}

// The range of frame rates available when Portrait Effect is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3875313-videoframeraterangeforportraitef?language=objc
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForPortraitEffect() FrameRateRange {
	rv := objc.Call[FrameRateRange](c_, objc.Sel("videoFrameRateRangeForPortraitEffect"))
	return rv
}

// A constant describing the media type of an AVCaptureDevice active or supported format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/1388503-mediatype?language=objc
func (c_ CaptureDeviceFormat) MediaType() MediaType {
	rv := objc.Call[MediaType](c_, objc.Sel("mediaType"))
	return rv
}

// A Boolean value that indicates whether the format supports the Portrait Effect feature. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3875312-portraiteffectsupported?language=objc
func (c_ CaptureDeviceFormat) IsPortraitEffectSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isPortraitEffectSupported"))
	return rv
}

// The list of color spaces the format supports for image and video capture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/1648611-supportedcolorspaces?language=objc
func (c_ CaptureDeviceFormat) SupportedColorSpaces() []foundation.Number {
	rv := objc.Call[[]foundation.Number](c_, objc.Sel("supportedColorSpaces"))
	return rv
}

// A Boolean value that indicates whether the format supports Center Stage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3738420-centerstagesupported?language=objc
func (c_ CaptureDeviceFormat) IsCenterStageSupported() bool {
	rv := objc.Call[bool](c_, objc.Sel("isCenterStageSupported"))
	return rv
}

// The range of frame rates available when Center Stage is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3738421-videoframeraterangeforcenterstag?language=objc
func (c_ CaptureDeviceFormat) VideoFrameRateRangeForCenterStage() FrameRateRange {
	rv := objc.Call[FrameRateRange](c_, objc.Sel("videoFrameRateRangeForCenterStage"))
	return rv
}

// The minimum zoom factor available when Center Stage is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/3738423-videominzoomfactorforcenterstage?language=objc
func (c_ CaptureDeviceFormat) VideoMinZoomFactorForCenterStage() float64 {
	rv := objc.Call[float64](c_, objc.Sel("videoMinZoomFactorForCenterStage"))
	return rv
}

// An object describing the capture format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcapturedeviceformat/1389445-formatdescription?language=objc
func (c_ CaptureDeviceFormat) FormatDescription() coremedia.FormatDescriptionRef {
	rv := objc.Call[coremedia.FormatDescriptionRef](c_, objc.Sel("formatDescription"))
	return rv
}