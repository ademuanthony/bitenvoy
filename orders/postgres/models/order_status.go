// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/friendsofgo/errors"
	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/queries/qm"
	"github.com/volatiletech/sqlboiler/queries/qmhelper"
	"github.com/volatiletech/sqlboiler/strmangle"
)

// OrderStatus is an object representing the database table.
type OrderStatus struct {
	ID      string `boil:"id" json:"id" toml:"id" yaml:"id"`
	OrderID string `boil:"order_id" json:"order_id" toml:"order_id" yaml:"order_id"`
	Status  string `boil:"status" json:"status" toml:"status" yaml:"status"`
	Date    int    `boil:"date" json:"date" toml:"date" yaml:"date"`

	R *orderStatusR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L orderStatusL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var OrderStatusColumns = struct {
	ID      string
	OrderID string
	Status  string
	Date    string
}{
	ID:      "id",
	OrderID: "order_id",
	Status:  "status",
	Date:    "date",
}

// Generated where

var OrderStatusWhere = struct {
	ID      whereHelperstring
	OrderID whereHelperstring
	Status  whereHelperstring
	Date    whereHelperint
}{
	ID:      whereHelperstring{field: "\"order_status\".\"id\""},
	OrderID: whereHelperstring{field: "\"order_status\".\"order_id\""},
	Status:  whereHelperstring{field: "\"order_status\".\"status\""},
	Date:    whereHelperint{field: "\"order_status\".\"date\""},
}

// OrderStatusRels is where relationship names are stored.
var OrderStatusRels = struct {
}{}

// orderStatusR is where relationships are stored.
type orderStatusR struct {
}

// NewStruct creates a new relationship struct
func (*orderStatusR) NewStruct() *orderStatusR {
	return &orderStatusR{}
}

// orderStatusL is where Load methods for each relationship are stored.
type orderStatusL struct{}

var (
	orderStatusAllColumns            = []string{"id", "order_id", "status", "date"}
	orderStatusColumnsWithoutDefault = []string{"id", "order_id", "status", "date"}
	orderStatusColumnsWithDefault    = []string{}
	orderStatusPrimaryKeyColumns     = []string{"id"}
)

type (
	// OrderStatusSlice is an alias for a slice of pointers to OrderStatus.
	// This should generally be used opposed to []OrderStatus.
	OrderStatusSlice []*OrderStatus

	orderStatusQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	orderStatusType                 = reflect.TypeOf(&OrderStatus{})
	orderStatusMapping              = queries.MakeStructMapping(orderStatusType)
	orderStatusPrimaryKeyMapping, _ = queries.BindMapping(orderStatusType, orderStatusMapping, orderStatusPrimaryKeyColumns)
	orderStatusInsertCacheMut       sync.RWMutex
	orderStatusInsertCache          = make(map[string]insertCache)
	orderStatusUpdateCacheMut       sync.RWMutex
	orderStatusUpdateCache          = make(map[string]updateCache)
	orderStatusUpsertCacheMut       sync.RWMutex
	orderStatusUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single orderStatus record from the query.
func (q orderStatusQuery) One(ctx context.Context, exec boil.ContextExecutor) (*OrderStatus, error) {
	o := &OrderStatus{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: failed to execute a one query for order_status")
	}

	return o, nil
}

// All returns all OrderStatus records from the query.
func (q orderStatusQuery) All(ctx context.Context, exec boil.ContextExecutor) (OrderStatusSlice, error) {
	var o []*OrderStatus

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "models: failed to assign all query results to OrderStatus slice")
	}

	return o, nil
}

// Count returns the count of all OrderStatus records in the query.
func (q orderStatusQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to count order_status rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q orderStatusQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "models: failed to check if order_status exists")
	}

	return count > 0, nil
}

// OrderStatuses retrieves all the records using an executor.
func OrderStatuses(mods ...qm.QueryMod) orderStatusQuery {
	mods = append(mods, qm.From("\"order_status\""))
	return orderStatusQuery{NewQuery(mods...)}
}

// FindOrderStatus retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindOrderStatus(ctx context.Context, exec boil.ContextExecutor, iD string, selectCols ...string) (*OrderStatus, error) {
	orderStatusObj := &OrderStatus{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"order_status\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, orderStatusObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "models: unable to select from order_status")
	}

	return orderStatusObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *OrderStatus) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("models: no order_status provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(orderStatusColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	orderStatusInsertCacheMut.RLock()
	cache, cached := orderStatusInsertCache[key]
	orderStatusInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			orderStatusAllColumns,
			orderStatusColumnsWithDefault,
			orderStatusColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(orderStatusType, orderStatusMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(orderStatusType, orderStatusMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"order_status\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"order_status\" %sDEFAULT VALUES%s"
		}

		var queryOutput, queryReturning string

		if len(cache.retMapping) != 0 {
			queryReturning = fmt.Sprintf(" RETURNING \"%s\"", strings.Join(returnColumns, "\",\""))
		}

		cache.query = fmt.Sprintf(cache.query, queryOutput, queryReturning)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	vals := queries.ValuesFromMapping(value, cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}

	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(queries.PtrsFromMapping(value, cache.retMapping)...)
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}

	if err != nil {
		return errors.Wrap(err, "models: unable to insert into order_status")
	}

	if !cached {
		orderStatusInsertCacheMut.Lock()
		orderStatusInsertCache[key] = cache
		orderStatusInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the OrderStatus.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *OrderStatus) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	orderStatusUpdateCacheMut.RLock()
	cache, cached := orderStatusUpdateCache[key]
	orderStatusUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			orderStatusAllColumns,
			orderStatusPrimaryKeyColumns,
		)

		if len(wl) == 0 {
			return 0, errors.New("models: unable to update order_status, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"order_status\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, orderStatusPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(orderStatusType, orderStatusMapping, append(wl, orderStatusPrimaryKeyColumns...))
		if err != nil {
			return 0, err
		}
	}

	values := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), cache.valueMapping)

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, values)
	}
	var result sql.Result
	result, err = exec.ExecContext(ctx, cache.query, values...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update order_status row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by update for order_status")
	}

	if !cached {
		orderStatusUpdateCacheMut.Lock()
		orderStatusUpdateCache[key] = cache
		orderStatusUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q orderStatusQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all for order_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected for order_status")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o OrderStatusSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("models: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"order_status\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, orderStatusPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to update all in orderStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to retrieve rows affected all in update all orderStatus")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *OrderStatus) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("models: no order_status provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(orderStatusColumnsWithDefault, o)

	// Build cache key in-line uglily - mysql vs psql problems
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
	buf.WriteString(strconv.Itoa(updateColumns.Kind))
	for _, c := range updateColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	buf.WriteString(strconv.Itoa(insertColumns.Kind))
	for _, c := range insertColumns.Cols {
		buf.WriteString(c)
	}
	buf.WriteByte('.')
	for _, c := range nzDefaults {
		buf.WriteString(c)
	}
	key := buf.String()
	strmangle.PutBuffer(buf)

	orderStatusUpsertCacheMut.RLock()
	cache, cached := orderStatusUpsertCache[key]
	orderStatusUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			orderStatusAllColumns,
			orderStatusColumnsWithDefault,
			orderStatusColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			orderStatusAllColumns,
			orderStatusPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("models: unable to upsert order_status, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(orderStatusPrimaryKeyColumns))
			copy(conflict, orderStatusPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"order_status\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(orderStatusType, orderStatusMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(orderStatusType, orderStatusMapping, ret)
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

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, cache.query)
		fmt.Fprintln(writer, vals)
	}
	if len(cache.retMapping) != 0 {
		err = exec.QueryRowContext(ctx, cache.query, vals...).Scan(returns...)
		if err == sql.ErrNoRows {
			err = nil // Postgres doesn't return anything when there's no update
		}
	} else {
		_, err = exec.ExecContext(ctx, cache.query, vals...)
	}
	if err != nil {
		return errors.Wrap(err, "models: unable to upsert order_status")
	}

	if !cached {
		orderStatusUpsertCacheMut.Lock()
		orderStatusUpsertCache[key] = cache
		orderStatusUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single OrderStatus record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *OrderStatus) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("models: no OrderStatus provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), orderStatusPrimaryKeyMapping)
	sql := "DELETE FROM \"order_status\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete from order_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by delete for order_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q orderStatusQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("models: no orderStatusQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from order_status")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for order_status")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o OrderStatusSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"order_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderStatusPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "models: unable to delete all from orderStatus slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "models: failed to get rows affected by deleteall for order_status")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *OrderStatus) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindOrderStatus(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *OrderStatusSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := OrderStatusSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), orderStatusPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"order_status\".* FROM \"order_status\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, orderStatusPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "models: unable to reload all in OrderStatusSlice")
	}

	*o = slice

	return nil
}

// OrderStatusExists checks if the OrderStatus row exists.
func OrderStatusExists(ctx context.Context, exec boil.ContextExecutor, iD string) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"order_status\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "models: unable to check if order_status exists")
	}

	return exists, nil
}
