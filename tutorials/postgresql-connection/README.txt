
# start postgre db using docker, forward local port 
docker run --name postgres-db -e POSTGRES_PASSWORD=docker -p 5432:5432 -d postgres

# check ps 
$ docker ps
CONTAINER ID    IMAGE                                COMMAND                   CREATED           STATUS    PORTS                     NAMES
97460373705c    docker.io/library/postgres:latest    "docker-entrypoint.s…"    22 seconds ago    Up        0.0.0.0:5432->5432/tcp    postgres-db

# install psql if not installed 
brew install libpq
brew link --force libpq

# connect to postgres db, enter password "docker"
$ psql -h localhost -p 5432 -U postgres
Password for user postgres: 
psql (15.1)
Type "help" for help.

postgres=# help
You are using psql, the command-line interface to PostgreSQL.
Type:  \copyright for distribution terms
       \h for help with SQL commands
       \? for help with psql commands
       \g or terminate with semicolon to execute query
       \q to quit

postgres=# \du
                                   List of roles
 Role name |                         Attributes                         | Member of 
-----------+------------------------------------------------------------+-----------
 postgres  | Superuser, Create role, Create DB, Replication, Bypass RLS | {}

# change password 
postgres=# \password postgres
Enter new password for user "postgres": 
Enter it again: 
postgres=# 

postgres=# create database demo;
CREATE DATABASE

postgres-# \l
                                                List of databases
   Name    |  Owner   | Encoding |  Collate   |   Ctype    | ICU Locale | Locale Provider |   Access privileges   
-----------+----------+----------+------------+------------+------------+-----------------+-----------------------
 demo      | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | 
 postgres  | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | 
 template0 | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | =c/postgres          +
           |          |          |            |            |            |                 | postgres=CTc/postgres
 template1 | postgres | UTF8     | en_US.utf8 | en_US.utf8 |            | libc            | =c/postgres          +
           |          |          |            |            |            |                 | postgres=CTc/postgres
(4 rows)

postgres=# \c demo
You are now connected to database "demo" as user "postgres".
demo=# 


demo=# select * from times;
 id |      datetime       
----+---------------------
  0 | 2009-11-10 23:00:00
(1 row)



## go programming code 

go mod init app
go get github.com/jinzhu/gorm

$ go run main.go 
&{<nil> 0xc000316080}
Id: 0, Datetime: 0001-01-01 00:00:00 +0000 UTC
