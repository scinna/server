main:
	docker-compose up

install:
	docker run -it --rm -v "$(PWD)/frontend/app":/app -w /app node:13-alpine yarn
	docker run -it --rm -v "$(PWD)/frontend/setup":/app -w /app node:13-alpine yarn
