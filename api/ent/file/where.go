// Code generated by ent, DO NOT EDIT.

package file

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/suyuan32/simple-admin-file/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uint64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uint64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uint64) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uint64) predicate.File {
	return predicate.File(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uint64) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uint64) predicate.File {
	return predicate.File(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uint64) predicate.File {
	return predicate.File(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uint64) predicate.File {
	return predicate.File(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uint64) predicate.File {
	return predicate.File(sql.FieldLTE(FieldID, id))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUpdatedAt, v))
}

// Status applies equality check predicate on the "status" field. It's identical to StatusEQ.
func Status(v uint8) predicate.File {
	return predicate.File(sql.FieldEQ(FieldStatus, v))
}

// UUID applies equality check predicate on the "uuid" field. It's identical to UUIDEQ.
func UUID(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUUID, v))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// FileType applies equality check predicate on the "file_type" field. It's identical to FileTypeEQ.
func FileType(v uint8) predicate.File {
	return predicate.File(sql.FieldEQ(FieldFileType, v))
}

// Size applies equality check predicate on the "size" field. It's identical to SizeEQ.
func Size(v uint64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// Path applies equality check predicate on the "path" field. It's identical to PathEQ.
func Path(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldPath, v))
}

// UserUUID applies equality check predicate on the "user_uuid" field. It's identical to UserUUIDEQ.
func UserUUID(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUserUUID, v))
}

// Md5 applies equality check predicate on the "md5" field. It's identical to Md5EQ.
func Md5(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldMd5, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.File {
	return predicate.File(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.File {
	return predicate.File(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUpdatedAt, v))
}

// StatusEQ applies the EQ predicate on the "status" field.
func StatusEQ(v uint8) predicate.File {
	return predicate.File(sql.FieldEQ(FieldStatus, v))
}

// StatusNEQ applies the NEQ predicate on the "status" field.
func StatusNEQ(v uint8) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldStatus, v))
}

// StatusIn applies the In predicate on the "status" field.
func StatusIn(vs ...uint8) predicate.File {
	return predicate.File(sql.FieldIn(FieldStatus, vs...))
}

// StatusNotIn applies the NotIn predicate on the "status" field.
func StatusNotIn(vs ...uint8) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldStatus, vs...))
}

// StatusGT applies the GT predicate on the "status" field.
func StatusGT(v uint8) predicate.File {
	return predicate.File(sql.FieldGT(FieldStatus, v))
}

// StatusGTE applies the GTE predicate on the "status" field.
func StatusGTE(v uint8) predicate.File {
	return predicate.File(sql.FieldGTE(FieldStatus, v))
}

// StatusLT applies the LT predicate on the "status" field.
func StatusLT(v uint8) predicate.File {
	return predicate.File(sql.FieldLT(FieldStatus, v))
}

// StatusLTE applies the LTE predicate on the "status" field.
func StatusLTE(v uint8) predicate.File {
	return predicate.File(sql.FieldLTE(FieldStatus, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.File {
	return predicate.File(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.File {
	return predicate.File(sql.FieldNotNull(FieldStatus))
}

// UUIDEQ applies the EQ predicate on the "uuid" field.
func UUIDEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUUID, v))
}

// UUIDNEQ applies the NEQ predicate on the "uuid" field.
func UUIDNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUUID, v))
}

// UUIDIn applies the In predicate on the "uuid" field.
func UUIDIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldUUID, vs...))
}

// UUIDNotIn applies the NotIn predicate on the "uuid" field.
func UUIDNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUUID, vs...))
}

// UUIDGT applies the GT predicate on the "uuid" field.
func UUIDGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldUUID, v))
}

// UUIDGTE applies the GTE predicate on the "uuid" field.
func UUIDGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUUID, v))
}

// UUIDLT applies the LT predicate on the "uuid" field.
func UUIDLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldUUID, v))
}

// UUIDLTE applies the LTE predicate on the "uuid" field.
func UUIDLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUUID, v))
}

// UUIDContains applies the Contains predicate on the "uuid" field.
func UUIDContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldUUID, v))
}

// UUIDHasPrefix applies the HasPrefix predicate on the "uuid" field.
func UUIDHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldUUID, v))
}

// UUIDHasSuffix applies the HasSuffix predicate on the "uuid" field.
func UUIDHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldUUID, v))
}

// UUIDEqualFold applies the EqualFold predicate on the "uuid" field.
func UUIDEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldUUID, v))
}

// UUIDContainsFold applies the ContainsFold predicate on the "uuid" field.
func UUIDContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldUUID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldName, v))
}

// FileTypeEQ applies the EQ predicate on the "file_type" field.
func FileTypeEQ(v uint8) predicate.File {
	return predicate.File(sql.FieldEQ(FieldFileType, v))
}

// FileTypeNEQ applies the NEQ predicate on the "file_type" field.
func FileTypeNEQ(v uint8) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldFileType, v))
}

// FileTypeIn applies the In predicate on the "file_type" field.
func FileTypeIn(vs ...uint8) predicate.File {
	return predicate.File(sql.FieldIn(FieldFileType, vs...))
}

// FileTypeNotIn applies the NotIn predicate on the "file_type" field.
func FileTypeNotIn(vs ...uint8) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldFileType, vs...))
}

// FileTypeGT applies the GT predicate on the "file_type" field.
func FileTypeGT(v uint8) predicate.File {
	return predicate.File(sql.FieldGT(FieldFileType, v))
}

// FileTypeGTE applies the GTE predicate on the "file_type" field.
func FileTypeGTE(v uint8) predicate.File {
	return predicate.File(sql.FieldGTE(FieldFileType, v))
}

// FileTypeLT applies the LT predicate on the "file_type" field.
func FileTypeLT(v uint8) predicate.File {
	return predicate.File(sql.FieldLT(FieldFileType, v))
}

// FileTypeLTE applies the LTE predicate on the "file_type" field.
func FileTypeLTE(v uint8) predicate.File {
	return predicate.File(sql.FieldLTE(FieldFileType, v))
}

// SizeEQ applies the EQ predicate on the "size" field.
func SizeEQ(v uint64) predicate.File {
	return predicate.File(sql.FieldEQ(FieldSize, v))
}

// SizeNEQ applies the NEQ predicate on the "size" field.
func SizeNEQ(v uint64) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldSize, v))
}

// SizeIn applies the In predicate on the "size" field.
func SizeIn(vs ...uint64) predicate.File {
	return predicate.File(sql.FieldIn(FieldSize, vs...))
}

// SizeNotIn applies the NotIn predicate on the "size" field.
func SizeNotIn(vs ...uint64) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldSize, vs...))
}

// SizeGT applies the GT predicate on the "size" field.
func SizeGT(v uint64) predicate.File {
	return predicate.File(sql.FieldGT(FieldSize, v))
}

// SizeGTE applies the GTE predicate on the "size" field.
func SizeGTE(v uint64) predicate.File {
	return predicate.File(sql.FieldGTE(FieldSize, v))
}

// SizeLT applies the LT predicate on the "size" field.
func SizeLT(v uint64) predicate.File {
	return predicate.File(sql.FieldLT(FieldSize, v))
}

// SizeLTE applies the LTE predicate on the "size" field.
func SizeLTE(v uint64) predicate.File {
	return predicate.File(sql.FieldLTE(FieldSize, v))
}

// PathEQ applies the EQ predicate on the "path" field.
func PathEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldPath, v))
}

// PathNEQ applies the NEQ predicate on the "path" field.
func PathNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldPath, v))
}

// PathIn applies the In predicate on the "path" field.
func PathIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldPath, vs...))
}

// PathNotIn applies the NotIn predicate on the "path" field.
func PathNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldPath, vs...))
}

// PathGT applies the GT predicate on the "path" field.
func PathGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldPath, v))
}

// PathGTE applies the GTE predicate on the "path" field.
func PathGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldPath, v))
}

// PathLT applies the LT predicate on the "path" field.
func PathLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldPath, v))
}

// PathLTE applies the LTE predicate on the "path" field.
func PathLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldPath, v))
}

// PathContains applies the Contains predicate on the "path" field.
func PathContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldPath, v))
}

// PathHasPrefix applies the HasPrefix predicate on the "path" field.
func PathHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldPath, v))
}

// PathHasSuffix applies the HasSuffix predicate on the "path" field.
func PathHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldPath, v))
}

// PathEqualFold applies the EqualFold predicate on the "path" field.
func PathEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldPath, v))
}

// PathContainsFold applies the ContainsFold predicate on the "path" field.
func PathContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldPath, v))
}

// UserUUIDEQ applies the EQ predicate on the "user_uuid" field.
func UserUUIDEQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldUserUUID, v))
}

// UserUUIDNEQ applies the NEQ predicate on the "user_uuid" field.
func UserUUIDNEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldUserUUID, v))
}

// UserUUIDIn applies the In predicate on the "user_uuid" field.
func UserUUIDIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldUserUUID, vs...))
}

// UserUUIDNotIn applies the NotIn predicate on the "user_uuid" field.
func UserUUIDNotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldUserUUID, vs...))
}

// UserUUIDGT applies the GT predicate on the "user_uuid" field.
func UserUUIDGT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldUserUUID, v))
}

// UserUUIDGTE applies the GTE predicate on the "user_uuid" field.
func UserUUIDGTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldUserUUID, v))
}

// UserUUIDLT applies the LT predicate on the "user_uuid" field.
func UserUUIDLT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldUserUUID, v))
}

// UserUUIDLTE applies the LTE predicate on the "user_uuid" field.
func UserUUIDLTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldUserUUID, v))
}

// UserUUIDContains applies the Contains predicate on the "user_uuid" field.
func UserUUIDContains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldUserUUID, v))
}

// UserUUIDHasPrefix applies the HasPrefix predicate on the "user_uuid" field.
func UserUUIDHasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldUserUUID, v))
}

// UserUUIDHasSuffix applies the HasSuffix predicate on the "user_uuid" field.
func UserUUIDHasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldUserUUID, v))
}

// UserUUIDEqualFold applies the EqualFold predicate on the "user_uuid" field.
func UserUUIDEqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldUserUUID, v))
}

// UserUUIDContainsFold applies the ContainsFold predicate on the "user_uuid" field.
func UserUUIDContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldUserUUID, v))
}

// Md5EQ applies the EQ predicate on the "md5" field.
func Md5EQ(v string) predicate.File {
	return predicate.File(sql.FieldEQ(FieldMd5, v))
}

// Md5NEQ applies the NEQ predicate on the "md5" field.
func Md5NEQ(v string) predicate.File {
	return predicate.File(sql.FieldNEQ(FieldMd5, v))
}

// Md5In applies the In predicate on the "md5" field.
func Md5In(vs ...string) predicate.File {
	return predicate.File(sql.FieldIn(FieldMd5, vs...))
}

// Md5NotIn applies the NotIn predicate on the "md5" field.
func Md5NotIn(vs ...string) predicate.File {
	return predicate.File(sql.FieldNotIn(FieldMd5, vs...))
}

// Md5GT applies the GT predicate on the "md5" field.
func Md5GT(v string) predicate.File {
	return predicate.File(sql.FieldGT(FieldMd5, v))
}

// Md5GTE applies the GTE predicate on the "md5" field.
func Md5GTE(v string) predicate.File {
	return predicate.File(sql.FieldGTE(FieldMd5, v))
}

// Md5LT applies the LT predicate on the "md5" field.
func Md5LT(v string) predicate.File {
	return predicate.File(sql.FieldLT(FieldMd5, v))
}

// Md5LTE applies the LTE predicate on the "md5" field.
func Md5LTE(v string) predicate.File {
	return predicate.File(sql.FieldLTE(FieldMd5, v))
}

// Md5Contains applies the Contains predicate on the "md5" field.
func Md5Contains(v string) predicate.File {
	return predicate.File(sql.FieldContains(FieldMd5, v))
}

// Md5HasPrefix applies the HasPrefix predicate on the "md5" field.
func Md5HasPrefix(v string) predicate.File {
	return predicate.File(sql.FieldHasPrefix(FieldMd5, v))
}

// Md5HasSuffix applies the HasSuffix predicate on the "md5" field.
func Md5HasSuffix(v string) predicate.File {
	return predicate.File(sql.FieldHasSuffix(FieldMd5, v))
}

// Md5EqualFold applies the EqualFold predicate on the "md5" field.
func Md5EqualFold(v string) predicate.File {
	return predicate.File(sql.FieldEqualFold(FieldMd5, v))
}

// Md5ContainsFold applies the ContainsFold predicate on the "md5" field.
func Md5ContainsFold(v string) predicate.File {
	return predicate.File(sql.FieldContainsFold(FieldMd5, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.File) predicate.File {
	return predicate.File(func(s *sql.Selector) {
		p(s.Not())
	})
}
