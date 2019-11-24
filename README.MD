# Scinna

## Installation
1. Download the .tar.gz archive in the releases tab
2. Extract it wherever you want
3. Install postgresql and the postgresql-contrib package for your distribution and create a user + database for scinna
4. Install the pgcrypto extension 
```
$ sudo su postgres
$ psql
postgres=# \c [NOM BD SCINNA];
> You are now connected to database "scinna" as user "postgres".
scinna=# CREATE EXTENSION "pgcrypto";
> CREATE EXTENSION
scinna=# \q
$ psql -U [DB User for scinna] -d [DB for scinna] -a -f Initialize.sql
> Password for user scinna:
$ exit
```

5. Copy the .env.dist to .env and fill it accordingly
6. `$ ./scinna`