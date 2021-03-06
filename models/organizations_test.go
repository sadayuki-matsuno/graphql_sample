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

func testOrganizations(t *testing.T) {
	t.Parallel()

	query := Organizations(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testOrganizationsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organization.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organizations(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrganizationsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganizationSlice{organization}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testOrganizationsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := OrganizationExists(tx, organization.ID)
	if err != nil {
		t.Errorf("Unable to check if Organization exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrganizationExistsG to return true, but got false.")
	}
}
func testOrganizationsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	organizationFound, err := FindOrganization(tx, organization.ID)
	if err != nil {
		t.Error(err)
	}

	if organizationFound == nil {
		t.Error("want a record, got nil")
	}
}
func testOrganizationsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Organizations(tx).Bind(organization); err != nil {
		t.Error(err)
	}
}

func testOrganizationsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Organizations(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrganizationsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organizationOne := &Organization{}
	organizationTwo := &Organization{}
	if err = randomize.Struct(seed, organizationOne, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationTwo, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organizationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organizationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organizations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrganizationsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	organizationOne := &Organization{}
	organizationTwo := &Organization{}
	if err = randomize.Struct(seed, organizationOne, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}
	if err = randomize.Struct(seed, organizationTwo, organizationDBTypes, false, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organizationOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = organizationTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func organizationBeforeInsertHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationAfterInsertHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationAfterSelectHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationBeforeUpdateHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationAfterUpdateHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationBeforeDeleteHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationAfterDeleteHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationBeforeUpsertHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func organizationAfterUpsertHook(e boil.Executor, o *Organization) error {
	*o = Organization{}
	return nil
}

func testOrganizationsHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Organization{}
	o := &Organization{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, organizationDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Organization object: %s", err)
	}

	AddOrganizationHook(boil.BeforeInsertHook, organizationBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	organizationBeforeInsertHooks = []OrganizationHook{}

	AddOrganizationHook(boil.AfterInsertHook, organizationAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	organizationAfterInsertHooks = []OrganizationHook{}

	AddOrganizationHook(boil.AfterSelectHook, organizationAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	organizationAfterSelectHooks = []OrganizationHook{}

	AddOrganizationHook(boil.BeforeUpdateHook, organizationBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	organizationBeforeUpdateHooks = []OrganizationHook{}

	AddOrganizationHook(boil.AfterUpdateHook, organizationAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	organizationAfterUpdateHooks = []OrganizationHook{}

	AddOrganizationHook(boil.BeforeDeleteHook, organizationBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	organizationBeforeDeleteHooks = []OrganizationHook{}

	AddOrganizationHook(boil.AfterDeleteHook, organizationAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	organizationAfterDeleteHooks = []OrganizationHook{}

	AddOrganizationHook(boil.BeforeUpsertHook, organizationBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	organizationBeforeUpsertHooks = []OrganizationHook{}

	AddOrganizationHook(boil.AfterUpsertHook, organizationAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	organizationAfterUpsertHooks = []OrganizationHook{}
}
func testOrganizationsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx, organizationColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrganizationsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = organization.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := OrganizationSlice{organization}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testOrganizationsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Organizations(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	organizationDBTypes = map[string]string{`Address`: `text`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `Name`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_                   = bytes.MinRead
)

func testOrganizationsUpdate(t *testing.T) {
	t.Parallel()

	if len(organizationColumns) == len(organizationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	if err = organization.Update(tx); err != nil {
		t.Error(err)
	}
}

func testOrganizationsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(organizationColumns) == len(organizationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	organization := &Organization{}
	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, organization, organizationDBTypes, true, organizationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(organizationColumns, organizationPrimaryKeyColumns) {
		fields = organizationColumns
	} else {
		fields = strmangle.SetComplement(
			organizationColumns,
			organizationPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(organization))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := OrganizationSlice{organization}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testOrganizationsUpsert(t *testing.T) {
	t.Parallel()

	if len(organizationColumns) == len(organizationPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	organization := Organization{}
	if err = randomize.Struct(seed, &organization, organizationDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = organization.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organization: %s", err)
	}

	count, err := Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &organization, organizationDBTypes, false, organizationPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Organization struct: %s", err)
	}

	if err = organization.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Organization: %s", err)
	}

	count, err = Organizations(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
