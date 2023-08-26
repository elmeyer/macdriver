// AUTO-GENERATED CODE, DO NOT MODIFY

package avfoundation

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Composition] class.
var CompositionClass = _CompositionClass{objc.GetClass("AVComposition")}

type _CompositionClass struct {
	objc.Class
}

// An interface definition for the [Composition] class.
type IComposition interface {
	IAsset
	TracksWithMediaType(mediaType MediaType) []CompositionTrack
	TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic) []CompositionTrack
	TrackWithTrackID(trackID objc.IObject) CompositionTrack
	URLAssetInitializationOptions() map[string]objc.Object
	NaturalSize() coregraphics.Size
}

// An object that combines and arranges media from multiple assets into a single composite asset that you can play or process. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition?language=objc
type Composition struct {
	Asset
}

func CompositionFrom(ptr unsafe.Pointer) Composition {
	return Composition{
		Asset: AssetFrom(ptr),
	}
}

func (cc _CompositionClass) Alloc() Composition {
	rv := objc.Call[Composition](cc, objc.Sel("alloc"))
	return rv
}

func Composition_Alloc() Composition {
	return CompositionClass.Alloc()
}

func (cc _CompositionClass) New() Composition {
	rv := objc.Call[Composition](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewComposition() Composition {
	return CompositionClass.New()
}

func (c_ Composition) Init() Composition {
	rv := objc.Call[Composition](c_, objc.Sel("init"))
	return rv
}

func (cc _CompositionClass) AssetWithURL(URL foundation.IURL) Composition {
	rv := objc.Call[Composition](cc, objc.Sel("assetWithURL:"), objc.Ptr(URL))
	return rv
}

// Creates an asset that models the media at the specified URL. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avasset/1389943-assetwithurl?language=objc
func Composition_AssetWithURL(URL foundation.IURL) Composition {
	return CompositionClass.AssetWithURL(URL)
}

// Returns tracks that contain media of a specified type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/1386534-trackswithmediatype?language=objc
func (c_ Composition) TracksWithMediaType(mediaType MediaType) []CompositionTrack {
	rv := objc.Call[[]CompositionTrack](c_, objc.Sel("tracksWithMediaType:"), mediaType)
	return rv
}

// Returns tracks that contain media of a specified characteristic. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/1387525-trackswithmediacharacteristic?language=objc
func (c_ Composition) TracksWithMediaCharacteristic(mediaCharacteristic MediaCharacteristic) []CompositionTrack {
	rv := objc.Call[[]CompositionTrack](c_, objc.Sel("tracksWithMediaCharacteristic:"), mediaCharacteristic)
	return rv
}

// Returns a track that contains the specified identifier. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/1388473-trackwithtrackid?language=objc
func (c_ Composition) TrackWithTrackID(trackID objc.IObject) CompositionTrack {
	rv := objc.Call[CompositionTrack](c_, objc.Sel("trackWithTrackID:"), objc.Ptr(trackID))
	return rv
}

// The options you used to create a composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/1387080-urlassetinitializationoptions?language=objc
func (c_ Composition) URLAssetInitializationOptions() map[string]objc.Object {
	rv := objc.Call[map[string]objc.Object](c_, objc.Sel("URLAssetInitializationOptions"))
	return rv
}

// The authored size of the visual portion of the composition. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/avfoundation/avcomposition/1387247-naturalsize?language=objc
func (c_ Composition) NaturalSize() coregraphics.Size {
	rv := objc.Call[coregraphics.Size](c_, objc.Sel("naturalSize"))
	return rv
}