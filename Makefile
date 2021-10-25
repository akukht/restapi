vegeta-load:
	vegeta attack -duration=10s -rate=10 -targets=vegeta/target.list -output=vegeta/attack-5.bin
	vegeta plot -title=Attack vegeta/attack-5.bin > vegeta/results.html

vegeta-simple:
	vegeta attack -duration=5s -rate=10 -targets=vegeta/target.list | tee results.bin | vegeta report

vegeta-create:
	vegeta attack -duration=10s -rate=10 -targets=vegeta/target-create.list -output=vegeta/attack-create.bin
	vegeta plot -title=Create vegeta/attack-create.bin > vegeta/results-create.html
	
lint:
	golangci-lint run

testdb:
	docker exec -it restapi_database_1 createdb --username=unicorn_user --owner=unicorn_user test_events
	migrate -path ./schema_test -database 'postgres://unicorn_user:magical_password@localhost:5432/test_events?sslmode=disable' up	