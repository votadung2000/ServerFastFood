uat:
	docker rmi -f mysql:8.0
	docker rmi -f server-fast-food:1.0
	docker-compose build --no-cache
	docker-compose up -d