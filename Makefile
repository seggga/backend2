run:
	docker-compose --file docker-compose.yaml up -d

# connect to DB in container
# docker exec -ti mysql-container mysql --host=127.0.0.1 --port=3306 -u root -p test
# password = test


# add data to the database (POST)
# curl -X POST -F 'token=admin_secret_token' -F 'id=3' -F 'data=data_three' http://127.0.0.1:8002/entity
# curl -X POST -F 'token=admin_secret_token' -F 'id=4' -F 'data=data_four' http://127.0.0.1:8002/entity

# read entities (GET) in JSON
# curl http://127.0.0.1:8002/entities


# load test
# wrk -t1 -c1 -d 5m http://localhost:8002/entities
# wrk -t1 -c1 -d5m -s ./load_test/wrk.lua http://localhost:8002

