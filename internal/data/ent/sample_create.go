// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"go-starter/internal/data/ent/sample"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// SampleCreate is the builder for creating a Sample entity.
type SampleCreate struct {
	config
	mutation *SampleMutation
	hooks    []Hook
}

// SetName sets the "name" field.
func (sc *SampleCreate) SetName(s string) *SampleCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetNillableName sets the "name" field if the given value is not nil.
func (sc *SampleCreate) SetNillableName(s *string) *SampleCreate {
	if s != nil {
		sc.SetName(*s)
	}
	return sc
}

// SetContent sets the "content" field.
func (sc *SampleCreate) SetContent(s string) *SampleCreate {
	sc.mutation.SetContent(s)
	return sc
}

// SetNillableContent sets the "content" field if the given value is not nil.
func (sc *SampleCreate) SetNillableContent(s *string) *SampleCreate {
	if s != nil {
		sc.SetContent(*s)
	}
	return sc
}

// SetCreatedAt sets the "created_at" field.
func (sc *SampleCreate) SetCreatedAt(t time.Time) *SampleCreate {
	sc.mutation.SetCreatedAt(t)
	return sc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (sc *SampleCreate) SetNillableCreatedAt(t *time.Time) *SampleCreate {
	if t != nil {
		sc.SetCreatedAt(*t)
	}
	return sc
}

// SetUpdatedAt sets the "updated_at" field.
func (sc *SampleCreate) SetUpdatedAt(t time.Time) *SampleCreate {
	sc.mutation.SetUpdatedAt(t)
	return sc
}

// SetNillableUpdatedAt sets the "updated_at" field if the given value is not nil.
func (sc *SampleCreate) SetNillableUpdatedAt(t *time.Time) *SampleCreate {
	if t != nil {
		sc.SetUpdatedAt(*t)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SampleCreate) SetID(s string) *SampleCreate {
	sc.mutation.SetID(s)
	return sc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (sc *SampleCreate) SetNillableID(s *string) *SampleCreate {
	if s != nil {
		sc.SetID(*s)
	}
	return sc
}

// Mutation returns the SampleMutation object of the builder.
func (sc *SampleCreate) Mutation() *SampleMutation {
	return sc.mutation
}

// Save creates the Sample in the database.
func (sc *SampleCreate) Save(ctx context.Context) (*Sample, error) {
	sc.defaults()
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SampleCreate) SaveX(ctx context.Context) *Sample {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SampleCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SampleCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SampleCreate) defaults() {
	if _, ok := sc.mutation.CreatedAt(); !ok {
		v := sample.DefaultCreatedAt()
		sc.mutation.SetCreatedAt(v)
	}
	if _, ok := sc.mutation.UpdatedAt(); !ok {
		v := sample.DefaultUpdatedAt()
		sc.mutation.SetUpdatedAt(v)
	}
	if _, ok := sc.mutation.ID(); !ok {
		v := sample.DefaultID()
		sc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (sc *SampleCreate) check() error {
	return nil
}

func (sc *SampleCreate) sqlSave(ctx context.Context) (*Sample, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected Sample.ID type: %T", _spec.ID.Value)
		}
	}
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SampleCreate) createSpec() (*Sample, *sqlgraph.CreateSpec) {
	var (
		_node = &Sample{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(sample.Table, sqlgraph.NewFieldSpec(sample.FieldID, field.TypeString))
	)
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(sample.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Content(); ok {
		_spec.SetField(sample.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := sc.mutation.CreatedAt(); ok {
		_spec.SetField(sample.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := sc.mutation.UpdatedAt(); ok {
		_spec.SetField(sample.FieldUpdatedAt, field.TypeTime, value)
		_node.UpdatedAt = value
	}
	return _node, _spec
}

// SampleCreateBulk is the builder for creating many Sample entities in bulk.
type SampleCreateBulk struct {
	config
	err      error
	builders []*SampleCreate
}

// Save creates the Sample entities in the database.
func (scb *SampleCreateBulk) Save(ctx context.Context) ([]*Sample, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Sample, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SampleMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SampleCreateBulk) SaveX(ctx context.Context) []*Sample {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SampleCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SampleCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}
