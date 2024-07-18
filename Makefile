uat:
	docker rmi -f mysql:8.0
	docker-compose build --no-cache
	docker-compose up -d