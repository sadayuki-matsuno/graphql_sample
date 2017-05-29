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

func testJVNS(t *testing.T) {
	t.Parallel()

	query := JVNS(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testJVNSDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jvn.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJVNSQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = JVNS(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testJVNSSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JVNSlice{jvn}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testJVNSExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := JVNExists(tx, jvn.ID)
	if err != nil {
		t.Errorf("Unable to check if JVN exists: %s", err)
	}
	if !e {
		t.Errorf("Expected JVNExistsG to return true, but got false.")
	}
}
func testJVNSFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	jvnFound, err := FindJVN(tx, jvn.ID)
	if err != nil {
		t.Error(err)
	}

	if jvnFound == nil {
		t.Error("want a record, got nil")
	}
}
func testJVNSBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = JVNS(tx).Bind(jvn); err != nil {
		t.Error(err)
	}
}

func testJVNSOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := JVNS(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testJVNSAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvnOne := &JVN{}
	jvnTwo := &JVN{}
	if err = randomize.Struct(seed, jvnOne, jvnDBTypes, false, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}
	if err = randomize.Struct(seed, jvnTwo, jvnDBTypes, false, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvnOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jvnTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := JVNS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testJVNSCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	jvnOne := &JVN{}
	jvnTwo := &JVN{}
	if err = randomize.Struct(seed, jvnOne, jvnDBTypes, false, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}
	if err = randomize.Struct(seed, jvnTwo, jvnDBTypes, false, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvnOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = jvnTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func jvnBeforeInsertHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnAfterInsertHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnAfterSelectHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnBeforeUpdateHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnAfterUpdateHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnBeforeDeleteHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnAfterDeleteHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnBeforeUpsertHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func jvnAfterUpsertHook(e boil.Executor, o *JVN) error {
	*o = JVN{}
	return nil
}

func testJVNSHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &JVN{}
	o := &JVN{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, jvnDBTypes, false); err != nil {
		t.Errorf("Unable to randomize JVN object: %s", err)
	}

	AddJVNHook(boil.BeforeInsertHook, jvnBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	jvnBeforeInsertHooks = []JVNHook{}

	AddJVNHook(boil.AfterInsertHook, jvnAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	jvnAfterInsertHooks = []JVNHook{}

	AddJVNHook(boil.AfterSelectHook, jvnAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	jvnAfterSelectHooks = []JVNHook{}

	AddJVNHook(boil.BeforeUpdateHook, jvnBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	jvnBeforeUpdateHooks = []JVNHook{}

	AddJVNHook(boil.AfterUpdateHook, jvnAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	jvnAfterUpdateHooks = []JVNHook{}

	AddJVNHook(boil.BeforeDeleteHook, jvnBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	jvnBeforeDeleteHooks = []JVNHook{}

	AddJVNHook(boil.AfterDeleteHook, jvnAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	jvnAfterDeleteHooks = []JVNHook{}

	AddJVNHook(boil.BeforeUpsertHook, jvnBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	jvnBeforeUpsertHooks = []JVNHook{}

	AddJVNHook(boil.AfterUpsertHook, jvnAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	jvnAfterUpsertHooks = []JVNHook{}
}
func testJVNSInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJVNSInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx, jvnColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testJVNSReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = jvn.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testJVNSReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := JVNSlice{jvn}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testJVNSSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := JVNS(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	jvnDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `CveDetailID`: `integer`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `JVNID`: `text`, `JVNLink`: `text`, `LastModifiedDate`: `timestamp with time zone`, `PublishedDate`: `timestamp with time zone`, `Score`: `numeric`, `Severity`: `text`, `Summary`: `character varying`, `Title`: `text`, `UpdatedAt`: `timestamp with time zone`, `Vector`: `text`}
	_          = bytes.MinRead
)

func testJVNSUpdate(t *testing.T) {
	t.Parallel()

	if len(jvnColumns) == len(jvnPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	if err = jvn.Update(tx); err != nil {
		t.Error(err)
	}
}

func testJVNSSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(jvnColumns) == len(jvnPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	jvn := &JVN{}
	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, jvn, jvnDBTypes, true, jvnPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(jvnColumns, jvnPrimaryKeyColumns) {
		fields = jvnColumns
	} else {
		fields = strmangle.SetComplement(
			jvnColumns,
			jvnPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(jvn))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := JVNSlice{jvn}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testJVNSUpsert(t *testing.T) {
	t.Parallel()

	if len(jvnColumns) == len(jvnPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	jvn := JVN{}
	if err = randomize.Struct(seed, &jvn, jvnDBTypes, true); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = jvn.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert JVN: %s", err)
	}

	count, err := JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &jvn, jvnDBTypes, false, jvnPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize JVN struct: %s", err)
	}

	if err = jvn.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert JVN: %s", err)
	}

	count, err = JVNS(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}