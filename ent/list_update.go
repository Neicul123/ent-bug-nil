// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/bug/ent/list"
	"entgo.io/bug/ent/predicate"
	"entgo.io/bug/ent/settings"
	"entgo.io/bug/ent/user"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ListUpdate is the builder for updating List entities.
type ListUpdate struct {
	config
	hooks     []Hook
	mutation  *ListMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the ListUpdate builder.
func (lu *ListUpdate) Where(ps ...predicate.List) *ListUpdate {
	lu.mutation.Where(ps...)
	return lu
}

// SetUserID sets the "user_id" field.
func (lu *ListUpdate) SetUserID(i int) *ListUpdate {
	lu.mutation.ResetUserID()
	lu.mutation.SetUserID(i)
	return lu
}

// AddUserID adds i to the "user_id" field.
func (lu *ListUpdate) AddUserID(i int) *ListUpdate {
	lu.mutation.AddUserID(i)
	return lu
}

// AddStudyingIDs adds the "studying" edge to the User entity by IDs.
func (lu *ListUpdate) AddStudyingIDs(ids ...int) *ListUpdate {
	lu.mutation.AddStudyingIDs(ids...)
	return lu
}

// AddStudying adds the "studying" edges to the User entity.
func (lu *ListUpdate) AddStudying(u ...*User) *ListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lu.AddStudyingIDs(ids...)
}

// AddSettingIDs adds the "settings" edge to the Settings entity by IDs.
func (lu *ListUpdate) AddSettingIDs(ids ...int) *ListUpdate {
	lu.mutation.AddSettingIDs(ids...)
	return lu
}

// AddSettings adds the "settings" edges to the Settings entity.
func (lu *ListUpdate) AddSettings(s ...*Settings) *ListUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return lu.AddSettingIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (lu *ListUpdate) Mutation() *ListMutation {
	return lu.mutation
}

// ClearStudying clears all "studying" edges to the User entity.
func (lu *ListUpdate) ClearStudying() *ListUpdate {
	lu.mutation.ClearStudying()
	return lu
}

// RemoveStudyingIDs removes the "studying" edge to User entities by IDs.
func (lu *ListUpdate) RemoveStudyingIDs(ids ...int) *ListUpdate {
	lu.mutation.RemoveStudyingIDs(ids...)
	return lu
}

// RemoveStudying removes "studying" edges to User entities.
func (lu *ListUpdate) RemoveStudying(u ...*User) *ListUpdate {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return lu.RemoveStudyingIDs(ids...)
}

// ClearSettings clears all "settings" edges to the Settings entity.
func (lu *ListUpdate) ClearSettings() *ListUpdate {
	lu.mutation.ClearSettings()
	return lu
}

// RemoveSettingIDs removes the "settings" edge to Settings entities by IDs.
func (lu *ListUpdate) RemoveSettingIDs(ids ...int) *ListUpdate {
	lu.mutation.RemoveSettingIDs(ids...)
	return lu
}

// RemoveSettings removes "settings" edges to Settings entities.
func (lu *ListUpdate) RemoveSettings(s ...*Settings) *ListUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return lu.RemoveSettingIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (lu *ListUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(lu.hooks) == 0 {
		affected, err = lu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			lu.mutation = mutation
			affected, err = lu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(lu.hooks) - 1; i >= 0; i-- {
			if lu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = lu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, lu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (lu *ListUpdate) SaveX(ctx context.Context) int {
	affected, err := lu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (lu *ListUpdate) Exec(ctx context.Context) error {
	_, err := lu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (lu *ListUpdate) ExecX(ctx context.Context) {
	if err := lu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (lu *ListUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ListUpdate {
	lu.modifiers = append(lu.modifiers, modifiers...)
	return lu
}

func (lu *ListUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: list.FieldID,
			},
		},
	}
	if ps := lu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := lu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: list.FieldUserID,
		})
	}
	if value, ok := lu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: list.FieldUserID,
		})
	}
	if lu.mutation.StudyingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   list.StudyingTable,
			Columns: list.StudyingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedStudyingIDs(); len(nodes) > 0 && !lu.mutation.StudyingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   list.StudyingTable,
			Columns: list.StudyingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.StudyingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   list.StudyingTable,
			Columns: list.StudyingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if lu.mutation.SettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   list.SettingsTable,
			Columns: []string{list.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settings.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.RemovedSettingsIDs(); len(nodes) > 0 && !lu.mutation.SettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   list.SettingsTable,
			Columns: []string{list.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settings.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := lu.mutation.SettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   list.SettingsTable,
			Columns: []string{list.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settings.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = lu.modifiers
	if n, err = sqlgraph.UpdateNodes(ctx, lu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{list.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// ListUpdateOne is the builder for updating a single List entity.
type ListUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ListMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUserID sets the "user_id" field.
func (luo *ListUpdateOne) SetUserID(i int) *ListUpdateOne {
	luo.mutation.ResetUserID()
	luo.mutation.SetUserID(i)
	return luo
}

// AddUserID adds i to the "user_id" field.
func (luo *ListUpdateOne) AddUserID(i int) *ListUpdateOne {
	luo.mutation.AddUserID(i)
	return luo
}

// AddStudyingIDs adds the "studying" edge to the User entity by IDs.
func (luo *ListUpdateOne) AddStudyingIDs(ids ...int) *ListUpdateOne {
	luo.mutation.AddStudyingIDs(ids...)
	return luo
}

// AddStudying adds the "studying" edges to the User entity.
func (luo *ListUpdateOne) AddStudying(u ...*User) *ListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return luo.AddStudyingIDs(ids...)
}

// AddSettingIDs adds the "settings" edge to the Settings entity by IDs.
func (luo *ListUpdateOne) AddSettingIDs(ids ...int) *ListUpdateOne {
	luo.mutation.AddSettingIDs(ids...)
	return luo
}

// AddSettings adds the "settings" edges to the Settings entity.
func (luo *ListUpdateOne) AddSettings(s ...*Settings) *ListUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return luo.AddSettingIDs(ids...)
}

// Mutation returns the ListMutation object of the builder.
func (luo *ListUpdateOne) Mutation() *ListMutation {
	return luo.mutation
}

// ClearStudying clears all "studying" edges to the User entity.
func (luo *ListUpdateOne) ClearStudying() *ListUpdateOne {
	luo.mutation.ClearStudying()
	return luo
}

// RemoveStudyingIDs removes the "studying" edge to User entities by IDs.
func (luo *ListUpdateOne) RemoveStudyingIDs(ids ...int) *ListUpdateOne {
	luo.mutation.RemoveStudyingIDs(ids...)
	return luo
}

// RemoveStudying removes "studying" edges to User entities.
func (luo *ListUpdateOne) RemoveStudying(u ...*User) *ListUpdateOne {
	ids := make([]int, len(u))
	for i := range u {
		ids[i] = u[i].ID
	}
	return luo.RemoveStudyingIDs(ids...)
}

// ClearSettings clears all "settings" edges to the Settings entity.
func (luo *ListUpdateOne) ClearSettings() *ListUpdateOne {
	luo.mutation.ClearSettings()
	return luo
}

// RemoveSettingIDs removes the "settings" edge to Settings entities by IDs.
func (luo *ListUpdateOne) RemoveSettingIDs(ids ...int) *ListUpdateOne {
	luo.mutation.RemoveSettingIDs(ids...)
	return luo
}

// RemoveSettings removes "settings" edges to Settings entities.
func (luo *ListUpdateOne) RemoveSettings(s ...*Settings) *ListUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return luo.RemoveSettingIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (luo *ListUpdateOne) Select(field string, fields ...string) *ListUpdateOne {
	luo.fields = append([]string{field}, fields...)
	return luo
}

// Save executes the query and returns the updated List entity.
func (luo *ListUpdateOne) Save(ctx context.Context) (*List, error) {
	var (
		err  error
		node *List
	)
	if len(luo.hooks) == 0 {
		node, err = luo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ListMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			luo.mutation = mutation
			node, err = luo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(luo.hooks) - 1; i >= 0; i-- {
			if luo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = luo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, luo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*List)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from ListMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (luo *ListUpdateOne) SaveX(ctx context.Context) *List {
	node, err := luo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (luo *ListUpdateOne) Exec(ctx context.Context) error {
	_, err := luo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (luo *ListUpdateOne) ExecX(ctx context.Context) {
	if err := luo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (luo *ListUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ListUpdateOne {
	luo.modifiers = append(luo.modifiers, modifiers...)
	return luo
}

func (luo *ListUpdateOne) sqlSave(ctx context.Context) (_node *List, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   list.Table,
			Columns: list.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: list.FieldID,
			},
		},
	}
	id, ok := luo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "List.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := luo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, list.FieldID)
		for _, f := range fields {
			if !list.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != list.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := luo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := luo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: list.FieldUserID,
		})
	}
	if value, ok := luo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: list.FieldUserID,
		})
	}
	if luo.mutation.StudyingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   list.StudyingTable,
			Columns: list.StudyingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedStudyingIDs(); len(nodes) > 0 && !luo.mutation.StudyingCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   list.StudyingTable,
			Columns: list.StudyingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.StudyingIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   list.StudyingTable,
			Columns: list.StudyingPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if luo.mutation.SettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   list.SettingsTable,
			Columns: []string{list.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settings.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.RemovedSettingsIDs(); len(nodes) > 0 && !luo.mutation.SettingsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   list.SettingsTable,
			Columns: []string{list.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settings.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := luo.mutation.SettingsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   list.SettingsTable,
			Columns: []string{list.SettingsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: settings.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_spec.Modifiers = luo.modifiers
	_node = &List{config: luo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, luo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{list.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}