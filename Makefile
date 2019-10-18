all: analysis unit_test build-artifact integration_test acceptance_test

analysis:
	golangci-lint run

unit_test:
	go test ./cmd/... --cover && go test ./internal/timesheet/... --cover

build-artifact:
	docker-compose build

integration_test:
	docker-compose up -d
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	go test ./internal/repository/... --cover
	docker-compose down
	
acceptance_test:
	docker-compose up -d
	newman run atdd/api/showSummaryTimesheetSuccess.json.json
	robot atdd/ui/timesheet.robot
	docker-compose down

down:
	docker-compose down