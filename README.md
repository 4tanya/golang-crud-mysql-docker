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

**2. Copy file .env.dist to .env and set up your local database settings there:**

```
$ cp .env.dist .env
```

**3. Go to the main package folder and run it:**

```
$ cd main
```
```
$ go get .
```
```
$ go run .
```
