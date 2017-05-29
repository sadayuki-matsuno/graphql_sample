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

func testReferences(t *testing.T) {
	t.Parallel()

	query := References(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testReferencesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = reference.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReferencesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = References(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testReferencesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ReferenceSlice{reference}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testReferencesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ReferenceExists(tx, reference.ID)
	if err != nil {
		t.Errorf("Unable to check if Reference exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ReferenceExistsG to return true, but got false.")
	}
}
func testReferencesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	referenceFound, err := FindReference(tx, reference.ID)
	if err != nil {
		t.Error(err)
	}

	if referenceFound == nil {
		t.Error("want a record, got nil")
	}
}
func testReferencesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = References(tx).Bind(reference); err != nil {
		t.Error(err)
	}
}

func testReferencesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := References(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testReferencesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	referenceOne := &Reference{}
	referenceTwo := &Reference{}
	if err = randomize.Struct(seed, referenceOne, referenceDBTypes, false, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}
	if err = randomize.Struct(seed, referenceTwo, referenceDBTypes, false, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = referenceOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = referenceTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := References(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testReferencesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	referenceOne := &Reference{}
	referenceTwo := &Reference{}
	if err = randomize.Struct(seed, referenceOne, referenceDBTypes, false, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}
	if err = randomize.Struct(seed, referenceTwo, referenceDBTypes, false, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = referenceOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = referenceTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func referenceBeforeInsertHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceAfterInsertHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceAfterSelectHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceBeforeUpdateHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceAfterUpdateHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceBeforeDeleteHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceAfterDeleteHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceBeforeUpsertHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func referenceAfterUpsertHook(e boil.Executor, o *Reference) error {
	*o = Reference{}
	return nil
}

func testReferencesHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Reference{}
	o := &Reference{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, referenceDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Reference object: %s", err)
	}

	AddReferenceHook(boil.BeforeInsertHook, referenceBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	referenceBeforeInsertHooks = []ReferenceHook{}

	AddReferenceHook(boil.AfterInsertHook, referenceAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	referenceAfterInsertHooks = []ReferenceHook{}

	AddReferenceHook(boil.AfterSelectHook, referenceAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	referenceAfterSelectHooks = []ReferenceHook{}

	AddReferenceHook(boil.BeforeUpdateHook, referenceBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	referenceBeforeUpdateHooks = []ReferenceHook{}

	AddReferenceHook(boil.AfterUpdateHook, referenceAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	referenceAfterUpdateHooks = []ReferenceHook{}

	AddReferenceHook(boil.BeforeDeleteHook, referenceBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	referenceBeforeDeleteHooks = []ReferenceHook{}

	AddReferenceHook(boil.AfterDeleteHook, referenceAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	referenceAfterDeleteHooks = []ReferenceHook{}

	AddReferenceHook(boil.BeforeUpsertHook, referenceBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	referenceBeforeUpsertHooks = []ReferenceHook{}

	AddReferenceHook(boil.AfterUpsertHook, referenceAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	referenceAfterUpsertHooks = []ReferenceHook{}
}
func testReferencesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReferencesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx, referenceColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testReferencesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = reference.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testReferencesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ReferenceSlice{reference}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testReferencesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := References(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	referenceDBTypes = map[string]string{`CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `JVNID`: `integer`, `Link`: `character varying`, `NVDID`: `integer`, `Source`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_                = bytes.MinRead
)

func testReferencesUpdate(t *testing.T) {
	t.Parallel()

	if len(referenceColumns) == len(referencePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	if err = reference.Update(tx); err != nil {
		t.Error(err)
	}
}

func testReferencesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(referenceColumns) == len(referencePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	reference := &Reference{}
	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referenceColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, reference, referenceDBTypes, true, referencePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(referenceColumns, referencePrimaryKeyColumns) {
		fields = referenceColumns
	} else {
		fields = strmangle.SetComplement(
			referenceColumns,
			referencePrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(reference))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ReferenceSlice{reference}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testReferencesUpsert(t *testing.T) {
	t.Parallel()

	if len(referenceColumns) == len(referencePrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	reference := Reference{}
	if err = randomize.Struct(seed, &reference, referenceDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = reference.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Reference: %s", err)
	}

	count, err := References(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &reference, referenceDBTypes, false, referencePrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Reference struct: %s", err)
	}

	if err = reference.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Reference: %s", err)
	}

	count, err = References(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}