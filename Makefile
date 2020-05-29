.PHONY: build run drop_carts drop_orders drop_products drop_products++ drop_tables drop_users drop_counterparties drop_counterpartydocuments

build:
	go build -v ./cmd/efclient

run:
	go build -v ./cmd/efclient; ./efclient

drop_orders:
	psql -d b2b -a -f ./sql_scripts/drop_orders.sql

drop_products:
	psql -d b2b -a -f ./sql_scripts/drop_products.sql

drop_products++:
	psql -d b2b -a -f ./sql_scripts/drop_products++.sql

drop_tables:
	psql -d b2b -a -f ./sql_scripts/drop_tables.sql && rm -rf ./data/*.pdf

drop_users:
	psql -d b2b -a -f ./sql_scripts/drop_users.sql

drop_counterparties:
	psql -d b2b -a -f ./sql_scripts/drop_counterparties.sql

drop_counterpartydocuments:
	psql -d b2b -a -f ./sql_scripts/drop_counterpartydocuments.sql && rm -rf ./data/*.pdf

.DEFAULT_GOAL := build
