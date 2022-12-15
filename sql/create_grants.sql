-- create grants table

SET client_encoding TO 'UTF8';

-- comment if not use psql !!
\set ON_ERROR_STOP ON

GRANT ALL PRIVILEGES ON DATABASE "authdb" TO auth;
