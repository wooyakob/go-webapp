docker compose up -d //start couchbase server
docker compose ps //check status
docker compose logs -f couchbase //view logs
docker compose stop //stop container, keep data
docker compose down // rmv container, keep volume/data
docker compose down -v //rmv container and delete all CB data



UI: http://localhost:8091

