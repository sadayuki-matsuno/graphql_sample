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

func testTasks(t *testing.T) {
	t.Parallel()

	query := Tasks(nil)

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}
func testTasksDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = task.Delete(tx); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tasks(tx).DeleteAll(); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testTasksSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TaskSlice{task}

	if err = slice.DeleteAll(tx); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}
func testTasksExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	e, err := TaskExists(tx, task.ID)
	if err != nil {
		t.Errorf("Unable to check if Task exists: %s", err)
	}
	if !e {
		t.Errorf("Expected TaskExistsG to return true, but got false.")
	}
}
func testTasksFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	taskFound, err := FindTask(tx, task.ID)
	if err != nil {
		t.Error(err)
	}

	if taskFound == nil {
		t.Error("want a record, got nil")
	}
}
func testTasksBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = Tasks(tx).Bind(task); err != nil {
		t.Error(err)
	}
}

func testTasksOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	if x, err := Tasks(tx).One(); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testTasksAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	taskOne := &Task{}
	taskTwo := &Task{}
	if err = randomize.Struct(seed, taskOne, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTwo, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = taskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testTasksCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	taskOne := &Task{}
	taskTwo := &Task{}
	if err = randomize.Struct(seed, taskOne, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}
	if err = randomize.Struct(seed, taskTwo, taskDBTypes, false, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = taskOne.Insert(tx); err != nil {
		t.Error(err)
	}
	if err = taskTwo.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}
func taskBeforeInsertHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterInsertHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterSelectHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeUpdateHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterUpdateHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeDeleteHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterDeleteHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskBeforeUpsertHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func taskAfterUpsertHook(e boil.Executor, o *Task) error {
	*o = Task{}
	return nil
}

func testTasksHooks(t *testing.T) {
	t.Parallel()

	var err error

	empty := &Task{}
	o := &Task{}

	seed := randomize.NewSeed()
	if err = randomize.Struct(seed, o, taskDBTypes, false); err != nil {
		t.Errorf("Unable to randomize Task object: %s", err)
	}

	AddTaskHook(boil.BeforeInsertHook, taskBeforeInsertHook)
	if err = o.doBeforeInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeInsertHook function to empty object, but got: %#v", o)
	}
	taskBeforeInsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterInsertHook, taskAfterInsertHook)
	if err = o.doAfterInsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterInsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterInsertHook function to empty object, but got: %#v", o)
	}
	taskAfterInsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterSelectHook, taskAfterSelectHook)
	if err = o.doAfterSelectHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterSelectHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterSelectHook function to empty object, but got: %#v", o)
	}
	taskAfterSelectHooks = []TaskHook{}

	AddTaskHook(boil.BeforeUpdateHook, taskBeforeUpdateHook)
	if err = o.doBeforeUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpdateHook function to empty object, but got: %#v", o)
	}
	taskBeforeUpdateHooks = []TaskHook{}

	AddTaskHook(boil.AfterUpdateHook, taskAfterUpdateHook)
	if err = o.doAfterUpdateHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpdateHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpdateHook function to empty object, but got: %#v", o)
	}
	taskAfterUpdateHooks = []TaskHook{}

	AddTaskHook(boil.BeforeDeleteHook, taskBeforeDeleteHook)
	if err = o.doBeforeDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeDeleteHook function to empty object, but got: %#v", o)
	}
	taskBeforeDeleteHooks = []TaskHook{}

	AddTaskHook(boil.AfterDeleteHook, taskAfterDeleteHook)
	if err = o.doAfterDeleteHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterDeleteHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterDeleteHook function to empty object, but got: %#v", o)
	}
	taskAfterDeleteHooks = []TaskHook{}

	AddTaskHook(boil.BeforeUpsertHook, taskBeforeUpsertHook)
	if err = o.doBeforeUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doBeforeUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected BeforeUpsertHook function to empty object, but got: %#v", o)
	}
	taskBeforeUpsertHooks = []TaskHook{}

	AddTaskHook(boil.AfterUpsertHook, taskAfterUpsertHook)
	if err = o.doAfterUpsertHooks(nil); err != nil {
		t.Errorf("Unable to execute doAfterUpsertHooks: %s", err)
	}
	if !reflect.DeepEqual(o, empty) {
		t.Errorf("Expected AfterUpsertHook function to empty object, but got: %#v", o)
	}
	taskAfterUpsertHooks = []TaskHook{}
}
func testTasksInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTasksInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx, taskColumnsWithoutDefault...); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testTasksReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	if err = task.Reload(tx); err != nil {
		t.Error(err)
	}
}

func testTasksReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	slice := TaskSlice{task}

	if err = slice.ReloadAll(tx); err != nil {
		t.Error(err)
	}
}
func testTasksSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	slice, err := Tasks(tx).All()
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	taskDBTypes = map[string]string{`Comments`: `text`, `CreatedAt`: `timestamp with time zone`, `CveID`: `text`, `DeletedAt`: `timestamp with time zone`, `ID`: `integer`, `OrganizationID`: `text`, `ServerName`: `text`, `UpdatedAt`: `timestamp with time zone`}
	_           = bytes.MinRead
)

func testTasksUpdate(t *testing.T) {
	t.Parallel()

	if len(taskColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err = task.Update(tx); err != nil {
		t.Error(err)
	}
}

func testTasksSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(taskColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	task := &Task{}
	if err = randomize.Struct(seed, task, taskDBTypes, true, taskColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Insert(tx); err != nil {
		t.Error(err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, task, taskDBTypes, true, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(taskColumns, taskPrimaryKeyColumns) {
		fields = taskColumns
	} else {
		fields = strmangle.SetComplement(
			taskColumns,
			taskPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(task))
	updateMap := M{}
	for _, col := range fields {
		updateMap[col] = value.FieldByName(strmangle.TitleCase(col)).Interface()
	}

	slice := TaskSlice{task}
	if err = slice.UpdateAll(tx, updateMap); err != nil {
		t.Error(err)
	}
}
func testTasksUpsert(t *testing.T) {
	t.Parallel()

	if len(taskColumns) == len(taskPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	task := Task{}
	if err = randomize.Struct(seed, &task, taskDBTypes, true); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	tx := MustTx(boil.Begin())
	defer tx.Rollback()
	if err = task.Upsert(tx, false, nil, nil); err != nil {
		t.Errorf("Unable to upsert Task: %s", err)
	}

	count, err := Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &task, taskDBTypes, false, taskPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize Task struct: %s", err)
	}

	if err = task.Upsert(tx, true, nil, nil); err != nil {
		t.Errorf("Unable to upsert Task: %s", err)
	}

	count, err = Tasks(tx).Count()
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
