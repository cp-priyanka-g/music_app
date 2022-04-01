# music_app
Music app using Golang, Gin Framework with mysql database

Requirements
Go 1.17

1.Run below command to install go dependencies
    -go mod tidy


2.Run go server
    -go run main.go

3. Start the docker to connect with database
    -sudo docker start local-mysql
    -sudo docker exec -it local-mysql bash 
        -mysql -uroot -p [Enter Password]
