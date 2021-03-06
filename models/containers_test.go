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

func testContainers(t *testing.T) {
	t.Parallel()

	query := Containers(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testContainersDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = container.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContainersQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Containers(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testContainersSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContainerSlice{container}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testContainersExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := ContainerExists(tx, container.ID)
	if err != nil {
		t.Errorf("Unable to check if Container exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ContainerExistsG to return true, but got false.")
	}
}
func testContainersFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	containerFound, err := FindContainer(tx, container.ID)
	if err != nil {
		t.Error(err)
	}

	if containerFound == nil {
		t.Error("want a record, got nil")
	}
}
func testContainersBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Containers(tx).Bind(container); err != nil {
		t.Error(err)
	}
}

func testContainersOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Containers(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testContainersAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	containerOne := &Container{}
	containerTwo := &Container{}
	if err = randomize.Struct(seed, containerOne, containerDBTypes, false, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	if err = randomize.Struct(seed, containerTwo, containerDBTypes, false, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = containerOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = containerTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Containers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testContainersCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	containerOne := &Container{}
	containerTwo := &Container{}
	if err = randomize.Struct(seed, containerOne, containerDBTypes, false, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}
	if err = randomize.Struct(seed, containerTwo, containerDBTypes, false, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = containerOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = containerTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func containerBeforeInsertHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerAfterInsertHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerAfterSelectHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerBeforeUpdateHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerAfterUpdateHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerBeforeDeleteHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerAfterDeleteHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerBeforeUpsertHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func containerAfterUpsertHook(e boil.Executor, o *Container) error {
	*o = Container{}
	return nil
}

func testContainersHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Container{}
	o := &Container{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, containerDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Container object: %s", err)
	}

	AddContainerHook(boil.BeforeInsertHook, containerBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	containerBeforeInsertHooks = []ContainerHook{}

	AddContainerHook(boil.AfterInsertHook, containerAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	containerAfterInsertHooks = []ContainerHook{}

	AddContainerHook(boil.AfterSelectHook, containerAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	containerAfterSelectHooks = []ContainerHook{}

	AddContainerHook(boil.BeforeUpdateHook, containerBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	containerBeforeUpdateHooks = []ContainerHook{}

	AddContainerHook(boil.AfterUpdateHook, containerAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	containerAfterUpdateHooks = []ContainerHook{}

	AddContainerHook(boil.BeforeDeleteHook, containerBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	containerBeforeDeleteHooks = []ContainerHook{}

	AddContainerHook(boil.AfterDeleteHook, containerAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	containerAfterDeleteHooks = []ContainerHook{}

	AddContainerHook(boil.BeforeUpsertHook, containerBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	containerBeforeUpsertHooks = []ContainerHook{}

	AddContainerHook(boil.AfterUpsertHook, containerAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	containerAfterUpsertHooks = []ContainerHook{}
}
func testContainersInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContainersInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx, containerColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testContainersReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = container.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testContainersReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := ContainerSlice{container}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testContainersSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Containers(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	containerDBTypes = map[string]string{`ContainerID`: `text`, `CreatedAt`: `timestamp with time zone`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `Image`: `text`, `Name`: `text`, `Type`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_                = bytes.MinRead
)

func testContainersUpdate(t *testing.T) {
	t.Parallel()

	if len(containerColumns) == len(containerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	if err = container.Update(tx); err != nil {
		t.Error(err)
	}
}

func testContainersSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(containerColumns) == len(containerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	container := &Container{}
	if err = randomize.Struct(seed, container, containerDBTypes, true, containerColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, container, containerDBTypes, true, containerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(containerColumns, containerPrimaryKeyColumns) {
		fields = containerColumns
	} else {
		fields = strmangle.SetComplement(
			containerColumns,
			containerPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(container))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := ContainerSlice{container}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testContainersUpsert(t *testing.T) {
	t.Parallel()

	if len(containerColumns) == len(containerPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	container := Container{}
	if err = randomize.Struct(seed, &container, containerDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = container.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Container: %s", err)
	}

	count, err := Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &container, containerDBTypes, false, containerPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Container struct: %s", err)
	}

	if err = container.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Container: %s", err)
	}

	count, err = Containers(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
