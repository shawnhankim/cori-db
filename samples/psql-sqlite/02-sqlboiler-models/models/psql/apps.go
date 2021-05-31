// Code generated by SQLBoiler 4.5.0 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package portal

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
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
	"github.com/volatiletech/sqlboiler/v4/queries/qmhelper"
	"github.com/volatiletech/strmangle"
)

// App is an object representing the database table.
type App struct {
	ID     int      `boil:"id" json:"id" toml:"id" yaml:"id"`
	UserID int      `boil:"user_id" json:"user_id" toml:"user_id" yaml:"user_id"`
	OrgID  null.Int `boil:"org_id" json:"org_id,omitempty" toml:"org_id" yaml:"org_id,omitempty"`

	R *appR `boil:"-" json:"-" toml:"-" yaml:"-"`
	L appL  `boil:"-" json:"-" toml:"-" yaml:"-"`
}

var AppColumns = struct {
	ID     string
	UserID string
	OrgID  string
}{
	ID:     "id",
	UserID: "user_id",
	OrgID:  "org_id",
}

// Generated where

type whereHelperint struct{ field string }

func (w whereHelperint) EQ(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.EQ, x) }
func (w whereHelperint) NEQ(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.NEQ, x) }
func (w whereHelperint) LT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.LT, x) }
func (w whereHelperint) LTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.LTE, x) }
func (w whereHelperint) GT(x int) qm.QueryMod  { return qmhelper.Where(w.field, qmhelper.GT, x) }
func (w whereHelperint) GTE(x int) qm.QueryMod { return qmhelper.Where(w.field, qmhelper.GTE, x) }
func (w whereHelperint) IN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereIn(fmt.Sprintf("%s IN ?", w.field), values...)
}
func (w whereHelperint) NIN(slice []int) qm.QueryMod {
	values := make([]interface{}, 0, len(slice))
	for _, value := range slice {
		values = append(values, value)
	}
	return qm.WhereNotIn(fmt.Sprintf("%s NOT IN ?", w.field), values...)
}

type whereHelpernull_Int struct{ field string }

func (w whereHelpernull_Int) EQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, false, x)
}
func (w whereHelpernull_Int) NEQ(x null.Int) qm.QueryMod {
	return qmhelper.WhereNullEQ(w.field, true, x)
}
func (w whereHelpernull_Int) IsNull() qm.QueryMod    { return qmhelper.WhereIsNull(w.field) }
func (w whereHelpernull_Int) IsNotNull() qm.QueryMod { return qmhelper.WhereIsNotNull(w.field) }
func (w whereHelpernull_Int) LT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LT, x)
}
func (w whereHelpernull_Int) LTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.LTE, x)
}
func (w whereHelpernull_Int) GT(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GT, x)
}
func (w whereHelpernull_Int) GTE(x null.Int) qm.QueryMod {
	return qmhelper.Where(w.field, qmhelper.GTE, x)
}

var AppWhere = struct {
	ID     whereHelperint
	UserID whereHelperint
	OrgID  whereHelpernull_Int
}{
	ID:     whereHelperint{field: "\"apps\".\"id\""},
	UserID: whereHelperint{field: "\"apps\".\"user_id\""},
	OrgID:  whereHelpernull_Int{field: "\"apps\".\"org_id\""},
}

// AppRels is where relationship names are stored.
var AppRels = struct {
	Org  string
	User string
}{
	Org:  "Org",
	User: "User",
}

// appR is where relationships are stored.
type appR struct {
	Org  *Org  `boil:"Org" json:"Org" toml:"Org" yaml:"Org"`
	User *User `boil:"User" json:"User" toml:"User" yaml:"User"`
}

// NewStruct creates a new relationship struct
func (*appR) NewStruct() *appR {
	return &appR{}
}

// appL is where Load methods for each relationship are stored.
type appL struct{}

var (
	appAllColumns            = []string{"id", "user_id", "org_id"}
	appColumnsWithoutDefault = []string{"id", "user_id", "org_id"}
	appColumnsWithDefault    = []string{}
	appPrimaryKeyColumns     = []string{"id"}
)

type (
	// AppSlice is an alias for a slice of pointers to App.
	// This should generally be used opposed to []App.
	AppSlice []*App

	appQuery struct {
		*queries.Query
	}
)

// Cache for insert, update and upsert
var (
	appType                 = reflect.TypeOf(&App{})
	appMapping              = queries.MakeStructMapping(appType)
	appPrimaryKeyMapping, _ = queries.BindMapping(appType, appMapping, appPrimaryKeyColumns)
	appInsertCacheMut       sync.RWMutex
	appInsertCache          = make(map[string]insertCache)
	appUpdateCacheMut       sync.RWMutex
	appUpdateCache          = make(map[string]updateCache)
	appUpsertCacheMut       sync.RWMutex
	appUpsertCache          = make(map[string]insertCache)
)

var (
	// Force time package dependency for automated UpdatedAt/CreatedAt.
	_ = time.Second
	// Force qmhelper dependency for where clause generation (which doesn't
	// always happen)
	_ = qmhelper.Where
)

// One returns a single app record from the query.
func (q appQuery) One(ctx context.Context, exec boil.ContextExecutor) (*App, error) {
	o := &App{}

	queries.SetLimit(q.Query, 1)

	err := q.Bind(ctx, exec, o)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "portal: failed to execute a one query for apps")
	}

	return o, nil
}

// All returns all App records from the query.
func (q appQuery) All(ctx context.Context, exec boil.ContextExecutor) (AppSlice, error) {
	var o []*App

	err := q.Bind(ctx, exec, &o)
	if err != nil {
		return nil, errors.Wrap(err, "portal: failed to assign all query results to App slice")
	}

	return o, nil
}

// Count returns the count of all App records in the query.
func (q appQuery) Count(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return 0, errors.Wrap(err, "portal: failed to count apps rows")
	}

	return count, nil
}

// Exists checks if the row exists in the table.
func (q appQuery) Exists(ctx context.Context, exec boil.ContextExecutor) (bool, error) {
	var count int64

	queries.SetSelect(q.Query, nil)
	queries.SetCount(q.Query)
	queries.SetLimit(q.Query, 1)

	err := q.Query.QueryRowContext(ctx, exec).Scan(&count)
	if err != nil {
		return false, errors.Wrap(err, "portal: failed to check if apps exists")
	}

	return count > 0, nil
}

// Org pointed to by the foreign key.
func (o *App) Org(mods ...qm.QueryMod) orgQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.OrgID),
	}

	queryMods = append(queryMods, mods...)

	query := Orgs(queryMods...)
	queries.SetFrom(query.Query, "\"orgs\"")

	return query
}

// User pointed to by the foreign key.
func (o *App) User(mods ...qm.QueryMod) userQuery {
	queryMods := []qm.QueryMod{
		qm.Where("\"id\" = ?", o.UserID),
	}

	queryMods = append(queryMods, mods...)

	query := Users(queryMods...)
	queries.SetFrom(query.Query, "\"users\"")

	return query
}

// LoadOrg allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (appL) LoadOrg(ctx context.Context, e boil.ContextExecutor, singular bool, maybeApp interface{}, mods queries.Applicator) error {
	var slice []*App
	var object *App

	if singular {
		object = maybeApp.(*App)
	} else {
		slice = *maybeApp.(*[]*App)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &appR{}
		}
		if !queries.IsNil(object.OrgID) {
			args = append(args, object.OrgID)
		}

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &appR{}
			}

			for _, a := range args {
				if queries.Equal(a, obj.OrgID) {
					continue Outer
				}
			}

			if !queries.IsNil(obj.OrgID) {
				args = append(args, obj.OrgID)
			}

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`orgs`),
		qm.WhereIn(`orgs.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load Org")
	}

	var resultSlice []*Org
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice Org")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for orgs")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for orgs")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.Org = foreign
		if foreign.R == nil {
			foreign.R = &orgR{}
		}
		foreign.R.App = object
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if queries.Equal(local.OrgID, foreign.ID) {
				local.R.Org = foreign
				if foreign.R == nil {
					foreign.R = &orgR{}
				}
				foreign.R.App = local
				break
			}
		}
	}

	return nil
}

// LoadUser allows an eager lookup of values, cached into the
// loaded structs of the objects. This is for an N-1 relationship.
func (appL) LoadUser(ctx context.Context, e boil.ContextExecutor, singular bool, maybeApp interface{}, mods queries.Applicator) error {
	var slice []*App
	var object *App

	if singular {
		object = maybeApp.(*App)
	} else {
		slice = *maybeApp.(*[]*App)
	}

	args := make([]interface{}, 0, 1)
	if singular {
		if object.R == nil {
			object.R = &appR{}
		}
		args = append(args, object.UserID)

	} else {
	Outer:
		for _, obj := range slice {
			if obj.R == nil {
				obj.R = &appR{}
			}

			for _, a := range args {
				if a == obj.UserID {
					continue Outer
				}
			}

			args = append(args, obj.UserID)

		}
	}

	if len(args) == 0 {
		return nil
	}

	query := NewQuery(
		qm.From(`users`),
		qm.WhereIn(`users.id in ?`, args...),
	)
	if mods != nil {
		mods.Apply(query)
	}

	results, err := query.QueryContext(ctx, e)
	if err != nil {
		return errors.Wrap(err, "failed to eager load User")
	}

	var resultSlice []*User
	if err = queries.Bind(results, &resultSlice); err != nil {
		return errors.Wrap(err, "failed to bind eager loaded slice User")
	}

	if err = results.Close(); err != nil {
		return errors.Wrap(err, "failed to close results of eager load for users")
	}
	if err = results.Err(); err != nil {
		return errors.Wrap(err, "error occurred during iteration of eager loaded relations for users")
	}

	if len(resultSlice) == 0 {
		return nil
	}

	if singular {
		foreign := resultSlice[0]
		object.R.User = foreign
		if foreign.R == nil {
			foreign.R = &userR{}
		}
		foreign.R.Apps = append(foreign.R.Apps, object)
		return nil
	}

	for _, local := range slice {
		for _, foreign := range resultSlice {
			if local.UserID == foreign.ID {
				local.R.User = foreign
				if foreign.R == nil {
					foreign.R = &userR{}
				}
				foreign.R.Apps = append(foreign.R.Apps, local)
				break
			}
		}
	}

	return nil
}

// SetOrg of the app to the related item.
// Sets o.R.Org to related.
// Adds o to related.R.App.
func (o *App) SetOrg(ctx context.Context, exec boil.ContextExecutor, insert bool, related *Org) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"apps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"org_id"}),
		strmangle.WhereClause("\"", "\"", 2, appPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	queries.Assign(&o.OrgID, related.ID)
	if o.R == nil {
		o.R = &appR{
			Org: related,
		}
	} else {
		o.R.Org = related
	}

	if related.R == nil {
		related.R = &orgR{
			App: o,
		}
	} else {
		related.R.App = o
	}

	return nil
}

// RemoveOrg relationship.
// Sets o.R.Org to nil.
// Removes o from all passed in related items' relationships struct (Optional).
func (o *App) RemoveOrg(ctx context.Context, exec boil.ContextExecutor, related *Org) error {
	var err error

	queries.SetScanner(&o.OrgID, nil)
	if _, err = o.Update(ctx, exec, boil.Whitelist("org_id")); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	if o.R != nil {
		o.R.Org = nil
	}
	if related == nil || related.R == nil {
		return nil
	}

	related.R.App = nil
	return nil
}

// SetUser of the app to the related item.
// Sets o.R.User to related.
// Adds o to related.R.Apps.
func (o *App) SetUser(ctx context.Context, exec boil.ContextExecutor, insert bool, related *User) error {
	var err error
	if insert {
		if err = related.Insert(ctx, exec, boil.Infer()); err != nil {
			return errors.Wrap(err, "failed to insert into foreign table")
		}
	}

	updateQuery := fmt.Sprintf(
		"UPDATE \"apps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, []string{"user_id"}),
		strmangle.WhereClause("\"", "\"", 2, appPrimaryKeyColumns),
	)
	values := []interface{}{related.ID, o.ID}

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, updateQuery)
		fmt.Fprintln(writer, values)
	}
	if _, err = exec.ExecContext(ctx, updateQuery, values...); err != nil {
		return errors.Wrap(err, "failed to update local table")
	}

	o.UserID = related.ID
	if o.R == nil {
		o.R = &appR{
			User: related,
		}
	} else {
		o.R.User = related
	}

	if related.R == nil {
		related.R = &userR{
			Apps: AppSlice{o},
		}
	} else {
		related.R.Apps = append(related.R.Apps, o)
	}

	return nil
}

// Apps retrieves all the records using an executor.
func Apps(mods ...qm.QueryMod) appQuery {
	mods = append(mods, qm.From("\"apps\""))
	return appQuery{NewQuery(mods...)}
}

// FindApp retrieves a single record by ID with an executor.
// If selectCols is empty Find will return all columns.
func FindApp(ctx context.Context, exec boil.ContextExecutor, iD int, selectCols ...string) (*App, error) {
	appObj := &App{}

	sel := "*"
	if len(selectCols) > 0 {
		sel = strings.Join(strmangle.IdentQuoteSlice(dialect.LQ, dialect.RQ, selectCols), ",")
	}
	query := fmt.Sprintf(
		"select %s from \"apps\" where \"id\"=$1", sel,
	)

	q := queries.Raw(query, iD)

	err := q.Bind(ctx, exec, appObj)
	if err != nil {
		if errors.Cause(err) == sql.ErrNoRows {
			return nil, sql.ErrNoRows
		}
		return nil, errors.Wrap(err, "portal: unable to select from apps")
	}

	return appObj, nil
}

// Insert a single record using an executor.
// See boil.Columns.InsertColumnSet documentation to understand column list inference for inserts.
func (o *App) Insert(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) error {
	if o == nil {
		return errors.New("portal: no apps provided for insertion")
	}

	var err error

	nzDefaults := queries.NonZeroDefaultSet(appColumnsWithDefault, o)

	key := makeCacheKey(columns, nzDefaults)
	appInsertCacheMut.RLock()
	cache, cached := appInsertCache[key]
	appInsertCacheMut.RUnlock()

	if !cached {
		wl, returnColumns := columns.InsertColumnSet(
			appAllColumns,
			appColumnsWithDefault,
			appColumnsWithoutDefault,
			nzDefaults,
		)

		cache.valueMapping, err = queries.BindMapping(appType, appMapping, wl)
		if err != nil {
			return err
		}
		cache.retMapping, err = queries.BindMapping(appType, appMapping, returnColumns)
		if err != nil {
			return err
		}
		if len(wl) != 0 {
			cache.query = fmt.Sprintf("INSERT INTO \"apps\" (\"%s\") %%sVALUES (%s)%%s", strings.Join(wl, "\",\""), strmangle.Placeholders(dialect.UseIndexPlaceholders, len(wl), 1, 1))
		} else {
			cache.query = "INSERT INTO \"apps\" %sDEFAULT VALUES%s"
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
		return errors.Wrap(err, "portal: unable to insert into apps")
	}

	if !cached {
		appInsertCacheMut.Lock()
		appInsertCache[key] = cache
		appInsertCacheMut.Unlock()
	}

	return nil
}

// Update uses an executor to update the App.
// See boil.Columns.UpdateColumnSet documentation to understand column list inference for updates.
// Update does not automatically update the record in case of default values. Use .Reload() to refresh the records.
func (o *App) Update(ctx context.Context, exec boil.ContextExecutor, columns boil.Columns) (int64, error) {
	var err error
	key := makeCacheKey(columns, nil)
	appUpdateCacheMut.RLock()
	cache, cached := appUpdateCache[key]
	appUpdateCacheMut.RUnlock()

	if !cached {
		wl := columns.UpdateColumnSet(
			appAllColumns,
			appPrimaryKeyColumns,
		)

		if !columns.IsWhitelist() {
			wl = strmangle.SetComplement(wl, []string{"created_at"})
		}
		if len(wl) == 0 {
			return 0, errors.New("portal: unable to update apps, could not build whitelist")
		}

		cache.query = fmt.Sprintf("UPDATE \"apps\" SET %s WHERE %s",
			strmangle.SetParamNames("\"", "\"", 1, wl),
			strmangle.WhereClause("\"", "\"", len(wl)+1, appPrimaryKeyColumns),
		)
		cache.valueMapping, err = queries.BindMapping(appType, appMapping, append(wl, appPrimaryKeyColumns...))
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
		return 0, errors.Wrap(err, "portal: unable to update apps row")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "portal: failed to get rows affected by update for apps")
	}

	if !cached {
		appUpdateCacheMut.Lock()
		appUpdateCache[key] = cache
		appUpdateCacheMut.Unlock()
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values.
func (q appQuery) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	queries.SetUpdate(q.Query, cols)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to update all for apps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to retrieve rows affected for apps")
	}

	return rowsAff, nil
}

// UpdateAll updates all rows with the specified column values, using an executor.
func (o AppSlice) UpdateAll(ctx context.Context, exec boil.ContextExecutor, cols M) (int64, error) {
	ln := int64(len(o))
	if ln == 0 {
		return 0, nil
	}

	if len(cols) == 0 {
		return 0, errors.New("portal: update all requires at least one column argument")
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
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := fmt.Sprintf("UPDATE \"apps\" SET %s WHERE %s",
		strmangle.SetParamNames("\"", "\"", 1, colNames),
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), len(colNames)+1, appPrimaryKeyColumns, len(o)))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to update all in app slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to retrieve rows affected all in update all app")
	}
	return rowsAff, nil
}

// Upsert attempts an insert using an executor, and does an update or ignore on conflict.
// See boil.Columns documentation for how to properly use updateColumns and insertColumns.
func (o *App) Upsert(ctx context.Context, exec boil.ContextExecutor, updateOnConflict bool, conflictColumns []string, updateColumns, insertColumns boil.Columns) error {
	if o == nil {
		return errors.New("portal: no apps provided for upsert")
	}

	nzDefaults := queries.NonZeroDefaultSet(appColumnsWithDefault, o)

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

	appUpsertCacheMut.RLock()
	cache, cached := appUpsertCache[key]
	appUpsertCacheMut.RUnlock()

	var err error

	if !cached {
		insert, ret := insertColumns.InsertColumnSet(
			appAllColumns,
			appColumnsWithDefault,
			appColumnsWithoutDefault,
			nzDefaults,
		)
		update := updateColumns.UpdateColumnSet(
			appAllColumns,
			appPrimaryKeyColumns,
		)

		if updateOnConflict && len(update) == 0 {
			return errors.New("portal: unable to upsert apps, could not build update column list")
		}

		conflict := conflictColumns
		if len(conflict) == 0 {
			conflict = make([]string, len(appPrimaryKeyColumns))
			copy(conflict, appPrimaryKeyColumns)
		}
		cache.query = buildUpsertQueryPostgres(dialect, "\"apps\"", updateOnConflict, ret, update, conflict, insert)

		cache.valueMapping, err = queries.BindMapping(appType, appMapping, insert)
		if err != nil {
			return err
		}
		if len(ret) != 0 {
			cache.retMapping, err = queries.BindMapping(appType, appMapping, ret)
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
		return errors.Wrap(err, "portal: unable to upsert apps")
	}

	if !cached {
		appUpsertCacheMut.Lock()
		appUpsertCache[key] = cache
		appUpsertCacheMut.Unlock()
	}

	return nil
}

// Delete deletes a single App record with an executor.
// Delete will match against the primary key column to find the record to delete.
func (o *App) Delete(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if o == nil {
		return 0, errors.New("portal: no App provided for delete")
	}

	args := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(o)), appPrimaryKeyMapping)
	sql := "DELETE FROM \"apps\" WHERE \"id\"=$1"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args...)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to delete from apps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "portal: failed to get rows affected by delete for apps")
	}

	return rowsAff, nil
}

// DeleteAll deletes all matching rows.
func (q appQuery) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if q.Query == nil {
		return 0, errors.New("portal: no appQuery provided for delete all")
	}

	queries.SetDelete(q.Query)

	result, err := q.Query.ExecContext(ctx, exec)
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to delete all from apps")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "portal: failed to get rows affected by deleteall for apps")
	}

	return rowsAff, nil
}

// DeleteAll deletes all rows in the slice, using an executor.
func (o AppSlice) DeleteAll(ctx context.Context, exec boil.ContextExecutor) (int64, error) {
	if len(o) == 0 {
		return 0, nil
	}

	var args []interface{}
	for _, obj := range o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "DELETE FROM \"apps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, appPrimaryKeyColumns, len(o))

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, args)
	}
	result, err := exec.ExecContext(ctx, sql, args...)
	if err != nil {
		return 0, errors.Wrap(err, "portal: unable to delete all from app slice")
	}

	rowsAff, err := result.RowsAffected()
	if err != nil {
		return 0, errors.Wrap(err, "portal: failed to get rows affected by deleteall for apps")
	}

	return rowsAff, nil
}

// Reload refetches the object from the database
// using the primary keys with an executor.
func (o *App) Reload(ctx context.Context, exec boil.ContextExecutor) error {
	ret, err := FindApp(ctx, exec, o.ID)
	if err != nil {
		return err
	}

	*o = *ret
	return nil
}

// ReloadAll refetches every row with matching primary key column values
// and overwrites the original object slice with the newly updated slice.
func (o *AppSlice) ReloadAll(ctx context.Context, exec boil.ContextExecutor) error {
	if o == nil || len(*o) == 0 {
		return nil
	}

	slice := AppSlice{}
	var args []interface{}
	for _, obj := range *o {
		pkeyArgs := queries.ValuesFromMapping(reflect.Indirect(reflect.ValueOf(obj)), appPrimaryKeyMapping)
		args = append(args, pkeyArgs...)
	}

	sql := "SELECT \"apps\".* FROM \"apps\" WHERE " +
		strmangle.WhereClauseRepeated(string(dialect.LQ), string(dialect.RQ), 1, appPrimaryKeyColumns, len(*o))

	q := queries.Raw(sql, args...)

	err := q.Bind(ctx, exec, &slice)
	if err != nil {
		return errors.Wrap(err, "portal: unable to reload all in AppSlice")
	}

	*o = slice

	return nil
}

// AppExists checks if the App row exists.
func AppExists(ctx context.Context, exec boil.ContextExecutor, iD int) (bool, error) {
	var exists bool
	sql := "select exists(select 1 from \"apps\" where \"id\"=$1 limit 1)"

	if boil.IsDebug(ctx) {
		writer := boil.DebugWriterFrom(ctx)
		fmt.Fprintln(writer, sql)
		fmt.Fprintln(writer, iD)
	}
	row := exec.QueryRowContext(ctx, sql, iD)

	err := row.Scan(&exists)
	if err != nil {
		return false, errors.Wrap(err, "portal: unable to check if apps exists")
	}

	return exists, nil
}
