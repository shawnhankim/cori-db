# Sample Codes w/ Postgres

This repo provides sample codes and how to start w/ SQLBoilder which is a tool to generate a Go ORM tailored to your database schema.

## Getting started

### Download

```bash
# Go 1.16 and above:
go install github.com/volatiletech/sqlboiler/v4@latest
go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest
```

To install sqlboiler as a dependency in your project use the commands below inside of your go module's directory tree. This will install the dependencies into your go.mod file at the correct version.

```bash
# Do not forget the trailing /v4 and /v8 in the following commands
go get github.com/volatiletech/sqlboiler/v4
# Assuming you're going to use the null package for its additional null types
go get github.com/volatiletech/null/v8
```

### Generate Model

```bash
sqlboiler psql -o portal -p portal -c portal.toml
sqlboiler psql -c portal-psql.toml
```

## Reference

- [SQLBoiler](https://github.com/volatiletech/sqlboiler)
- [SQLBoiler: Getting Started](https://www.youtube.com/watch?v=y5utRS9axfg)
- [SQLBoiler: What's New in v3](https://www.youtube.com/watch?v=-B-OPsYRZJA)
- [SQLBoiler: Advanced Queries & Relationships](https://www.youtube.com/watch?v=iiJuM9NR8No)
