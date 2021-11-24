// Code generated by entc, DO NOT EDIT.

package company

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/Wanted-Linx/linx-backend/api/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// BusinessNumber applies equality check predicate on the "business_number" field. It's identical to BusinessNumberEQ.
func BusinessNumber(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBusinessNumber), v))
	})
}

// Address applies equality check predicate on the "address" field. It's identical to AddressEQ.
func Address(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// ProfileImage applies equality check predicate on the "profile_image" field. It's identical to ProfileImageEQ.
func ProfileImage(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfileImage), v))
	})
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldName), v))
	})
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldName), v))
	})
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldName), v...))
	})
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldName), v...))
	})
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldName), v))
	})
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldName), v))
	})
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldName), v))
	})
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldName), v))
	})
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldName), v))
	})
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldName), v))
	})
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldName), v))
	})
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldName), v))
	})
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldName), v))
	})
}

// BusinessNumberEQ applies the EQ predicate on the "business_number" field.
func BusinessNumberEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberNEQ applies the NEQ predicate on the "business_number" field.
func BusinessNumberNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberIn applies the In predicate on the "business_number" field.
func BusinessNumberIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldBusinessNumber), v...))
	})
}

// BusinessNumberNotIn applies the NotIn predicate on the "business_number" field.
func BusinessNumberNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldBusinessNumber), v...))
	})
}

// BusinessNumberGT applies the GT predicate on the "business_number" field.
func BusinessNumberGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberGTE applies the GTE predicate on the "business_number" field.
func BusinessNumberGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberLT applies the LT predicate on the "business_number" field.
func BusinessNumberLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberLTE applies the LTE predicate on the "business_number" field.
func BusinessNumberLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberContains applies the Contains predicate on the "business_number" field.
func BusinessNumberContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberHasPrefix applies the HasPrefix predicate on the "business_number" field.
func BusinessNumberHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberHasSuffix applies the HasSuffix predicate on the "business_number" field.
func BusinessNumberHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberEqualFold applies the EqualFold predicate on the "business_number" field.
func BusinessNumberEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldBusinessNumber), v))
	})
}

// BusinessNumberContainsFold applies the ContainsFold predicate on the "business_number" field.
func BusinessNumberContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldBusinessNumber), v))
	})
}

// AddressEQ applies the EQ predicate on the "address" field.
func AddressEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAddress), v))
	})
}

// AddressNEQ applies the NEQ predicate on the "address" field.
func AddressNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAddress), v))
	})
}

// AddressIn applies the In predicate on the "address" field.
func AddressIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAddress), v...))
	})
}

// AddressNotIn applies the NotIn predicate on the "address" field.
func AddressNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAddress), v...))
	})
}

// AddressGT applies the GT predicate on the "address" field.
func AddressGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAddress), v))
	})
}

// AddressGTE applies the GTE predicate on the "address" field.
func AddressGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAddress), v))
	})
}

// AddressLT applies the LT predicate on the "address" field.
func AddressLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAddress), v))
	})
}

// AddressLTE applies the LTE predicate on the "address" field.
func AddressLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAddress), v))
	})
}

// AddressContains applies the Contains predicate on the "address" field.
func AddressContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAddress), v))
	})
}

// AddressHasPrefix applies the HasPrefix predicate on the "address" field.
func AddressHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAddress), v))
	})
}

// AddressHasSuffix applies the HasSuffix predicate on the "address" field.
func AddressHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAddress), v))
	})
}

// AddressIsNil applies the IsNil predicate on the "address" field.
func AddressIsNil() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldAddress)))
	})
}

// AddressNotNil applies the NotNil predicate on the "address" field.
func AddressNotNil() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldAddress)))
	})
}

// AddressEqualFold applies the EqualFold predicate on the "address" field.
func AddressEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAddress), v))
	})
}

// AddressContainsFold applies the ContainsFold predicate on the "address" field.
func AddressContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAddress), v))
	})
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldDescription), v))
	})
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldDescription), v))
	})
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldDescription), v...))
	})
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldDescription), v...))
	})
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldDescription), v))
	})
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldDescription), v))
	})
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldDescription), v))
	})
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldDescription), v))
	})
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldDescription), v))
	})
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldDescription), v))
	})
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldDescription), v))
	})
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldDescription)))
	})
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldDescription)))
	})
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldDescription), v))
	})
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldDescription), v))
	})
}

// ProfileImageEQ applies the EQ predicate on the "profile_image" field.
func ProfileImageEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldProfileImage), v))
	})
}

// ProfileImageNEQ applies the NEQ predicate on the "profile_image" field.
func ProfileImageNEQ(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldProfileImage), v))
	})
}

// ProfileImageIn applies the In predicate on the "profile_image" field.
func ProfileImageIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldProfileImage), v...))
	})
}

// ProfileImageNotIn applies the NotIn predicate on the "profile_image" field.
func ProfileImageNotIn(vs ...string) predicate.Company {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Company(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldProfileImage), v...))
	})
}

// ProfileImageGT applies the GT predicate on the "profile_image" field.
func ProfileImageGT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldProfileImage), v))
	})
}

// ProfileImageGTE applies the GTE predicate on the "profile_image" field.
func ProfileImageGTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldProfileImage), v))
	})
}

// ProfileImageLT applies the LT predicate on the "profile_image" field.
func ProfileImageLT(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldProfileImage), v))
	})
}

// ProfileImageLTE applies the LTE predicate on the "profile_image" field.
func ProfileImageLTE(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldProfileImage), v))
	})
}

// ProfileImageContains applies the Contains predicate on the "profile_image" field.
func ProfileImageContains(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldProfileImage), v))
	})
}

// ProfileImageHasPrefix applies the HasPrefix predicate on the "profile_image" field.
func ProfileImageHasPrefix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldProfileImage), v))
	})
}

// ProfileImageHasSuffix applies the HasSuffix predicate on the "profile_image" field.
func ProfileImageHasSuffix(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldProfileImage), v))
	})
}

// ProfileImageIsNil applies the IsNil predicate on the "profile_image" field.
func ProfileImageIsNil() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldProfileImage)))
	})
}

// ProfileImageNotNil applies the NotNil predicate on the "profile_image" field.
func ProfileImageNotNil() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldProfileImage)))
	})
}

// ProfileImageEqualFold applies the EqualFold predicate on the "profile_image" field.
func ProfileImageEqualFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldProfileImage), v))
	})
}

// ProfileImageContainsFold applies the ContainsFold predicate on the "profile_image" field.
func ProfileImageContainsFold(v string) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldProfileImage), v))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(UserInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
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
func Not(p predicate.Company) predicate.Company {
	return predicate.Company(func(s *sql.Selector) {
		p(s.Not())
	})
}