# Dockerfile of PostgreSQL

Docker image of the PostgreSQL10.17.

## Build and Run Container Image

```bash
$ docker build --no-cache -t test-postgres .
$ docker run --name test-postgres-1 -p 5432:5432 -d test-postgres
```

## Getting Started

### Installation

Automated builds of the image are available on [Dockerhub](https://hub.docker.com/r/sameersbn/postgresql) and is the recommended method of installation.

| Note: Builds are also available on [Quay.io](https://quay.io/repository/sameersbn/postgresql)

```bash
$ docker pull sameersbn/postgresql:12-20200524
```

Alternatively you can build the image yourself.

```bash
$ docker build --no-cache -t postgresql ./ubutu
```

### Quickstart

Start PostgreSQL using:

```bash
$ docker run --name postgresql -itd --restart always \
             --publish 5432:5432                     \
             --volume postgresql:/var/lib/postgresql \
             sameersbn/postgresql:12-20200524
```

Login to the PostgreSQL server using:

```bash
$ docker exec -it postgresql sudo -u postgres psql
```

Alternatively, you can use the sample [docker-compose.yml](https://github.com/shawnhankim/docker-postgresql/blob/master/docker-compose.yml) file to start the container using [Docker Compose](https://docs.docker.com/compose/).

### Persistence

For PostgreSQL to preserve its state across container shutdown and startup you should mount a volume at /var/lib/postgresql. If you don't like the default volume destination then you can change it

> *The [Quickstart](#Quickstart) command already mounts a volume for persistence.*

SELinux users should update the security context of the host mountpoint so that it plays nicely with Docker:

```bash
$ mkdir -p /srv/docker/postgresql
$ chcon -Rt svirt_sandbox_file_t /srv/docker/postgresql
```

### Trusting local connections

By default connections to the PostgreSQL server need to authenticated using a password. If desired you can trust connections from the local network using the `PG_TRUST_LOCALNET` variable.

```bash
docker run --name postgresql -itd --restart always \
  --env 'PG_TRUST_LOCALNET=true' \
  sameersbn/postgresql:12-20200524
```

> **Note**
>
> The local network here is network to which the container is attached. This has different meanings depending on the --net parameter specified while starting the container. In the default configuration, this parameter would trust connections from other containers on the docker0 bridge.

### Setting `postgres` user password

By default the `postgres` user is not assigned a password and as a result you can only login to the PostgreSQL server locally. If you wish to login remotely to the PostgreSQL server as the `postgres` user, you will need to assign a password for the user using the `PG_PASSWORD` variable.

```bash
docker run --name postgresql -itd --restart always \
  --env 'PG_PASSWORD=passw0rd' \
  sameersbn/postgresql:12-20200524
```

> **Note**
>
> - When [persistence](#Persistence) is in use, `PG_PASSWORD` is effective on the first run.
> - This feature is only available in the `latest` and versions > `9.4-10`.

### Creating database user

A new PostgreSQL database user can be created by specifying the `DB_USER` and `DB_PASS` variables while starting the container.

```bash
docker run --name postgresql -itd --restart always \
           --env 'DB_USER=dbuser' --env 'DB_PASS=dbuserpass' \
           sameersbn/postgresql:12-20200524
```

> **Note**
>
> - The created user can login remotely
> - The container will error out if a password is not specified for the user
> - No changes will be made if the user already exists
> - Only a single user can be created at each launch

### Creating databases

A new PostgreSQL database can be created by specifying the `DB_NAME` variable while starting the container.

```bash
docker run --name postgresql -itd --restart always \
           --env 'DB_NAME=dbname' \
           sameersbn/postgresql:12-20200524
```

By default databases are created by copying the standard system database named `template1`. You can specify a different template for your database using the `DB_TEMPLATE` parameter. Refer to [Template Databases](http://www.postgresql.org/docs/9.4/static/manage-ag-templatedbs.html) for further information.

Additionally, more than one database can be created by specifying a comma separated list of database names in `DB_NAME`. For example, the following command creates two new databases named `dbname1` and `dbname2`.

*This feature is only available in releases greater than 9.1-1.*

```bash
docker run --name postgresql -itd --restart always \
           --env 'DB_NAME=dbname1,dbname2' \
           sameersbn/postgresql:12-20200524
```

### Granting user access to a database

If the `DB_USER` and `DB_PASS` variables are specified along with the `DB_NAME` variable, then the user specified in `DB_USER` will be granted access to all the databases listed in `DB_NAME`. Note that if the user and/or databases do not exist, they will be created.

```bash
docker run --name postgresql -itd --restart always           \
           --env 'DB_USER=dbuser' --env 'DB_PASS=dbuserpass' \
           --env 'DB_NAME=dbname1,dbname2'                   \
  sameersbn/postgresql:12-20200524
```

## References

- [Dockerfile of PostgreSQL 10.17](https://github.com/docker-library/postgres/tree/master/10)
- [sameersbn/postgresql:12-20200524](https://github.com/sameersbn/docker-postgresql)