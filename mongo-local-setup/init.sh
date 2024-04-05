#!/bin/sh
set -e

echo "User: $DB_USERNAME"
echo "Password: $DB_PASSWORD"

mongoimport --host=mongo --db=it_store --collection=products --type=json --file=/products.json --jsonArray --username $MONGO_INITDB_ROOT_USERNAME --password $MONGO_INITDB_ROOT_PASSWORD --authenticationDatabase admin