ID=0
NAME=name
EMAIL=user@example.com

list:
	@curl -Ss http://localhost:8888/users

fetch:
	@curl -Ss http://localhost:8888/users/$(ID)

create:
	@curl -Ss -X POST -F 'name=$(NAME)' -F 'email=$(EMAIL)' http://localhost:8888/users

update:
	@curl -Ss -X PUT -F 'name=$(NAME)' -F 'email=$(EMAIL)' http://localhost:8888/users/$(ID)

delete:
	@curl -Ss -X DELETE http://localhost:8888/users/$(ID)
