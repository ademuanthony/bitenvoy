// Code generated by SQLBoiler 3.6.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import (
	"bytes"
	"context"
	"reflect"
	"testing"

	"github.com/volatiletech/sqlboiler/boil"
	"github.com/volatiletech/sqlboiler/queries"
	"github.com/volatiletech/sqlboiler/randomize"
	"github.com/volatiletech/sqlboiler/strmangle"
)

var (
	// Relationships sometimes use the reflection helper queries.Equal/queries.Assign
	// so force a package dependency in case they don't.
	_ = queries.Equal
)

func testOrderModels(t *testing.T) {
	t.Parallel()

	query := OrderModels()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testOrderModelsDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := o.Delete(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrderModelsQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := OrderModels().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrderModelsSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderModelSlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testOrderModelsExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := OrderModelExists(ctx, tx, o.ID)
	if err != nil {
		t.Errorf("Unable to check if OrderModel exists: %s", err)
	}
	if !e {
		t.Errorf("Expected OrderModelExists to return true, but got false.")
	}
}

func testOrderModelsFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	orderModelFound, err := FindOrderModel(ctx, tx, o.ID)
	if err != nil {
		t.Error(err)
	}

	if orderModelFound == nil {
		t.Error("want a record, got nil")
	}
}

func testOrderModelsBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = OrderModels().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testOrderModelsOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := OrderModels().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testOrderModelsAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	orderModelOne := &OrderModel{}
	orderModelTwo := &OrderModel{}
	if err = randomize.Struct(seed, orderModelOne, orderModelDBTypes, false, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}
	if err = randomize.Struct(seed, orderModelTwo, orderModelDBTypes, false, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderModelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderModelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrderModels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testOrderModelsCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	orderModelOne := &OrderModel{}
	orderModelTwo := &OrderModel{}
	if err = randomize.Struct(seed, orderModelOne, orderModelDBTypes, false, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}
	if err = randomize.Struct(seed, orderModelTwo, orderModelDBTypes, false, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = orderModelOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = orderModelTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testOrderModelsInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrderModelsInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(orderModelColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testOrderModelsReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = o.Reload(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrderModelsReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := OrderModelSlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testOrderModelsSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := OrderModels().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	orderModelDBTypes = map[string]string{`ID`: `character varying`, `Username`: `character varying`, `Product`: `character varying`, `Status`: `character varying`, `Details`: `text`, `Date`: `integer`, `Amount`: `double precision`}
	_                 = bytes.MinRead
)

func testOrderModelsUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(orderModelPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(orderModelAllColumns) == len(orderModelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testOrderModelsSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(orderModelAllColumns) == len(orderModelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &OrderModel{}
	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, orderModelDBTypes, true, orderModelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(orderModelAllColumns, orderModelPrimaryKeyColumns) {
		fields = orderModelAllColumns
	} else {
		fields = strmangle.SetComplement(
			orderModelAllColumns,
			orderModelPrimaryKeyColumns,
		)
	}

	value := reflect.Indirect(reflect.ValueOf(o))
	typ := reflect.TypeOf(o).Elem()
	n := typ.NumField()

	updateMap := M{}
	for _, col := range fields {
		for i := 0; i < n; i++ {
			f := typ.Field(i)
			if f.Tag.Get("boil") == col {
				updateMap[col] = value.Field(i).Interface()
			}
		}
	}

	slice := OrderModelSlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testOrderModelsUpsert(t *testing.T) {
	t.Parallel()

	if len(orderModelAllColumns) == len(orderModelPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := OrderModel{}
	if err = randomize.Struct(seed, &o, orderModelDBTypes, true); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrderModel: %s", err)
	}

	count, err := OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, orderModelDBTypes, false, orderModelPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize OrderModel struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert OrderModel: %s", err)
	}

	count, err = OrderModels().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
