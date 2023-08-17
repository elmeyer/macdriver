// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PresentationIntent] class.
var PresentationIntentClass = _PresentationIntentClass{objc.GetClass("NSPresentationIntent")}

type _PresentationIntentClass struct {
	objc.Class
}

// An interface definition for the [PresentationIntent] class.
type IPresentationIntent interface {
	objc.IObject
	IsEquivalentToPresentationIntent(other IPresentationIntent) bool
	IntentKind() PresentationIntentKind
	Ordinal() int
	IndentationLevel() int
	ColumnCount() int
	Column() int
	HeaderLevel() int
	Identity() int
	Row() int
	ColumnAlignments() []Number
	LanguageHint() string
	ParentIntent() PresentationIntent
}

// The intended presentation for blocks of text, like paragraphs, lists, code blocsk, and parts of tables. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent?language=objc
type PresentationIntent struct {
	objc.Object
}

func PresentationIntentFrom(ptr unsafe.Pointer) PresentationIntent {
	return PresentationIntent{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PresentationIntentClass) Alloc() PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("alloc"))
	return rv
}

func PresentationIntent_Alloc() PresentationIntent {
	return PresentationIntentClass.Alloc()
}

func (pc _PresentationIntentClass) New() PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPresentationIntent() PresentationIntent {
	return PresentationIntentClass.New()
}

func (p_ PresentationIntent) Init() PresentationIntent {
	rv := objc.Call[PresentationIntent](p_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793753-tablecellintentwithidentity?language=objc
func (pc _PresentationIntentClass) TableCellIntentWithIdentityColumnNestedInsideIntent(identity int, column int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("tableCellIntentWithIdentity:column:nestedInsideIntent:"), identity, column, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793753-tablecellintentwithidentity?language=objc
func PresentationIntent_TableCellIntentWithIdentityColumnNestedInsideIntent(identity int, column int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.TableCellIntentWithIdentityColumnNestedInsideIntent(identity, column, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793755-tableintentwithidentity?language=objc
func (pc _PresentationIntentClass) TableIntentWithIdentityColumnCountAlignmentsNestedInsideIntent(identity int, columnCount int, alignments []INumber, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("tableIntentWithIdentity:columnCount:alignments:nestedInsideIntent:"), identity, columnCount, alignments, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793755-tableintentwithidentity?language=objc
func PresentationIntent_TableIntentWithIdentityColumnCountAlignmentsNestedInsideIntent(identity int, columnCount int, alignments []INumber, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.TableIntentWithIdentityColumnCountAlignmentsNestedInsideIntent(identity, columnCount, alignments, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793736-codeblockintentwithidentity?language=objc
func (pc _PresentationIntentClass) CodeBlockIntentWithIdentityLanguageHintNestedInsideIntent(identity int, languageHint string, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("codeBlockIntentWithIdentity:languageHint:nestedInsideIntent:"), identity, languageHint, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793736-codeblockintentwithidentity?language=objc
func PresentationIntent_CodeBlockIntentWithIdentityLanguageHintNestedInsideIntent(identity int, languageHint string, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.CodeBlockIntentWithIdentityLanguageHintNestedInsideIntent(identity, languageHint, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793754-tableheaderrowintentwithidentity?language=objc
func (pc _PresentationIntentClass) TableHeaderRowIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("tableHeaderRowIntentWithIdentity:nestedInsideIntent:"), identity, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793754-tableheaderrowintentwithidentity?language=objc
func PresentationIntent_TableHeaderRowIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.TableHeaderRowIntentWithIdentityNestedInsideIntent(identity, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793747-listitemintentwithidentity?language=objc
func (pc _PresentationIntentClass) ListItemIntentWithIdentityOrdinalNestedInsideIntent(identity int, ordinal int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("listItemIntentWithIdentity:ordinal:nestedInsideIntent:"), identity, ordinal, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793747-listitemintentwithidentity?language=objc
func PresentationIntent_ListItemIntentWithIdentityOrdinalNestedInsideIntent(identity int, ordinal int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.ListItemIntentWithIdentityOrdinalNestedInsideIntent(identity, ordinal, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793757-thematicbreakintentwithidentity?language=objc
func (pc _PresentationIntentClass) ThematicBreakIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("thematicBreakIntentWithIdentity:nestedInsideIntent:"), identity, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793757-thematicbreakintentwithidentity?language=objc
func PresentationIntent_ThematicBreakIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.ThematicBreakIntentWithIdentityNestedInsideIntent(identity, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793750-paragraphintentwithidentity?language=objc
func (pc _PresentationIntentClass) ParagraphIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("paragraphIntentWithIdentity:nestedInsideIntent:"), identity, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793750-paragraphintentwithidentity?language=objc
func PresentationIntent_ParagraphIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.ParagraphIntentWithIdentityNestedInsideIntent(identity, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793758-unorderedlistintentwithidentity?language=objc
func (pc _PresentationIntentClass) UnorderedListIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("unorderedListIntentWithIdentity:nestedInsideIntent:"), identity, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793758-unorderedlistintentwithidentity?language=objc
func PresentationIntent_UnorderedListIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.UnorderedListIntentWithIdentityNestedInsideIntent(identity, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793745-isequivalenttopresentationintent?language=objc
func (p_ PresentationIntent) IsEquivalentToPresentationIntent(other IPresentationIntent) bool {
	rv := objc.Call[bool](p_, objc.Sel("isEquivalentToPresentationIntent:"), objc.Ptr(other))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793735-blockquoteintentwithidentity?language=objc
func (pc _PresentationIntentClass) BlockQuoteIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("blockQuoteIntentWithIdentity:nestedInsideIntent:"), identity, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793735-blockquoteintentwithidentity?language=objc
func PresentationIntent_BlockQuoteIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.BlockQuoteIntentWithIdentityNestedInsideIntent(identity, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793756-tablerowintentwithidentity?language=objc
func (pc _PresentationIntentClass) TableRowIntentWithIdentityRowNestedInsideIntent(identity int, row int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("tableRowIntentWithIdentity:row:nestedInsideIntent:"), identity, row, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793756-tablerowintentwithidentity?language=objc
func PresentationIntent_TableRowIntentWithIdentityRowNestedInsideIntent(identity int, row int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.TableRowIntentWithIdentityRowNestedInsideIntent(identity, row, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793740-headerintentwithidentity?language=objc
func (pc _PresentationIntentClass) HeaderIntentWithIdentityLevelNestedInsideIntent(identity int, level int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("headerIntentWithIdentity:level:nestedInsideIntent:"), identity, level, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793740-headerintentwithidentity?language=objc
func PresentationIntent_HeaderIntentWithIdentityLevelNestedInsideIntent(identity int, level int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.HeaderIntentWithIdentityLevelNestedInsideIntent(identity, level, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793748-orderedlistintentwithidentity?language=objc
func (pc _PresentationIntentClass) OrderedListIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	rv := objc.Call[PresentationIntent](pc, objc.Sel("orderedListIntentWithIdentity:nestedInsideIntent:"), identity, objc.Ptr(parent))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793748-orderedlistintentwithidentity?language=objc
func PresentationIntent_OrderedListIntentWithIdentityNestedInsideIntent(identity int, parent IPresentationIntent) PresentationIntent {
	return PresentationIntentClass.OrderedListIntentWithIdentityNestedInsideIntent(identity, parent)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793744-intentkind?language=objc
func (p_ PresentationIntent) IntentKind() PresentationIntentKind {
	rv := objc.Call[PresentationIntentKind](p_, objc.Sel("intentKind"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793749-ordinal?language=objc
func (p_ PresentationIntent) Ordinal() int {
	rv := objc.Call[int](p_, objc.Sel("ordinal"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793743-indentationlevel?language=objc
func (p_ PresentationIntent) IndentationLevel() int {
	rv := objc.Call[int](p_, objc.Sel("indentationLevel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793739-columncount?language=objc
func (p_ PresentationIntent) ColumnCount() int {
	rv := objc.Call[int](p_, objc.Sel("columnCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793737-column?language=objc
func (p_ PresentationIntent) Column() int {
	rv := objc.Call[int](p_, objc.Sel("column"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793741-headerlevel?language=objc
func (p_ PresentationIntent) HeaderLevel() int {
	rv := objc.Call[int](p_, objc.Sel("headerLevel"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793742-identity?language=objc
func (p_ PresentationIntent) Identity() int {
	rv := objc.Call[int](p_, objc.Sel("identity"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793752-row?language=objc
func (p_ PresentationIntent) Row() int {
	rv := objc.Call[int](p_, objc.Sel("row"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793738-columnalignments?language=objc
func (p_ PresentationIntent) ColumnAlignments() []Number {
	rv := objc.Call[[]Number](p_, objc.Sel("columnAlignments"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793746-languagehint?language=objc
func (p_ PresentationIntent) LanguageHint() string {
	rv := objc.Call[string](p_, objc.Sel("languageHint"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nspresentationintent/3793751-parentintent?language=objc
func (p_ PresentationIntent) ParentIntent() PresentationIntent {
	rv := objc.Call[PresentationIntent](p_, objc.Sel("parentIntent"))
	return rv
}