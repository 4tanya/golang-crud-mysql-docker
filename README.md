# golang-crud-mysql-docker

**1. Create a database and run a migration script:**

```
$ mysql -u root -p
Enter password:
```
```
mysql> create database library;
```
```
mysql> use library;
Database changed
```
```
mysql> source /path/to/migration.sql;
```
```
mysql> exit;
```

**2. Go to the main package folder and run it:**

```
$ cd main
```
```
$ go get .
```
```
$ go run .
```
