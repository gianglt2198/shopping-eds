migrate:
	@migrate -path scripts \ 
		-database "postgresql://postgres:itsasecret@127.0.0.1:5432/postgres?sslmode=disable" \ 
		-verbose up