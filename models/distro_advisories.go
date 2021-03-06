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

// DistroAdvisory is an object representing the database table.
type DistroAdvisory struct {
	ID         int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt  null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt  null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt  null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	AdvisoryID null.String `boil:"advisory_id" json:"advisory_id,omitempty" toml:"advisory_id" yaml:"advisory_id,omitempty"`
	Severity   null.String `boil:"severity" json:"severity,omitempty" toml:"severity" yaml:"severity,omitempty"`
	Issued     null.Time   `boil:"issued" json:"issued,omitempty" toml:"issued" yaml:"issued,omitempty"`
	Updated    null.Time   `boil:"updated" json:"updated,omitempty" toml:"updated" yaml:"updated,omitempty"`

	R *distroAdvisoryR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L distroAdvisoryL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// distroAdvisoryR is where relationships are stored.
type distroAdvisoryR struct {
}

// distroAdvisoryL is where Load methods for each relationship are stored.
type distroAdvisoryL struct{}

var (
	distroAdvisoryColumns               = []string{"id", "created_at", "updated_at", "deleted_at", "advisory_id", "severity", "issued", "updated"}
	distroAdvisoryColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "advisory_id", "severity", "issued", "updated"}
	distroAdvisoryColumnsWithDefault    = []string{"id"}
	distroAdvisoryPrimaryKeyColumns     = []string{"id"}
)

type (
	// DistroAdvisorySlice is an alias for a slice of pointers to DistroAdvisory.
	// This should generally be used opposed to []DistroAdvisory.
	DistroAdvisorySlice []*DistroAdvisory
	// DistroAdvisoryHook is the signature for custom DistroAdvisory hook methods
	DistroAdvisoryHook func(boil.Executor, *DistroAdvisory) error

	distroAdvisoryQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	distroAdvisoryType                 = reflect.TypeOf(&DistroAdvisory{})
	distroAdvisoryMapping              = queries.MakeStructMapping(distroAdvisoryType)
	distroAdvisoryPrimaryKeyMapping, _ = queries.BindMapping(distroAdvisoryType, distroAdvisoryMapping, distroAdvisoryPrimaryKeyColumns)
	distroAdvisoryInsertCacheMut       sync.RWMutex
	distroAdvisoryInsertCache          = make(map[string]insertCache)
	distroAdvisoryUpdateCacheMut       sync.RWMutex
	distroAdvisoryUpdateCache          = make(map[string]updateCache)
	distroAdvisoryUpsertCacheMut       sync.RWMutex
	distroAdvisoryUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var distroAdvisoryBeforeInsertHooks []DistroAdvisoryHook
var distroAdvisoryBeforeUpdateHooks []DistroAdvisoryHook
var distroAdvisoryBeforeDeleteHooks []DistroAdvisoryHook
var distroAdvisoryBeforeUpsertHooks []DistroAdvisoryHook

var distroAdvisoryAfterInsertHooks []DistroAdvisoryHook
var distroAdvisoryAfterSelectHooks []DistroAdvisoryHook
var distroAdvisoryAfterUpdateHooks []DistroAdvisoryHook
var distroAdvisoryAfterDeleteHooks []DistroAdvisoryHook
var distroAdvisoryAfterUpsertHooks []DistroAdvisoryHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *DistroAdvisory) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *DistroAdvisory) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *DistroAdvisory) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *DistroAdvisory) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *DistroAdvisory) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *DistroAdvisory) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *DistroAdvisory) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *DistroAdvisory) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *DistroAdvisory) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range distroAdvisoryAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddDistroAdvisoryHook registers your hook function for all future operations.
func AddDistroAdvisoryHook(hookPoint boil.HookPoint, distroAdvisoryHook DistroAdvisoryHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		distroAdvisoryBeforeInsertHooks = append(distroAdvisoryBeforeInsertHooks, distroAdvisoryHook)
	case boil.BeforeUpdateHook:
		distroAdvisoryBeforeUpdateHooks = append(distroAdvisoryBeforeUpdateHooks, distroAdvisoryHook)
	case boil.BeforeDeleteHook:
		distroAdvisoryBeforeDeleteHooks = append(distroAdvisoryBeforeDeleteHooks, distroAdvisoryHook)
	case boil.BeforeUpsertHook:
		distroAdvisoryBeforeUpsertHooks = append(distroAdvisoryBeforeUpsertHooks, distroAdvisoryHook)
	case boil.AfterInsertHook:
		distroAdvisoryAfterInsertHooks = append(distroAdvisoryAfterInsertHooks, distroAdvisoryHook)
	case boil.AfterSelectHook:
		distroAdvisoryAfterSelectHooks = append(distroAdvisoryAfterSelectHooks, distroAdvisoryHook)
	case boil.AfterUpdateHook:
		distroAdvisoryAfterUpdateHooks = append(distroAdvisoryAfterUpdateHooks, distroAdvisoryHook)
	case boil.AfterDeleteHook:
		distroAdvisoryAfterDeleteHooks = append(distroAdvisoryAfterDeleteHooks, distroAdvisoryHook)
	case boil.AfterUpsertHook:
		distroAdvisoryAfterUpsertHooks = append(distroAdvisoryAfterUpsertHooks, distroAdvisoryHook)
	}
}

// OneP returns a single distroAdvisory record from the query, and panics on error.
func (q distroAdvisoryQuery) OneP() *DistroAdvisory {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single distroAdvisory record from the query.
func (q distroAdvisoryQuery) One() (*DistroAdvisory, error) {
	o := &DistroAdvisory{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for distro_advisories")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all DistroAdvisory records from the query, and panics on error.
func (q distroAdvisoryQuery) AllP() DistroAdvisorySlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all DistroAdvisory records from the query.
func (q distroAdvisoryQuery) All() (DistroAdvisorySlice, error) {
	var o []*DistroAdvisory

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to DistroAdvisory slice")
	}

	if len(distroAdvisoryAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all DistroAdvisory records in the query, and panics on error.
func (q distroAdvisoryQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all DistroAdvisory records in the query.
func (q distroAdvisoryQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count distro_advisories rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q distroAdvisoryQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q distroAdvisoryQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if distro_advisories exists")
	}

	return count > 0, nil
}

// DistroAdvisoriesG retrieves all records.
func DistroAdvisoriesG(mods ...qm.QueryMod) distroAdvisoryQuery {
	return DistroAdvisories(boil.GetDB(), mods...)
}

// DistroAdvisories retrieves all the records using an executor.
func DistroAdvisories(exec boil.Executor, mods ...qm.QueryMod) distroAdvisoryQuery {
	mods = append(mods, qm.From("\"distro_advisories\""))
	return distroAdvisoryQuery{NewQuery(exec, mods...)}
}

// FindDistroAdvisoryG retrieves a single record by ID.
func FindDistroAdvisoryG(id int, selectCols ...string) (*DistroAdvisory, error) {
	return FindDistroAdvisory(boil.GetDB(), id, selectCols...)
}

// FindDistroAdvisoryGP retrieves a single record by ID, and panics on error.
func FindDistroAdvisoryGP(id int, selectCols ...string) *DistroAdvisory {
	retobj, err := FindDistroAdvisory(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindDistroAdvisory retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindDistroAdvisory(exec boil.Executor, id int, selectCols ...string) (*DistroAdvisory, error) {
	distroAdvisoryObj := &DistroAdvisory{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"distro_advisories\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(distroAdvisoryObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from distro_advisories")
	}

	return distroAdvisoryObj, nil
}

// FindDistroAdvisoryP retrieves a single record by ID with an executor, and panics on error.
func FindDistroAdvisoryP(exec boil.Executor, id int, selectCols ...string) *DistroAdvisory {
	retobj, err := FindDistroAdvisory(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *DistroAdvisory) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *DistroAdvisory) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *DistroAdvisory) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *DistroAdvisory) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no distro_advisories provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(distroAdvisoryColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	distroAdvisoryInsertCacheMut.RLock()
	cache, cached := distroAdvisoryInsertCache[key]
	distroAdvisoryInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			distroAdvisoryColumns,
			distroAdvisoryColumnsWithDefault,
			distroAdvisoryColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(distroAdvisoryType, distroAdvisoryMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(distroAdvisoryType, distroAdvisoryMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"distro_advisories\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"distro_advisories\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into distro_advisories")
	}

	if !cached {
		distroAdvisoryInsertCacheMut.Lock()
		distroAdvisoryInsertCache[key] = cache
		distroAdvisoryInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single DistroAdvisory record. See Update for
// whitelist behavior description.
func (o *DistroAdvisory) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single DistroAdvisory record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *DistroAdvisory) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the DistroAdvisory, and panics on error.
// See Update for whitelist behavior description.
func (o *DistroAdvisory) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the DistroAdvisory.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *DistroAdvisory) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	distroAdvisoryUpdateCacheMut.RLock()
	cache, cached := distroAdvisoryUpdateCache[key]
	distroAdvisoryUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			distroAdvisoryColumns,
			distroAdvisoryPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update distro_advisories, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"distro_advisories\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, distroAdvisoryPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(distroAdvisoryType, distroAdvisoryMapping, append(wl, distroAdvisoryPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update distro_advisories row")
	}

	if !cached {
		distroAdvisoryUpdateCacheMut.Lock()
		distroAdvisoryUpdateCache[key] = cache
		distroAdvisoryUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q distroAdvisoryQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q distroAdvisoryQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for distro_advisories")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o DistroAdvisorySlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o DistroAdvisorySlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o DistroAdvisorySlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o DistroAdvisorySlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), distroAdvisoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"distro_advisories\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, distroAdvisoryPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in distroAdvisory slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *DistroAdvisory) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *DistroAdvisory) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *DistroAdvisory) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *DistroAdvisory) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no distro_advisories provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(distroAdvisoryColumnsWithDefault, o)

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

	distroAdvisoryUpsertCacheMut.RLock()
	cache, cached := distroAdvisoryUpsertCache[key]
	distroAdvisoryUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			distroAdvisoryColumns,
			distroAdvisoryColumnsWithDefault,
			distroAdvisoryColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			distroAdvisoryColumns,
			distroAdvisoryPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert distro_advisories, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(distroAdvisoryPrimaryKeyColumns))
			copy(conflict, distroAdvisoryPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"distro_advisories\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(distroAdvisoryType, distroAdvisoryMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(distroAdvisoryType, distroAdvisoryMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert distro_advisories")
	}

	if !cached {
		distroAdvisoryUpsertCacheMut.Lock()
		distroAdvisoryUpsertCache[key] = cache
		distroAdvisoryUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single DistroAdvisory record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DistroAdvisory) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single DistroAdvisory record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *DistroAdvisory) DeleteG() error {
	if o == nil {
		return errors.New("models: no DistroAdvisory provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single DistroAdvisory record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *DistroAdvisory) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single DistroAdvisory record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *DistroAdvisory) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DistroAdvisory provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), distroAdvisoryPrimaryKeyMapping)
	sql := "DELETE FROM \"distro_advisories\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from distro_advisories")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q distroAdvisoryQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q distroAdvisoryQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no distroAdvisoryQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from distro_advisories")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o DistroAdvisorySlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o DistroAdvisorySlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no DistroAdvisory slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o DistroAdvisorySlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o DistroAdvisorySlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no DistroAdvisory slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(distroAdvisoryBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), distroAdvisoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"distro_advisories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, distroAdvisoryPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from distroAdvisory slice")
	}

	if len(distroAdvisoryAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *DistroAdvisory) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *DistroAdvisory) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *DistroAdvisory) ReloadG() error {
	if o == nil {
		return errors.New("models: no DistroAdvisory provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *DistroAdvisory) Reload(exec boil.Executor) error {
	ret, err := FindDistroAdvisory(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DistroAdvisorySlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *DistroAdvisorySlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DistroAdvisorySlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty DistroAdvisorySlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *DistroAdvisorySlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	distroAdvisories := DistroAdvisorySlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), distroAdvisoryPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"distro_advisories\".* FROM \"distro_advisories\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, distroAdvisoryPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&distroAdvisories)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in DistroAdvisorySlice")
	}

	*o = distroAdvisories

	return nil
}

// DistroAdvisoryExists checks if the DistroAdvisory row exists.
func DistroAdvisoryExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"distro_advisories\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if distro_advisories exists")
	}

	return exists, nil
}

// DistroAdvisoryExistsG checks if the DistroAdvisory row exists.
func DistroAdvisoryExistsG(id int) (bool, error) {
	return DistroAdvisoryExists(boil.GetDB(), id)
}

// DistroAdvisoryExistsGP checks if the DistroAdvisory row exists. Panics on error.
func DistroAdvisoryExistsGP(id int) bool {
	e, err := DistroAdvisoryExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// DistroAdvisoryExistsP checks if the DistroAdvisory row exists. Panics on error.
func DistroAdvisoryExistsP(exec boil.Executor, id int) bool {
	e, err := DistroAdvisoryExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
