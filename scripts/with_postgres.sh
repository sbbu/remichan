#!/bin/bash

echo "Starting Backup Cronjob..."
crontab -l | { cat; echo "0 * * * * bash /meguca/scripts/backup_postgres.sh"; } | crontab -
cron

service postgresql start

echo "Setting up database user and permissions..."
su postgres -c "psql -c \"CREATE USER meguca WITH SUPERUSER PASSWORD 'meguca';\" || true"
su postgres -c "createdb -O meguca meguca || true"

echo "Waiting on PostgreSQL server..."
until pg_isready > /dev/null
do
    sleep 1
done

eval $@
