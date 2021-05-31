# keinos-sqlite3

## Pull Docker Image
```bash
docker pull keinos/sqlite3:latest
```

## Interactive

  Running `sqlite3 --version` command for example.
  
  ```bash
  $ docker run --rm -it keinos/sqlite3
  SQLite version 3.28.0 2019-04-16 19:49:53
  Enter ".help" for usage hints.
  Connected to a transient in-memory database.
  Use ".open FILENAME" to reopen on a persistent database.
  
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

## Run Container Image

```bash
docker run --name sqlite-db -d keinos/sqlite3
```
  
## Reference

- [Create/Backup/Restore SQLite](https://devopsheaven.com/sqlite/backup/restore/dump/databases/docker/2017/10/10/sqlite-backup-restore-docker.html)  