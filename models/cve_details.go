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

// CveDetail is an object representing the database table.
type CveDetail struct {
	ID        int         `boil:"id" json:"id" toml:"id" yaml:"id"`
	CreatedAt null.Time   `boil:"created_at" json:"created_at,omitempty" toml:"created_at" yaml:"created_at,omitempty"`
	UpdatedAt null.Time   `boil:"updated_at" json:"updated_at,omitempty" toml:"updated_at" yaml:"updated_at,omitempty"`
	DeletedAt null.Time   `boil:"deleted_at" json:"deleted_at,omitempty" toml:"deleted_at" yaml:"deleted_at,omitempty"`
	CveInfoID null.Int    `boil:"cve_info_id" json:"cve_info_id,omitempty" toml:"cve_info_id" yaml:"cve_info_id,omitempty"`
	CveID     null.String `boil:"cve_id" json:"cve_id,omitempty" toml:"cve_id" yaml:"cve_id,omitempty"`

	R *cveDetailR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L cveDetailL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

// cveDetailR is where relationships are stored.
type cveDetailR struct {
}

// cveDetailL is where Load methods for each relationship are stored.
type cveDetailL struct{}

var (
	cveDetailColumns               = []string{"id", "created_at", "updated_at", "deleted_at", "cve_info_id", "cve_id"}
	cveDetailColumnsWithoutDefault = []string{"created_at", "updated_at", "deleted_at", "cve_info_id", "cve_id"}
	cveDetailColumnsWithDefault    = []string{"id"}
	cveDetailPrimaryKeyColumns     = []string{"id"}
)

type (
	// CveDetailSlice is an alias for a slice of pointers to CveDetail.
	// This should generally be used opposed to []CveDetail.
	CveDetailSlice []*CveDetail
	// CveDetailHook is the signature for custom CveDetail hook methods
	CveDetailHook func(boil.Executor, *CveDetail) error

	cveDetailQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	cveDetailType                 = reflect.TypeOf(&CveDetail{})
	cveDetailMapping              = queries.MakeStructMapping(cveDetailType)
	cveDetailPrimaryKeyMapping, _ = queries.BindMapping(cveDetailType, cveDetailMapping, cveDetailPrimaryKeyColumns)
	cveDetailInsertCacheMut       sync.RWMutex
	cveDetailInsertCache          = make(map[string]insertCache)
	cveDetailUpdateCacheMut       sync.RWMutex
	cveDetailUpdateCache          = make(map[string]updateCache)
	cveDetailUpsertCacheMut       sync.RWMutex
	cveDetailUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force bytes in case of primary key column that uses []byte (for relationship compares)
	_ = bytes.MinRead
)
var cveDetailBeforeInsertHooks []CveDetailHook
var cveDetailBeforeUpdateHooks []CveDetailHook
var cveDetailBeforeDeleteHooks []CveDetailHook
var cveDetailBeforeUpsertHooks []CveDetailHook

var cveDetailAfterInsertHooks []CveDetailHook
var cveDetailAfterSelectHooks []CveDetailHook
var cveDetailAfterUpdateHooks []CveDetailHook
var cveDetailAfterDeleteHooks []CveDetailHook
var cveDetailAfterUpsertHooks []CveDetailHook

// doBeforeInsertHooks executes all "before insert" hooks.
func (o *CveDetail) doBeforeInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailBeforeInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpdateHooks executes all "before Update" hooks.
func (o *CveDetail) doBeforeUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailBeforeUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeDeleteHooks executes all "before Delete" hooks.
func (o *CveDetail) doBeforeDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailBeforeDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doBeforeUpsertHooks executes all "before Upsert" hooks.
func (o *CveDetail) doBeforeUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailBeforeUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterInsertHooks executes all "after Insert" hooks.
func (o *CveDetail) doAfterInsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailAfterInsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterSelectHooks executes all "after Select" hooks.
func (o *CveDetail) doAfterSelectHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailAfterSelectHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpdateHooks executes all "after Update" hooks.
func (o *CveDetail) doAfterUpdateHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailAfterUpdateHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterDeleteHooks executes all "after Delete" hooks.
func (o *CveDetail) doAfterDeleteHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailAfterDeleteHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// doAfterUpsertHooks executes all "after Upsert" hooks.
func (o *CveDetail) doAfterUpsertHooks(exec boil.Executor) (err error) {
	for _, hook := range cveDetailAfterUpsertHooks {
		if err := hook(exec, o); err != nil {
			return err
		}
	}

	return nil
}

// AddCveDetailHook registers your hook function for all future operations.
func AddCveDetailHook(hookPoint boil.HookPoint, cveDetailHook CveDetailHook) {
	switch hookPoint {
	case boil.BeforeInsertHook:
		cveDetailBeforeInsertHooks = append(cveDetailBeforeInsertHooks, cveDetailHook)
	case boil.BeforeUpdateHook:
		cveDetailBeforeUpdateHooks = append(cveDetailBeforeUpdateHooks, cveDetailHook)
	case boil.BeforeDeleteHook:
		cveDetailBeforeDeleteHooks = append(cveDetailBeforeDeleteHooks, cveDetailHook)
	case boil.BeforeUpsertHook:
		cveDetailBeforeUpsertHooks = append(cveDetailBeforeUpsertHooks, cveDetailHook)
	case boil.AfterInsertHook:
		cveDetailAfterInsertHooks = append(cveDetailAfterInsertHooks, cveDetailHook)
	case boil.AfterSelectHook:
		cveDetailAfterSelectHooks = append(cveDetailAfterSelectHooks, cveDetailHook)
	case boil.AfterUpdateHook:
		cveDetailAfterUpdateHooks = append(cveDetailAfterUpdateHooks, cveDetailHook)
	case boil.AfterDeleteHook:
		cveDetailAfterDeleteHooks = append(cveDetailAfterDeleteHooks, cveDetailHook)
	case boil.AfterUpsertHook:
		cveDetailAfterUpsertHooks = append(cveDetailAfterUpsertHooks, cveDetailHook)
	}
}

// OneP returns a single cveDetail record from the query, and panics on error.
func (q cveDetailQuery) OneP() *CveDetail {
	o, err := q.One()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// One returns a single cveDetail record from the query.
func (q cveDetailQuery) One() (*CveDetail, error) {
	o := &CveDetail{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for cve_details")
	}

	if err := o.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
		return o, err
	}

	return o, nil
}

// AllP returns all CveDetail records from the query, and panics on error.
func (q cveDetailQuery) AllP() CveDetailSlice {
	o, err := q.All()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return o
}

// All returns all CveDetail records from the query.
func (q cveDetailQuery) All() (CveDetailSlice, error) {
	var o []*CveDetail

	err := q.Bind(&o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to CveDetail slice")
	}

	if len(cveDetailAfterSelectHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterSelectHooks(queries.GetExecutor(q.Query)); err != nil {
				return o, err
			}
		}
	}

	return o, nil
}

// CountP returns the count of all CveDetail records in the query, and panics on error.
func (q cveDetailQuery) CountP() int64 {
	c, err := q.Count()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return c
}

// Count returns the count of all CveDetail records in the query.
func (q cveDetailQuery) Count() (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count cve_details rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table, and panics on error.
func (q cveDetailQuery) ExistsP() bool {
	e, err := q.Exists()
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// Exists checks if the row exists in the table.
func (q cveDetailQuery) Exists() (bool, error) {
	var count int64

	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRow().Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if cve_details exists")
	}

	return count > 0, nil
}

// CveDetailsG retrieves all records.
func CveDetailsG(mods ...qm.QueryMod) cveDetailQuery {
	return CveDetails(boil.GetDB(), mods...)
}

// CveDetails retrieves all the records using an executor.
func CveDetails(exec boil.Executor, mods ...qm.QueryMod) cveDetailQuery {
	mods = append(mods, qm.From("\"cve_details\""))
	return cveDetailQuery{NewQuery(exec, mods...)}
}

// FindCveDetailG retrieves a single record by ID.
func FindCveDetailG(id int, selectCols ...string) (*CveDetail, error) {
	return FindCveDetail(boil.GetDB(), id, selectCols...)
}

// FindCveDetailGP retrieves a single record by ID, and panics on error.
func FindCveDetailGP(id int, selectCols ...string) *CveDetail {
	retobj, err := FindCveDetail(boil.GetDB(), id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// FindCveDetail retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindCveDetail(exec boil.Executor, id int, selectCols ...string) (*CveDetail, error) {
	cveDetailObj := &CveDetail{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"cve_details\" where \"id\"=$1", sel,
	)

	q := queries.Raw(exec, query, id)

	err := q.Bind(cveDetailObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from cve_details")
	}

	return cveDetailObj, nil
}

// FindCveDetailP retrieves a single record by ID with an executor, and panics on error.
func FindCveDetailP(exec boil.Executor, id int, selectCols ...string) *CveDetail {
	retobj, err := FindCveDetail(exec, id, selectCols...)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return retobj
}

// InsertG a single record. See Insert for whitelist behavior description.
func (o *CveDetail) InsertG(whitelist ...string) error {
	return o.Insert(boil.GetDB(), whitelist...)
}

// InsertGP a single record, and panics on error. See Insert for whitelist
// behavior description.
func (o *CveDetail) InsertGP(whitelist ...string) {
	if err := o.Insert(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// InsertP a single record using an executor, and panics on error. See Insert
// for whitelist behavior description.
func (o *CveDetail) InsertP(exec boil.Executor, whitelist ...string) {
	if err := o.Insert(exec, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Insert a single record using an executor.
// Whitelist behavior: If a whitelist is provided, only those columns supplied are inserted
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns without a default value are included (i.e. name, age)
// - All columns with a default, but non-zero are included (i.e. health = 75)
func (o *CveDetail) Insert(exec boil.Executor, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cve_details provided for insertion")
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

	nzDefaults := queries.NonZeroDefaultSet(cveDetailColumnsWithDefault, o)

	key := makeCacheKey(whitelist, nzDefaults)
	cveDetailInsertCacheMut.RLock()
	cache, cached := cveDetailInsertCache[key]
	cveDetailInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := strmangle.InsertColumnSet(
			cveDetailColumns,
			cveDetailColumnsWithDefault,
			cveDetailColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		cache.valueMapping, err = queries.BindMapping(cveDetailType, cveDetailMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(cveDetailType, cveDetailMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"cve_details\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.IndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"cve_details\" DEFAULT VALUES"
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
		return errors.Wrap(err, "models: unable to insert into cve_details")
	}

	if !cached {
		cveDetailInsertCacheMut.Lock()
		cveDetailInsertCache[key] = cache
		cveDetailInsertCacheMut.Unlock()
	}

	return o.doAfterInsertHooks(exec)
}

// UpdateG a single CveDetail record. See Update for
// whitelist behavior description.
func (o *CveDetail) UpdateG(whitelist ...string) error {
	return o.Update(boil.GetDB(), whitelist...)
}

// UpdateGP a single CveDetail record.
// UpdateGP takes a whitelist of column names that should be updated.
// Panics on error. See Update for whitelist behavior description.
func (o *CveDetail) UpdateGP(whitelist ...string) {
	if err := o.Update(boil.GetDB(), whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateP uses an executor to update the CveDetail, and panics on error.
// See Update for whitelist behavior description.
func (o *CveDetail) UpdateP(exec boil.Executor, whitelist ...string) {
	err := o.Update(exec, whitelist...)
	if err != nil {
		panic(boil.WrapErr(err))
	}
}

// Update uses an executor to update the CveDetail.
// Whitelist behavior: If a whitelist is provided, only the columns given are updated.
// No whitelist behavior: Without a whitelist, columns are inferred by the following rules:
// - All columns are inferred to start with
// - All primary keys are subtracted from this set
// Update does not automatically update the record in case of default values. Use .Reload()
// to refresh the records.
func (o *CveDetail) Update(exec boil.Executor, whitelist ...string) error {
	currTime := time.Now().In(boil.GetLocation())

	o.UpdatedAt.Time = currTime
	o.UpdatedAt.Valid = true

	var err error
	if err = o.doBeforeUpdateHooks(exec); err != nil {
		return err
	}
	key := makeCacheKey(whitelist, nil)
	cveDetailUpdateCacheMut.RLock()
	cache, cached := cveDetailUpdateCache[key]
	cveDetailUpdateCacheMut.RUnlock()

	if !cached {
		wl := strmangle.UpdateColumnSet(
			cveDetailColumns,
			cveDetailPrimaryKeyColumns,
			whitelist,
		)

		if len(whitelist) == 0 {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return errors.New("models: unable to update cve_details, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"cve_details\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, cveDetailPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(cveDetailType, cveDetailMapping, append(wl, cveDetailPrimaryKeyColumns...))
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
		return errors.Wrap(err, "models: unable to update cve_details row")
	}

	if !cached {
		cveDetailUpdateCacheMut.Lock()
		cveDetailUpdateCache[key] = cache
		cveDetailUpdateCacheMut.Unlock()
	}

	return o.doAfterUpdateHooks(exec)
}

// UpdateAllP updates all rows with matching column names, and panics on error.
func (q cveDetailQuery) UpdateAllP(cols M) {
	if err := q.UpdateAll(cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values.
func (q cveDetailQuery) UpdateAll(cols M) error {
	queries.SetUpdate(q.Query, cols)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to update all for cve_details")
	}

	return nil
}

// UpdateAllG updates all rows with the specified column values.
func (o CveDetailSlice) UpdateAllG(cols M) error {
	return o.UpdateAll(boil.GetDB(), cols)
}

// UpdateAllGP updates all rows with the specified column values, and panics on error.
func (o CveDetailSlice) UpdateAllGP(cols M) {
	if err := o.UpdateAll(boil.GetDB(), cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAllP updates all rows with the specified column values, and panics on error.
func (o CveDetailSlice) UpdateAllP(exec boil.Executor, cols M) {
	if err := o.UpdateAll(exec, cols); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o CveDetailSlice) UpdateAll(exec boil.Executor, cols M) error {
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cveDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"cve_details\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, cveDetailPrimaryKeyColumns, len(o)))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to update all in cveDetail slice")
	}

	return nil
}

// UpsertG attempts an insert, and does an update or ignore on conflict.
func (o *CveDetail) UpsertG(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	return o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...)
}

// UpsertGP attempts an insert, and does an update or ignore on conflict. Panics on error.
func (o *CveDetail) UpsertGP(updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(boil.GetDB(), updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// UpsertP attempts an insert using an executor, and does an update or ignore on conflict.
// UpsertP panics on error.
func (o *CveDetail) UpsertP(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) {
	if err := o.Upsert(exec, updateOnConflict, conflictColumns, updateColumns, whitelist...); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
func (o *CveDetail) Upsert(exec boil.Executor, updateOnConflict bool, conflictColumns []string, updateColumns []string, whitelist ...string) error {
	if o == nil {
		return errors.New("models: no cve_details provided for upsert")
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

	nzDefaults := queries.NonZeroDefaultSet(cveDetailColumnsWithDefault, o)

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

	cveDetailUpsertCacheMut.RLock()
	cache, cached := cveDetailUpsertCache[key]
	cveDetailUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := strmangle.InsertColumnSet(
			cveDetailColumns,
			cveDetailColumnsWithDefault,
			cveDetailColumnsWithoutDefault,
			nzDefaults,
			whitelist,
		)

		update := strmangle.UpdateColumnSet(
			cveDetailColumns,
			cveDetailPrimaryKeyColumns,
			updateColumns,
		)
		if len(update) == 0 {
			return errors.New("models: unable to upsert cve_details, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(cveDetailPrimaryKeyColumns))
			copy(conflict, cveDetailPrimaryKeyColumns)
		}
		cache.query = queries.BuildUpsertQueryPostgres(dialect, "\"cve_details\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(cveDetailType, cveDetailMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(cveDetailType, cveDetailMapping, ret)
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
		return errors.Wrap(err, "models: unable to upsert cve_details")
	}

	if !cached {
		cveDetailUpsertCacheMut.Lock()
		cveDetailUpsertCache[key] = cache
		cveDetailUpsertCacheMut.Unlock()
	}

	return o.doAfterUpsertHooks(exec)
}

// DeleteP deletes a single CveDetail record with an executor.
// DeleteP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CveDetail) DeleteP(exec boil.Executor) {
	if err := o.Delete(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteG deletes a single CveDetail record.
// DeleteG will match against the primary key column to find the record to delete.
func (o *CveDetail) DeleteG() error {
	if o == nil {
		return errors.New("models: no CveDetail provided for deletion")
	}

	return o.Delete(boil.GetDB())
}

// DeleteGP deletes a single CveDetail record.
// DeleteGP will match against the primary key column to find the record to delete.
// Panics on error.
func (o *CveDetail) DeleteGP() {
	if err := o.DeleteG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// Delete deletes a single CveDetail record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *CveDetail) Delete(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no CveDetail provided for delete")
	}

	if err := o.doBeforeDeleteHooks(exec); err != nil {
		return err
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cveDetailPrimaryKeyMapping)
	sql := "DELETE FROM \"cve_details\" WHERE \"id\"=$1"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args...)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete from cve_details")
	}

	if err := o.doAfterDeleteHooks(exec); err != nil {
		return err
	}

	return nil
}

// DeleteAllP deletes all rows, and panics on error.
func (q cveDetailQuery) DeleteAllP() {
	if err := q.DeleteAll(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all matching rows.
func (q cveDetailQuery) DeleteAll() error {
	if q.Query == nil {
		return errors.New("models: no cveDetailQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	_, err := q.Query.Exec()
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cve_details")
	}

	return nil
}

// DeleteAllGP deletes all rows in the slice, and panics on error.
func (o CveDetailSlice) DeleteAllGP() {
	if err := o.DeleteAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAllG deletes all rows in the slice.
func (o CveDetailSlice) DeleteAllG() error {
	if o == nil {
		return errors.New("models: no CveDetail slice provided for delete all")
	}
	return o.DeleteAll(boil.GetDB())
}

// DeleteAllP deletes all rows in the slice, using an executor, and panics on error.
func (o CveDetailSlice) DeleteAllP(exec boil.Executor) {
	if err := o.DeleteAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o CveDetailSlice) DeleteAll(exec boil.Executor) error {
	if o == nil {
		return errors.New("models: no CveDetail slice provided for delete all")
	}

	if len(o) == 0 {
		return nil
	}

	if len(cveDetailBeforeDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doBeforeDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cveDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"cve_details\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cveDetailPrimaryKeyColumns, len(o))

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, args)
	}

	_, err := exec.Exec(sql, args...)
	if err != nil {
		return errors.Wrap(err, "models: unable to delete all from cveDetail slice")
	}

	if len(cveDetailAfterDeleteHooks) != 0 {
		for _, obj := range o {
			if err := obj.doAfterDeleteHooks(exec); err != nil {
				return err
			}
		}
	}

	return nil
}

// ReloadGP refetches the object from the database and panics on error.
func (o *CveDetail) ReloadGP() {
	if err := o.ReloadG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadP refetches the object from the database with an executor. Panics on error.
func (o *CveDetail) ReloadP(exec boil.Executor) {
	if err := o.Reload(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadG refetches the object from the database using the primary keys.
func (o *CveDetail) ReloadG() error {
	if o == nil {
		return errors.New("models: no CveDetail provided for reload")
	}

	return o.Reload(boil.GetDB())
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *CveDetail) Reload(exec boil.Executor) error {
	ret, err := FindCveDetail(exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAllGP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CveDetailSlice) ReloadAllGP() {
	if err := o.ReloadAllG(); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllP refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
// Panics on error.
func (o *CveDetailSlice) ReloadAllP(exec boil.Executor) {
	if err := o.ReloadAll(exec); err != nil {
		panic(boil.WrapErr(err))
	}
}

// ReloadAllG refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CveDetailSlice) ReloadAllG() error {
	if o == nil {
		return errors.New("models: empty CveDetailSlice provided for reload all")
	}

	return o.ReloadAll(boil.GetDB())
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *CveDetailSlice) ReloadAll(exec boil.Executor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	cveDetails := CveDetailSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), cveDetailPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"cve_details\".* FROM \"cve_details\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, cveDetailPrimaryKeyColumns, len(*o))

	q := queries.Raw(exec, sql, args...)

	err := q.Bind(&cveDetails)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in CveDetailSlice")
	}

	*o = cveDetails

	return nil
}

// CveDetailExists checks if the CveDetail row exists.
func CveDetailExists(exec boil.Executor, id int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"cve_details\" where \"id\"=$1 limit 1)"

	if boil.DebugMode {
		fmt.Fprintln(boil.DebugWriter, sql)
		fmt.Fprintln(boil.DebugWriter, id)
	}

	row := exec.QueryRow(sql, id)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if cve_details exists")
	}

	return exists, nil
}

// CveDetailExistsG checks if the CveDetail row exists.
func CveDetailExistsG(id int) (bool, error) {
	return CveDetailExists(boil.GetDB(), id)
}

// CveDetailExistsGP checks if the CveDetail row exists. Panics on error.
func CveDetailExistsGP(id int) bool {
	e, err := CveDetailExists(boil.GetDB(), id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}

// CveDetailExistsP checks if the CveDetail row exists. Panics on error.
func CveDetailExistsP(exec boil.Executor, id int) bool {
	e, err := CveDetailExists(exec, id)
	if err != nil {
		panic(boil.WrapErr(err))
	}

	return e
}
