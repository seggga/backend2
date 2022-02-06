docker_build:
	docker build -f dockerfile -t seggga/static_server:1.0.1 .

docker_push:
	docker push seggga/static_server:1.0.1

docker_run:
	docker run -p 8080:8080 seggga/static_server