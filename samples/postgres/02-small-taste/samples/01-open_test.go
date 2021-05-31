package main

import (
	"context"
	"database/sql"
	"fmt"
	"testing"

	_ "github.com/lib/pq"
	"github.com/shawnhankim/cori-db/samples/postgres/02-small-taste/models/portal"
	"github.com/volatiletech/sqlboiler/v4/boil"
	pqdriver "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

func TestOpen(t *testing.T) {
	// Open handle to database like normal
	// connStr := "dbname=portal host=localhost port=5432 user=portal pass=portal sslmode=disable"
	var (
		user      = "portal"
		pass      = "portal"
		dbname    = "portal"
		host      = "localhost"
		port      = 5432
		sslmode   = "disable"
		nTestCase = 1
	)
	connStr := pqdriver.PSQLBuildQueryString(user, pass, dbname, host, port, sslmode)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("ERROR:", err)
		return
	}
	fmt.Println("DB:", db)

	// SELECT COUNT(*) FROM pilots;

	ctx := context.Background()
	count, err := portal.Pilots().Count(ctx, db)
	handleErr(err)
	fmt.Println(nTestCase, count)
	nTestCase++

	// SELECT * FROM "pilots" LIMIT 5;
	pilots, err := portal.Pilots(qm.Limit(5)).All(ctx, db)
	handleErr(err)
	fmt.Println(nTestCase, pilots)
	nTestCase++

	// DELETE FROM "pilots" WHERE "id"=$1;
	nPilots, err := portal.Pilots(qm.Where("id=?", 1)).DeleteAll(ctx, db)
	handleErr(err)
	fmt.Println(nTestCase, nPilots, err)
	nTestCase++

	// type safe version of above
	nPilots, err = portal.Pilots(portal.PilotWhere.ID.EQ(1)).DeleteAll(ctx, db)
	handleErr(err)
	fmt.Println(nTestCase, nPilots, err)
	nTestCase++

	boil.SetDB(db)
	var p1 portal.Pilot
	p1.Name = "Larry1"
	err = p1.Insert(ctx, db, boil.Infer()) // Insert the first pilot with name "Larry"
	handleErr(err)
	fmt.Println(nTestCase, nPilots, err)
	nTestCase++

	var p2 portal.Pilot
	p2.Name = "Larry2"
	err = p2.Insert(ctx, db, boil.Infer()) // Insert the first pilot with name "Larry"
	handleErr(err)
	fmt.Println(nTestCase, nPilots, err)
	nTestCase++
}

func handleErr(err error) {
	if err != nil {
		panic(err)
	}
}

//   dbname    = "portal"
//   host      = "localhost"
//   port      = 5432
//   user      = "portal"
//   pass      = "portal"
//   schema    = "public"
//   sslmode   = "disable"
