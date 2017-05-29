// This file is generated by SQLBoiler (https://github.com/vattle/sqlboiler)
// and is meant to be re-generated in place and/or deleted at any time.
// DO NOT EDIT

package models

import (
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/vattle/sqlboiler/boil"
	"github.com/vattle/sqlboiler/queries"
	"github.com/vattle/sqlboiler/queries/qm"
	"github.com/vattle/sqlboiler/strmangle"
	"gopkg.in/nullbio/null.v6"
)

// Reference is an object representing the database table.
type Reference struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	JVNID     null.Int    `boil:"jvn_id" json:"jvn_id,omitempty" toml:"jvn_id" yaml:"jvn_id,omitempty"`
	NVDID     null.Int    `boil:"nvd_id" json:"nvd_id,omitempty" toml:"nvd_id" yaml:"nvd_id,omitempty"`
	Source    null.String `boil:"source" json:"source,omitempty" toml:"source" yaml:"source,omitempty"`
	Link      null.String `boil:"link" json:"link,omitempty" toml:"link" yaml:"link,omitempty"`

	R *referenceR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L referenceL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// referenceR is where relationships are stored.
type referenceR struct {
}

// referenceL is where Load methods for each relationship are stored.
type referenceL struct{}

var (
	referenceColumns               = []string{"id", "created_at", "updated_at", "deleted_at", "jvn_id", "nvd_id", "source", "link"}
	referenceColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "jvn_id", "nvd_id", "source", "link"}
	referenceColumnsWithDefault    = []string{"id"}
	referencePrimaryKeyColumns     = []string{"id"}
)

type (
	// ReferenceSlice is an alias for a slice of pointers to Reference.
	// This should generally be used opposed to []Reference.
	ReferenceSlice []*Reference
	// ReferenceHook is the signature for custom Reference hook methods
	ReferenceHook func(boil.Executor, *Reference) error

	referenceQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	referenceType                 = reflect.TypeOf(&Reference{})
	referenceMapping              = queries.MakeStructMapping(referenceType)
	referencePrimaryKeyMapping, _ = queries.BindMapping(referenceType, referenceMapping, referencePrimaryKeyColumns)
	referenceInsertCacheMut       sync.RWMutex
	referenceInsertCache          = make(map[string]insertCache)
	referenceUpdateCacheMut       sync.RWMutex
	referenceUpdateCache          = make(map[string]updateCache)
	referenceUpsertCacheMut       sync.RWMutex
	referenceUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var referenceBeforeInsertHooks []ReferenceHook
var referenceBeforeUpdateHooks []ReferenceHook
var referenceBeforeDeleteHooks []ReferenceHook
var referenceBeforeUpsertHooks []ReferenceHook

var referenceAfterInsertHooks []ReferenceHook
var referenceAfterSelectHooks []ReferenceHook
var referenceAfterUpdateHooks []ReferenceHook
var referenceAfterDeleteHooks []ReferenceHook
var referenceAfterUpsertHooks []ReferenceHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Reference) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Reference) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Reference) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Reference) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Reference) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Reference) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Reference) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Reference) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Reference) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range referenceAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddReferenceHook registers your hook function for all future operations.
func AddReferenceHook(hookPoint boil.HookPoint, referenceHook ReferenceHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		referenceBeforeInsertHooks = append(referenceBeforeInsertHooks, referenceHook)
	case boil.BeforeUpdateHook:
		referenceBeforeUpdateHooks = append(referenceBeforeUpdateHooks, referenceHook)
	case boil.BeforeDeleteHook:
		referenceBeforeDeleteHooks = append(referenceBeforeDeleteHooks, referenceHook)
	case boil.BeforeUpsertHook:
		referenceBeforeUpsertHooks = append(referenceBeforeUpsertHooks, referenceHook)
	case boil.AfterInsertHook:
		referenceAfterInsertHooks = append(referenceAfterInsertHooks, referenceHook)
	case boil.AfterSelectHook:
		referenceAfterSelectHooks = append(referenceAfterSelectHooks, referenceHook)
	case boil.AfterUpdateHook:
		referenceAfterUpdateHooks = append(referenceAfterUpdateHooks, referenceHook)
	case boil.AfterDeleteHook:
		referenceAfterDeleteHooks = append(referenceAfterDeleteHooks, referenceHook)
	case boil.AfterUpsertHook:
		referenceAfterUpsertHooks = append(referenceAfterUpsertHooks, referenceHook)
	}
}

// OneP returns a single reference record from the query, and panics on error.
func (q referenceQuery) OneP() *Reference {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single reference record from the query.
func (q referenceQuery) One() (*Reference, error) {
	o := &Reference{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for references")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Reference records from the query, and panics on error.
func (q referenceQuery) AllP() ReferenceSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Reference records from the query.
func (q referenceQuery) All() (ReferenceSlice, error) {
	var o []*Reference

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Reference slice")
	}

	if len(referenceAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Reference records in the query, and panics on error.
func (q referenceQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Reference records in the query.
func (q referenceQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count references rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q referenceQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q referenceQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if references exists")
	}

	return count > 0, nil
}

// ReferencesG retrieves all records.
func ReferencesG(mods ...qm.QueryMod) referenceQuery {
	return References(boil.GetDB(), mods...)
}

// References retrieves all the records using an executor.
func References(exec boil.Executor, mods ...qm.QueryMod) referenceQuery {
	mods = append(mods, qm.From("\"references\""))
	return referenceQuery{NewQuery(exec, mods...)}
}

// FindReferenceG retrieves a single record by ID.
func FindReferenceG(id int, selectCols ...string) (*Reference, error) {
	return FindReference(boil.GetDB(), id, selectCols...)
}

// FindReferenceGP retrieves a single record by ID, and panics on error.
func FindReferenceGP(id int, selectCols ...string) *Reference {
	retobj, err := FindReference(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindReference retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindReference(exec boil.Executor, id int, selectCols ...string) (*Reference, error) {
	referenceObj := &Reference{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"references\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(referenceObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from references")
	}

	return referenceObj, nil
}

// FindReferenceP retrieves a single record by ID with an executor, and panics on error.
func FindReferenceP(exec boil.Executor, id int, selectCols ...string) *Reference {
	retobj, err := FindReference(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Reference) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Reference) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Reference) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Reference) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no references provided for insertion")
	}

	var err error
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	if o.UpdatedAt.Time.IsZero() {
		o.UpdatedAt.Time = currTime
		o.UpdatedAt.Valid = true
	}

	if err := o.doBeforeInsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(referenceColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	referenceInsertCacheMut.RLock()
	cache, cached := referenceInsertCache[key]
	referenceInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			referenceColumns,
			referenceColumnsWithDefault,
			referenceColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(referenceType, referenceMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(referenceType, referenceMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"references\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"references\" DEFAULT VALUES"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		if len(wl) != 0 {
			cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into references")
	}

	if !cached {
		referenceInsertCacheMut.Lock()
		referenceInsertCache[key] = cache
		referenceInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Reference record. See Update for
// whitelist behavior description.
func (o *Reference) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Reference record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Reference) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Reference, and panics on error.
// See Update for whitelist behavior description.
func (o *Reference) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Reference.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Reference) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	referenceUpdateCacheMut.RLock()
	cache, cached := referenceUpdateCache[key]
	referenceUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			referenceColumns,
			referencePrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update references, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"references\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, referencePrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(referenceType, referenceMapping, append(wl, referencePrimaryKeyColumns...))
		if err != nil {
			return err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, values)
	}

	_, err = exec.Exec(cache.query, values...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update references row")
	}

	if !cached {
		referenceUpdateCacheMut.Lock()
		referenceUpdateCache[key] = cache
		referenceUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q referenceQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q referenceQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for references")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ReferenceSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ReferenceSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ReferenceSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ReferenceSlice) UpdateAll(exec boil.Executor, cols M) error {
	ln := int64(len(o))
	if ln == 0 {
		return nil
	}

	if len(cols) == 0 {
		return errors.New("models: update all requires at least one column argument")
	}

	colNames := make([]string, len(cols))
	args := make([]interface{}, len(cols))

	i := 0
	for name, value := range cols {
		colNames[i] = name
		args[i] = value
		i++
	}

	// Append all of the primary key values for each column
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), referencePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"references\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, referencePrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in reference slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Reference) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Reference) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Reference) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Reference) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no references provided for upsert")
	}
	currTime := time.Now().In(boil.GetLocation())

	if o.CreatedAt.Time.IsZero() {
		o.CreatedAt.Time = currTime
		o.CreatedAt.Valid = true
	}
	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	if err := o.doBeforeUpsertHooks(exec); err != nil {
		return err
	}

	nzDefaults := queries.NonZeroDefaultSet(referenceColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs postgres problems
	buf := strmangle.GetBuffer()

	if updateOnConflict {
		buf.WriteByte('t')
	} else {
		buf.WriteByte('f')
	}
	buf.WriteByte('.')
	for _, c := range conflictColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range updateColumns {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range whitelist {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	referenceUpsertCacheMut.RLock()
	cache, cached := referenceUpsertCache[key]
	referenceUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			referenceColumns,
			referenceColumnsWithDefault,
			referenceColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			referenceColumns,
			referencePrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert references, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(referencePrimaryKeyColumns))
			copy(conflict, referencePrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"references\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(referenceType, referenceMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(referenceType, referenceMapping, ret)
			if err != nil {
				return err
			}
		}
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)
	var returns []interface{}
	if len(cache.retMapping) != 0 {
		returns = queries.PtrsFromMapping(value, cache.retMapping)
	}

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, cache.query)
		fmt.Fprintln(boil.DebugWriter, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRow(cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.Exec(cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert references")
	}

	if !cached {
		referenceUpsertCacheMut.Lock()
		referenceUpsertCache[key] = cache
		referenceUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Reference record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Reference) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Reference record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Reference) DeleteG() error {
	if o == nil {
		return errors.New("models: no Reference provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Reference record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Reference) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Reference record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Reference) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Reference provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), referencePrimaryKeyMapping)
	sql := "DELETE FROM \"references\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from references")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q referenceQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q referenceQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no referenceQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from references")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ReferenceSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ReferenceSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Reference slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ReferenceSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ReferenceSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Reference slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(referenceBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), referencePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"references\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, referencePrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from reference slice")
	}

	if len(referenceAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Reference) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Reference) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Reference) ReloadG() error {
	if o == nil {
		return errors.New("models: no Reference provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Reference) Reload(exec boil.Executor) error {
	ret, err := FindReference(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ReferenceSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ReferenceSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReferenceSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ReferenceSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ReferenceSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	references := ReferenceSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), referencePrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"references\".* FROM \"references\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, referencePrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&references)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ReferenceSlice")
	}

	*o = references

	return nil
}

// ReferenceExists checks if the Reference row exists.
func ReferenceExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"references\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if references exists")
	}

	return exists, nil
}

// ReferenceExistsG checks if the Reference row exists.
func ReferenceExistsG(id int) (bool, error) {
	return ReferenceExists(boil.GetDB(), id)
}

// ReferenceExistsGP checks if the Reference row exists. Panics on error.
func ReferenceExistsGP(id int) bool {
	e, err := ReferenceExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ReferenceExistsP checks if the Reference row exists. Panics on error.
func ReferenceExistsP(exec boil.Executor, id int) bool {
	e, err := ReferenceExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}