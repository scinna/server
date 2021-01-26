run:
	docker-compose run node yarn
	docker-compose up -d
build:
	go build
	docker build -t scinna/server:latest .
