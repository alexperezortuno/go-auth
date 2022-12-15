-- create a new database

SET client_encoding TO 'UTF8';

-- comment if not use psql !!
\set ON_ERROR_STOP ON
\set VERBOSITY verbose

-- create database if not exists
SELECT 'CREATE DATABASE authdb'
WHERE NOT EXISTS(SELECT FROM pg_database WHERE datname = 'authdb')
\gexec
