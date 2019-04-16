docker stop test-mysql
docker rm test-mysql
docker run --name test-mysql -p3306:3306 -e MYSQL_ROOT_PASSWORD=password -e MYSQL_DATABASE=test -e MYSQL_USER=testuser -e MYSQL_PASSWORD=testpasswd -d mysql:5.7

export ENVIRONMENT=DEV
go run main.go

#curl --header "Content-Type: application/json" --request POST --data '{"name":"product2","price": "20.15" }' http://localhost:8081/product
