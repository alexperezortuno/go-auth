-- user creation

SET client_encoding TO 'UTF8';

-- comment if not use psql !!
\set ON_ERROR_STOP ON
\set VERBOSITY verbose

\du

DO
$$
    BEGIN
        CREATE ROLE auth WITH SUPERUSER CREATEDB CREATEROLE LOGIN ENCRYPTED PASSWORD 'Me.123';
    EXCEPTION
        WHEN duplicate_object THEN RAISE NOTICE '%, skipping', SQLERRM USING ERRCODE = SQLSTATE;
    END
$$;
