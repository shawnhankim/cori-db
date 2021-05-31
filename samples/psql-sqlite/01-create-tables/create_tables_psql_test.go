// These tests assume there is a user sqlboiler_driver_user and a database
// by the name of sqlboiler_driver_test that it has full R/W rights to.
// In order to create this you can use the following steps from a root
// psql account:
//
//   create role sqlboiler_driver_user login nocreatedb nocreaterole nocreateuser password 'sqlboiler';
//   create database sqlboiler_driver_test owner = sqlboiler_driver_user;

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"testing"

	"github.com/volatiletech/sqlboiler/v4/drivers"
	pqdriver "github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql/driver"
)

var (
	envHostname = drivers.DefaultEnv("DRIVER_HOSTNAME", "localhost")
	envPort     = drivers.DefaultEnv("DRIVER_PORT", "5432")
	envUsername = drivers.DefaultEnv("DRIVER_USER", "portal")
	envPassword = drivers.DefaultEnv("DRIVER_PASS", "portal")
	envDatabase = drivers.DefaultEnv("DRIVER_DB", "portal")
)

func TestCreateTablesPostgres(t *testing.T) {
	b, err := ioutil.ReadFile("schema-pgres-1.sql")
	if err != nil {
		t.Fatal(err)
	}

	out := &bytes.Buffer{}
	createDB := exec.Command("/usr/local/opt/libpq/bin/psql", "-h", envHostname, "-U", envUsername, envDatabase)
	createDB.Env = append([]string{fmt.Sprintf("PGPASSWORD=%s", envPassword)}, os.Environ()...)
	createDB.Stdout = out
	createDB.Stderr = out
	createDB.Stdin = bytes.NewReader(b)

	if err := createDB.Run(); err != nil {
		fmt.Printf("psql output:\n%s\n", out.Bytes())
		//t.Fatal(err)
		panic(err)
	}
	fmt.Printf("psql output:\n%s\n", out.Bytes())

	config := drivers.Config{
		"user":    envUsername,
		"pass":    envPassword,
		"dbname":  envDatabase,
		"host":    envHostname,
		"port":    envPort,
		"sslmode": "disable",
		"schema":  "public",
	}

	p := &pqdriver.PostgresDriver{}
	info, err := p.Assemble(config)
	if err != nil {
		panic(err)
	}
	fmt.Println(info)

	// got, err := json.MarshalIndent(info, "", "\t")
	// if err != nil {
	// 	panic(err)
	// }

	// if *flagOverwriteGolden {
	// 	if err = ioutil.WriteFile("psql.golden.json", got, 0664); err != nil {
	// 		panic(err)
	// 	}
	// 	fmt.Println("wrote:", string(got))
	// 	return
	// }

	// want, err := ioutil.ReadFile("psql.golden.json")
	// if err != nil {
	// 	panic(err)
	// }

	// if bytes.Compare(want, got) != 0 {
	// 	fmt.Printf("want:\n%s\ngot:\n%s\n", want, got)
	// }
}
