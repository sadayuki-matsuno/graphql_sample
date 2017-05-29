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

func testChangelogs(t *testing.T) {
	t.Parallel()

	query := Changelogs(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testChangelogsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = changelog.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChangelogsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Changelogs(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testChangelogsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ChangelogSlice{changelog}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testChangelogsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ChangelogExists(tx, changelog.ID)
	if err != nil {
		t.Errorf("Unable to check if Changelog exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ChangelogExistsG to return true, but got false.")
	}
}
func testChangelogsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	changelogFound, err := FindChangelog(tx, changelog.ID)
	if err != nil {
		t.Error(err)
	}

	if changelogFound == nil {
		t.Error("want a record, got nil")
	}
}
func testChangelogsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Changelogs(tx).Bind(changelog); err != nil {
		t.Error(err)
	}
}

func testChangelogsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Changelogs(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testChangelogsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelogOne := &Changelog{}
	changelogTwo := &Changelog{}
	if err = randomize.Struct(seed, changelogOne, changelogDBTypes, false, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}
	if err = randomize.Struct(seed, changelogTwo, changelogDBTypes, false, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelogOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = changelogTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Changelogs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testChangelogsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	changelogOne := &Changelog{}
	changelogTwo := &Changelog{}
	if err = randomize.Struct(seed, changelogOne, changelogDBTypes, false, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}
	if err = randomize.Struct(seed, changelogTwo, changelogDBTypes, false, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelogOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = changelogTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func changelogBeforeInsertHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogAfterInsertHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogAfterSelectHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogBeforeUpdateHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogAfterUpdateHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogBeforeDeleteHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogAfterDeleteHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogBeforeUpsertHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func changelogAfterUpsertHook(e boil.Executor, o *Changelog) error {
	*o = Changelog{}
	return nil
}

func testChangelogsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Changelog{}
	o := &Changelog{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, changelogDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Changelog object: %s", err)
	}

	AddChangelogHook(boil.BeforeInsertHook, changelogBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	changelogBeforeInsertHooks = []ChangelogHook{}

	AddChangelogHook(boil.AfterInsertHook, changelogAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	changelogAfterInsertHooks = []ChangelogHook{}

	AddChangelogHook(boil.AfterSelectHook, changelogAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	changelogAfterSelectHooks = []ChangelogHook{}

	AddChangelogHook(boil.BeforeUpdateHook, changelogBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	changelogBeforeUpdateHooks = []ChangelogHook{}

	AddChangelogHook(boil.AfterUpdateHook, changelogAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	changelogAfterUpdateHooks = []ChangelogHook{}

	AddChangelogHook(boil.BeforeDeleteHook, changelogBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	changelogBeforeDeleteHooks = []ChangelogHook{}

	AddChangelogHook(boil.AfterDeleteHook, changelogAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	changelogAfterDeleteHooks = []ChangelogHook{}

	AddChangelogHook(boil.BeforeUpsertHook, changelogBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	changelogBeforeUpsertHooks = []ChangelogHook{}

	AddChangelogHook(boil.AfterUpsertHook, changelogAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	changelogAfterUpsertHooks = []ChangelogHook{}
}
func testChangelogsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChangelogsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx, changelogColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testChangelogsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = changelog.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testChangelogsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ChangelogSlice{changelog}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testChangelogsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Changelogs(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	changelogDBTypes = map[string]string{`Contents`: `text`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `Method`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_                = bytes.MinRead
)

func testChangelogsUpdate(t *testing.T) {
	t.Parallel()

	if len(changelogColumns) == len(changelogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	if err = changelog.Update(tx); err != nil {
		t.Error(err)
	}
}

func testChangelogsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(changelogColumns) == len(changelogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	changelog := &Changelog{}
	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, changelog, changelogDBTypes, true, changelogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(changelogColumns, changelogPrimaryKeyColumns) {
		fields = changelogColumns
	} else {
		fields = strmangle.SetComplement(
			changelogColumns,
			changelogPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(changelog))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ChangelogSlice{changelog}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testChangelogsUpsert(t *testing.T) {
	t.Parallel()

	if len(changelogColumns) == len(changelogPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	changelog := Changelog{}
	if err = randomize.Struct(seed, &changelog, changelogDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = changelog.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Changelog: %s", err)
	}

	count, err := Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &changelog, changelogDBTypes, false, changelogPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Changelog struct: %s", err)
	}

	if err = changelog.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Changelog: %s", err)
	}

	count, err = Changelogs(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
