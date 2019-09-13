# Set Up

*Run mysql db in container, map to host port 8081, password is "password".

docker run --name mysql -p 8081:3306 -e MYSQL_ROOT_PASSWORD=password -d mysql
