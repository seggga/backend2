run:
	docker-compose --file docker-compose.yaml up -d

# connect to DB in container
# docker exec -ti mysql-container mysql --host=127.0.0.1 --port=3306 -u root -p test
# password = test


# add data to the database (POST)
# curl -X POST -F 'token=admin_secret_token' -F 'id=3' -F 'data=data_three' http://172.26.177.139:8002/entity

# read entities (GET)
# curl http://172.26.177.139:8002/entities


