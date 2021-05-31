# Docker w/ DQLite

## Build Docker Image

```bash
docker build --no-cache -t dqlite .
```

## Persist DB file inside host OS folder `/home/dbfolder`

```bash
docker run -it -v ~/db/:/db dqlite
```

## Create a SQLite DB using Docker

Now it’s time to run the previous image to create and use a database called test.db:

```bash
docker run --name dqlite-db --rm -it -v `pwd`:/db dqlite test.db
docker run --name dqlite-db -v `pwd`:/db dqlite test.db
```

Once the container starts its entrypoint sqlite3 is executed and then you can run the following commands to create a table, insert values and select them:

```bash
sqlite> create table test_table(id int, description varchar(10));
sqlite> .tables
test_table
sqlite> insert into test_table values(1, 'foo');
sqlite> insert into test_table values(2, 'bar');
sqlite> select * from test_table;
1|foo
2|bar
sqlite> .exit
```

## Backup a SQLite DB using Docker

Command to backup test.db to a dump.sql file on host:

```bash
docker run --rm -it -v `pwd`:/db dqlite test.db .dump >> dump.sql
```

Finally, double check the dump.sql contains all SQLite queries:

```bash
$ cat dump.sql
PRAGMA foreign_keys=OFF;
BEGIN TRANSACTION;
CREATE TABLE test_table(id int, description varchar(10));
INSERT INTO test_table VALUES(1,'foo');
INSERT INTO test_table VALUES(2,'bar');
COMMIT;
```

## Restore a SQLite DB using Docker

Before restoring the DB make sure that the destination DB is empty (moving current DB to .old):

```bash
mv test.db test.db.old
```

Command to restore test.db:

```bash
cat dump.sql | docker run --rm -i -v `pwd`:/db some-sqlite test.db
```

## Reference

- [Create/Backup/Restore SQLite](https://devopsheaven.com/sqlite/backup/restore/dump/databases/docker/2017/10/10/sqlite-backup-restore-docker.html)