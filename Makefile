build:
	go build
	docker build -t scinna/server:latest .
