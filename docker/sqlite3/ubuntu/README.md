# SQLite Docker

## Dockerfile Example to install SQLite 3.0

```bash
FROM ubuntu:trusty
RUN sudo apt-get -y update
RUN sudo apt-get -y upgrade
RUN sudo apt-get install -y sqlite3 libsqlite3-dev
RUN mkdir /db
RUN /usr/bin/sqlite3 /db/test.db
CMD /bin/bash
```

## Persist DB file inside host OS folder `/home/dbfolder`

```bash
docker run -it -v /home/dbfolder/:/db imagename
```

## Build Docker Image

```bash
docker build --no-cache -t sqlite .
docker run   --name sqlite-db -v ~/db/:/db -d sqllite
```

