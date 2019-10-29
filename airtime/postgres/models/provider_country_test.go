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

func testProviderCountries(t *testing.T) {
	t.Parallel()

	query := ProviderCountries()

	if query.Query == nil {
		t.Error("expected a query, got nothing")
	}
}

func testProviderCountriesDelete(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
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

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProviderCountriesQueryDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if rowsAff, err := ProviderCountries().DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProviderCountriesSliceDeleteAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProviderCountrySlice{o}

	if rowsAff, err := slice.DeleteAll(ctx, tx); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only have deleted one row, but affected:", rowsAff)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 0 {
		t.Error("want zero records, got:", count)
	}
}

func testProviderCountriesExists(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	e, err := ProviderCountryExists(ctx, tx, o.ProviderID, o.CountryID)
	if err != nil {
		t.Errorf("Unable to check if ProviderCountry exists: %s", err)
	}
	if !e {
		t.Errorf("Expected ProviderCountryExists to return true, but got false.")
	}
}

func testProviderCountriesFind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	providerCountryFound, err := FindProviderCountry(ctx, tx, o.ProviderID, o.CountryID)
	if err != nil {
		t.Error(err)
	}

	if providerCountryFound == nil {
		t.Error("want a record, got nil")
	}
}

func testProviderCountriesBind(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if err = ProviderCountries().Bind(ctx, tx, o); err != nil {
		t.Error(err)
	}
}

func testProviderCountriesOne(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	if x, err := ProviderCountries().One(ctx, tx); err != nil {
		t.Error(err)
	} else if x == nil {
		t.Error("expected to get a non nil record")
	}
}

func testProviderCountriesAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	providerCountryOne := &ProviderCountry{}
	providerCountryTwo := &ProviderCountry{}
	if err = randomize.Struct(seed, providerCountryOne, providerCountryDBTypes, false, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}
	if err = randomize.Struct(seed, providerCountryTwo, providerCountryDBTypes, false, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = providerCountryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = providerCountryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProviderCountries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 2 {
		t.Error("want 2 records, got:", len(slice))
	}
}

func testProviderCountriesCount(t *testing.T) {
	t.Parallel()

	var err error
	seed := randomize.NewSeed()
	providerCountryOne := &ProviderCountry{}
	providerCountryTwo := &ProviderCountry{}
	if err = randomize.Struct(seed, providerCountryOne, providerCountryDBTypes, false, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}
	if err = randomize.Struct(seed, providerCountryTwo, providerCountryDBTypes, false, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = providerCountryOne.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}
	if err = providerCountryTwo.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 2 {
		t.Error("want 2 records, got:", count)
	}
}

func testProviderCountriesInsert(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProviderCountriesInsertWhitelist(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Whitelist(providerCountryColumnsWithoutDefault...)); err != nil {
		t.Error(err)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}
}

func testProviderCountriesReload(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
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

func testProviderCountriesReloadAll(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice := ProviderCountrySlice{o}

	if err = slice.ReloadAll(ctx, tx); err != nil {
		t.Error(err)
	}
}

func testProviderCountriesSelect(t *testing.T) {
	t.Parallel()

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	slice, err := ProviderCountries().All(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if len(slice) != 1 {
		t.Error("want one record, got:", len(slice))
	}
}

var (
	providerCountryDBTypes = map[string]string{`ProviderID`: `character varying`, `CountryID`: `character varying`}
	_                      = bytes.MinRead
)

func testProviderCountriesUpdate(t *testing.T) {
	t.Parallel()

	if 0 == len(providerCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with no primary key columns")
	}
	if len(providerCountryAllColumns) == len(providerCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	if rowsAff, err := o.Update(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("should only affect one row but affected", rowsAff)
	}
}

func testProviderCountriesSliceUpdateAll(t *testing.T) {
	t.Parallel()

	if len(providerCountryAllColumns) == len(providerCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	o := &ProviderCountry{}
	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryColumnsWithDefault...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Insert(ctx, tx, boil.Infer()); err != nil {
		t.Error(err)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}

	if count != 1 {
		t.Error("want one record, got:", count)
	}

	if err = randomize.Struct(seed, o, providerCountryDBTypes, true, providerCountryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	// Remove Primary keys and unique columns from what we plan to update
	var fields []string
	if strmangle.StringSliceMatch(providerCountryAllColumns, providerCountryPrimaryKeyColumns) {
		fields = providerCountryAllColumns
	} else {
		fields = strmangle.SetComplement(
			providerCountryAllColumns,
			providerCountryPrimaryKeyColumns,
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

	slice := ProviderCountrySlice{o}
	if rowsAff, err := slice.UpdateAll(ctx, tx, updateMap); err != nil {
		t.Error(err)
	} else if rowsAff != 1 {
		t.Error("wanted one record updated but got", rowsAff)
	}
}

func testProviderCountriesUpsert(t *testing.T) {
	t.Parallel()

	if len(providerCountryAllColumns) == len(providerCountryPrimaryKeyColumns) {
		t.Skip("Skipping table with only primary key columns")
	}

	seed := randomize.NewSeed()
	var err error
	// Attempt the INSERT side of an UPSERT
	o := ProviderCountry{}
	if err = randomize.Struct(seed, &o, providerCountryDBTypes, true); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	ctx := context.Background()
	tx := MustTx(boil.BeginTx(ctx, nil))
	defer func() { _ = tx.Rollback() }()
	if err = o.Upsert(ctx, tx, false, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProviderCountry: %s", err)
	}

	count, err := ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}

	// Attempt the UPDATE side of an UPSERT
	if err = randomize.Struct(seed, &o, providerCountryDBTypes, false, providerCountryPrimaryKeyColumns...); err != nil {
		t.Errorf("Unable to randomize ProviderCountry struct: %s", err)
	}

	if err = o.Upsert(ctx, tx, true, nil, boil.Infer(), boil.Infer()); err != nil {
		t.Errorf("Unable to upsert ProviderCountry: %s", err)
	}

	count, err = ProviderCountries().Count(ctx, tx)
	if err != nil {
		t.Error(err)
	}
	if count != 1 {
		t.Error("want one record, got:", count)
	}
}
