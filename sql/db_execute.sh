#!/bin/bash

usage() { echo "Usage: $0 [-U <user>] [-h <host>] [-p <port>]" 1>&2; exit 1; }

read -s -p "Enter database password: " PGPASSWORD

while getopts ":h:p:U:" o; do
    case "${o}" in
        h)
            h=${OPTARG}
            ;;
        p)
            p=${OPTARG}
            ;;
        U)
            u=${OPTARG}
            ;;
        *)
            usage
            ;;
    esac
done
shift $((OPTIND-1))

if [ -z "${h}" ]
then
  HOST="localhost"
else
  HOST="${h}"
fi

if [ -z "${p}" ]
then
  PORT="5432"
else
  PORT="${p}"
fi

if [ -z "${u}" ]
then
  USER="postgres"
else
  USER="${u}"
fi

echo "HOST = $HOST"
echo "PORT = $PORT"
echo "USER = $USER"


psql -U $USER -h $HOST -p $PORT -f $PWD/create_all.sql
