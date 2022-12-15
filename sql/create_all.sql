SET client_encoding TO 'UTF8';

\echo *********************** starting general data loading... `date`
\set ON_ERROR_STOP ON
set session authorization 'postgres';

\ir create_user.sql
\ir create_db.sql
\ir create_grants.sql;

\echo *********************** ending general data loading... `date`
ALTER ROLE thebox SET search_path TO 'public','authdb';
CREATE SCHEMA IF NOT EXISTS "public" AUTHORIZATION "auth";

\echo ***************************** Create db_version...

set schema 'public';
CREATE TABLE IF NOT EXISTS db_version
(
    version  VARCHAR,
    date_str timestamp
);
INSERT INTO db_version (version, date_str)
VALUES ('1.0.0', now());
