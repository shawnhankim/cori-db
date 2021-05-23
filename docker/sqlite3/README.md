# Dockerfile of SQLite3

Docker image of the latest SQLite3 version.

- Docker image: keinos/sqlite3:latest
- Repositories:
  - Image: https://hub.docker.com/r/keinos/sqlite3 @ DockerHub
  - Source: https://github.com/KEINOS/Dockerfile_of_SQLite3 @ GitHub
  - Issues: https://github.com/KEINOS/Dockerfile_of_SQLite3/issues @ GitHub

## Usage

- Pull the latest image

  ```bash
  docker pull keinos/sqlite3:latest
  ```

- Interactive

  Running `sqlite3 --version` command for example.
  
  ```bash
  $ docker run --rm -it keinos/sqlite3
  SQLite version 3.28.0 2019-04-16 19:49:53
  Enter ".help" for usage hints.
  Connected to a transient in-memory database.
  Use ".open FILENAME" to reopen on a persistent database.
  sqlite>
  sqlite> .quit
  ```

- Command

  Running sqlite3 --version command for example.
  
  ```bash
  $ docker run --rm keinos/sqlite3 sqlite3 --version
  3.28.0 2019-04-16 19:49:53 884b4b7e502b4e991677b53971277adfaf0a04a284f8e483e2553d0f83156b50
  ```

- Run test
  
  This container includes a [simple test script](./run-test.sh).
  
  You can run the script to see if the container and `sqlite3` binary is working.

  ```bash
  $ docker run --rm keinos/sqlite3 /run-test.sh
    /usr/bin/sqlite3
    - Creating test DB:
    rowid  timestamp            description            
    -----  -------------------  -----------------------
    1      2021-05-22 17:01:15  First sample data. Hoo 
    2      2021-05-22 17:01:15  Second sample data. Hoo
    - Query to run:
    select * from table_sample;
    - Result of query:
    2021-05-22 17:01:15|First sample data. Hoo
    2021-05-22 17:01:15|Second sample data. Hoo
    - Test result:
    success  
  $ docker run --name test-sqlite-1 -d keinos/sqlite3
  ```

## Build and Run Container Image

```bash
$ docker build --no-cache -t test-sqlite .
$ docker run --name test-sqlite-1 -p 81:80 -d test-sqlite
```

## References

- [Dockerfile of SQLite3](https://hub.docker.com/r/keinos/sqlite3)