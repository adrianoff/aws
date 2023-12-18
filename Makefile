up:
	cd docker && docker-compose up -d

build:
	cd docker && docker-compose build

up-force:
	cd docker && docker-compose up -d --force-recreate

stop:
	cd docker && docker-compose stop

down:
	cd docker && docker-compose down

restart: stop up

recreate: down up-force

aws_backend:
	docker exec -ti aws_backend bash
