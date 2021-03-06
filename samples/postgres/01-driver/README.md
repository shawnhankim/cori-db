# Sample Code w/ PostgreSQL Driver

## Prerequisites

- First, update brew. From your command line, run the following command:

```bash
brew doctor
brew update
brew install libpq
```

- Second, symlink psql (and other libpq tools) into /usr/local/bin

```bash
brew link --force libpq ail
```

## Build PostgreSQL Docket

```bash
$ docker build --no-cache -t postgresql ./ubuntu
$ docker run --name portal-db -itd --restart always \
  --env 'DB_USER=portal' --env 'DB_PASS=portal' --env 'DB_NAME=portal' \
  -p 5432:5432 -d postgresql
```


## Reference

- [SQLBoiler: PostgreSQL Driver](https://github.com/volatiletech/sqlboiler/tree/master/drivers/sqlboiler-psql)