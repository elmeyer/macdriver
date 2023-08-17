// AUTO-GENERATED CODE, DO NOT MODIFY

package coredata

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [EntityMapping] class.
var EntityMappingClass = _EntityMappingClass{objc.GetClass("NSEntityMapping")}

type _EntityMappingClass struct {
	objc.Class
}

// An interface definition for the [EntityMapping] class.
type IEntityMapping interface {
	objc.IObject
	AttributeMappings() []PropertyMapping
	SetAttributeMappings(value []IPropertyMapping)
	EntityMigrationPolicyClassName() string
	SetEntityMigrationPolicyClassName(value string)
	Name() string
	SetName(value string)
	DestinationEntityName() string
	SetDestinationEntityName(value string)
	UserInfo() foundation.Dictionary
	SetUserInfo(value foundation.Dictionary)
	SourceEntityName() string
	SetSourceEntityName(value string)
	SourceEntityVersionHash() []byte
	SetSourceEntityVersionHash(value []byte)
	MappingType() EntityMappingType
	SetMappingType(value EntityMappingType)
	RelationshipMappings() []PropertyMapping
	SetRelationshipMappings(value []IPropertyMapping)
	DestinationEntityVersionHash() []byte
	SetDestinationEntityVersionHash(value []byte)
	SourceExpression() foundation.Expression
	SetSourceExpression(value foundation.IExpression)
}

// A mapping instance that specifies how to map an entity from a source to a destination managed object model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping?language=objc
type EntityMapping struct {
	objc.Object
}

func EntityMappingFrom(ptr unsafe.Pointer) EntityMapping {
	return EntityMapping{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _EntityMappingClass) Alloc() EntityMapping {
	rv := objc.Call[EntityMapping](ec, objc.Sel("alloc"))
	return rv
}

func EntityMapping_Alloc() EntityMapping {
	return EntityMappingClass.Alloc()
}

func (ec _EntityMappingClass) New() EntityMapping {
	rv := objc.Call[EntityMapping](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewEntityMapping() EntityMapping {
	return EntityMappingClass.New()
}

func (e_ EntityMapping) Init() EntityMapping {
	rv := objc.Call[EntityMapping](e_, objc.Sel("init"))
	return rv
}

// The array of attribute mappings for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443193-attributemappings?language=objc
func (e_ EntityMapping) AttributeMappings() []PropertyMapping {
	rv := objc.Call[[]PropertyMapping](e_, objc.Sel("attributeMappings"))
	return rv
}

// The array of attribute mappings for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443193-attributemappings?language=objc
func (e_ EntityMapping) SetAttributeMappings(value []IPropertyMapping) {
	objc.Call[objc.Void](e_, objc.Sel("setAttributeMappings:"), value)
}

// The class name of the migration policy for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443171-entitymigrationpolicyclassname?language=objc
func (e_ EntityMapping) EntityMigrationPolicyClassName() string {
	rv := objc.Call[string](e_, objc.Sel("entityMigrationPolicyClassName"))
	return rv
}

// The class name of the migration policy for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443171-entitymigrationpolicyclassname?language=objc
func (e_ EntityMapping) SetEntityMigrationPolicyClassName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setEntityMigrationPolicyClassName:"), value)
}

// The name of the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443167-name?language=objc
func (e_ EntityMapping) Name() string {
	rv := objc.Call[string](e_, objc.Sel("name"))
	return rv
}

// The name of the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443167-name?language=objc
func (e_ EntityMapping) SetName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setName:"), value)
}

// The destination entity name for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443176-destinationentityname?language=objc
func (e_ EntityMapping) DestinationEntityName() string {
	rv := objc.Call[string](e_, objc.Sel("destinationEntityName"))
	return rv
}

// The destination entity name for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443176-destinationentityname?language=objc
func (e_ EntityMapping) SetDestinationEntityName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setDestinationEntityName:"), value)
}

// The user info dictionary for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443184-userinfo?language=objc
func (e_ EntityMapping) UserInfo() foundation.Dictionary {
	rv := objc.Call[foundation.Dictionary](e_, objc.Sel("userInfo"))
	return rv
}

// The user info dictionary for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443184-userinfo?language=objc
func (e_ EntityMapping) SetUserInfo(value foundation.Dictionary) {
	objc.Call[objc.Void](e_, objc.Sel("setUserInfo:"), value)
}

// The source entity name for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443187-sourceentityname?language=objc
func (e_ EntityMapping) SourceEntityName() string {
	rv := objc.Call[string](e_, objc.Sel("sourceEntityName"))
	return rv
}

// The source entity name for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443187-sourceentityname?language=objc
func (e_ EntityMapping) SetSourceEntityName(value string) {
	objc.Call[objc.Void](e_, objc.Sel("setSourceEntityName:"), value)
}

// The version hash of the source entity for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443182-sourceentityversionhash?language=objc
func (e_ EntityMapping) SourceEntityVersionHash() []byte {
	rv := objc.Call[[]byte](e_, objc.Sel("sourceEntityVersionHash"))
	return rv
}

// The version hash of the source entity for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443182-sourceentityversionhash?language=objc
func (e_ EntityMapping) SetSourceEntityVersionHash(value []byte) {
	objc.Call[objc.Void](e_, objc.Sel("setSourceEntityVersionHash:"), value)
}

// The mapping type for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443194-mappingtype?language=objc
func (e_ EntityMapping) MappingType() EntityMappingType {
	rv := objc.Call[EntityMappingType](e_, objc.Sel("mappingType"))
	return rv
}

// The mapping type for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443194-mappingtype?language=objc
func (e_ EntityMapping) SetMappingType(value EntityMappingType) {
	objc.Call[objc.Void](e_, objc.Sel("setMappingType:"), value)
}

// The array of relationship mappings for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443163-relationshipmappings?language=objc
func (e_ EntityMapping) RelationshipMappings() []PropertyMapping {
	rv := objc.Call[[]PropertyMapping](e_, objc.Sel("relationshipMappings"))
	return rv
}

// The array of relationship mappings for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443163-relationshipmappings?language=objc
func (e_ EntityMapping) SetRelationshipMappings(value []IPropertyMapping) {
	objc.Call[objc.Void](e_, objc.Sel("setRelationshipMappings:"), value)
}

// The version hash for the destination entity for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443169-destinationentityversionhash?language=objc
func (e_ EntityMapping) DestinationEntityVersionHash() []byte {
	rv := objc.Call[[]byte](e_, objc.Sel("destinationEntityVersionHash"))
	return rv
}

// The version hash for the destination entity for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443169-destinationentityversionhash?language=objc
func (e_ EntityMapping) SetDestinationEntityVersionHash(value []byte) {
	objc.Call[objc.Void](e_, objc.Sel("setDestinationEntityVersionHash:"), value)
}

// The source expression for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443180-sourceexpression?language=objc
func (e_ EntityMapping) SourceExpression() foundation.Expression {
	rv := objc.Call[foundation.Expression](e_, objc.Sel("sourceExpression"))
	return rv
}

// The source expression for the entity mapping. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coredata/nsentitymapping/1443180-sourceexpression?language=objc
func (e_ EntityMapping) SetSourceExpression(value foundation.IExpression) {
	objc.Call[objc.Void](e_, objc.Sel("setSourceExpression:"), objc.Ptr(value))
}