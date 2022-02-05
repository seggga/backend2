docker_build:
	docker build -f dockerfile -t seggga/static_server .

docker_push:
	docker push seggga/static_server

docker_run:
	docker run -p 8080:8080 seggga/static_server