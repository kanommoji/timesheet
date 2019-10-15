all: analysis unit_test build-artifact integration_test acceptance_test

prepare:
	go get -u github.com/gin-gonic/gin

analysis:
	cd internal && golangci-lint run

unit_test:
	go test ./... --cover

build-artifact:
	docker-compose build

integration_test:
	docker-compose up -d
	docker exec -i my-mariadb mysql --user=root --password=root timesheet < atdd/data/prepare_timesheet.sql
	docker-compose down
	
acceptance_test:
	docker-compose up -d
	robot atdd/ui/timesheet.robot
	docker-compose down

down:
	docker-compose down