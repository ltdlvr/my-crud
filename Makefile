build:
	go build -o app .

run:
	go run main.go db.go

test:
	go test -v 

docker-up:
	docker-compose up --build

docker-down:
	docker-compose down