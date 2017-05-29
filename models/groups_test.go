// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"reflect"
	"testing"

	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/randomize"
	"github.com/vattle/sqlboiler/strmangle"
)

func testGroups(t *testing.T) {
	t.Parallel()

	query := Groups(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testGroupsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = group.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Groups(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testGroupsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GroupSlice{group}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testGroupsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := GroupExists(tx, group.ID)
	if err != nil {
		t.Errorf("Unable to check if Group exists: %s", err)
	}
	if !e {
		t.Errorf("Expected GroupExistsG to return true, but got false.")
	}
}
func testGroupsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	groupFound, err := FindGroup(tx, group.ID)
	if err != nil {
		t.Error(err)
	}

	if groupFound == nil {
		t.Error("want a record, got nil")
	}
}
func testGroupsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Groups(tx).Bind(group); err != nil {
		t.Error(err)
	}
}

func testGroupsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Groups(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testGroupsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	groupOne := &Group{}
	groupTwo := &Group{}
	if err = randomize.Struct(seed, groupOne, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}
	if err = randomize.Struct(seed, groupTwo, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = groupTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Groups(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testGroupsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	groupOne := &Group{}
	groupTwo := &Group{}
	if err = randomize.Struct(seed, groupOne, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}
	if err = randomize.Struct(seed, groupTwo, groupDBTypes, false, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = groupOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = groupTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func groupBeforeInsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterInsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterSelectHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupBeforeUpdateHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterUpdateHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupBeforeDeleteHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterDeleteHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupBeforeUpsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func groupAfterUpsertHook(e boil.Executor, o *Group) error {
	*o = Group{}
	return nil
}

func testGroupsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Group{}
	o := &Group{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, groupDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Group object: %s", err)
	}

	AddGroupHook(boil.BeforeInsertHook, groupBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	groupBeforeInsertHooks = []GroupHook{}

	AddGroupHook(boil.AfterInsertHook, groupAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	groupAfterInsertHooks = []GroupHook{}

	AddGroupHook(boil.AfterSelectHook, groupAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	groupAfterSelectHooks = []GroupHook{}

	AddGroupHook(boil.BeforeUpdateHook, groupBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	groupBeforeUpdateHooks = []GroupHook{}

	AddGroupHook(boil.AfterUpdateHook, groupAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	groupAfterUpdateHooks = []GroupHook{}

	AddGroupHook(boil.BeforeDeleteHook, groupBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	groupBeforeDeleteHooks = []GroupHook{}

	AddGroupHook(boil.AfterDeleteHook, groupAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	groupAfterDeleteHooks = []GroupHook{}

	AddGroupHook(boil.BeforeUpsertHook, groupBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	groupBeforeUpsertHooks = []GroupHook{}

	AddGroupHook(boil.AfterUpsertHook, groupAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	groupAfterUpsertHooks = []GroupHook{}
}
func testGroupsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGroupsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx, groupColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testGroupsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = group.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testGroupsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := GroupSlice{group}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testGroupsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Groups(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	groupDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `Name`: `text`, `OrganizationID`: `integer`, `UpdatedAt`: `timestamp with time zone`}
	_            = bytes.MinRead
)

func testGroupsUpdate(t *testing.T) {
	t.Parallel()

	if len(groupColumns) == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if err = group.Update(tx); err != nil {
		t.Error(err)
	}
}

func testGroupsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(groupColumns) == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	group := &Group{}
	if err = randomize.Struct(seed, group, groupDBTypes, true, groupColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, group, groupDBTypes, true, groupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(groupColumns, groupPrimaryKeyColumns) {
		fields = groupColumns
	} else {
		fields = strmangle.SetComplement(
			groupColumns,
			groupPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(group))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := GroupSlice{group}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testGroupsUpsert(t *testing.T) {
	t.Parallel()

	if len(groupColumns) == len(groupPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	group := Group{}
	if err = randomize.Struct(seed, &group, groupDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = group.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Group: %s", err)
	}

	count, err := Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &group, groupDBTypes, false, groupPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Group struct: %s", err)
	}

	if err = group.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Group: %s", err)
	}

	count, err = Groups(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
