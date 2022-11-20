source : https://www.youtube.com/watch?v=1E_YycpCsXw&list=PL5dTjWUk_cPYztKD7WxVFluHvpBNM28N9&index=3
GO And MYSQL 

diagram tool - https://app.diagrams.net/

bookstore apis 
database mysql 
gorm 
json marshall, unmarshall
project structure 
gorilla mux 

cmd / main.go
pkg 
    config  app.go 
    controlles  book-controller.go 
    models  book.go
    routes  bookstore-routes 
    utils   utils.go 

routes / controller funcs 
GET BOOKS /book/ POST 
CREATE BOOK /book/ GET 
GET BOOK BY ID /book/{bookid}   GET 
UPDATE BOOK /book/{bookid}  PUT 
DELETE BOOK /book/{bookid}  DELETE 






go mod init dev/bookstore
go get "github.com/jinzhu/gorm"
go get "github.com/jinzhu/gorm/dialects/mysql"
go get "github.com/gorilla/mux"

first write bookstore-routes.go file to define routes 


// start mysql server 
docker pull mysql/mysql-server:latest
docker images 
//start mysql on localhost
docker run -p 13306:3306 --name mysql -eMYSQL_ROOT_PASSWORD=pass -d  mysql/mysql-server:latest

docker ps
docker logs mysql   
...
[Entrypoint]   A random onetime password will be generated.
[Entrypoint] GENERATED ROOT PASSWORD: I6=aqv:Kh&kNP*6E98=_bH.5p4jD2v,6
run/mysqld/mysqlx.sock
2022-11-20T03:51:49.611933Z 0 [System] [MY-010931] [Server] /usr/sbin/mysqld: ready for connections. Version: '8.0.31'  socket: '/var/lib/mysql/mysql.sock'  port: 3306  MySQL Community Server - GPL.
...

docker inspect mysql

# connect to docker mysql server 
docker exec -it mysql -- bash

bash-4.4# mysql -u root -p
Enter password: 
Welcome to the MySQL monitor.  Commands end with ; or \g.
Your MySQL connection id is 18
Server version: 8.0.31 MySQL Community Server - GPL
Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.

mysql> 

# create new user and password 
mysql> create user 'demo'@'%' IDENTIFIED BY 'pass';
Query OK, 0 rows affected (0.02 sec)

mysql> GRANT ALL PRIVILEGES ON *.* to 'demo'@'%' WITH GRANT OPTION;
Query OK, 0 rows affected (0.01 sec)

mysql> FLUSH PRIVILEGES;
Query OK, 0 rows affected (0.00 sec)

mysql> CREATE DATABASE demo;
Query OK, 1 row affected (0.01 sec)

mysql> show databases;
+--------------------+
| Database           |
+--------------------+
| demo               |
...

mysql> use demo;
Database changed

DROP TABLE IF EXISTS `demo`;
CREATE TABLE demo (
    ID BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    UpdateAt TIMESTAMP, 
    DeleteAt TIMESTAMP, 
    name varchar(255),
    author varchar(255),
    publication varchar(255),
    PRIMARY KEY (`ID`)
)ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;

# below table is an example, but "gorm" package auto create table 
mysql> DROP TABLE IF EXISTS `demo`;
Query OK, 0 rows affected, 1 warning (0.01 sec)

mysql> CREATE TABLE demo (
    ->     ID BIGINT(20) UNSIGNED NOT NULL AUTO_INCREMENT,
    ->     CreatedAt TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    ->     UpdateAt TIMESTAMP, 
    ->     DeleteAt TIMESTAMP, 
    ->     name varchar(255),
    ->     author varchar(255),
    ->     publication varchar(255),
    ->     PRIMARY KEY (`ID`)
    -> )ENGINE=INNODB AUTO_INCREMENT=1 DEFAULT CHARSET=utf8;
Query OK, 0 rows affected, 2 warnings (0.06 sec)



# after this you can connect to mysql @ localhost using new user 
# use this user in go code 
(base) user:cmd user$ mysql --host=127.0.0.1 --port=13306 -u demo -p
Enter password: 
...
Server version: 8.0.31 MySQL Community Server - GPL
...
Type 'help;' or '\h' for help. Type '\c' to clear the current input statement.
mysql> 


# go code create a new table "books"
	// AutoMigrate will create tables
	db.AutoMigrate(&Book{})


# check table after insert 
mysql> show tables;
+----------------+
| Tables_in_demo |
+----------------+
| books          |
+----------------+
1 row in set (0.01 sec)

mysql> select * from books;
+----+---------------------+---------------------+------------+------+--------+-------------+
| id | created_at          | updated_at          | deleted_at | name | author | publication |
+----+---------------------+---------------------+------------+------+--------+-------------+
|  1 | 2022-11-20 05:11:44 | 2022-11-20 05:11:44 | NULL       | name | author | book1       |
+----+---------------------+---------------------+------------+------+--------+-------------+
1 row in set (0.00 sec)

mysql> 


# shutdown app 
$ go run main.go 
connected to mysql ...
^Csignal: interrupt

# stop mysql docker 
$ docker stop mysql
mysql
