all: analysis unit_test build-artifact integration_test acceptance_test

analysis:
	golangci-lint run

unit_test:
	go test ./cmd/... && go test ./internal/timesheet/...

build-artifact:
	docker-compose build

integration_test:
	docker-compose up -d
	sleep 30
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	go test ./internal/repository
	docker-compose down
	
acceptance_test:
	docker-compose up -d
	sleep 30
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	sleep 10
	newman run atdd/api/timesheetSuccess.json
	docker-compose down

down:
	docker-compose down