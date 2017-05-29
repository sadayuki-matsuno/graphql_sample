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

// Changelog is an object representing the database table.
type Changelog struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	Contents  null.String `boil:"contents" json:"contents,omitempty" toml:"contents" yaml:"contents,omitempty"`
	Method    null.String `boil:"method" json:"method,omitempty" toml:"method" yaml:"method,omitempty"`

	R *changelogR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L changelogL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// changelogR is where relationships are stored.
type changelogR struct {
}

// changelogL is where Load methods for each relationship are stored.
type changelogL struct{}

var (
	changelogColumns               = []string{"id", "created_at", "updated_at", "deleted_at", "contents", "method"}
	changelogColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "contents", "method"}
	changelogColumnsWithDefault    = []string{"id"}
	changelogPrimaryKeyColumns     = []string{"id"}
)

type (
	// ChangelogSlice is an alias for a slice of pointers to Changelog.
	// This should generally be used opposed to []Changelog.
	ChangelogSlice []*Changelog
	// ChangelogHook is the signature for custom Changelog hook methods
	ChangelogHook func(boil.Executor, *Changelog) error

	changelogQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	changelogType                 = reflect.TypeOf(&Changelog{})
	changelogMapping              = queries.MakeStructMapping(changelogType)
	changelogPrimaryKeyMapping, _ = queries.BindMapping(changelogType, changelogMapping, changelogPrimaryKeyColumns)
	changelogInsertCacheMut       sync.RWMutex
	changelogInsertCache          = make(map[string]insertCache)
	changelogUpdateCacheMut       sync.RWMutex
	changelogUpdateCache          = make(map[string]updateCache)
	changelogUpsertCacheMut       sync.RWMutex
	changelogUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var changelogBeforeInsertHooks []ChangelogHook
var changelogBeforeUpdateHooks []ChangelogHook
var changelogBeforeDeleteHooks []ChangelogHook
var changelogBeforeUpsertHooks []ChangelogHook

var changelogAfterInsertHooks []ChangelogHook
var changelogAfterSelectHooks []ChangelogHook
var changelogAfterUpdateHooks []ChangelogHook
var changelogAfterDeleteHooks []ChangelogHook
var changelogAfterUpsertHooks []ChangelogHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *Changelog) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *Changelog) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *Changelog) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *Changelog) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *Changelog) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *Changelog) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *Changelog) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *Changelog) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *Changelog) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range changelogAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddChangelogHook registers your hook function for all future operations.
func AddChangelogHook(hookPoint boil.HookPoint, changelogHook ChangelogHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		changelogBeforeInsertHooks = append(changelogBeforeInsertHooks, changelogHook)
	case boil.BeforeUpdateHook:
		changelogBeforeUpdateHooks = append(changelogBeforeUpdateHooks, changelogHook)
	case boil.BeforeDeleteHook:
		changelogBeforeDeleteHooks = append(changelogBeforeDeleteHooks, changelogHook)
	case boil.BeforeUpsertHook:
		changelogBeforeUpsertHooks = append(changelogBeforeUpsertHooks, changelogHook)
	case boil.AfterInsertHook:
		changelogAfterInsertHooks = append(changelogAfterInsertHooks, changelogHook)
	case boil.AfterSelectHook:
		changelogAfterSelectHooks = append(changelogAfterSelectHooks, changelogHook)
	case boil.AfterUpdateHook:
		changelogAfterUpdateHooks = append(changelogAfterUpdateHooks, changelogHook)
	case boil.AfterDeleteHook:
		changelogAfterDeleteHooks = append(changelogAfterDeleteHooks, changelogHook)
	case boil.AfterUpsertHook:
		changelogAfterUpsertHooks = append(changelogAfterUpsertHooks, changelogHook)
	}
}

// OneP returns a single changelog record from the query, and panics on error.
func (q changelogQuery) OneP() *Changelog {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single changelog record from the query.
func (q changelogQuery) One() (*Changelog, error) {
	o := &Changelog{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for changelogs")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all Changelog records from the query, and panics on error.
func (q changelogQuery) AllP() ChangelogSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all Changelog records from the query.
func (q changelogQuery) All() (ChangelogSlice, error) {
	var o []*Changelog

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to Changelog slice")
	}

	if len(changelogAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all Changelog records in the query, and panics on error.
func (q changelogQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all Changelog records in the query.
func (q changelogQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count changelogs rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q changelogQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q changelogQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if changelogs exists")
	}

	return count > 0, nil
}

// ChangelogsG retrieves all records.
func ChangelogsG(mods ...qm.QueryMod) changelogQuery {
	return Changelogs(boil.GetDB(), mods...)
}

// Changelogs retrieves all the records using an executor.
func Changelogs(exec boil.Executor, mods ...qm.QueryMod) changelogQuery {
	mods = append(mods, qm.From("\"changelogs\""))
	return changelogQuery{NewQuery(exec, mods...)}
}

// FindChangelogG retrieves a single record by ID.
func FindChangelogG(id int, selectCols ...string) (*Changelog, error) {
	return FindChangelog(boil.GetDB(), id, selectCols...)
}

// FindChangelogGP retrieves a single record by ID, and panics on error.
func FindChangelogGP(id int, selectCols ...string) *Changelog {
	retobj, err := FindChangelog(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindChangelog retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindChangelog(exec boil.Executor, id int, selectCols ...string) (*Changelog, error) {
	changelogObj := &Changelog{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"changelogs\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(changelogObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from changelogs")
	}

	return changelogObj, nil
}

// FindChangelogP retrieves a single record by ID with an executor, and panics on error.
func FindChangelogP(exec boil.Executor, id int, selectCols ...string) *Changelog {
	retobj, err := FindChangelog(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *Changelog) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *Changelog) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *Changelog) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *Changelog) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no changelogs provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(changelogColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	changelogInsertCacheMut.RLock()
	cache, cached := changelogInsertCache[key]
	changelogInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			changelogColumns,
			changelogColumnsWithDefault,
			changelogColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(changelogType, changelogMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(changelogType, changelogMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"changelogs\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"changelogs\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into changelogs")
	}

	if !cached {
		changelogInsertCacheMut.Lock()
		changelogInsertCache[key] = cache
		changelogInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single Changelog record. See Update for
// whitelist behavior description.
func (o *Changelog) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single Changelog record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *Changelog) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the Changelog, and panics on error.
// See Update for whitelist behavior description.
func (o *Changelog) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the Changelog.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *Changelog) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	changelogUpdateCacheMut.RLock()
	cache, cached := changelogUpdateCache[key]
	changelogUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			changelogColumns,
			changelogPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update changelogs, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"changelogs\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, changelogPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(changelogType, changelogMapping, append(wl, changelogPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update changelogs row")
	}

	if !cached {
		changelogUpdateCacheMut.Lock()
		changelogUpdateCache[key] = cache
		changelogUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q changelogQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q changelogQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for changelogs")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o ChangelogSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o ChangelogSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o ChangelogSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o ChangelogSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), changelogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"changelogs\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, changelogPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in changelog slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *Changelog) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *Changelog) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *Changelog) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *Changelog) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no changelogs provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(changelogColumnsWithDefault, o)

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

	changelogUpsertCacheMut.RLock()
	cache, cached := changelogUpsertCache[key]
	changelogUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			changelogColumns,
			changelogColumnsWithDefault,
			changelogColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			changelogColumns,
			changelogPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert changelogs, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(changelogPrimaryKeyColumns))
			copy(conflict, changelogPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"changelogs\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(changelogType, changelogMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(changelogType, changelogMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert changelogs")
	}

	if !cached {
		changelogUpsertCacheMut.Lock()
		changelogUpsertCache[key] = cache
		changelogUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single Changelog record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Changelog) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single Changelog record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *Changelog) DeleteG() error {
	if o == nil {
		return errors.New("models: no Changelog provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single Changelog record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *Changelog) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single Changelog record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *Changelog) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Changelog provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), changelogPrimaryKeyMapping)
	sql := "DELETE FROM \"changelogs\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from changelogs")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q changelogQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q changelogQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no changelogQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from changelogs")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o ChangelogSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o ChangelogSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no Changelog slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o ChangelogSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o ChangelogSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no Changelog slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(changelogBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), changelogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"changelogs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, changelogPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from changelog slice")
	}

	if len(changelogAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *Changelog) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *Changelog) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *Changelog) ReloadG() error {
	if o == nil {
		return errors.New("models: no Changelog provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *Changelog) Reload(exec boil.Executor) error {
	ret, err := FindChangelog(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChangelogSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *ChangelogSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChangelogSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty ChangelogSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *ChangelogSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	changelogs := ChangelogSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), changelogPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"changelogs\".* FROM \"changelogs\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, changelogPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&changelogs)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in ChangelogSlice")
	}

	*o = changelogs

	return nil
}

// ChangelogExists checks if the Changelog row exists.
func ChangelogExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"changelogs\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if changelogs exists")
	}

	return exists, nil
}

// ChangelogExistsG checks if the Changelog row exists.
func ChangelogExistsG(id int) (bool, error) {
	return ChangelogExists(boil.GetDB(), id)
}

// ChangelogExistsGP checks if the Changelog row exists. Panics on error.
func ChangelogExistsGP(id int) bool {
	e, err := ChangelogExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// ChangelogExistsP checks if the Changelog row exists. Panics on error.
func ChangelogExistsP(exec boil.Executor, id int) bool {
	e, err := ChangelogExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
