package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"os/exec"
	"testing"
	"time"

	_ "github.com/mattn/go-sqlite3"
	"github.com/volatiletech/sqlboiler/v4/drivers"

	sqdriver "github.com/volatiletech/sqlboiler-sqlite3/driver"
)

var (
	flagOverwriteGolden = flag.Bool("overwrite-golden",
		false, "Overwrite the golden file with the current execution results")
)

func TestDriver(t *testing.T) {
	var (
		sqlFile = "testdatabase.sql"
		dbFile  = os.Getenv("HOME") + "/db/test.db"
		// #dbFile = os.Getenv("HOME") + ""
	)

	rand.Seed(time.Now().Unix())
	b, err := ioutil.ReadFile(sqlFile)
	if err != nil {
		t.Fatal(err)
	}

	// tmpName := filepath.Join(os.TempDir(), fmt.Sprintf("sqlboiler-sqlite3-%d.sql", rand.Int()))
	os.Remove(dbFile)
	out := &bytes.Buffer{}
	createDB := exec.Command("sqlite3", dbFile) //tmpName)
	createDB.Stdout = out
	createDB.Stderr = out
	createDB.Stdin = bytes.NewReader(b)

	// t.Log("sqlite file:", tmpName)
	fmt.Println("SQLite Temp File: ", dbFile)
	if err := createDB.Run(); err != nil {
		t.Logf("sqlite output:\n%s\n", out.Bytes())
		t.Fatal(err)
	}
	t.Logf("sqlite output:\n%s\n", out.Bytes())

	config := drivers.Config{
		"dbname": dbFile,
	}

	s := &sqdriver.SQLiteDriver{}
	info, err := s.Assemble(config)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(info)
	// got, err := json.Marshal(info)
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if *flagOverwriteGolden {
	// 	if err = ioutil.WriteFile("sqlite3.golden.json", got, 0664); err != nil {
	// 		t.Fatal(err)
	// 	}
	// 	t.Log("wrote:", string(got))
	// 	return
	// }

	// want, err := ioutil.ReadFile("sqlite3.golden.json")
	// if err != nil {
	// 	t.Fatal(err)
	// }

	// if bytes.Compare(want, got) != 0 {
	// 	t.Errorf("want:\n%s\ngot:\n%s\n", want, got)
	// }
}
