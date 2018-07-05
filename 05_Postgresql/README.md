01. ls /etc/postgresql/10/main
02. service postgresql
03. sudo su postgres
03. psql
04. \l -> Show List of databases
05. \du -> Show List of roles
06. ALTER USER postgres WITH PASSWORD 'test123'; -> Change Password for default user 'postgres'
07. CREATE USER user_1 WITH PASSWORD 'test123';
08. ALTER USER user_1 WITH SUPERUSER;

09. CREATE USER user_2 WITH PASSWORD 'test123';
10. DROP USER user_2

CREATE DATABASE employees;
11. \c <database name> -> connect to a database
12. \c postgres -> switch back to postgres database
13. SELECT current_user; -> see current user
14. SELECT current_database(); -> see current database
15. DROP DATABASE <database name>; -> drop (remove, delete) database

12. \q to quit

nn. terminal: man psql

Ubuntu Software: Search pgAdmin
Connection:
- Name: localhost
- Host: 127.0.0.1
- Post: 5432
- Username: postgres (default)
- Password: test123

# ALTER
ALTER TABLE users ALTER COLUMN fullname DROP  NOT NULL;

ALTER TABLE users ADD CONSTRAINT users_username_unique UNIQUE (username);
ALTER TABLE users DROP CONSTRAINT users_username_unique;

CREATE EXTENSION pgcrypto;
