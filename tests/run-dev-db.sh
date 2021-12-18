echo "running and creating database..."
docker run \
  --name bloc-db \
  -e POSTGRES_USER=postgres \
  -e POSTGRES_PASSWORD=12345 \
  -e POSTGRES_DB=bloc \
  -p 5432:5432 \
  -d postgres

echo "waiting for container booting"
sleep 4

echo "creating tables"
docker container exec \
  -i bloc-db psql -U postgres -d bloc < ./database/sql/tables.sql

echo "insert fake test datas"
#docker container exec \
#  -i bloc-db psql -U postgres -d bloc < ./tests/dataset.sql

echo "done!"
