// Code generated by entc, DO NOT EDIT.

package book

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/wonesy/bookfahrt/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// Author applies equality check predicate on the "author" field. It's identical to AuthorEQ.
func Author(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// Slug applies equality check predicate on the "slug" field. It's identical to SlugEQ.
func Slug(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldTitle), v))
	})
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldTitle), v))
	})
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldTitle), v...))
	})
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldTitle), v...))
	})
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldTitle), v))
	})
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldTitle), v))
	})
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldTitle), v))
	})
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldTitle), v))
	})
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldTitle), v))
	})
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldTitle), v))
	})
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldTitle), v))
	})
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldTitle), v))
	})
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldTitle), v))
	})
}

// AuthorEQ applies the EQ predicate on the "author" field.
func AuthorEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldAuthor), v))
	})
}

// AuthorNEQ applies the NEQ predicate on the "author" field.
func AuthorNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldAuthor), v))
	})
}

// AuthorIn applies the In predicate on the "author" field.
func AuthorIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldAuthor), v...))
	})
}

// AuthorNotIn applies the NotIn predicate on the "author" field.
func AuthorNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldAuthor), v...))
	})
}

// AuthorGT applies the GT predicate on the "author" field.
func AuthorGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldAuthor), v))
	})
}

// AuthorGTE applies the GTE predicate on the "author" field.
func AuthorGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldAuthor), v))
	})
}

// AuthorLT applies the LT predicate on the "author" field.
func AuthorLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldAuthor), v))
	})
}

// AuthorLTE applies the LTE predicate on the "author" field.
func AuthorLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldAuthor), v))
	})
}

// AuthorContains applies the Contains predicate on the "author" field.
func AuthorContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldAuthor), v))
	})
}

// AuthorHasPrefix applies the HasPrefix predicate on the "author" field.
func AuthorHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldAuthor), v))
	})
}

// AuthorHasSuffix applies the HasSuffix predicate on the "author" field.
func AuthorHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldAuthor), v))
	})
}

// AuthorEqualFold applies the EqualFold predicate on the "author" field.
func AuthorEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldAuthor), v))
	})
}

// AuthorContainsFold applies the ContainsFold predicate on the "author" field.
func AuthorContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldAuthor), v))
	})
}

// SlugEQ applies the EQ predicate on the "slug" field.
func SlugEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldSlug), v))
	})
}

// SlugNEQ applies the NEQ predicate on the "slug" field.
func SlugNEQ(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldSlug), v))
	})
}

// SlugIn applies the In predicate on the "slug" field.
func SlugIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldSlug), v...))
	})
}

// SlugNotIn applies the NotIn predicate on the "slug" field.
func SlugNotIn(vs ...string) predicate.Book {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Book(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldSlug), v...))
	})
}

// SlugGT applies the GT predicate on the "slug" field.
func SlugGT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldSlug), v))
	})
}

// SlugGTE applies the GTE predicate on the "slug" field.
func SlugGTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldSlug), v))
	})
}

// SlugLT applies the LT predicate on the "slug" field.
func SlugLT(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldSlug), v))
	})
}

// SlugLTE applies the LTE predicate on the "slug" field.
func SlugLTE(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldSlug), v))
	})
}

// SlugContains applies the Contains predicate on the "slug" field.
func SlugContains(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldSlug), v))
	})
}

// SlugHasPrefix applies the HasPrefix predicate on the "slug" field.
func SlugHasPrefix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldSlug), v))
	})
}

// SlugHasSuffix applies the HasSuffix predicate on the "slug" field.
func SlugHasSuffix(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldSlug), v))
	})
}

// SlugEqualFold applies the EqualFold predicate on the "slug" field.
func SlugEqualFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldSlug), v))
	})
}

// SlugContainsFold applies the ContainsFold predicate on the "slug" field.
func SlugContainsFold(v string) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldSlug), v))
	})
}

// HasGenres applies the HasEdge predicate on the "genres" edge.
func HasGenres() predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenresTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GenresTable, GenresPrimaryKey...),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasGenresWith applies the HasEdge predicate on the "genres" edge with a given conditions (other predicates).
func HasGenresWith(preds ...predicate.Genre) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(GenresInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, GenresTable, GenresPrimaryKey...),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
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
func Not(p predicate.Book) predicate.Book {
	return predicate.Book(func(s *sql.Selector) {
		p(s.Not())
	})
}
