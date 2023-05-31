// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"github.com/adshao/ordinals-indexer/internal/data/ent/inscription"
	"github.com/adshao/ordinals-indexer/internal/data/ent/predicate"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// InscriptionDelete is the builder for deleting a Inscription entity.
type InscriptionDelete struct {
	config
	hooks    []Hook
	mutation *InscriptionMutation
}

// Where appends a list predicates to the InscriptionDelete builder.
func (id *InscriptionDelete) Where(ps ...predicate.Inscription) *InscriptionDelete {
	id.mutation.Where(ps...)
	return id
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (id *InscriptionDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, id.sqlExec, id.mutation, id.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (id *InscriptionDelete) ExecX(ctx context.Context) int {
	n, err := id.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (id *InscriptionDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(inscription.Table, sqlgraph.NewFieldSpec(inscription.FieldID, field.TypeInt))
	if ps := id.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, id.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	id.mutation.done = true
	return affected, err
}

// InscriptionDeleteOne is the builder for deleting a single Inscription entity.
type InscriptionDeleteOne struct {
	id *InscriptionDelete
}

// Where appends a list predicates to the InscriptionDelete builder.
func (ido *InscriptionDeleteOne) Where(ps ...predicate.Inscription) *InscriptionDeleteOne {
	ido.id.mutation.Where(ps...)
	return ido
}

// Exec executes the deletion query.
func (ido *InscriptionDeleteOne) Exec(ctx context.Context) error {
	n, err := ido.id.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{inscription.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (ido *InscriptionDeleteOne) ExecX(ctx context.Context) {
	if err := ido.Exec(ctx); err != nil {
		panic(err)
	}
}
