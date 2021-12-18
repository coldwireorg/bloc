#!/bin/bash
echo "loading tables..."
psql -U postgres -d bloc < /docker-entrypoint-initdb.d/tables.sql