#!bin/bash
set -e


echo "runningg query multiple-databases"

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "postgres" <<-EOSQL
    CREATE DATABASE moneytest;
EOSQL

echo "success query multiple-databases"