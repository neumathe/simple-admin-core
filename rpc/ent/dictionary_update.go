// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/neumathe/simple-admin-core/rpc/ent/dictionary"
	"github.com/neumathe/simple-admin-core/rpc/ent/dictionarydetail"
	"github.com/neumathe/simple-admin-core/rpc/ent/predicate"
)

// DictionaryUpdate is the builder for updating Dictionary entities.
type DictionaryUpdate struct {
	config
	hooks     []Hook
	mutation  *DictionaryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the DictionaryUpdate builder.
func (du *DictionaryUpdate) Where(ps ...predicate.Dictionary) *DictionaryUpdate {
	du.mutation.Where(ps...)
	return du
}

// SetUpdatedAt sets the "updated_at" field.
func (du *DictionaryUpdate) SetUpdatedAt(t time.Time) *DictionaryUpdate {
	du.mutation.SetUpdatedAt(t)
	return du
}

// SetStatus sets the "status" field.
func (du *DictionaryUpdate) SetStatus(u uint8) *DictionaryUpdate {
	du.mutation.ResetStatus()
	du.mutation.SetStatus(u)
	return du
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (du *DictionaryUpdate) SetNillableStatus(u *uint8) *DictionaryUpdate {
	if u != nil {
		du.SetStatus(*u)
	}
	return du
}

// AddStatus adds u to the "status" field.
func (du *DictionaryUpdate) AddStatus(u int8) *DictionaryUpdate {
	du.mutation.AddStatus(u)
	return du
}

// ClearStatus clears the value of the "status" field.
func (du *DictionaryUpdate) ClearStatus() *DictionaryUpdate {
	du.mutation.ClearStatus()
	return du
}

// SetTitle sets the "title" field.
func (du *DictionaryUpdate) SetTitle(s string) *DictionaryUpdate {
	du.mutation.SetTitle(s)
	return du
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (du *DictionaryUpdate) SetNillableTitle(s *string) *DictionaryUpdate {
	if s != nil {
		du.SetTitle(*s)
	}
	return du
}

// SetName sets the "name" field.
func (du *DictionaryUpdate) SetName(s string) *DictionaryUpdate {
	du.mutation.SetName(s)
	return du
}

// SetNillableName sets the "name" field if the given value is not nil.
func (du *DictionaryUpdate) SetNillableName(s *string) *DictionaryUpdate {
	if s != nil {
		du.SetName(*s)
	}
	return du
}

// SetDesc sets the "desc" field.
func (du *DictionaryUpdate) SetDesc(s string) *DictionaryUpdate {
	du.mutation.SetDesc(s)
	return du
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (du *DictionaryUpdate) SetNillableDesc(s *string) *DictionaryUpdate {
	if s != nil {
		du.SetDesc(*s)
	}
	return du
}

// ClearDesc clears the value of the "desc" field.
func (du *DictionaryUpdate) ClearDesc() *DictionaryUpdate {
	du.mutation.ClearDesc()
	return du
}

// AddDictionaryDetailIDs adds the "dictionary_details" edge to the DictionaryDetail entity by IDs.
func (du *DictionaryUpdate) AddDictionaryDetailIDs(ids ...uint64) *DictionaryUpdate {
	du.mutation.AddDictionaryDetailIDs(ids...)
	return du
}

// AddDictionaryDetails adds the "dictionary_details" edges to the DictionaryDetail entity.
func (du *DictionaryUpdate) AddDictionaryDetails(d ...*DictionaryDetail) *DictionaryUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.AddDictionaryDetailIDs(ids...)
}

// Mutation returns the DictionaryMutation object of the builder.
func (du *DictionaryUpdate) Mutation() *DictionaryMutation {
	return du.mutation
}

// ClearDictionaryDetails clears all "dictionary_details" edges to the DictionaryDetail entity.
func (du *DictionaryUpdate) ClearDictionaryDetails() *DictionaryUpdate {
	du.mutation.ClearDictionaryDetails()
	return du
}

// RemoveDictionaryDetailIDs removes the "dictionary_details" edge to DictionaryDetail entities by IDs.
func (du *DictionaryUpdate) RemoveDictionaryDetailIDs(ids ...uint64) *DictionaryUpdate {
	du.mutation.RemoveDictionaryDetailIDs(ids...)
	return du
}

// RemoveDictionaryDetails removes "dictionary_details" edges to DictionaryDetail entities.
func (du *DictionaryUpdate) RemoveDictionaryDetails(d ...*DictionaryDetail) *DictionaryUpdate {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return du.RemoveDictionaryDetailIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (du *DictionaryUpdate) Save(ctx context.Context) (int, error) {
	du.defaults()
	return withHooks(ctx, du.sqlSave, du.mutation, du.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (du *DictionaryUpdate) SaveX(ctx context.Context) int {
	affected, err := du.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (du *DictionaryUpdate) Exec(ctx context.Context) error {
	_, err := du.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (du *DictionaryUpdate) ExecX(ctx context.Context) {
	if err := du.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (du *DictionaryUpdate) defaults() {
	if _, ok := du.mutation.UpdatedAt(); !ok {
		v := dictionary.UpdateDefaultUpdatedAt()
		du.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (du *DictionaryUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DictionaryUpdate {
	du.modifiers = append(du.modifiers, modifiers...)
	return du
}

func (du *DictionaryUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(dictionary.Table, dictionary.Columns, sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64))
	if ps := du.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := du.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionary.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := du.mutation.Status(); ok {
		_spec.SetField(dictionary.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := du.mutation.AddedStatus(); ok {
		_spec.AddField(dictionary.FieldStatus, field.TypeUint8, value)
	}
	if du.mutation.StatusCleared() {
		_spec.ClearField(dictionary.FieldStatus, field.TypeUint8)
	}
	if value, ok := du.mutation.Title(); ok {
		_spec.SetField(dictionary.FieldTitle, field.TypeString, value)
	}
	if value, ok := du.mutation.Name(); ok {
		_spec.SetField(dictionary.FieldName, field.TypeString, value)
	}
	if value, ok := du.mutation.Desc(); ok {
		_spec.SetField(dictionary.FieldDesc, field.TypeString, value)
	}
	if du.mutation.DescCleared() {
		_spec.ClearField(dictionary.FieldDesc, field.TypeString)
	}
	if du.mutation.DictionaryDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.RemovedDictionaryDetailsIDs(); len(nodes) > 0 && !du.mutation.DictionaryDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := du.mutation.DictionaryDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(du.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, du.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	du.mutation.done = true
	return n, nil
}

// DictionaryUpdateOne is the builder for updating a single Dictionary entity.
type DictionaryUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DictionaryMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdatedAt sets the "updated_at" field.
func (duo *DictionaryUpdateOne) SetUpdatedAt(t time.Time) *DictionaryUpdateOne {
	duo.mutation.SetUpdatedAt(t)
	return duo
}

// SetStatus sets the "status" field.
func (duo *DictionaryUpdateOne) SetStatus(u uint8) *DictionaryUpdateOne {
	duo.mutation.ResetStatus()
	duo.mutation.SetStatus(u)
	return duo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (duo *DictionaryUpdateOne) SetNillableStatus(u *uint8) *DictionaryUpdateOne {
	if u != nil {
		duo.SetStatus(*u)
	}
	return duo
}

// AddStatus adds u to the "status" field.
func (duo *DictionaryUpdateOne) AddStatus(u int8) *DictionaryUpdateOne {
	duo.mutation.AddStatus(u)
	return duo
}

// ClearStatus clears the value of the "status" field.
func (duo *DictionaryUpdateOne) ClearStatus() *DictionaryUpdateOne {
	duo.mutation.ClearStatus()
	return duo
}

// SetTitle sets the "title" field.
func (duo *DictionaryUpdateOne) SetTitle(s string) *DictionaryUpdateOne {
	duo.mutation.SetTitle(s)
	return duo
}

// SetNillableTitle sets the "title" field if the given value is not nil.
func (duo *DictionaryUpdateOne) SetNillableTitle(s *string) *DictionaryUpdateOne {
	if s != nil {
		duo.SetTitle(*s)
	}
	return duo
}

// SetName sets the "name" field.
func (duo *DictionaryUpdateOne) SetName(s string) *DictionaryUpdateOne {
	duo.mutation.SetName(s)
	return duo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (duo *DictionaryUpdateOne) SetNillableName(s *string) *DictionaryUpdateOne {
	if s != nil {
		duo.SetName(*s)
	}
	return duo
}

// SetDesc sets the "desc" field.
func (duo *DictionaryUpdateOne) SetDesc(s string) *DictionaryUpdateOne {
	duo.mutation.SetDesc(s)
	return duo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (duo *DictionaryUpdateOne) SetNillableDesc(s *string) *DictionaryUpdateOne {
	if s != nil {
		duo.SetDesc(*s)
	}
	return duo
}

// ClearDesc clears the value of the "desc" field.
func (duo *DictionaryUpdateOne) ClearDesc() *DictionaryUpdateOne {
	duo.mutation.ClearDesc()
	return duo
}

// AddDictionaryDetailIDs adds the "dictionary_details" edge to the DictionaryDetail entity by IDs.
func (duo *DictionaryUpdateOne) AddDictionaryDetailIDs(ids ...uint64) *DictionaryUpdateOne {
	duo.mutation.AddDictionaryDetailIDs(ids...)
	return duo
}

// AddDictionaryDetails adds the "dictionary_details" edges to the DictionaryDetail entity.
func (duo *DictionaryUpdateOne) AddDictionaryDetails(d ...*DictionaryDetail) *DictionaryUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.AddDictionaryDetailIDs(ids...)
}

// Mutation returns the DictionaryMutation object of the builder.
func (duo *DictionaryUpdateOne) Mutation() *DictionaryMutation {
	return duo.mutation
}

// ClearDictionaryDetails clears all "dictionary_details" edges to the DictionaryDetail entity.
func (duo *DictionaryUpdateOne) ClearDictionaryDetails() *DictionaryUpdateOne {
	duo.mutation.ClearDictionaryDetails()
	return duo
}

// RemoveDictionaryDetailIDs removes the "dictionary_details" edge to DictionaryDetail entities by IDs.
func (duo *DictionaryUpdateOne) RemoveDictionaryDetailIDs(ids ...uint64) *DictionaryUpdateOne {
	duo.mutation.RemoveDictionaryDetailIDs(ids...)
	return duo
}

// RemoveDictionaryDetails removes "dictionary_details" edges to DictionaryDetail entities.
func (duo *DictionaryUpdateOne) RemoveDictionaryDetails(d ...*DictionaryDetail) *DictionaryUpdateOne {
	ids := make([]uint64, len(d))
	for i := range d {
		ids[i] = d[i].ID
	}
	return duo.RemoveDictionaryDetailIDs(ids...)
}

// Where appends a list predicates to the DictionaryUpdate builder.
func (duo *DictionaryUpdateOne) Where(ps ...predicate.Dictionary) *DictionaryUpdateOne {
	duo.mutation.Where(ps...)
	return duo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (duo *DictionaryUpdateOne) Select(field string, fields ...string) *DictionaryUpdateOne {
	duo.fields = append([]string{field}, fields...)
	return duo
}

// Save executes the query and returns the updated Dictionary entity.
func (duo *DictionaryUpdateOne) Save(ctx context.Context) (*Dictionary, error) {
	duo.defaults()
	return withHooks(ctx, duo.sqlSave, duo.mutation, duo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (duo *DictionaryUpdateOne) SaveX(ctx context.Context) *Dictionary {
	node, err := duo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (duo *DictionaryUpdateOne) Exec(ctx context.Context) error {
	_, err := duo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (duo *DictionaryUpdateOne) ExecX(ctx context.Context) {
	if err := duo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (duo *DictionaryUpdateOne) defaults() {
	if _, ok := duo.mutation.UpdatedAt(); !ok {
		v := dictionary.UpdateDefaultUpdatedAt()
		duo.mutation.SetUpdatedAt(v)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (duo *DictionaryUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DictionaryUpdateOne {
	duo.modifiers = append(duo.modifiers, modifiers...)
	return duo
}

func (duo *DictionaryUpdateOne) sqlSave(ctx context.Context) (_node *Dictionary, err error) {
	_spec := sqlgraph.NewUpdateSpec(dictionary.Table, dictionary.Columns, sqlgraph.NewFieldSpec(dictionary.FieldID, field.TypeUint64))
	id, ok := duo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Dictionary.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := duo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, dictionary.FieldID)
		for _, f := range fields {
			if !dictionary.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != dictionary.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := duo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := duo.mutation.UpdatedAt(); ok {
		_spec.SetField(dictionary.FieldUpdatedAt, field.TypeTime, value)
	}
	if value, ok := duo.mutation.Status(); ok {
		_spec.SetField(dictionary.FieldStatus, field.TypeUint8, value)
	}
	if value, ok := duo.mutation.AddedStatus(); ok {
		_spec.AddField(dictionary.FieldStatus, field.TypeUint8, value)
	}
	if duo.mutation.StatusCleared() {
		_spec.ClearField(dictionary.FieldStatus, field.TypeUint8)
	}
	if value, ok := duo.mutation.Title(); ok {
		_spec.SetField(dictionary.FieldTitle, field.TypeString, value)
	}
	if value, ok := duo.mutation.Name(); ok {
		_spec.SetField(dictionary.FieldName, field.TypeString, value)
	}
	if value, ok := duo.mutation.Desc(); ok {
		_spec.SetField(dictionary.FieldDesc, field.TypeString, value)
	}
	if duo.mutation.DescCleared() {
		_spec.ClearField(dictionary.FieldDesc, field.TypeString)
	}
	if duo.mutation.DictionaryDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.RemovedDictionaryDetailsIDs(); len(nodes) > 0 && !duo.mutation.DictionaryDetailsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := duo.mutation.DictionaryDetailsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   dictionary.DictionaryDetailsTable,
			Columns: []string{dictionary.DictionaryDetailsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(dictionarydetail.FieldID, field.TypeUint64),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.AddModifiers(duo.modifiers...)
	_node = &Dictionary{config: duo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, duo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{dictionary.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	duo.mutation.done = true
	return _node, nil
}
