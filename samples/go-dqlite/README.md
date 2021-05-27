# go-dqlite
This repository provides sample code using `go-dqlite` Go package that containing bindings for the [dqlite](https://github.com/canonical/dqlite) C library and a pure-Go client for the dqlite wire [protocol](https://github.com/canonical/dqlite/blob/master/doc/protocol.md).

## Usage
The best way to understand how to use the go-dqlite package is probably by looking at the source code of the [demo program](https://github.com/canonical/go-dqlite/blob/master/cmd/dqlite-demo/dqlite-demo.go) and use it as example.

In general your application will use code such as:


```Go
dir := "/path/to/data/directory"
address := "1.2.3.4:666" // Unique node address
cluster := []string{...} // Optional list of existing nodes, when starting a new node
app, err := app.New(dir, app.WithAddress(address), app.WithCluster(cluster))
if err != nil {
        // ...
}

db, err := app.Open(context.Background(), "my-database")
if err != nil {
        // ...
}

// db is a *sql.DB object
if _, err := db.Exec("CREATE TABLE my_table (n INT)"); err != nil
        // ...
}
```

## Build
In order to use the go-dqlite package in your application, you'll need to have the [dqlite](https://github.com/canonical/dqlite) C library installed on your system, along with its dependencies. You then need to pass the -tags argument to the Go tools when building or testing your packages, for example:

```
go build -tags libsqlite3
go test  -tags libsqlite3
```

## Reference
- [go-dqlite](https://github.com/canonical/go-dqlite)